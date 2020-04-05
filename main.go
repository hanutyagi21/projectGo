package main

import (
	//"fmt"
	"net/http"

	"github.com/hanutyagi21/projectGo/controllers"
	//"github.com/hanutyagi21/projectGo/models"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000",nil )
}