package luoqiangMark

import (
	"fmt"
	"github.com/bits-and-blooms/bitset"
	"testing"
)

const DEFAULT_SIZE = 2 << 24              //默认布隆过滤器大小
var seeds = []uint{7, 11, 13, 31, 37, 61} //种子

type BloomFilterInterface interface {
	Add(value string)
	Contains(value string) bool
}

//SimpleHash 哈希用
type SimpleHash struct {
	Cap  uint
	Seed uint
}

//BloomFilter 布隆过滤器模型
type BloomFilter struct {
	Set  *bitset.BitSet
	Func [6]SimpleHash
}

//包测试用
func TestBloomFilter(t *testing.T) {
	var b bitset.BitSet            // 定义一个BitSet对象
	b.Set(1).Set(2).Set(3).Set(10) //添加4个元素
	if b.Test(2) {
		fmt.Println("2已经存在")
	}
	fmt.Println("总数：", b.Count())

	b.Clear(2)
	if !b.Test(2) {
		fmt.Println("2不存在")
	}
	fmt.Println("总数：", b.Count())

	i, e := b.NextSet(1)
	fmt.Println(i, e)

	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("The following bit is set:", i)
	}

	//查看2bitset交集
	b2 := bitset.New(100).Set(10)
	if b.Intersection(b2).Count() == 1 {
		fmt.Println("Intersection works.")
	} else {
		fmt.Println("Intersection doesn't work???")
	}
}

//实现一个简单的布隆过滤器
func TestBloomFilterUse(t *testing.T) {
	bloomFilter := NewBloomFilter()
	bloomFilter.Add("1")
	bloomFilter.Add("222")
	bloomFilter.Add("1111")
	res := bloomFilter.Contains("222")
	fmt.Println(res)
}

/*
NewBloomFilter
@Desc new一个布隆过滤器
*/
func NewBloomFilter() *BloomFilter {
	b := new(BloomFilter)
	for i := 0; i < len(b.Func); i++ {
		b.Func[i] = SimpleHash{DEFAULT_SIZE, seeds[i]}
	}
	b.Set = bitset.New(DEFAULT_SIZE)
	return b
}

/*
Add
@Desc 	添加值
@Param	value string
*/
func (b *BloomFilter) Add(value string) {
	for _, f := range b.Func {
		b.Set.Set(f.hash(value))
	}
}

/*
Contains
@Desc 	比对值
@Param	value string
*/
func (b *BloomFilter) Contains(value string) bool {
	if value == "" {
		return false
	}

	ret := true
	for _, f := range b.Func {
		ret = ret && b.Set.Test(f.hash(value))
	}

	return ret
}

/*
Contains
@Desc 	hash
@Param	value string
*/
func (s *SimpleHash) hash(value string) uint {
	var result uint = 0
	for i := 0; i < len(value); i++ {
		result = result*s.Seed + uint(value[i])
	}
	return (s.Cap - 1) & result
}
