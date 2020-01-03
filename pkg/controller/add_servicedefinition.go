package controller

import (
	"servicedefinition/pkg/controller/servicedefinition"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, servicedefinition.Add)
}
