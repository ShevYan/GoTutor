# cobra
---
cobra是一个命令行框架, 它可以生成和解析命令行工程.

## 安装
cobra的源工程地址在github.com/spf13/cobra/cobra, 同其他Go开源项目一样,我们通过go get来拉取:
```sh
$ go get -v github.com/spf13/cobra/cobra
```

安装完成之后, 通过下面语句import到工程:
```Go
import "github.com/spf13/cobra"
```

## 通过cobra命令生成工程
```sh
$ cobra init cobra-example
Your Cobra application is ready at
/Users/dongyan/Documents/go/src/github.com/ShevYan/GoTutor/cobra-example
Give it a try by going there and running `go run main.go`
Add commands to it by running `cobra add [cmdname]`
```

生成后的工程目录结构如下:
```sh
$ tree cobra-example
cobra-example
├── LICENSE
├── cmd
│   └── root.go
├── cobra-example.iml
├── main.go
└── readme.md

1 directory, 5 files
```

cobra生成的工程, 默认文件都放在cmd目录下, 以root.go作为根命令. cobra可以嵌套[flag](../flag-example), [pflag](../pflag-example), [viper](../viper-example)一起使用. 其中init()里面调用了initConfig(), 并可以调用pflag:

```Go
func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-example.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
```
flag的操作中, 分为全局的flag和当前子命令的flag. 什么是子命令呢? cobra支持多级子命令, 比如: mycmd sub1 sub2, 其中mycmd是根命令, sub1是mycmd的子命令, sub2是sub1的子命令. 因此, RootCmd.PersistentFlags()是一个全局flag, RootCmd.Flags()是当前命令的flag. 由于是根命令,其实二者是相同的, 但是如果在子命令中, 二者是不同的.

下面是使用viper的地方:
```Go
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".cobra-example") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
```

cobra对象初始化:
```Go
// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cobra-example",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
// Uncomment the following line if your bare application
// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { },
}
```
上面这个对象就是cobra的根命令对象, 默认生成的时候Run字段是被注释的, 打开以后就可以在函数里面写入自己的处理函数. 下面是运行结果:
```sh
$ go run main.go -h
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra-example [flags]

Flags:
      --config="": config file (default is $HOME/.cobra-example.yaml)
  -t, --toggle[=false]: Help message for toggle

```

##通过cobra命令生成子命令
我们执行下列命令:
```sh
cobra add serve
cobra add config
cobra add create -p 'configCmd'
```
之后,工程结构变为:
```sh
$ tree cobra-example
cobra-example
├── LICENSE
├── cmd
│   ├── config.go
│   ├── create.go
│   ├── root.go
│   └── serve.go
├── cobra-example.iml
├── main.go
└── readme.md

1 directory, 8 files

```

通过上述命令,我们生成了3个子命令, config, serve, create:
```sh
$ go run main.go -h
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra-example [flags]
  cobra-example [command]

Available Commands:
  config      A brief description of your command
  serve       A brief description of your command

Flags:
      --config="": config file (default is $HOME/.cobra-example.yaml)
  -t, --toggle[=false]: Help message for toggle
```

执行config子命令
```sh
$ go run main.go config
config called
```

执行config的create子命令
```sh
$ go run main.go config create
create called
```

我们可以看到代码里面config是作为根命令的子命令:
```Go
func init() {
	RootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
```

## 综合使用flag和viper
* 添加.cobra-example.yaml
```
author: ShevYan
```
* 添加flag解析
```Go
func init() {
	...
	RootCmd.PersistentFlags().Parsed()
}
```
* 添加author变量
```Go
var author string
// initConfig reads in config file and ENV variables if set.
func initConfig() {
    ...
    viper.AddConfigPath(".")
	...
	author = viper.GetString("author")
}
```
* 打印参数
```Go
// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	...
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config file:", cfgFile)
		fmt.Println("author:", author)
	},
}
```
