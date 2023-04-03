package main

import (
	"fmt"
	"adzuna/source"
)


func main() {
	 
	fmt.Println("SERVER STARTED->")

	var send_data source.RequestData
	send_data.What = "Nurse"
	//send_data.Where = "New York"
	result := source.Adzuna(send_data)

	for key, value := range result.Results {
		fmt.Println(key, ":", value.Title)
	}
}

