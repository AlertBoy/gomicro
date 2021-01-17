package main

import "net/http"

// run micro service --name httpTest --endpoint http://localhost:9090 go run httpServiceMain.go
// to register service
//run micro call -o raw httpTest / to run
//run micro call -o raw httpTest /

func main() {
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world 11"))
	})
	http.ListenAndServe(":9090", nil)
}
