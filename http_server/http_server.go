package main

import "fmt"
import "net/http"
import "net"

type MyHandler struct{}

func (h MyHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("User-Agent"))
	fmt.Println("Ahoj ahoj: ", r.RequestURI, "")
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	http.Serve(listener, new(MyHandler))

}
