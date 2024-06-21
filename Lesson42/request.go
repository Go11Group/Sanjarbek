package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	fmt.Println("Received URL:", url)

	fmt.Fprintf(w, "URL: %s", url)

	client := http.Client{}

	switch r.Method {
	case "GET":
		res := Get(url)
		fmt.Fprintf(w, res)
	case "POST":
		str, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		res := Post(str, url, &client)
		fmt.Fprint(w, res)
	case "PUT":
		str, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		Put(&client, str, url)
	case "DELETE":
		str,err := io.ReadAll(r.Body)
		if err != nil{
			panic(err)
		}
		Delete(&client,url,str)
	}
}

func Get(url string) string {
	URL := "http://localhost:8080" + url

	resp, err := http.Get(URL)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func Post(info []byte, url string, client *http.Client) string {
	URL := "http://localhost:8080" + url
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(info))
	if err != nil {
		panic(err)
	}
	res, _ := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func Put(client *http.Client, str []byte, url string){
	URL := "http://localhost:8080" + url
	req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(str))

	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/text")
	_,err = client.Do(req)

	if err != nil {
		panic(err)
	}
}

func Delete(client *http.Client,url string,str []byte){
	URL := "http://localhost:8080" + url
	req,err := http.NewRequest("DELETE",URL,bytes.NewBuffer(str))

	if err != nil{
		panic(err)
	}
	req.Header.Set("Content-Type","application/json")
	_,err = client.Do(req)

	if err != nil{
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("Server started at :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

}