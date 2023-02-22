package cas

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestCas(t *testing.T) {
	a := uint32(1)
	for i := 0; i < 10; i++ {
		go func() {
			v := atomic.LoadUint32(&a)
			res := atomic.CompareAndSwapUint32(&a, v, a+1)
			fmt.Println(i, res, a)
		}()

	}
	time.Sleep(time.Second)
}
