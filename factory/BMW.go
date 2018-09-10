package factory

import "fmt"

type BMW struct {
	car
}

func (a *BMW) sale() {
	fmt.Printf("宝马车销售价格：%f\t", a.price)
}
func (a *BMW) service() {
	fmt.Println("宝马车售后服务")
}
