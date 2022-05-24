package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Sex  byte
	Age  int
}

type Student struct {
	Person
	Id      int
	Score   float64
	Friends []string `json:"friends"`
}

func main() {
	//TODO 看视频 json 属性大小写的问题
	s1 := Student{Person{"yorick", 'm', 24}, 10, 98.23, []string{"tom", "jerry", "smith"}}
	res, err := json.Marshal(s1)
	//res,err := json.MarshalIndent(s1,"","	")
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(string(res))

	var s2 Student
	err = json.Unmarshal(res, &s2)
	if err != nil {
		fmt.Println("json err: ", err)
		return
	}
	fmt.Printf("%+v\n", s2)
}
