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