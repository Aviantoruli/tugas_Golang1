package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var json_string = `
	[
	{
		"name":"Muhamad Edo Syahputra",
		"student_id":"GLNG018ONL001",
		"address":"Depok",
		"job":"backend"
	},
	{
		"name":"Muhamad Edo Syahputra",
		"student_id":"GLNG018ONL001",
		"address":"Depok",
		"job":"backend"
	},
	{
		"name":"Muhamad Edo Syahputra",
		"student_id":"GLNG018ONL001",
		"address":"Depok",
		"job":"backend"
	}
	]
	`
	var result student.Student

	var err = json.Unmarshal([]byte(json_string), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)

}

// func biodata()  {

// }
