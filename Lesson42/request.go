package main

import "module/service/handler"

func main(){
	h := handler.NewHandler()
	server := handler.CreateServer(h)
	server.ListenAndServe()
}