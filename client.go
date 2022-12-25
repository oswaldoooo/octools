package main
import(
	"fmt"
	"os"
	"strings"
	"bufio"
	"net"
	"github.com/go-ini/ini"
)
const(
	CONFPATH="conf/site-conf.ini"
)
func checkerror(er error){
 	if er!=nil{
 		panic(er)
 	}
}
func Client(){
	cfg,err:=ini.Load(CONFPATH)
	checkerror(err)
	Port,_:=cfg.Section("Server").Key("Port").Int()
	Address:=cfg.Section("Server").Key("Address").String()
	Address=fmt.Sprintf("%v:%v",Address,Port)
	con,err:=net.Dial("tcp",Address)
	checkerror(err)
	for{
		fmt.Print("console-> ")
		read:=bufio.NewReader(os.Stdin)
		msg,_:=read.ReadString('\n')
		msg=strings.TrimSpace(msg)
		_,err:=con.Write([]byte(msg))
		checkerror(err)
		var buff [256]byte
		n,err:=con.Read(buff[:])
		checkerror(err)
		fmt.Printf("%v -> %v",Address,string(buff[:n]))
	}
}
func main(){
	Client()
}
