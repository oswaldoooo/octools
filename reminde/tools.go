package remindmetools

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/go-ini/ini"
)
const(
	CONFPATH="conf/"
)
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
	date string
	time string
	tasks string
	emailadd string
}
var origin_email string
var mailpwd string
func checkerror(er error){
	if er!=nil{
		panic(er)
	}
}
func Readpwd(mail string){
	filename:=fmt.Sprintf("%vsite-conf.ini",CONFPATH)
	f,err:=ini.Load(filename)
	checkerror(err)
	mailpwd=f.Section("User").Key(origin_email).String()
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
		foo:=New(origin_email,mailpwd)
		mess:=fmt.Sprintf("Your task is time:\n%v\nSend by CodeLab-ToDo",res.tasks)
		foo.Send("Your task arrived!",mess,res.emailadd)
	}
}
func Found(value string,db *DB) (*Task){
	d,err:=sql.Open("mysql",db.dbaddr)
	checkerror(err)
	esql:=fmt.Sprintf("select user,target_date,target_time,tasks from todo where taskid='%v'",value)
	var task Task
	err=d.QueryRow(esql,1).Scan(&task.userid,&task.date,&task.time,&task.tasks)
	checkerror(err)
	esql=fmt.Sprintf("select origin_email,to_email from users where user=%v",&task.userid)
	err=d.QueryRow(esql,1).Scan(&origin_email,&task.emailadd)
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
	s.dbaddr=fmt.Sprintf("%v:%v@%v:%v/%v?charset=utf8",usr,pwd,host,port,db)
}