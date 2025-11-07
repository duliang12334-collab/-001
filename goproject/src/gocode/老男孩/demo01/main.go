package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "maysql ... -u root123 -p qwe123321.."

	uIdnex := strings.Index(s, "-u")
	pIdnex := strings.Index(s, "-p")
	fmt.Println(uIdnex)
	fmt.Println(s[uIdnex+3 : pIdnex-1])
	fmt.Println(s[pIdnex+3:])
	fmt.Println(s[pIdnex+3:])

}
