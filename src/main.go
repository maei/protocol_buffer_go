package main

import (
	"fmt"
	"github.com/maei/protocol_buffer_go/src/messages/simplepb"

	"github.com/maei/protocol_buffer_go/src/service"
	"log"
)

func main() {
	// Create a struct out of the protocol-buffer file
	erg := service.ProtoExampleService.DoSimple()
	fmt.Println(erg)

	// Make []byte out of the struct with inbuild marshal function
	bs, err := service.ProtoExampleService.MarshalProtoBuff(erg)
	if err != nil {
		log.Println(err)
	}
	// write byte-slice into a textfile
	writeErr := service.ProtoExampleService.WriteFile(bs)
	if writeErr != nil {
		log.Println(writeErr)
	}

	// Read out of a text-file
	data, readErr := service.ProtoExampleService.ReadFile("output.txt")
	if readErr != nil {
		log.Println(readErr)
	}

	// unmarshal byte-slice into a protocol-struct
	var m1 simplepb.SimpleMessage
	// m1 := &simplepb.SimpleMessage{}
	err2 := service.ProtoExampleService.UnmarshalProtoBuff(data, &m1)
	if err2 != nil {
		log.Println(err2)
	}

	// make a sting/JSON out of the protocol-buffer struct
	resultJSON, jsonErr := service.ProtoExampleService.ProtoBuffToJSON(&m1)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	fmt.Println(resultJSON)

	// JSON to Protobuff structure
	// var m2 simplepb.SimpleMessage
	m2 := &simplepb.SimpleMessage{}
	//service.ProtoExampleService.JSONtoProtoBuff(resultJSON, &m2)
	service.ProtoExampleService.JSONtoProtoBuff(resultJSON, m2)
	fmt.Println(m2)
	// fmt.Println(&m2)

}
