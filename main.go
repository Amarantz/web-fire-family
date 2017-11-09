package main

import (
	"fmt"
	"net/http"

	//"github.com/Xero67/web-fire-family/models"
	//"github.com/Xero67/web-fire-family/routes"
	"./models"
	"./routes"

)

func main() {
	db, err := models.InitDB()
	if err != nil {
		fmt.Errorf("Database wasn't initialized!")
	}

	env := models.Env{Db: db}

	http.ListenAndServe(":8000", routes.InitRoutes(env))
}
