package main
import (
	"flag"
	"fmt"
)

func main() {
	var (
		str *string
		i *int
		vStr string
		vI int
	)
	str = flag.String("mystring", "default-value", "mystring is my test flag of String")
	i = flag.Int("myint", 123, "test int flag")
	flag.StringVar(&vStr, "mystring2", "default-value2", "mystring2 is my test flag of String")
	flag.IntVar(&vI, "myint2", 123, "test int flag")
	flag.Parse()

	flag.Usage()
	fmt.Println("str:", *str)
	fmt.Println("int:", *i)
	fmt.Println("vStr:", vStr)
	fmt.Println("vI:", vI)
}