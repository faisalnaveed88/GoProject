// client.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fastorhandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[8:])

	res, err :=
		http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintf(w, string(robots))
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", " ", robots)

}

func main() {

	http.HandleFunc("/fastor/", fastorhandler)
	http.ListenAndServe(":8080", nil)

	/*conn, err := net.Dial("tcp", "localhost:6000")
	if err != nil {
		// handle error
	}
	fmt.Println("Connection successful!!", conn.RemoteAddr())
	recvdSlice := make([]byte, 11)
	conn.Read(recvdSlice)
	fmt.Println(string(recvdSlice))*/
}
