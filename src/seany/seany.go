package main

import (
  "fmt"
  "net/http"
)

func scan(url string){
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
		return
	}
	defer resp.Body.Close()
  fmt.Println(resp.Status)
}

func main(){
  scan("http://pepsico.s3.amazonaws.com/")
  fmt.Println("yay!")
}
