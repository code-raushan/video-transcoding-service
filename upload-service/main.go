package main

import (
	"fmt"

	"github.com/code-raushan/video-transcoding-service/upload-service/database"
)

func main(){

	fmt.Println("connecting to database")
	database.ConnectDB()
	fmt.Println("connected to database")


	fmt.Println("upload service implementation with user management")
}