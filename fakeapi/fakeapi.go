package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9527",
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/bettersun", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bettersun/hello", helloBS)

	log.Println("ListenAndServe 127.0.0.1:9527")
	server.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, bettersun")
}

func hello(w http.ResponseWriter, r *http.Request) {

	// 读取请求的Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("Request Body: %v", string(body)))

	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "[GET] Hello, world.")
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "[Post] Hello, world.")
		return
	}
	if r.Method == http.MethodPut {
		fmt.Fprintf(w, "[Put] Hello, world.")
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, "[Delete] Hello, world.")
		return
	}

	fmt.Fprintf(w, "unsupport http method.[hello]")
}

func helloBS(w http.ResponseWriter, r *http.Request) {

	// 读取请求的Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("Request Body: %v", string(body)))

	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "[GET] Hello, bettersun.")
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "[Post] Hello, bettersun.")
		return
	}
	if r.Method == http.MethodPut {
		fmt.Fprintf(w, "[Put] Hello, bettersun.")
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, "[Delete] Hello, bettersun.")
		return
	}

	fmt.Fprintf(w, "unsupport http method.[bettersun]")
}
