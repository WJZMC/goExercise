package factory

type storeInterface interface {
	sale()    //销售
	service() //售后服务
}
type car struct {
	name  string
	price float64
}

func sale(i storeInterface) {
	i.sale()
}
func service(i storeInterface) {
	i.service()
}
