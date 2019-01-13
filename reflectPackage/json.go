package reflectPackage

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Id   uint64 `json:"id"`//添加Tag
	Size uint64 `json:"size"`
	Data string `json:"data"`
}

func (m *Message)ZsgInit(id,size uint64,data string)  {
	m.Id=id
	m.Size=size
	m.Data=data
}
func (m Message)Print()  {
	fmt.Println("--Print--",m.Id,"----",m.Size,"----",m.Data)
}
func (m Message)GetMsgId() {
	fmt.Println("--GetMsgId----",m.Id)
}
func (m Message)PrintInput(input1,input2 int) int {
	fmt.Println("PrintInput---input1=",input1,"---input2=",input2)
	return input1+input2
}


func JsonEncodeDecode() {
	fmt.Println("-----------------")
	m1 := Message{3, 1024, "json"}
	var buf []byte
	var err error
	if buf, err = json.Marshal(m1); err != nil {
		fmt.Println("json marshal error:", err)
	}
	fmt.Println(string(buf))

	var m2 Message
	if err = json.Unmarshal(buf, &m2); err != nil {
		fmt.Println("json unmarshal error:", err)
	}
	fmt.Println(m2)
}