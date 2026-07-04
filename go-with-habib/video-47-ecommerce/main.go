package main

import (
	"fmt"
	"mains/cmd"
	"mains/config"
	"mains/util"
)






func main() {
	
	cmd.Serve();



}

func init() {
	config.LoadConfig()

	fmt.Println(util.CreateJwtToken("your_secret_key", util.Payload{Sub: "1", Name: "tia", Email: "tia@example.com"}))

}