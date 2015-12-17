# viper
viper是一个通用的程序配置解析工具包,其功能包括:
* 解析JSON, TOML, YAML 和 HCL配置文件
* 监控配置文件变化
* 解析环境变量
* 解析远程配置系统(etcd 或 Consul),并监控
* 解析命令行flags
* 解析内存配置

## viper解析顺序
由于viper支持多种配置, 因此存在一个解析顺序:
1. 通过Set方法的显示设置
2. flag参数设置
3. 环境变量
4. 配置文件
5. 键值对存储
6. 默认参数

## 使用viper
---
### 默认参数
```Go
    viper.SetDefault("ip", "127.0.0.1")
	ip := viper.GetString("ip")
	fmt.Println("ip:", ip)
```

### 读取配置文件
viper 支持JSON, TOML, YAML 和 HCL的文件,注意配置文件必须以小写的扩展名结尾. 同时,可以指定ConfigName, 添加ConfigPath来进行搜索.
```Go
    // config file
	viper.SetConfigName("test") // name of config file (without extension)
	viper.AddConfigPath("conf1")   // path to look for the config file in
	viper.AddConfigPath("conf2")   // path to look for the config file in
	viper.AddConfigPath("$HOME/")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("name:", viper.GetString("name"))
```

### 解析环境变量
环境变量支持设置前缀. 需要指出环境变量必须在bind之后才能Get.
```Go
    // env variable
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("id")
	os.Setenv("SPF_ID", "13") // typically done outside of the app
	fmt.Println("id:", viper.Get("id"))
```

### 和pflag一起使用
pflag是flag的扩展,可参考[pflag-example](../pflag-example)
```Go
    // work with pflag
	pflag.Int("port", 1138, "Port to run Application server on")
	viper.BindPFlag("port", pflag.Lookup("port"))
	fmt.Println("port:", viper.Get("port"))
```

### 和io.Reader一起使用
```Go
    // work with io.Reader
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)
	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	fmt.Println("age", viper.Get("age"))
```

### 访问嵌套对象
```Go
    // access nested object
	fmt.Println("clothing.jacket", viper.Get("clothing.jacket"))
```

### 所有支持的类型
* Get(key string) : interface{}
* GetBool(key string) : bool
* GetFloat64(key string) : float64
* GetInt(key string) : int
* GetString(key string) : string
* GetStringMap(key string) : map[string]interface{}
* GetStringMapString(key string) : map[string]string
* GetStringSlice(key string) : []string
* GetTime(key string) : time.Time
* GetDuration(key string) : time.Duration
* IsSet(key string) : bool

### unmarshal
viper可以将字符串反序列化为对象.
```Go
    Unmarshal(rawVal interface{}) : error
    UnmarshalKey(key string, rawVal interface{}) : error
```