package main

import (
	"e-commerce-cart-system/controllers"
	"e-commerce-cart-system/database"
	"e-commerce-cart-system/middleware"
	"e-commerce-cart-system/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication( database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":"+ port))

}