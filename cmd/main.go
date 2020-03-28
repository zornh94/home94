package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tndevr/pkgg"
)

func main(){
	r := gin.Default()
	r.GET("/",pkgg.Greeting)
	fmt.Printf("Server listen Default %s",pkgg.ADDR)
	r.Run(pkgg.ADDR)
}
