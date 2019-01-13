package extision

import (
	"flag"
	"fmt"
	"os"
)

func Test() {
	////默认flag
	//var name string
	//flag.StringVar(&name, "name", "name", "填写name")
	//var pwd string
	//flag.StringVar(&pwd, "pwd", "pwd", "pwd")
	//
	//flag.Parse()
	//
	//fmt.Println("name=", name)
	//fmt.Println("pwd=", pwd)

	//自定义flag
	myAccountfalg := flag.NewFlagSet("account", flag.ExitOnError)
	myJobfalg := flag.NewFlagSet("job", flag.ExitOnError)

	switch os.Args[1] {
	case "account":
		var name string
		myAccountfalg.StringVar(&name, "name", "name", "填写name")
		var pwd string
		myAccountfalg.StringVar(&pwd, "pwd", "pwd", "pwd")

		myAccountfalg.Parse(os.Args[2:])
		fmt.Println("name==", name)
		fmt.Println("pwd==", pwd)

	case "job":
		var jobname string
		myJobfalg.StringVar(&jobname, "jn", "name", "job")
		var jobSalery float64
		myJobfalg.Float64Var(&jobSalery, "s", 10.0, "sa")

		myJobfalg.Parse(os.Args[2:])
		fmt.Println("name==", jobname)
		fmt.Println("sa==", jobSalery)
	default:

	}

}
