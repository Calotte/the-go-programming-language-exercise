package main

import (
	"code.byted.org/gopkg/pkg/log"
	"fmt"
	"net/http"
	"os"
)

const port  =  "8000"

func sayHi(w http.ResponseWriter,r *http.Request){

	hostName,_:=os.Hostname()
	fmt.Fprintf(w,"Hi,host name: %s",hostName)
}

func main() {
	http.HandleFunc("/hi",sayHi)

	fmt.Println("running on port: "+port)

	log.Fatal(http.ListenAndServe(":"+port,nil))

}
