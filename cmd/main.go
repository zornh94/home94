package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tndevr/packages"
//	"net/http"
)
/*
func main(){
	r := gin.Default()
	r.GET("/",func (ctx *gin.Context){
		ctx.String(http.StatusOK, "hello gin"
	})
	fmt.Println("Server listen Default 8080")
	r.Run()
}*/

func main(){
	r := gin.Default()
	r.GET("/",packages.Greeting)
	fmt.Printf("Server listen Default %s",packages.ADDR)
	r.Run(packages.ADDR)
}
