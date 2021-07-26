package main

import (
	"GOProject/controller"
	//"GOProject/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"fmt"
	"bytes"
)

func Test_CreateProduct(t *testing.T) {
	uc := controller.NewUserController(getSession())
	
	var jsonStr = []byte(`{"pname":"test product", "pprice": 11.22, "pqty":10}`)

	create_prod_req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonStr))
	create_prod_req.Header.Set("Content-Type", "application/json")
	
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router:=mux.NewRouter()
	router.HandleFunc("/product", uc.CreateProduct)
	router.ServeHTTP(rr, create_prod_req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	/*
	var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    if m["name"] != "test product" {
        t.Errorf("Expected product name to be 'test product'. Got '%v'", m["name"])
    }
	*/
}

func Test_GetAllProduct(t *testing.T) {
	
	uc := controller.NewUserController(getSession())
	
	get_allprod_req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	//get_prod_req,err := http.NewRequest("GET","/product",)

	rr := httptest.NewRecorder()
	router:=mux.NewRouter()
	router.HandleFunc("/products", uc.GetAllProducts)
	router.ServeHTTP(rr, get_allprod_req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	
	var m []map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &m)
	fmt.Println(m[0])
}

func Test_GetProduct(t *testing.T) {
	
	uc := controller.NewUserController(getSession())
	
	get_prod_req, err := http.NewRequest("GET", "/product", nil)
	get_prod_req.Header.Set("Content-Type", "application/json")
	get_prod_req.Header.Set("id", "60f521b8a1875b06b86c620d")
	if err != nil {
		t.Fatal(err)
	}

	//get_prod_req,err := http.NewRequest("GET","/product",)

	rr := httptest.NewRecorder()
	router:=mux.NewRouter()
	router.HandleFunc("/product", uc.GetProduct)
	router.ServeHTTP(rr, get_prod_req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	
	var m map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &m)
	fmt.Println(m)
}

/*
func Test_DeleteProduct(t *testing.T) {
	
	uc := controller.NewUserController(getSession())
	
	var foundProduct model.Product
	if err := uc.session.DB("go-web-dev-db").C("products").Find(bson.M{"pname": "test product"}).One(&foundProduct); err != nil {
		w.WriteHeader(404)
		return
	}

	delete_prod_req, err := http.NewRequest("DELETE", "/product", nil)
	delete_prod_req.Header.Set("Content-Type", "application/json")

	delete_prod_req.Header.Set("id", "60f521b8a1875b06b86c620d")
	if err != nil {
		t.Fatal(err)
	}

	//get_prod_req,err := http.NewRequest("GET","/product",)

	rr := httptest.NewRecorder()
	router:=mux.NewRouter()
	router.HandleFunc("/product", uc.GetProduct)
	router.ServeHTTP(rr, get_prod_req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	
	var m map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &m)
	fmt.Println(m)
}*/