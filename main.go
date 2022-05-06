package main

import (
	"fmt"

	"github.com/feloy/test-k8s-gen/pack"
)

func main() {
	myStructure := pack.MyStructure{
		Name: "my name",
		Sub: pack.SubStructure{
			Address: "my address",
		},
	}
	myCopy := myStructure.DeepCopy()
	fmt.Printf("%+v\n", myCopy)
}
