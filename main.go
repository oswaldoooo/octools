package main
import(
	"fmt"
	"bufio"
	"github.com/go-ini/ini"
	"os"
	"strings"
)
func checkerror(er error){
	if er!=nil{
		panic(er)
	}
}
func Iniread(){
	fmt.Print("conf name: ")
	read:=bufio.NewReader(os.Stdin)
	msg,err:=read.ReadString('\n')
	checkerror(err)
	msg=strings.TrimSpace(msg)
	msg="conf/"+msg
	cfg,err:=ini.Load(msg)
	checkerror(err)
	serverport,er:=cfg.Section("Server").Key("Port").Int()
	checkerror(er)
	fmt.Printf("Server Port:%v",serverport)
}

