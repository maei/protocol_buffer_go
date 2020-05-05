package service

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	address "github.com/maei/protocol_buffer_go/src/messages/addresspb"
	"github.com/maei/shared_utils_go/logger"
	"github.com/maei/shared_utils_go/rest_errors"
	"io/ioutil"
)

var AddressService addressServiceInterface = &addressService{}

var addressBook address.AddressBook

type addressServiceInterface interface {
	WriteAddress([]byte) (*address.AddressBook, rest_errors.RestErr)
	ReadAddress() (*address.AddressBook, rest_errors.RestErr)
}

type addressService struct{}

func (a *addressService) WriteAddress(bs []byte) (*address.AddressBook, rest_errors.RestErr) {
	p := &address.Person{}
	err := jsonpb.UnmarshalString(string(bs), p)
	if err != nil {
		logger.Error("sth went wrong make json to proto", err)
		return nil, rest_errors.NewInternalServerError("sth went wrong make json to proto", err)
	}
	addressBook.People = append(addressBook.People, p)

	bs, errMar := proto.Marshal(&addressBook)
	if errMar != nil {
		logger.Error("sth went wrong make proto to byte", errMar)
		return nil, rest_errors.NewInternalServerError("sth went wrong make proto to byte", errMar)
	}

	writeErr := ioutil.WriteFile("addressbook.txt", bs, 0644)
	if writeErr != nil {
		logger.Error("sth went wrong writing to file", writeErr)
		return nil, rest_errors.NewInternalServerError("sth went wrong writing to file", writeErr)
	}
	return &addressBook, nil
}

func (a *addressService) ReadAddress() (*address.AddressBook, rest_errors.RestErr) {
	bs, readErr := ioutil.ReadFile("addressbook.txt")
	if readErr != nil {
		logger.Error("sth went wrong reading file", readErr)
		return nil, rest_errors.NewInternalServerError("sth went wrong reading file", readErr)
	}
	ab := &address.AddressBook{}
	protoErr := proto.Unmarshal(bs, ab)
	if protoErr != nil {
		logger.Error("sth went wrong creating protobuf", protoErr)
		return nil, rest_errors.NewInternalServerError("sth went wrong creating protobuf", protoErr)
	}

	/*	marshaller := jsonpb.Marshaler{}
		result, marshErr := marshaller.MarshalToString(ab)
		if marshErr != nil {
			logger.Error("sth went wrong marshal protobuf", marshErr)
			return "", rest_errors.NewInternalServerError("sth went wrong marshal protobuf", marshErr)
		}
		//fmt.Println(&addressBook)*/
	return ab, nil
}
