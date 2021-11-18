package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Vegetables struct {
	ProductId   string
	ProductName string
	Quantity    int
	Price       string
}
type Fruits struct {
	Id       string
	Name     string
	Quantity int
	Price    string
}
type Grains struct {
	ItemId   string
	ItemName string
	Quantity int
	Price    string
}

const Fruits_Url = "https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b"
const Grain_Url = "https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148"
const Vegetable_Url = "https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c"

var mg sync.WaitGroup
var pulses []Grains
var veg []Vegetables
var fruit []Fruits

func Getbyname(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Buy item")
	w.Header().Set("Content-Type", "application/json")

	value := mux.Vars(r)

	response, err1 := http.Get(Vegetable_Url)
	if err1 != nil {
		panic(err1)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err2 := json.Unmarshal(content, &veg)
	if err2 != nil {
		panic(err2)
	}

	for _, m := range veg {

		if m.ProductName == value["itemname"] {
			json.NewEncoder(w).Encode(m)
			return
		}

	}

	response1, er := http.Get(Fruits_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err3 := json.Unmarshal(content1, &fruit)
	if err3 != nil {
		panic(err3)
	}

	response2, er := http.Get(Grain_Url)
	if er != nil {
		panic(er)
	}

	for _, v := range fruit {
		if v.Name == value["itemname"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err4 := json.Unmarshal(content2, &pulses)
	if err4 != nil {
		panic(err4)
	}

	for _, p := range pulses {
		if p.ItemName == value["itemname"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!! Item Not Found")

}

func Getbynamequantity(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Buy item")
	w.Header().Set("Content-Type", "application/json")

	value := mux.Vars(r)
	k, _ := strconv.Atoi(value["quantity"])
	response, err1 := http.Get(Vegetable_Url)
	if err1 != nil {
		panic(err1)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err2 := json.Unmarshal(content, &veg)
	if err2 != nil {
		panic(err2)
	}

	for _, m := range veg {

		if m.Quantity >= k && m.ProductName == value["itemname"] {
			json.NewEncoder(w).Encode(m)
			return
		}

	}

	response1, er := http.Get(Fruits_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err3 := json.Unmarshal(content1, &fruit)
	if err3 != nil {
		panic(err3)
	}

	response2, er := http.Get(Grain_Url)
	if er != nil {
		panic(er)
	}

	for _, g := range fruit {
		if g.Quantity >= k && g.Name == value["itemname"] {
			json.NewEncoder(w).Encode(g)
			return
		}
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err4 := json.Unmarshal(content2, &pulses)
	if err4 != nil {
		panic(err4)
	}

	for _, r := range pulses {
		if r.Quantity >= k && r.ItemName == value["itemname"] {
			json.NewEncoder(w).Encode(r)
			return
		}
	}

	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!! Item Not Found")

}
func Getbynamequantityprice(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" getBuy name qty price")
	w.Header().Set("Content-Type", "application/json")

	value := mux.Vars(r)
	str := "$" + value["price"]
	k, _ := strconv.Atoi(value["quantity"])
	response, err1 := http.Get(Vegetable_Url)
	if err1 != nil {
		panic(err1)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err2 := json.Unmarshal(content, &veg)
	if err2 != nil {
		panic(err2)
	}

	for _, a := range veg {

		if a.Quantity >= k && a.ProductName == value["itemname"] && a.Price == str {
			json.NewEncoder(w).Encode(a)
			return
		}

	}

	response1, er := http.Get(Fruits_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err3 := json.Unmarshal(content1, &fruit)
	if err3 != nil {
		panic(err3)
	}

	response2, er := http.Get(Grain_Url)
	if er != nil {
		panic(er)
	}

	for _, h := range fruit {
		if h.Quantity >= k && h.Name == value["itemname"] && h.Price == str {
			json.NewEncoder(w).Encode(h)
			return
		}
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err4 := json.Unmarshal(content2, &pulses)
	if err4 != nil {
		panic(err4)
	}

	for _, b := range pulses {
		if b.Quantity >= k && b.ItemName == value["itemname"] && b.Price == str {
			json.NewEncoder(w).Encode(b)
		}
	}

	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!! Item Not Found")

}
func getveg() {
	response, err1 := http.Get(Vegetable_Url)
	if err1 != nil {
		panic(err1)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err2 := json.Unmarshal(content, &veg)
	if err2 != nil {
		panic(err2)
	}
	defer mg.Done()
	defer response.Body.Close()

}
func getfruit() {
	response1, er := http.Get(Fruits_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err3 := json.Unmarshal(content1, &fruit)
	if err3 != nil {
		panic(err3)
	}
	defer mg.Done()
	defer response1.Body.Close()
}
func getgrain() {
	response2, er := http.Get(Grain_Url)
	if er != nil {
		panic(er)
	}
	content2, _ := ioutil.ReadAll(response2.Body)

	err4 := json.Unmarshal(content2, &pulses)
	if err4 != nil {
		panic(err4)
	}
	defer mg.Done()
	defer response2.Body.Close()
}
func Getbynamefast(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Buy item")
	w.Header().Set("Content-Type", "application/json")

	value := mux.Vars(r)
	go getgrain()
	go getveg()
	go getfruit()

	mg.Add(3)
	for _, m := range veg {

		if m.ProductName == value["itemname"] {
			json.NewEncoder(w).Encode(m)
			return
		}

	}
	for _, a := range fruit {
		if a.Name == value["itemname"] {
			json.NewEncoder(w).Encode(a)
			return
		}
	}
	for _, p := range pulses {
		if p.ItemName == value["itemname"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	json.NewEncoder(w).Encode("!! Item Not Found")

}

func main() {
	c := mux.NewRouter()
	c.HandleFunc("/buy-item/{itemname}", Getbyname).Methods("GET")
	c.HandleFunc("/buy-item-qty/{itemname}/{quantity}", Getbynamequantity).Methods("GET")
	c.HandleFunc("/buy-item-qty-price/{itemname}/{quantity}/{price}", Getbynamequantityprice).Methods("GET")
	c.HandleFunc("/fast-buy-item/{itemname}", Getbynamefast).Methods("GET")
	log.Fatal(http.ListenAndServe(":9090", c))
	mg.Wait()
}
