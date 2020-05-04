package service

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	address "github.com/maei/protocol_buffer_go/src/messages/addresspb"
	"github.com/maei/shared_utils_go/logger"
	"github.com/maei/shared_utils_go/rest_errors"
)

var AddressService addressServiceInterface = &addressService{}

var addressBook address.AddressBook

type addressServiceInterface interface {
	WriteAddress([]byte) (*address.AddressBook, rest_errors.RestErr)
	ReadAddress() rest_errors.RestErr
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
	return &addressBook, nil
}

func (a *addressService) ReadAddress() rest_errors.RestErr {
	fmt.Println(&addressBook)
	return nil
}
