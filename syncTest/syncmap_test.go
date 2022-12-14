package syncTest

import (
	"fmt"
	"sync"
	"testing"
)

type TestMapData struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func TestMap(t *testing.T) {
	var m sync.Map
	fmt.Println(m)

	//存数据
	m.Store("1", 1)
	m.Store("2", 2)
	m.Store("3", 3)
	m.Store("4", 4)
	m.Store("5", 5)
	m.Store("6", 6)
	m.Store("7", 7)
	m.Store("data", TestMapData{
		ID:   1,
		Name: "test",
	})
	fmt.Println(m)

	//循环数据
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true //继续循环
		//return false//跳出循环
	})

	//读取
	data, ok := m.Load("qcrao")
	fmt.Println(ok, data)

	//删除
	m.Delete(1111)

	//不存在则写入,否则读取
	data, ok = m.LoadOrStore("1", 100)
	fmt.Println(ok, data)
}
