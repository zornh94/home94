package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tndevr/packages"
)

func main(){
	r := gin.Default()
	r.GET("/",packages.Greeting)
	fmt.Printf("Server listen Default %s",packages.ADDR)
	r.Run(packages.ADDR)
}
