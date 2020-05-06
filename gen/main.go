package main

import (
	"fmt"
	"github.com/yangtizi/cz88"
)

func main() {
	fmt.Println(cz88.GetAddress("47.56.100.100"))  
	fmt.Println(cz88.GetAddressShort("47.56.100.100"))  
}