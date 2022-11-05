package main

import (
	"doce-panda/infra/api"
	"fmt"
)

func main() {
	err := api.Server()
	if err != nil {
		fmt.Errorf("Error when started server")
	}
}
