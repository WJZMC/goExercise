package factory

import "fmt"

type Audi struct {
	car
}

func (a *Audi) sale() {
	fmt.Printf("奥迪车销售价格：%f\t", a.price)
}
func (a *Audi) service() {
	fmt.Println("奥迪车售后服务")
}
