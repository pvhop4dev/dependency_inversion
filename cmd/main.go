package main

import (
	"dependency_inversion/internal/interfaces"
	"dependency_inversion/internal/service/servicea"
	"dependency_inversion/internal/service/serviceb"
	"dependency_inversion/internal/service/servicec"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(multiWriter)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		// DisableColors:   false,
		DisableQuote:    true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := filepath.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Println("Initializing")
}

func main() {
	a := servicea.NewStructA()
	b := serviceb.NewStructB()
	c := servicec.NewStructC()

	coreModels := interfaces.GetList()
	for _, model := range coreModels {
		logrus.Println(reflect.TypeOf(model))
	}

	for _, model := range coreModels {
		for _, callback := range coreModels {
			model.InitCallback(callback)
		}
	}

	logrus.Println(a.GetA1())
	logrus.Println(a.GetA2())
	logrus.Println(b.GetB1())
	logrus.Println(b.GetB2())
	logrus.Println(c.GetC1())
	logrus.Println(c.GetC2())
	logrus.Debugln("Debug")
	logrus.Infoln("Info")
	logrus.Warnln("Warn")
	logrus.Errorln("Error")
	logrus.Debugln("Debug")
	logrus.Fatalf("Fatal")
}
