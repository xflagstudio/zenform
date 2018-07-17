package command

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type StepExecutor struct {
	indentLevel int
	padding     string
}

func NewStepExecutor() *StepExecutor {
	return &StepExecutor{
		indentLevel: 0,
		padding:     "    ", // padding left spaces
	}
}

func (exe *StepExecutor) Step(msgOnStart string, stepFunc func() error) error {
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

func (exe *StepExecutor) StepIf(condition bool, msgOnStart string, stepFunc func() error) error {
	if !condition {
		return nil
	}
	return exe.Step(msgOnStart, stepFunc)
}

func (exe *StepExecutor) Info(msg string) {
	fmt.Println(strings.Repeat(exe.padding, exe.indentLevel) + msg)
}

func (exe *StepExecutor) Error(msg string) {
	color.Red(strings.Repeat(exe.padding, exe.indentLevel) + msg)
}

func (exe *StepExecutor) Success(msg string) {
	color.Green(strings.Repeat(exe.padding, exe.indentLevel) + msg)
}
