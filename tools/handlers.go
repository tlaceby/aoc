package main

func HelpHandler (cmd Command, args []string) error {
	println("\n------------------  Help  ------------------")
	for _, command := range cli_commands {
		command.Display()
	}

	println("--------------------------------------------")
	return nil
}

func ScaffoldHelper (cmd Command, args []string) error {

	return nil
}
