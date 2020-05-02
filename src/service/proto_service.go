package service

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/maei/protocol_buffer_go/src/messages"
	"io/ioutil"
	"log"
)

var ProtoExampleService protoExampleInterface = &protoExampleService{}

type protoExampleInterface interface {
	DoSimple() *simplepb.SimpleMessage
	MarshalProtoBuff(proto.Message) ([]byte, error)
	UnmarshalProtoBuff([]byte, proto.Message) error
	WriteFile([]byte) error
	ReadFile(string) ([]byte, error)
}

type protoExampleService struct{}

func (p *protoExampleService) DoSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1,
		IsSimple:   true,
		Name:       "Matthias",
		SampleList: []int32{1, 2, 3, 4, 5},
	}
	return &sm
}

func (p *protoExampleService) MarshalProtoBuff(message proto.Message) ([]byte, error) {
	bs, err := proto.Marshal(message)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bs, nil
}
func (p *protoExampleService) UnmarshalProtoBuff(bs []byte, message proto.Message) error {
	// das hier nehmen, wenn ich hier die struct initialisieren will und von der main abkapseln möchte
	//m1 := &simplepb.SimpleMessage{}
	err := proto.Unmarshal(bs, message)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(message)
	return nil
}

func (p *protoExampleService) WriteFile(bs []byte) error {
	err := ioutil.WriteFile("output.txt", bs, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *protoExampleService) ReadFile(filename string) ([]byte, error) {
	out, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return out, nil

}
