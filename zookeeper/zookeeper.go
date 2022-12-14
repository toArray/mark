package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func main() {
	//连接地址
	hosts := []string{"127.0.0.1:2181", "127.0.0.1:2182", "127.0.0.1:2183"}

	//node
	path := "/watch_get"

	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	defer conn.Close()
	if err != nil {
		fmt.Printf("zookeeper connect is faield. err:%v\n", err)
		return
	}

	//创建节点
	var data = []byte("test value")
	acls := zk.WorldACL(zk.PermAll)
	nodePath, err := conn.Create(path, data, zk.FlagEphemeral, acls)
	if err != nil {
		fmt.Printf("zookeeper create node is faield. err:%v\n", err)
		return
	}

	//获得节点信息
	nodeValue, sate, err := conn.Get(path)
	if err != nil {
		fmt.Printf("zookeeper get node is faield. err:%v\n", err)
		return
	}
	fmt.Printf("nodeValue:%v, nodeSate:%+v", string(nodeValue), sate)

	//设置节点信息
	_, err = conn.Set(nodePath, []byte("hello zookeeper"), sate.Version)
	if err != nil {
		fmt.Printf("zookeeper set node is faield. err:%v\n", err)
		return
	}

	//获得节点信息
	nodeValue, sate, err = conn.Get(path)
	if err != nil {
		fmt.Printf("zookeeper get node is faield. err:%v\n", err)
		return
	}
	fmt.Printf("----nodeValue:%v, nodeSate:%+v", string(nodeValue), sate)

	//删除节点
	err = conn.Delete(nodePath, sate.Version)
	if err != nil {
		fmt.Printf("zookeeper delete node is faield: %v\n", err)
		return
	}

	//监听节点创建/删除/修改
	go watchExistsW(conn, path)

	//监听节点数据变化
	go watchGetW(conn, path)

	//监听节点孩子变化
	go childrenW(conn, path)

	for {
		select {
		default:

		}
	}

}

/*
watchExistsW
@Desc 监听节点创建/删除/修改
*/
func watchExistsW(conn *zk.Conn, path string) {
	for {
		b, sate, e, err := conn.ExistsW(path)
		if err != nil {
			fmt.Printf("watchExistsW is failed. err:%s\n", err.Error())
			return
		}

		event := <-e
		fmt.Println("监听到节点变化")
		fmt.Println("b: ", b)
		fmt.Println("sate: ", sate)
		fmt.Println("path: ", event.Path)
		fmt.Println("type: ", event.Type.String())
		fmt.Println("state: ", event.State.String())
		fmt.Println("---------------------------")
	}
}

/*
watchExistsW
@Desc 监听节点创建/删除/修改
*/
func watchGetW(conn *zk.Conn, path string) {
	for {
		b, sate, e, err := conn.GetW(path)
		if err != nil {
			fmt.Printf("watchGetW is failed. err:%s\n", err.Error())
			return
		}

		event := <-e
		fmt.Println("监听到节点数据修改")
		fmt.Println("value: ", string(b))
		fmt.Println("sate: ", sate)
		fmt.Println("path: ", event.Path)
		fmt.Println("type: ", event.Type.String())
		fmt.Println("state: ", event.State.String())
		fmt.Println("---------------------------")
	}
}

/*
childrenW
@Desc 监听节点创建/删除/修改
*/
func childrenW(conn *zk.Conn, path string) {
	for {
		b, sate, e, err := conn.ChildrenW(path)
		if err != nil {
			fmt.Printf("watchGetW is failed. err:%s\n", err.Error())
			return
		}

		event := <-e
		fmt.Println("监听到节点children变化")
		fmt.Println("children: ", b)
		fmt.Println("sate: ", sate)
		fmt.Println("path: ", event.Path)
		fmt.Println("type: ", event.Type.String())
		fmt.Println("state: ", event.State.String())
		fmt.Println("---------------------------")
	}
}
