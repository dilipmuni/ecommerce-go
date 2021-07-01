package main

import (
	"GOProject/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controller.NewUserController(getSession())

	// r.GET("/users", uc.GetAllUsers)
	// r.GET("/user/:id", uc.GetUser)
	// r.POST("/user", uc.CreateUser)
	// r.DELETE("/user/:id", uc.DeleteUser)

	r.GET("/products", uc.GetAllProducts)
	r.GET("/product/:id", uc.GetProduct)
	r.POST("/product", uc.CreateProduct)
	r.DELETE("/product/:id", uc.DeleteProduct)

	r.PUT("/user/:id/cart", uc.RedirectToAddCart)

	// r.GET("/carts", uc.GetAllCarts)
	// r.POST("/carts", uc.CreateCart)
	// //r.DELETE("/user/:id/cart", uc.DeleteCart)

	// r.GET("/user/:id/cart", uc.GetCartUser)
	// r.PUT("/user/:id/cart", uc.AddToCart)
	// r.DELETE("/user/:id/cart", uc.DeleteItemInCart)

	// r.GET("/user/:id/payment", uc.GetPayment)
	// r.POST("/user/:id/payment", uc.PostPayment)

	// r.POST("/user/:id/order", uc.PlaceOrder)

	http.ListenAndServe("localhost:8082", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
