package service

import (
	"github.com/maei/protocol_buffer_go/src/domain"
	"strings"
)

var TestService testServiceInterface = &testService{}

type testServiceInterface interface {
	CheckName(domain.Person) bool
}
type testService struct{}

func (t *testService) CheckName(p domain.Person) bool {
	if strings.Contains("M", p.Name) {
		return true
	}
	return false
}
