package main
import "github.com/spf13/pflag"

func main() {
	var (
		p1 string
		p2 int
	)

	pflag.StringVarP(&p1, "stringflag", "s", "test-p-flag-string", "test pflag for string.")
	pflag.IntVarP(&p2, "intflag", "i", 12345, "test pflag for int.")
	pflag.Usage()
}
