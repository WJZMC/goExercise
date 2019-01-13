package reflectPackage

import (
	"reflect"
	"fmt"
	"math"
)

func ReflectDemo()  {


	//num:=10
	//reflectTest(num)//反射获取基本数据类型


	//msg:=Message{1,2,"3"}
	//reflectTest(msg)//反射获取结构体数据类型

	//changeValue()//反射修改变量的值

	//msg:=Message{}
	//msg.ZsgInit(1,200,"我是data")
	//msg.Print()
	//msg.GetMsgId()
	////structTest(msg)//反射解析结构体
	//structTestPoint(&msg)//反射解析结构体指针

	//pubFuncTest(add,10,20)
	//pubFuncTest(square,100.0)

	newStruct()
}

func reflectTest(obj interface{})  {
	rtyp:=reflect.TypeOf(obj)
	fmt.Printf("-1---obj =%v  type=%T\n",rtyp,rtyp)
	rvalue:=reflect.ValueOf(obj)
	fmt.Printf("-2---obj=%v  type=%T kind=%v\n",rvalue,rvalue,rvalue.Kind())

	tmp:=rvalue.Interface()
	fmt.Printf("-3---obj=%v  type=%T \n",tmp,tmp)

	switch tmp.(type) {
	case int:
		fmt.Println("类型为int value=",rvalue.Int())
	case Message:
		data,_ :=tmp.(Message)
		fmt.Println("类型为Message,id=",data.Id)
	default:

	}

}

func changeValue()  {

	//num:=10
	//rvalue:=reflect.ValueOf(num)
	//rvalue.SetInt(20)
	//fmt.Println(num)
	num:=10
	rvalue:=reflect.ValueOf(&num)
	rvalue.Elem().SetInt(20)//SetXX需要接收一个 地址指向的值
	fmt.Println(num)

}

func structTest(tmp interface{})  {

	rvalue:=reflect.ValueOf(tmp)

	rkind:=rvalue.Kind()
	if rkind != reflect.Struct{
		return
	}
	fmt.Println("kind=",rkind)

	fieldNum:=rvalue.NumField()
	fmt.Println("--字段个数---",fieldNum)
	for i:=0;i<fieldNum;i++{
		fmt.Printf("--第 %d 个字段 --%s, tag=%s--value=%v----\n",
			i,rvalue.Type().Field(i).Name,rvalue.Type().Field(i).Tag.Get("json"),rvalue.Field(i))
	}
	methodNum:=rvalue.NumMethod()
	fmt.Println("--方法个数---",methodNum)

	//方法顺序是按方法方法名字(ASCII)排序
	//根据传入的变量类型，返回对应的方法
	//如果是指针方法，传入对象则无法获取
	for i:=0;i<methodNum;i++{
		fmt.Printf("--第 %d 个方法 --%s\n", i,rvalue.Type().Method(i).Name)
	}

	rvalue.Method(0).Call(nil)

	rvalue.Method(1).Call(nil)

	//call() 参数时  []Value
	var param []reflect.Value
	param= append(param, reflect.ValueOf(10))
	param= append(param, reflect.ValueOf(20))
	res:=rvalue.Method(2).Call(param)
	fmt.Printf("---res=%v,type=%T,value=%v\n",res,res,res[0].Interface())


}

func structTestPoint(tmp interface{})  {


	rvalue:=reflect.ValueOf(tmp)

	rkind:=rvalue.Kind()
	if rkind != reflect.Struct&&rkind!=reflect.Ptr{
		return
	}
	fmt.Println("kind=",rkind)

	fieldNum:=rvalue.Elem().NumField()
	fmt.Println("--字段个数---",fieldNum)
	for i:=0;i<fieldNum;i++{
		fmt.Printf("--第 %d 个字段 --%s, tag=%s--value=%v----\n",
			i,rvalue.Elem().Type().Field(i).Name,rvalue.Elem().Type().Field(i).Tag.Get("json"),rvalue.Elem().Field(i))
	}
	methodNum:=rvalue.NumMethod()
	fmt.Println("--方法个数---",methodNum)

	//方法顺序是按方法方法名字(ASCII)排序
	//根据传入的变量类型，返回对应的方法
	//如果是指针方法，传入对象则无法获取
	for i:=0;i<methodNum;i++{
		fmt.Printf("--第 %d 个方法 --%s\n", i,rvalue.Type().Method(i).Name)
	}

	rvalue.Method(0).Call(nil)
	rvalue.Method(1).Call(nil)


	//call() 参数是  []Value  返回值也是[]Value
	var param []reflect.Value
	param= append(param, reflect.ValueOf(10))
	param= append(param, reflect.ValueOf(20))
	res:=rvalue.Method(2).Call(param)
	fmt.Printf("---res=%v,type=%T,value=%v\n",res,res,res[0].Interface())


	//结构体初始化
	var paramInit []reflect.Value
	var idTmp uint64 = 9
	var sizeTmp uint64 =100
	var dataTmp string = "我是结构体指针的data"
	paramInit= append(paramInit, reflect.ValueOf(idTmp))
	paramInit= append(paramInit, reflect.ValueOf(sizeTmp))
	paramInit= append(paramInit, reflect.ValueOf(dataTmp))

	rvalue.Method(3).Call(paramInit)

	rvalue.Method(1).Call(nil)

	//修改字段值
	rvalue.Elem().Field(0).SetUint(999)
	rvalue.Method(1).Call(nil)

}
func add(x,y int) (int ,int) {
	return x,y
}
func square(x float64) float64 {
	return math.Sqrt(x)
}
//使用反射解析任意类型方法和参数替代（switch case ）
func pubFuncTest(funtion interface{},args...interface{})  {
	rvalue:=reflect.ValueOf(funtion)
	fmt.Printf("-rValue--%v,%T\n",rvalue,rvalue)
	rtype:=reflect.TypeOf(funtion)
	fmt.Printf("-rType--%v,%T\n",rtype,rtype)

	param:=make([]reflect.Value,0)
	for i:=0;i< len(args);i++ {
		param= append(param, reflect.ValueOf(args[i]))
	}
	res:=rvalue.Call(param)

	for i,v:=range res{
		fmt.Printf("---func RetrunValue==[%d]=%v\n",i,v)
	}

}
//使用反射创建结构体
func newStruct(){
	var (
		model *Message
		elem reflect.Value
		rType reflect.Type
	)

	rType=reflect.TypeOf(model)

	elem=reflect.New(rType.Elem())

	model=elem.Interface().(*Message)

	elem.Elem().Field(0).SetUint(99)
	elem.Elem().Field(1).SetUint(1024)
	elem.Elem().Field(2).SetString("reflect test")

	fmt.Println(*model)


}