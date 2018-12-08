package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/jfrogtraining/project-examples/go-module-example/demo"
)

var args struct {
	Text string `arg:"required"`
}

func main() {
	arg.MustParse(&args)
	result := demo.DecorateString(args.Text)
	fmt.Println(result)
}
