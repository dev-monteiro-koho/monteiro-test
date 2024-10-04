package main

import (
	"fmt"

	"github.com/dev-monteiro-koho/test/common"
)

func main() {
	fmt.Println("Module B depends on Module A")
	common.Main()
}
