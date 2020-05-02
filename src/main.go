package main

import (
	"fmt"
	simplepb "github.com/maei/protocol_buffer_go/src/messages"
	"github.com/maei/protocol_buffer_go/src/service"
	"log"
)

func main() {
	erg := service.ProtoExampleService.DoSimple()
	fmt.Println(erg)

	test, err := service.ProtoExampleService.MarshalProtoBuff(erg)
	if err != nil {
		log.Println(err)
	}

	writeErr := service.ProtoExampleService.WriteFile(test)
	if writeErr != nil {
		log.Println(writeErr)
	}

	data, readErr := service.ProtoExampleService.ReadFile("output.txt")
	if readErr != nil {
		log.Println(readErr)
	}

	var m1 simplepb.SimpleMessage
	err2 := service.ProtoExampleService.UnmarshalProtoBuff(data, &m1)
	if err2 != nil {
		log.Println(err2)
	}

}
