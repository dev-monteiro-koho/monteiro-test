package main

import (
	"fmt"

	"github.com/dev-monteiro-koho/monteiro-test/common"
)

func main() {
	fmt.Println("Module service depends on module common")
	common.Main()
}
