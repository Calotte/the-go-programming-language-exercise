package search

import (
	"fmt"
	"net/http"
	"the-go-programming-language-exercise/ch12/params"
)

func Search(resp http.ResponseWriter, req *http.Request){
	var data struct {
		Labels []string `http:"l"`
		MaxResults int `http:"max"`
		Exact bool `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}
	fmt.Fprintf(resp, "Search: %+v\n", data)
}
