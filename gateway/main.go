package main

import (
	"fmt"

	"github.com/Sudarshan-PR/playground/gateway/controllers"
	"github.com/Sudarshan-PR/playground/gateway/setup"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/compile", controllers.CompileHandler)

	return r
}
func main() {
	if err := setup.Setup(); err != nil {
		fmt.Println("Error during setup")
		fmt.Println(err)
		return
	}
	
	r := SetupRouter()

	r.Run(":8000")
}
