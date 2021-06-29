package main 

import (
	"fmt"
	"net/http"
)


func rootHandler(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprint(w, "hellow world!")
}

func main()  {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
	
}