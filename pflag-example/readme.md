# pflag
pflag是flag扩展,主要是加入了对短参数的支持.

## 安装

```sh
    go get github.com/ogier/pflag
```

## 测试
```sh
    go test github.com/ogier/pflag
```

## 使用
如果你之前已经import了flag包,那么直接把那行改为:
```Go
    import flag "github.com/ogier/pflag"
```

##短参数支持
在函数后面加上P即可:
```Go
    pflag.StringVarP(&p1, "stringflag", "s", "test-p-flag-string", "test pflag for string.")
	pflag.IntVarP(&p2, "intflag", "i", 12345, "test pflag for int.")
```

```sh
Usage of /private/var/folders/98/pmb6b_8x5w5319m9ts6h6z6m0000gn/T/Build main.go and run1go:
  -i, --intflag=12345: test pflag for int.
  -s, --stringflag="test-p-flag-string": test pflag for string.
```