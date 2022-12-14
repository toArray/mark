package syncTest

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestA(t *testing.T) {
	//我们创建⼀个Pool，并实现New()函数
	sp := sync.Pool{}
	item := sp.Get()
	//打印可以看到，我们通过New返回的⼤⼩为16的[]int
	fmt.Println("item : ", item)
	//然后我们对item进⾏操作
	//New()返回的是interface{}，我们需要通过类型断⾔来转换
	for i := 0; i < len(item.([]int)); i++ {
		item.([]int)[i] = i
	}
	fmt.Println("item : ", item)
	//使⽤完后，我们把item放回池中，让对象可以重⽤
	sp.Put(item)
	runtime.GC()
	//再次从池中获取对象
	item2 := sp.Get()
	//注意这⾥获取的对象就是上⾯我们放回池中的对象
	fmt.Println("item2 : ", item2)
	//我们再次获取对象
	item3 := sp.Get()
	//因为池中的对象已经没有了，所以⼜重新通过New()创建⼀个新对象，放⼊池中，然后返回
	//所以item3是⼤⼩为16的空[]int
	fmt.Println("item3 : ", item3)
	//测试sync.Pool保存socket长连接池
}
