package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect([]string{"192.168.1.12"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	for {
		children, stat, ch, err := c.ChildrenW("/")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v %+v\n", children, stat)
		e := <-ch
		fmt.Printf("%+v\n", e)
	}
}