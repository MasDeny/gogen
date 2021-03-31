package main

import (
	"flag"
	"fmt"

	"github.com/mirzaakhena/gogen2/gogencommand"
)

func main() {

	flag.Parse()

	cmds := map[string]func() (gogencommand.Commander, error){
		"usecase":     gogencommand.NewUsecaseModel,
		"entity":      gogencommand.NewEntityModel,
		"method":      gogencommand.NewMethodModel,
		"enum":        gogencommand.NewEnumModel,
		"error":       gogencommand.NewErrorModel,
		"valueobject": gogencommand.NewValueObjectModel,
		"valuestring": gogencommand.NewValueStringModel,
		"repository":  gogencommand.NewRepositoryModel,
		"service":     gogencommand.NewServiceModel,
		"controller":  gogencommand.NewControllerModel,
		"gateway":     gogencommand.NewGatewayModel,
	}

	var obj gogencommand.Commander
	var err error

	f, ok := cmds[flag.Arg(0)]
	if !ok {
		fmt.Printf("ERROR : %v", "Command is not recognized")
	}

	obj, err = f()
	if err != nil {
		fmt.Printf("ERROR : %v", err.Error())
		return
	}

	err = obj.Run()
	if err != nil {
		fmt.Printf("ERROR : %v", err.Error())
	}

}
