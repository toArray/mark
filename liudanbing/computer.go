package main

import "fmt"

//=====抽象层=====
type CPU interface {
	CPURun()
}

type Mem interface {
	MemRun()
}

type Show interface {
	ShowRun()
}

type Computer struct {
	cpu  CPU
	mem  Mem
	show Show
}

func NewComputer(cpu CPU, mem Mem, show Show) *Computer {
	return &Computer{
		cpu:  cpu,
		mem:  mem,
		show: show,
	}
}

func (c *Computer) Work() {
	c.cpu.CPURun()
	c.mem.MemRun()
	c.show.ShowRun()
}

//实现层
type InterCpu struct {
	CPU
}

func (i *InterCpu) CPURun() {
	fmt.Println("inter cpu runing")
}

type InterMem struct{}

func (i *InterMem) MemRun() {
	fmt.Println("inter mem runing")
}

type InterShow struct{}

func (i *InterShow) ShowRun() {
	fmt.Println("inter show runing")
}

//业务逻辑层
func main() {
	computer := NewComputer(&InterCpu{}, &InterMem{}, &InterShow{})
	computer.Work()
}
