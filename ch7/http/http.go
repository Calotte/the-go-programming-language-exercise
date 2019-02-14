package myhttp

import (
	"fmt"
	"net/http"
)

type dollors float32
func (d dollors)String()string{return fmt.Sprintf("$%.2f",d)}
type Database map[string]dollors

func (db Database)ServeHTTP(w http.ResponseWriter, req *http.Request){
	switch req.URL.Path {
	case "/list":
		for item,price := range db{
			fmt.Fprintf(w,"%s:%s\n",item,price)
		}
	case "/price":
		item:=req.URL.Query().Get("item")
		price,ok:=db[item]
		if !ok{
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w,"%s\n",price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func (db Database)List(w http.ResponseWriter,req *http.Request){
	for item,price := range db{
		fmt.Fprintf(w,"%s:%s\n",item,price)
	}
}

func (db Database)Price(w http.ResponseWriter,req *http.Request){
	item := req.URL.Query().Get("item")
	price,ok := db[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"no sum item:%q\n",item)
		return
	}
	fmt.Fprintf(w,"%s\n",price)
}


