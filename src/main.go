package main

import (
	"fmt"
	"github.com/maei/protocol_buffer_go/src/domain"
)

func main() {
	p1 := domain.Person{"Matthias"}
	fmt.Println(p1)
}
