# Octools
## stand lib
get max and min from target array
```go
max,min:=tools.MaxandMin(your_array_name)
```
reverse string
```go
yourstringname=tools.Reverse(yourstringname)
```
sort array
```go
//normal sort
yourarrayname=tools.Sort(yourarrayname)
//sort plus
yourarrayname=tools.SortPlus(yourarrayname,sort_postion)
```
read conf from ini file.default conf file path:conf/site-conf.ini
```go
keyvaluemap:=remindmetools.ReadConfPlus(keyarray) //input a array that should include your all key you want,and it will return a key : value map to you
```
## **DataBase Tutorial**
#### You can use mysql database easier than before.It's based on github.com/jmoiron/sqlx
* Insert into table (...)values(...)
    ```go
    // insert into user_info (id,name,password)values("oc awesome","oswaldoooo","it's great!")
    err := dbcontroller.Insert(map[string]string{"id": "'" + template.HTMLEscapeString("oc awesome") + "'", "name": "'" + template.HTMLEscapeString("oswaldoooo") + "'", "password": "'" + template.HTMLEscapeString("it's great!") + "'"})
    //or use inserteasily.
    err:=dbcontroller.InsertEasily(map[string]string{"id":"oc awesome","name":"oswaldoooo","password":"it's great!"})
    ```
* select ... from table where ...=...
    ```go
    //select id,name,password from user_info where id='oc awesome'
    userinfo := struct{ Id, Name, Password string }{}
    err = dbcontroller.Get(&userinfo, "id", "'oc awesome'", "id", "name", "password")
    ```
* update table set ...... where ....
    ```go
    //update user_info set name='oswaldo' where id='oc awesome'
    err = dbcontroller.Update(map[string]string{"name": "'oswaldo'"}, "id", "'oc awesome'")
    // or use UpdateEasily
    err=dbcontroller.UpdateEasily(map[string]string{"name":"oswaldo","id":"oc awesome"})
    ```
* delete from table where ...
    ```go
    err = dbcontroller.Delete("id", "'oc awesome'")
    ```
## **JwtToken Tutorial**
It's based on github.com/golang-jwt/jwt

* **Register jwttoken service**
    ```go
    //this method will use default claims to create token
    jt := jwttoken.NewJwt()
    ```
* **Generate A New Token**
    ```go
    //you can put your user's unimportant data into the map,then generate a new token for user
    token,err:=jt.GenerateToken(map[string]string{})
    ```
* **Parse The Token**
    ```go
    //claim.Args is the map you put in when you generate the token
    claim, err := jt.ParseToken(token)
    ```
**BinarySearch**
```go
//the default match rate is 50%
datasotre.BinarySearch(content,origin_array)[]string
```
## **Pluginer Tutorial**
quick start your plugin mode in your program
* **prepare your plugin configuration file.It should look like this**
    ```xml
    <?xml version="1.0" encoding="UTF-8"?>
    <plugin_info>
            <!-- the name is your plugin file's name.you can contain '.so',or not -->
        <plugin classname="test" name="greet"/>
        <!-- here is your own path that your project.and your project directory should also contains plugin directory.the directory name is 'plugin'-->
        <rootpath>/Users/oswaldo/dev/golang/examples</rootpath>
    </plugin_info>
    ```
* **set the pluginer in your project**
    ```go
    var coremap = map[string]func(*plugin.Plugin) error{"test": loadnormal}
    //here is the path of your plugin configuration file.And the coremap(map[classname]parsing method)
    pluginer, err := pluginer.CreatePluginer("/Users/oswaldo/dev/golang/examples/site.xml", coremap)
    ```
* **and your parsing method should like this**
    ```go
    //input *plugin.Plugin,return error
    func(*plugin.Plugin)error
    //this is an example
    func loadnormal(pluginer *plugin.Plugin) (err error) {
        srm, err := pluginer.Lookup("Pattern")
        if err == nil {
            pattern := *srm.(*string)
            srm, err = pluginer.Lookup("Greet")
            if err == nil {
                resfunc := srm.(func(string, *int) error)
                testfunc[pattern] = resfunc
            }
        }
        return
    }
    ```
* **finally,set your plugin method as global variable**
## **Terminal**
**Process Bar**
```go
ProBar:=toolsbox.ProcessInit('=')//init process bar,and fill by char '='
go ProcessRun(ProBar)//listen the program's process
for i:=0;i<100;i++{
    ProBar.Pos<-uint(i)//tell the program's process
}
toolsbox.ProcessFinished(ProBar)//finished the process,dont care finished more,it'll be finished when it not finished,or do nothing there
```