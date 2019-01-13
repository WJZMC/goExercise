package encoding

import (
	"encoding/json"
	"fmt"
)

type Job struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

func TestJson() {
	job := Job{1, "golang", 25000}
	jsonStruct, err := json.Marshal(&job)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStructStr := string(jsonStruct)
	fmt.Printf("Job Struct Encode json:%s\n", jsonStructStr)

	var jobTemp Job
	err = json.Unmarshal([]byte(jsonStructStr), &jobTemp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Job Struct Dncode json:%v\n", jobTemp)

	part := []Job{Job{1, "golang", 25000}, Job{2, "ios", 25000}}
	jsonPart, err := json.Marshal(&part)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonPartStr := string(jsonPart)
	fmt.Printf("part Encode json:%s\n", jsonPartStr)

	var part1 []Job
	err = json.Unmarshal([]byte(jsonPartStr), &part1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("part Dncode json:%v\n", part1)

	mapTest := map[string][]Job{"section": []Job{job, job}}
	jsonMap, err := json.Marshal(&mapTest)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonMapStr := string(jsonMap)
	fmt.Printf("map Encode json:%s\n", jsonMapStr)

	var mapTestttt map[string]Job
	err = json.Unmarshal([]byte(jsonMapStr), &mapTestttt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("map Dncode json:%v\n", mapTestttt)

}
