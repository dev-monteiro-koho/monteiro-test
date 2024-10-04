package main

import (
	"fmt"

	"github.com/dev-monteiro-koho/monteiro-test/common"
	"github.com/go-faker/faker/v4"
)

type Info struct {
	paramA string
	paramB string
	paramC string
}

func main() {
	var i Info
	err := faker.FakeData(&i)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Data but that's okay cause so are you: %v\n", i)
	common.Main()
}
