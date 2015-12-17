package main
import (
	"github.com/spf13/viper"
	"fmt"
	"os"
	"github.com/spf13/pflag"
	"bytes"
)

func main() {
	// default value
	viper.SetDefault("ip", "127.0.0.1")
	ip := viper.GetString("ip")
	fmt.Println("ip:", ip)

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

	// env variable
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("id")
	os.Setenv("SPF_ID", "13") // typically done outside of the app
	fmt.Println("id:", viper.Get("id"))

	// work with pflag
	pflag.Int("port", 1138, "Port to run Application server on")
	viper.BindPFlag("port", pflag.Lookup("port"))
	fmt.Println("port:", viper.Get("port"))

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
	// access nested object
	fmt.Println("clothing.jacket", viper.Get("clothing.jacket"))
}
