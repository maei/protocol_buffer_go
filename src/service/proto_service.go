package service

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/maei/protocol_buffer_go/src/messages/complexpb"
	"github.com/maei/protocol_buffer_go/src/messages/simplepb"
	"github.com/maei/shared_utils_go/rest_errors"
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
	ProtoBuffToJSON(proto.Message) (string, error)
	JSONtoProtoBuff(string, proto.Message) error
	DoComplex() *complexpb.ComplexMessage
	WriteAddress([]byte) (*simplepb.SimpleMessage, rest_errors.RestErr)
}

type protoExampleService struct{}

func (p *protoExampleService) WriteAddress(bs []byte) (*simplepb.SimpleMessage, rest_errors.RestErr) {
	sm := &simplepb.SimpleMessage{}
	err := jsonpb.UnmarshalString(string(bs), sm)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("sth went wrong make json to proto", err)
	}
	return sm, nil
}

func (p *protoExampleService) DoSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:           1,
		IsSimple:     true,
		Name:         "Matthias",
		SampleList:   []int32{1, 2, 3, 4, 5},
		DayOfTheWeek: 4,
	}
	return &sm
}

func (p *protoExampleService) DoComplex() *complexpb.ComplexMessage {
	dm := complexpb.DummyMessage{
		Id:   2,
		Name: "Amphitheater",
	}

	cm := complexpb.ComplexMessage{
		OneDummy:      &dm,
		MultipleDummy: []*complexpb.DummyMessage{{Id: 5, Name: "Matthias"}, {Id: 6, Name: "Sonia"}},
	}
	return &cm
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

func (p *protoExampleService) ProtoBuffToJSON(message proto.Message) (string, error) {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(message)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return out, nil
}

func (p *protoExampleService) JSONtoProtoBuff(jsonInput string, message proto.Message) error {
	err := jsonpb.UnmarshalString(jsonInput, message)
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
