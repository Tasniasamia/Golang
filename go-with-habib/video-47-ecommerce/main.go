package main

import (

	"mains/cmd"
	"mains/config"
	
)






func main() {
	
	cmd.Serve();



}

func init() {
	config.LoadConfig()


}