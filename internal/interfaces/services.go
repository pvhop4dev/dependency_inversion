package interfaces

import "github.com/sirupsen/logrus"

var listModel []CoreModel

type CoreModel interface {
	Register()
	InitCallback(coreModel CoreModel)
}

type IServiceA interface {
	CoreModel
	GetA1() string
	GetA2() string
}

type IServiceB interface {
	CoreModel
	GetB1() string
	GetB2() string
}

type IServiceC interface {
	CoreModel
	GetC1() string
	GetC2() string
}

func Register(m CoreModel) {
	logrus.Println("Register")
	listModel = append(listModel, m)
}

func GetList() []CoreModel {
	return listModel
}
