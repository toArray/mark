package luoqiangMark

import (
	"fmt"
	"testing"
	"unsafe"
)

type M struct {
	A int8   //1字节
	B int64  //8字节
	C bool   //1字节
	D int16  //2字节
	E string //16字节
	F rune   //4字节
}

/*
对齐系数8
当前偏移量0
1-8 得偏移量1,8是1的整数倍  A

当前偏移量1
8-8 得偏移量8,8是8的整数倍  B
Axxx xxxx BBBB BBBB

当前偏移量16
1-8 得偏移量1,8是1的整数倍  C
Axxx xxxx BBBB BBBB C

当前偏移量17
2-8 得偏移量2 到24的位置 D
Axxx xxxx BBBB BBBB CxDD

当前偏移量40
1-8 得偏移量1,40是1的整数倍  E
Axxx xxxx BBBB BBBB Cxxx xxxx DDDD DDDD DDDD DDDD E

当前偏移量41
4-8 得偏移量4  F
Axxx xxxx BBBB BBBB Cxxx xxxx DDDD DDDD DDDD DDDD Exxx FFFF


//总共32位


*/

/*
TestMemory
内存对齐

*/
func TestMemory(t *testing.T) {
	var a M

	fmt.Println(unsafe.Sizeof(a))
}
