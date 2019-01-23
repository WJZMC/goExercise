package encoding

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

//package 名字叫做 lm
//定义了一个消息 helloworld
//该消息有三个成员，类型为 int32 的 id，
//另一个为类型为 string 的成员 str。
//opt 是一个可选的成员，即消息中可以不包含该成员

func protobuftest() {
	job := &JobProto{
		Id:   1,
		Name: "jack",
		Age:  19,
		Emails: []string{
			"test@sin.com",
			"tes@sina.com",
		},
		Phones: []*PhoneNum{
			&PhoneNum{
				Number: "123455",
				Type:   PhoneType_Mobile,
			},
			&PhoneNum{
				Number: "123455",
				Type:   PhoneType_Mobile,
			},
		},
	}
	fmt.Println(*job)

	str, err := proto.Marshal(job)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(str))
	}
	var jobb JobProto
	err = proto.Unmarshal(str, &jobb)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(jobb)
	}

}
