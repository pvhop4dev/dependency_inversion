package servicec

import (
	"dependency_inversion/internal/interfaces"
	"fmt"
	"log"
)

type structC struct {
	serviceACallBack interfaces.IServiceA
	serviceBCallBack interfaces.IServiceB
}

// GetC1 implements interfaces.IServiceC.
func (s *structC) GetC1() string {
	a1 := s.serviceACallBack.GetA1()
	return fmt.Sprintf("GetC1 from A1 %s", a1)
}

// GetC2 implements interfaces.IServiceC.
func (s *structC) GetC2() string {
	b2 := s.serviceBCallBack.GetB2()
	return fmt.Sprintf("GetC2 from B2 %s", b2)
}

func (s *structC) Register() {
	log.Println("Register")
	interfaces.Register(s)
}

func (s *structC) InitCallback(coreModel interfaces.CoreModel) {
	switch coreModel := coreModel.(type) {
	case interfaces.IServiceA:
		s.serviceACallBack = coreModel
	case interfaces.IServiceB:
		s.serviceBCallBack = coreModel
	}
}

var s *structC

func NewStructC() interfaces.IServiceC {
	if s == nil {
		s = &structC{}
	}
	s.Register()
	return s
}
