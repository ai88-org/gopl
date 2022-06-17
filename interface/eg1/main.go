package main

import "fmt"

// 创建一个接口类型，有两个方法
type Usb interface {
	Start()
	Stop()
}

// 手机结构体
type Phone struct {
}

// 手机实现了Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机启动了...")
}

func (p Phone) Stop() {
	fmt.Println("手机拔出来了...")
}

// 相机结构体
type Camera struct {
}

// 相机实现了start方法
func (c Camera) Start() {
	fmt.Println("相机开机了...")
}

// 相机实现了stop方法
func (c Camera) Stop() {
	fmt.Println("相机关机了...")
}

// 电脑结构体
type Computer struct {
}

// 电脑的工作方法
// 参数是一个USB接口
func (c Computer) Working(u Usb) {
	u.Start()
	u.Stop()
}

func main() {
	c := Computer{}
	p := Phone{}
	cm := Camera{}

	c.Working(p)
	c.Working(cm)
}
