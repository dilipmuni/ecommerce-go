package main

import (
	"GOProject/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	l := httprouter.New()
	m := httprouter.New()
	uc := controller.NewUserController(getSession())

	r.GET("/users", uc.GetAllUsers)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	l.GET("/products", uc.GetAllProducts)
	l.GET("/product/:id", uc.GetProduct)
	l.POST("/product", uc.CreateProduct)
	// r.DELETE("/product/:id", uc.DeleteProduct)

	m.GET("/carts", uc.GetAllCarts)
	m.POST("/carts", uc.CreateCart)
	//r.DELETE("/user/:id/cart", uc.DeleteCart)

	m.GET("/user/:id/cart", uc.GetCartUser)
	m.PUT("/user/:id/cart", uc.AddToCart)
	m.DELETE("/user/:id/cart", uc.DeleteItemInCart)

	r.GET("/user/:id/cart2", uc.AddToCart2)

	// r.GET("/user/:id/payment", uc.GetPayment)
	// r.POST("/user/:id/payment", uc.PostPayment)

	// r.POST("/user/:id/order", uc.PlaceOrder)

	go http.ListenAndServe("172.18.0.3:9012", r)
	go http.ListenAndServe("172.18.0.3:9013", l)
	http.ListenAndServe("172.18.0.3:9014", m)
}
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://172.18.0.2")

	if err != nil {
		panic(err)
	}
	return s
}
