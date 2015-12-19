# zookeeper go client
zookeeper是由Apache基金会支持的高可用中间件,是随着hadoop的发展而诞生的, 它是Google Chubby的开源实现, 内部的主要算法是paxos. 项目源地址为:https://zookeeper.apache.org. 本工程主要是简单介绍zookper在Go语言环境下的使用.

## 安装
我们选取ubuntu-14.04作为测试环境, 我们可以通过网络包进行下载安装,但是也可以采用docker进行安装, 为了方便我们使用docker进行安装:

* 搜索docker镜像
```sh
root@gctest:~# docker search zookeeper
```

* 拉取mesos的zookeeper包
```sh
root@gctest:~# docker pull mesoscloud/zookeeper
```

* 启动zookeeper container
```sh
root@gctest:~# docker run --rm -p 2181:2181 mesoscloud/zookeeper
JMX enabled by default
Using config: /opt/zookeeper/bin/../conf/zoo.cfg
2015-12-19 10:37:16,730 [myid:] - INFO  [main:QuorumPeerConfig@103] - Reading configuration from: /opt/zookeeper/bin/../conf/zoo.cfg
2015-12-19 10:37:16,738 [myid:] - INFO  [main:DatadirCleanupManager@78] - autopurge.snapRetainCount set to 3
2015-12-19 10:37:16,738 [myid:] - INFO  [main:DatadirCleanupManager@79] - autopurge.purgeInterval set to 0
2015-12-19 10:37:16,739 [myid:] - INFO  [main:DatadirCleanupManager@101] - Purge task is not scheduled.
2015-12-19 10:37:16,740 [myid:] - WARN  [main:QuorumPeerMain@113] - Either no config or no quorum defined in config, running  in standalone mode
2015-12-19 10:37:16,771 [myid:] - INFO  [main:QuorumPeerConfig@103] - Reading configuration from: /opt/zookeeper/bin/../conf/zoo.cfg
2015-12-19 10:37:16,777 [myid:] - INFO  [main:ZooKeeperServerMain@95] - Starting server
2015-12-19 10:37:16,802 [myid:] - INFO  [main:Environment@100] - Server environment:zookeeper.version=3.4.6-1569965, built on 02/20/2014 09:09 GMT
2015-12-19 10:37:16,802 [myid:] - INFO  [main:Environment@100] - Server environment:host.name=a021b7861447
2015-12-19 10:37:16,803 [myid:] - INFO  [main:Environment@100] - Server environment:java.version=1.7.0_91
2015-12-19 10:37:16,803 [myid:] - INFO  [main:Environment@100] - Server environment:java.vendor=Oracle Corporation
2015-12-19 10:37:16,804 [myid:] - INFO  [main:Environment@100] - Server environment:java.home=/usr/lib/jvm/java-1.7.0-openjdk-1.7.0.91-2.6.2.1.el7_1.x86_64/jre
2015-12-19 10:37:16,804 [myid:] - INFO  [main:Environment@100] - Server environment:java.class.path=/opt/zookeeper/bin/../build/classes:/opt/zookeeper/bin/../build/lib/*.jar:/opt/zookeeper/bin/../lib/slf4j-log4j12-1.6.1.jar:/opt/zookeeper/bin/../lib/slf4j-api-1.6.1.jar:/opt/zookeeper/bin/../lib/netty-3.7.0.Final.jar:/opt/zookeeper/bin/../lib/log4j-1.2.16.jar:/opt/zookeeper/bin/../lib/jline-0.9.94.jar:/opt/zookeeper/bin/../zookeeper-3.4.6.jar:/opt/zookeeper/bin/../src/java/lib/*.jar:/opt/zookeeper/bin/../conf:
2015-12-19 10:37:16,805 [myid:] - INFO  [main:Environment@100] - Server environment:java.library.path=/usr/java/packages/lib/amd64:/usr/lib64:/lib64:/lib:/usr/lib
2015-12-19 10:37:16,805 [myid:] - INFO  [main:Environment@100] - Server environment:java.io.tmpdir=/tmp
2015-12-19 10:37:16,806 [myid:] - INFO  [main:Environment@100] - Server environment:java.compiler=<NA>
2015-12-19 10:37:16,810 [myid:] - INFO  [main:Environment@100] - Server environment:os.name=Linux
2015-12-19 10:37:16,810 [myid:] - INFO  [main:Environment@100] - Server environment:os.arch=amd64
2015-12-19 10:37:16,811 [myid:] - INFO  [main:Environment@100] - Server environment:os.version=3.13.0-32-generic
2015-12-19 10:37:16,811 [myid:] - INFO  [main:Environment@100] - Server environment:user.name=root
2015-12-19 10:37:16,812 [myid:] - INFO  [main:Environment@100] - Server environment:user.home=/root
2015-12-19 10:37:16,812 [myid:] - INFO  [main:Environment@100] - Server environment:user.dir=/
2015-12-19 10:37:16,822 [myid:] - INFO  [main:ZooKeeperServer@755] - tickTime set to 2000
2015-12-19 10:37:16,822 [myid:] - INFO  [main:ZooKeeperServer@764] - minSessionTimeout set to -1
2015-12-19 10:37:16,823 [myid:] - INFO  [main:ZooKeeperServer@773] - maxSessionTimeout set to -1
2015-12-19 10:37:16,852 [myid:] - INFO  [main:NIOServerCnxnFactory@94] - binding to port 0.0.0.0/0.0.0.0:2181
```

## 下载zookeeper安装包, 并通过命令行连接测试zookeeper
* 安装jre:
```sh
root@gctest:~# apt-get install openjdk-7-jre
...
root@gctest:~# java -version
java version "1.7.0_91"
OpenJDK Runtime Environment (IcedTea 2.6.3) (7u91-2.6.3-0ubuntu0.12.04.1)
OpenJDK 64-Bit Server VM (build 24.91-b01, mixed mode)
```

* 下载地址为: http://www.apache.org/dyn/closer.cgi/zookeeper/

* 解压包
```sh
root@gctest:~# tar xzvf zookeeper-3.4.7.tar.gz
```

* 通过命令行连接,并测试
```sh
root@gctest:~# bin/zkCli.sh -server 192.168.1.12:2181
...
WATCHER::

WatchedEvent state:SyncConnected type:None path:null

[zk: 192.168.1.12:2181(CONNECTED) 0] ls /
[zookeeper]
[zk: 192.168.1.12:2181(CONNECTED) 1] create /abc datatest
Created /abc
[zk: 192.168.1.12:2181(CONNECTED) 2] ls /
[abc, zookeeper]
[zk: 192.168.1.12:2181(CONNECTED) 3]
```

## 运行Go client
* 下载go zookeeper client. 我们还是通过go get来获取包. 这个包的地址是:https://github.com/samuel/go-zookeeper, 是samuel来进行开发的.
```sh
$ go get -v github.com/samuel/go-zookeeper/zk
```

* 新建一个main.go文件进行测试
```Go
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
```
我们在main中, 新建了一个Zookeeper的连接,并对根目录进行了持续监控.
zkClient端的操作:
```sh
[zk: 192.168.1.12:2181(CONNECTED) 3] create /t1 t1data
Created /t1
[zk: 192.168.1.12:2181(CONNECTED) 4] create /t2 t1data
Created /t2
```

Go 程序的输出
```sh
$ go run main.go
2015/12/19 18:47:24 Connected to 192.168.1.12:2181
2015/12/19 18:47:24 Authenticated: id=95061372886974465, timeout=4000
[abc zookeeper] &{Czxid:0 Mzxid:0 Ctime:0 Mtime:0 Version:0 Cversion:0 Aversion:0 EphemeralOwner:0 DataLength:0 NumChildren:2 Pzxid:2}
{Type:EventNodeChildrenChanged State:Unknown Path:/ Err:<nil> Server:}
[t1 abc zookeeper] &{Czxid:0 Mzxid:0 Ctime:0 Mtime:0 Version:0 Cversion:1 Aversion:0 EphemeralOwner:0 DataLength:0 NumChildren:3 Pzxid:4}
{Type:EventNodeChildrenChanged State:Unknown Path:/ Err:<nil> Server:}
[t2 t1 abc zookeeper] &{Czxid:0 Mzxid:0 Ctime:0 Mtime:0 Version:0 Cversion:2 Aversion:0 EphemeralOwner:0 DataLength:0 NumChildren:4 Pzxid:5}
```
从上面我们可以看到, 其中一个客户端创建了/t1和/t2文件, Go程序端接收到了消息,并打印出了日志.