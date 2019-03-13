package main

import (
	"fmt"
)

type BulbInterface interface {
	turnOn()
	turnOff()
	isOnFun() bool
}

const (
	SMALL = iota
	MEDIUM
	LARGE
)

type Bulb struct {
	isOn bool
	size int
}

func (b *Bulb) getBulfSize() int {
	return b.size
}

func (b *Bulb) setBulbSize(s int) {
	b.size = s
}

func (b *Bulb) turnOn() {
	b.isOn = true
}

func (b *Bulb) turnOff() {
	b.isOn = false
}

func (b *Bulb) isOnFun() bool {
	return b.isOn
}

type AdvanceBulb struct {
	Bulb
	intensity int
}

func (b *AdvanceBulb) setIntersity(i int) {
	b.intensity = i
}

func (b *AdvanceBulb) getIntersity() int {
	return b.intensity
}

func getBulbStatus(bi BulbInterface) {
	fmt.Println("bulb is on: ", bi.isOnFun())
}

func main() {
	b := Bulb{false, SMALL}
	fmt.Println("bulb is on return : ", b.isOnFun())
	b.turnOn()
	fmt.Println("bulb is on return : ", b.isOnFun())
	ab := AdvanceBulb{Bulb{false, MEDIUM}, 10}
	fmt.Println("bulb is on return : ", ab.isOnFun(), " and its intensity is : ", ab.getIntersity())
	fmt.Println(ab.isOnFun())
	
	getBulbStatus(&b)
	getBulbStatus(&ab)
}