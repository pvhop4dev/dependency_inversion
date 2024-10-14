package servicea

import (
	"dependency_inversion/internal/interfaces"
	"fmt"
	"log"
)

type structA struct {
	serviceBCallBack interfaces.IServiceB
}

// GetA1 implements interfaces.IServiceA.
func (s *structA) GetA1() string {
	b1 := s.serviceBCallBack.GetB1()
	return fmt.Sprintf("GetA1  from B1 %s", b1)
}

// GetA2 implements interfaces.IServiceA.
func (s *structA) GetA2() string {
	return "A2"
}

func (s *structA) Register() {
	log.Println("Register")
	interfaces.Register(s)
}

func (s *structA) InitCallback(coreModel interfaces.CoreModel) {
	switch coreModel := coreModel.(type) {
	case interfaces.IServiceB:
		s.serviceBCallBack = coreModel
	}
}

var s *structA

func NewStructA() interfaces.IServiceA {
	if s == nil {
		s = &structA{}
	}
	s.Register()
	return s
}
