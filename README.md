#Octools
##stand lib
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