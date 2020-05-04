package main

import (
	"fmt"
	"github.com/maei/protocol_buffer_go/src/app"
	"github.com/maei/protocol_buffer_go/src/messages/simplepb"
	"github.com/maei/protocol_buffer_go/src/service"
	"github.com/maei/shared_utils_go/logger"
	"log"
)

func main() {
	logger.Info("starting protobuff-service")
	app.StartApplication()

	// Create a struct out of the protocol-buffer file
	erg := service.ProtoExampleService.DoSimple()
	fmt.Println(erg)

	// Make []byte out of the struct with in build marshal function
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

	/*	p1 := address.Person{Name: "Matthias", Id: 1, Email: "eiletz@oecg.de", Phones: []*address.Person_PhoneNumber{{Number: "0162215225", Type: 2},},}
		ab2 := service.AddressService.WriteAddress(p1)
		fmt.Println(ab2)

		service.AddressService.ReadAddress()
		p2 := address.Person{Name: "Sonia", Id: 1, Email: "sonia@oecg.de", Phones: []*address.Person_PhoneNumber{{Number: "1337", Type: 1},},}
		ab2 = service.AddressService.WriteAddress(p2)
		fmt.Println(ab2)*/
}
