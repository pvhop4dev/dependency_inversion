package main

import (
	"dependency_inversion/internal/interfaces"
	"dependency_inversion/internal/service/servicea"
	"dependency_inversion/internal/service/serviceb"
	"dependency_inversion/internal/service/servicec"
	"log"
	"reflect"
)

func init() {
	log.Printf("Initializing")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}
func main() {
	a := servicea.NewStructA()
	b := serviceb.NewStructB()
	c := servicec.NewStructC()

	coreModels := interfaces.GetList()
	for _, model := range coreModels {
		log.Println(reflect.TypeOf(model))
	}

	for _, model := range coreModels {
		for _, callback := range coreModels {
			model.InitCallback(callback)
		}
	}

	log.Println(a.GetA1())
	log.Println(a.GetA2())
	log.Println(b.GetB1())
	log.Println(b.GetB2())
	log.Println(c.GetC1())
	log.Println(c.GetC2())
}
