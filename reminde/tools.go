package remindmetools

import (
	"database/sql"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)
const(
	CONFPATH="conf"
)
var ostype=runtime.GOOS
type DB struct{
	// host string
	// port int
	// database string
	// usr string
	// pwd string
	dbaddr string
}
type Task struct{
	userid string
	date sql.NullString
	time sql.NullString
	tasks string
	emailadd string
}
var origin_email string
// var mailpwd string
func checkerror(er error){
	if er!=nil{
		panic(er)
	}
}
// READ CONF FILE from path
func ReadConfPlus(parent string,args []string,filepath string) map[string]string{
	res:=make(map[string]string)
	f,err:=ini.Load(filepath)
	checkerror(err)
	parentnode:=f.Section(parent)
	for _,v:=range args{
		res[v]=parentnode.Key(v).String()
	}
	return res
}
func Readpwd(mail string) string{
	filename:=fmt.Sprintf("%vsite-conf.ini",CONFPATH)
	f,err:=ini.Load(filename)
	checkerror(err)
	mailpwd:=f.Section("User").Key(mail).String()
	return mailpwd
}
func Reminde(cmd string,db *DB) {
	cmdarr:=strings.Split(cmd, "::")
	cmdmap:=make(map[string]string)
	for _,v:=range cmdarr{
		varr:=strings.Split(v, "--")
		cmdmap[varr[0]]=varr[1]
	}
	if _,ok:=cmdmap["taskid"];ok{
		res:=Found(cmdmap["taskid"],db)

		Readpwd(origin_email)
		fmt.Printf("You got it!...%+v\n",res)
		// foo:=New(origin_email,mailpwd)
		// mess:=fmt.Sprintf("Your task is time:\n%v\nSend by CodeLab-ToDo",res.tasks)
		// foo.Send("Your task arrived!",mess,res.emailadd)
	}
}
func Found(value string,db *DB) (*Task){
	d,err:=sql.Open("mysql",db.dbaddr)
	defer d.Close()
	checkerror(err)
	esql:=fmt.Sprintf("select user,target_date,target_time,tasks from todo where taskid='%v'",value)
	var task Task
	err=d.QueryRow(esql).Scan(&task.userid,&task.date,&task.time,&task.tasks)
	checkerror(err)

	if !task.date.Valid{
		go func (db *sql.DB)  {
			esql:=fmt.Sprintf("update todo set target_date='%v' where taskid='%v'",time.Now().Format("2006-01-02"),value)
			_,err:=db.Exec(esql)
			checkerror(err)
		}(d)
	}
	// fmt.Printf("%+v\n",task)
	esql=fmt.Sprintf("select origin_email,to_email from users where user='%v'",task.userid)
	err=d.QueryRow(esql).Scan(&origin_email,&task.emailadd)
	// debug line
	// fmt.Printf("esql:%v\norigin_email:%v,to_email:%v\n",esql,origin_email,task.emailadd)
	checkerror(err)
	esql=fmt.Sprintf("update todo set remind=%v where taskid='%v'",true,value)
	_,err=d.Exec(esql)
	checkerror(err)
	return &task
}

func SetReminder(omail string,tmail string){

}
func (s *DB) SetDB(host string,port int,usr string,pwd string,db string){
	// s.host=host
	// s.database=db
	// s.port=port
	// s.pwd=pwd
	// s.usr=usr
	s.dbaddr=fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",usr,pwd,host,port,db)
	fmt.Printf("Set mysql url:%v\n",s.dbaddr)
}
type remindthing struct{
	usr string
	task string
	taskid string
	omail string
	tmail string
	pwd string
}



func CheckTasks(db *DB,mailcore func(ac,pwd,title,body,to string)bool){
	d,err:=sql.Open("mysql",db.dbaddr)
	defer d.Close()
	checkerror(err)
	datenow:=time.Now().Format("2006-01-02")
	timenow:=time.Now().Format("15:04")
	esql:=fmt.Sprintf("select user,tasks,taskid from todo where remind=%v and target_date='%v' and target_time='%v'",true,datenow,timenow)
	rows,err:=d.Query(esql)
	checkerror(err)
	esql=fmt.Sprintf("update todo set remind=%v where remind=%v and target_date='%v' and target_time='%v'",false,true,datenow,timenow)
	_,err=d.Exec(esql)
	checkerror(err)
	for rows.Next(){
		var usrs remindthing
		err=rows.Scan(&usrs.usr,&usrs.task,&usrs.taskid)
		usrs.pwd=Readpwd(usrs.usr)
		checkerror(err)
		go SendMail(&usrs,db,mailcore)
	}
}
func SendMail(value *remindthing,db *DB,mailcore func(ac,pwd,title,body,to string)bool){
	d,err:=sql.Open("mysql",db.dbaddr)
	checkerror(err)
	esql:=fmt.Sprintf("select origin_email,to_email from users where user='%v'",value.usr)
	err=d.QueryRow(esql).Scan(&value.omail,&value.tmail)
	checkerror(err)
	value.pwd=Readpwd(value.usr)
	// fmt.Printf("User:%v\nTask:%v\nSendMail:%v\nAcceptMail:%v\n",value.usr,value.task,value.omail,value.tmail)
	f,err:=ini.Load(fmt.Sprintf("%v%v",CONFPATH,"site-conf.ini"))
	checkerror(err)
	title:=f.Section("Email").Key("title").String()
	body:=f.Section("Email").Key("body").String()
	body=fmt.Sprintf("%v%v",body,value.task)
	if ok:=mailcore(value.omail,value.pwd,title,body,value.tmail);!ok{
		fmt.Println("Mission failed...")
	}else{
		fmt.Println("Mission success")
	}
}