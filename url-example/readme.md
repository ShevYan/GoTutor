# url-example
go 自带了url的包, 提供了基本的url封装

## url结构
```Go
type URL struct {
        Scheme   string
        Opaque   string    // encoded opaque data
        User     *Userinfo // username and password information
        Host     string    // host or host:port
        Path     string
        RawPath  string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
        RawQuery string // encoded query values, without '?'
        Fragment string // fragment for references, without '#'
}
```
如上所示, 一个URL是一个经过封装的对象, 其内部会对字符串进行解析. 相反, URI只是一个字符串, 并没有特殊的封装. URL包含了几个部分:
```sh
scheme://[userinfo@]host/path[?query][#fragment]
```

## URL编码及反编码
有些时候我们需要在http的请求地址中, 存放一些特殊字符, 比如空格, =, ?, /等. 为了安全起见, 我们需要对这些字符进行转码(escape), go的url包提供了全局的编码和反编码函数:
```Go
func QueryEscape
    QueryEscape escapes the string so it can be safely placed inside a URL query.
```

```Go
func QueryUnescape
    QueryUnescape does the inverse transformation of QueryEscape, converting %AB into the byte 0xAB and '+' into ' ' (space). It returns an error if any % is not followed by two hexadecimal digits.
```

### 请求参数
通过
```Go
type Values map[string][]string

func (u *URL) Query() Values
```
可以得到一个字符串的map, 可以很方便的进行操.