package object

import "fmt"

type people struct {
	name string
	age  int
}
type student struct {
	people //匿名字段
	stid   int
	class  int
	name   string //重名字段 遵循就近原则
}

type teacher struct {
	*people
	tid   int
	class int
}

type teacherLeader struct {
	teacher
	lid    int
	leavel int
}

func (peo people) printfMsg() {
	fmt.Printf("name=%s,age=%d\n", peo.name, peo.age)
}
func (peo *people) test() {
	fmt.Printf("name=%s,age=%d\n", peo.name, peo.age)
}
func (tea teacherLeader) printfMsg() {
	fmt.Printf("name=%s,age=%d,tid=%d,class=%d,lid=%d,leavel=%d\n", tea.teacher.people.name, tea.age, tea.tid, tea.class, tea.lid, tea.leavel)
}
func (tea teacherLeader) editor(name string, age, tid, class, lid, leavel int) {
	tea.name = name
	tea.age = age
	tea.tid = tid
	tea.class = class
	tea.lid = lid
	tea.leavel = leavel
}

func (tea *teacherLeader) PointEditor(name string, age, tid, class, lid, leavel int) {
	tea.name = name
	tea.age = age
	tea.tid = tid
	tea.class = class
	tea.lid = lid
	tea.leavel = leavel
}
func Object() {
	////初始化
	//var stu student = student{people{"jack", 18}, 101, 2, "test"}
	//fmt.Println(stu)
	//
	////匿名字段遵循就近原则，如果子类中有，父类中值未发生变化
	//stu.name = "her"
	//fmt.Println(stu)
	////stu.name = "test"
	//fmt.Println(stu, stu.people.name)
	//
	////匿名指针,遵循指针的规则，需要先定义后使用
	//var tea teacher
	//tea.people = new(people)
	//tea.people.name = "jack"
	////tea := teacher{&people{"jack", 18}, 111, 102}
	//fmt.Println(*tea.people, tea)

	////多重继承
	//var leader teacherLeader
	//leader.teacher.people = new(people)
	//leader.teacher.people.name = "jack"
	//fmt.Println(*leader.teacher.people)

	//结构体方法
	//leader := teacherLeader{teacher{&people{"jack", 18}, 1111, 101}, 100, 1}
	//leader.printfMsg()

	////结构体指针方法
	//(&leader).PointEditor("test", 20, 1111, 101, 100, 1)
	//(&leader).printfMsg()

	//leader.PointEditor("test", 20, 1111, 101, 100, 1)
	//leader.editor("test", 20, 111, 11, 10, 2)
	//leader.printfMsg()

	//stu := Student{"jack", "男", 26, 99, 88, 80, 0, 0}
	//stu.sayHello()
	//stu.complateScore()
	//fmt.Printf("我叫%s,这次考试总成绩为%2f分,平均成绩为%2f分\n", stu.name, stu.sum, stu.avg)

	//方法继承和重写，遵循就近原则，如果子类重写了该方法，则优先执行子类方法
	//stu := student{people{"jack", 18}, 110, 100, "test"}
	//stu.printfMsg()
	//leader := teacherLeader{teacher{&people{"jack", 18}, 1111, 101}, 100, 1}
	//leader.printfMsg()

	//rep := reporter{Human{"记者", 48}, "偷拍"}
	//dev := develop{Human{"jack", 26}, 5}
	//rep.say()
	//dev.say()
	//
	////方法存储在代码区，子类和父类是不同的代码区地址
	//fmt.Printf("%p,%p", rep.say, rep.Human.say)

	//方法值和方法表达式
	//隐式，通过方法值隐藏接受者
	//var stu student = student{people{"jack", 18}, 110, 100, "test"}
	//f := stu.printfMsg
	//f()
	//
	////方法表达式，显式把接受者传递过去
	//f1 := (student).printfMsg
	//f1(stu)
	//
	//f2 := (*student).test
	//f2(&stu)

}

//定义一个学生类,有六个属性,分别为姓名、性别、年龄、语文、数学、英语成绩。
//有2个方法：
//一个打招呼的方法：介绍自己叫XX，今年几岁了。是男同学还是女同学。
//两个计算自己总分数和平均分的方法。{显示:我叫XX,这次考试总成绩为X分,平均成绩为X分}
type Student struct {
	name    string
	gender  string
	age     int
	China   float64
	math    float64
	English float64
	sum     float64
	avg     float64
}

func (stu Student) sayHello() {
	fmt.Printf("我叫%s，今年%d岁了，我是%s同学\n", stu.name, stu.age, stu.gender)
}
func (stu *Student) complateScore() {
	stu.sum = stu.China + stu.math + stu.English
	stu.avg = stu.sum / 3.0
}

type Human struct {
	name string
	age  int
}
type reporter struct {
	Human
	hobby string
}
type develop struct {
	Human
	worktime int
}

func (h *Human) say() {
	fmt.Printf("%s,%d\n", h.name, h.age)
}

func (r *reporter) say() {
	fmt.Printf("%s,%d,职业：%s\n", r.name, r.age, r.hobby)
}

func (d *develop) say() {
	fmt.Printf("%s,%d,工作年限：%d\n", d.name, d.age, d.worktime)
}
