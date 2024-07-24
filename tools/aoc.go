package main

import "fmt"

var cli_commands []Command

type Command struct {
	Arity 			int
	Command 		string
	Description string
	Examples 		[]string
	Handler 		func(Command, []string)error
}

func CreateCommand (name string, desc string, handler func(Command, []string)error) Command {
	return Command{
		Command: name,
		Description: desc,
		Handler: handler,
	}
}

func (self Command) SetArity (arity int) Command {
	self.Arity = arity
	return self
}

func (self Command) AddExample (example string) Command {
	self.Examples = append(self.Examples, example)
	return self
}

func (self *Command) Display () *Command {
	fmt.Printf("/%s |%d|\n - %s\n"	, self.Command, self.Arity, self.Description)

	if len(self.Examples) > 0 {
		for _, example := range self.Examples {
			fmt.Printf(" ex: `/%s %s`\n",self.Command, example)
		}
	}

	println("")
	return self
}

func main () {
	cli_commands = []Command{
		CreateCommand("help", "Display's helpful information regarding CLI usage", HelpHandler).SetArity(0),
		CreateCommand("scaffold", "Scaffold a Go project for an entire year's AOC. Downloads all required and available file's & puzzle-inputs if possible.", ScaffoldHelper).SetArity(2).AddExample("2024 3"),
	}

	HelpHandler(cli_commands[0], []string{})
}
