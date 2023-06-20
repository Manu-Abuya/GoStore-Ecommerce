package main

import (
	"os"

	"github.com/Manu-Abuya/GoStore-Ecommerce/controllers"
	"github.com/Manu-Abuya/GoStore-Ecommerce/database"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

}
