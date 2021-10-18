package main

import (
	"fmt"
	infra "github.com/hasegit/practice-go/infra"
	adapter "github.com/hasegit/practice-go/adapter"
)

func main(){
	client := infra.GetClient("api_key")
	info := adapter.Access(client)
	fmt.Println(info)
}
