package object

import "fmt"

//接口
func (stu student) run() {
	fmt.Printf("%s 跑步\n", stu.name)
}
func runAction(i peopler) {
	i.run()
}

type peopler interface {
	run()
}
type teacherleaderer interface {
	peopler
	test()
}
type teacherLead struct {
	student
	level int
}

func (tea teacherLead) run() {
	fmt.Printf("%d 开车\n", tea.level)
}
func (tl teacherLead) test() {
	fmt.Printf("%d 测试\n", tl.level)
}

type nilInterfacer interface {
}

func Interface() {
	//stu := student{1001, "jack", 18, "beijing"}
	////stu.run()
	//teaLeader := teacherLead{student{1001, "jack", 18, "beijing"}, 10}
	////teaLeader.run()
	//
	//runAction(stu)
	//runAction(teaLeader)

	//var p peopler
	//var stuer teacherleaderer = teacherLead{1}
	////接口继承
	//stuer.run()
	//stuer.test()
	//
	////父接口接收 子接口
	//p = stuer
	//p.run()

	//类型断言
	//var test1 nilInterfacer = 1
	//var test2 nilInterfacer = "test"
	//fmt.Println(test1, test2)
	//
	//var arr [2]nilInterfacer = [2]nilInterfacer{test2, test1}
	//for _, v := range arr {
	//	if value, ok := v.(string); ok == true {
	//		fmt.Println("string", value)
	//	} else if value, ok := v.(int); ok == true {
	//		fmt.Println("int", value)
	//	} else {
	//		fmt.Println("other")
	//	}
	//}

}
