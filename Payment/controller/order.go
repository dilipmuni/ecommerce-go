//-------------------------------------//
//Order products

package controller

import (
	"encoding/json"
	"fmt"
	"time"

	//"C:/Users/Dell/Desktop/GOProject/model"
	"GOProject/model"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

func (uc UserController) PlaceOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	cartproducts := model.CartProduct{}
	json.NewDecoder(r.Body).Decode(&cartproducts)

	u := model.User{}

	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	fmt.Println(u.Name)

	cart := model.Cart{}
	if err := uc.session.DB("go-web-dev-db").C("carts").Find(bson.M{"uname": u.Name}).One(&cart); err != nil {
		fmt.Println("Error we arer in")
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	order := model.Order{}
	order.OrderProducts = cart.CartProducts
	order.TotalPrice = cart.TotalPrice
	order.DeliveryStatus = "To be shipped!"
	order.OrderDate = time.Now()
	order.Username = u.Name

	order.Id = bson.NewObjectId()

	//json.NewDecoder(r.Body).Decode(&cart)

	uc.session.DB("go-web-dev-db").C("orders").Insert(order)

	uc.UpdateProductsAfterOrder(order.OrderProducts)

	uj, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) UpdateProductsAfterOrder(slice []model.CartProduct) {
	for i, _ := range slice {
		product := model.Product{}
		if err := uc.session.DB("go-web-dev-db").C("products").Find(bson.M{"pname": slice[i].ProductName}).One(&product); err != nil {
			fmt.Println("Error we arer in")
			fmt.Println(err)
			//w.WriteHeader(404)
			return
		}
		product.ProductQty -= slice[i].ProductQty

		if err := uc.session.DB("go-web-dev-db").C("products").Update(bson.M{"_id": product.Id}, product); err != nil {
			fmt.Println("Error we arer in")
			fmt.Println(err)
		}
	}
}
