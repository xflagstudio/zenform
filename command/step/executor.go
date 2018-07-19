package step

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Executor runs steps sequentially by Step method and manage step nest level
//   to pretty print the result.
type Executor struct {
	indentLevel int
	padding     string
}

// NewExecutor creates initialized Executor instance.
func NewExecutor() *Executor {
	return &Executor{
		indentLevel: 0,
		padding:     "    ", // padding left spaces
	}
}

// Step executes stepFunc and returns its result.
func (exe *Executor) Step(msgOnStart string, stepFunc func() error) error {
	// switch arrow type
	var arrow string
	if exe.indentLevel == 0 {
		arrow = "==> "
	} else {
		arrow = "--> "
	}

	fmt.Println(strings.Repeat(exe.padding, exe.indentLevel) + arrow + msgOnStart)

	// increment indent level and decrement when finish stepFunc
	exe.indentLevel++
	defer func() { exe.indentLevel-- }()

	// execute stepFunc and display result
	err := stepFunc()
	if err != nil {
		return err
	}
	return nil
}

// StepIf is conditional execution of step. If condition is false, this step will be skipped.
func (exe *Executor) StepIf(condition bool, msgOnStart string, stepFunc func() error) error {
	if !condition {
		return nil
	}
	return exe.Step(msgOnStart, stepFunc)
}

func (exe *Executor) Info(msg string) {
	fmt.Println(strings.Repeat(exe.padding, exe.indentLevel) + msg)
}

func (exe *Executor) Error(msg string) {
	color.Red(strings.Repeat(exe.padding, exe.indentLevel) + msg)
}

func (exe *Executor) Success(msg string) {
	color.Green(strings.Repeat(exe.padding, exe.indentLevel) + msg)
}
