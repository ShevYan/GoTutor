#flag-example

flag-example是一个使用Golang默认flag包的例子. flag跟以往的传统命令行解析不同,它将参数定义, 参数默认值, 命令帮助, 参数读取集成到一起,摆脱以前繁琐的命令判断流程.

##返回指针
通过flag.String, flag.Int, flag.Bool等可以直接返回一个参数指针,因此后续使用的时候需要加*进行取值:
```Go
    str = flag.String("mystring", "default-value", "mystring is my test flag of String")
    i = flag.Int("myint", 123, "test int flag")
```

##序列化到变量
通过flag.StringVar, flag.IntVar, flag.BoolVar等,可以直接把值序列化到变量:
```Go
    flag.StringVar(&vStr, "mystring2", "default-value2", "mystring2 is my test flag of String")
	flag.IntVar(&vI, "myint2", 123, "test int flag")
```

##程序运行事例
```Bash
go run main.go
$ go run main.go
Usage of /var/folders/98/pmb6b_8x5w5319m9ts6h6z6m0000gn/T/go-build744639267/command-line-arguments/_obj/exe/main:
  -myint int
        test int flag (default 123)
  -myint2 int
        test int flag (default 123)
  -mystring string
        mystring is my test flag of String (default "default-value")
  -mystring2 string
        mystring2 is my test flag of String (default "default-value2")
str: default-value
int: 123
vStr: default-value2
vI: 123
```