package main
import (
	"net/url"
	"fmt"
)

//type URL struct {
//	Scheme   string
//	Opaque   string    // encoded opaque data
//	User     *Userinfo // username and password information
//	Host     string    // host or host:port
//	Path     string
//	RawPath  string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
//	RawQuery string // encoded query values, without '?'
//	Fragment string // fragment for references, without '#'
//}

func main() {
	var uri string
	var myurl *url.URL

	uri = "http://abc.com?p1=中国&p2=日本"
	myurl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	fmt.Println("URI:", myurl.RequestURI())
	fmt.Println("Values:", myurl.Query())
	fmt.Println("url编码(全局):", url.QueryEscape(uri))

	str, err := url.QueryUnescape(url.QueryEscape(uri))
	fmt.Println("url反编码(全局):", str)
}