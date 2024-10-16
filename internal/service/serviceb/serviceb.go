package serviceb

import (
	"dependency_inversion/internal/interfaces"
	"fmt"

	"github.com/sirupsen/logrus"
)

type structB struct {
	serviceACallBack interfaces.IServiceA
}

// GetB1 implements interfaces.IServiceB.
func (s *structB) GetB1() string {
	return "B1"
}

// GetB2 implements interfaces.IServiceB.
func (s *structB) GetB2() string {
	a2 := s.serviceACallBack.GetA2()
	return fmt.Sprintf("GetB2 from A2 %s", a2)
}

func (s *structB) Register() {
	logrus.Println("Register")
	interfaces.Register(s)
}

func (s *structB) InitCallback(coreModel interfaces.CoreModel) {
	switch coreModel := coreModel.(type) {
	case interfaces.IServiceA:
		s.serviceACallBack = coreModel
	}
}

var s *structB

func NewStructB() interfaces.IServiceB {
	if s == nil {

		s = &structB{}
	}
	s.Register()
	return s
}
