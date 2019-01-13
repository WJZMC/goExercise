package fileOperation

import (
	"fmt"
	"time"
)

func ErrorExercise() {
	//error := errors.New("你遇到了一个错误")
	//fmt.Println(error.Error())
	//
	//error1 := fmt.Errorf("fmt提供的错误")
	//fmt.Println(error1.Error())
	//

	test()

	for {
		fmt.Println("1111")
		time.Sleep(time.Second)
	}

}

func test() {
	defer func() {
		fmt.Println(recover())
	}()
	var p *int
	fmt.Println(*p)
}
