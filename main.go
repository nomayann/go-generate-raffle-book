package main

import (
	"encoding/json"
	"fmt"
	c "generate-raffle-book/color"
	d "generate-raffle-book/document"
	p "generate-raffle-book/parameters"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	parametersFilePath string = "assets/parameters.json"
	customsFilePath    string = "customs.json"
)

func main() {
	MainMenu()
}

func MainMenu() {
	answer := DisplayInstructions()

	switch answer {
	case 1:
		Generate()
		MainMenu()
	case 2:
		OutputCustomFile()
		MainMenu()
	case 3:
		DisplayHelp()
		MainMenu()
	case 4:
		Quit()
		os.Exit(0)
	}
}

func Quit() {
	fmt.Println("\n\n", c.Green, "Bye", c.Reset)
}

func clear() {
	startTime := time.Now()
	fmt.Println("Start time:", startTime)

	timer := time.After(1 * time.Second)

	fmt.Println("Waiting for 1 second...")
	endTime := <-timer

	fmt.Println("Action triggered after 1 second at:", endTime)

	command := exec.Command("command")
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		fmt.Println(c.Red, "Fail to clear the terminal", c.Reset)
	}
}

func Generate() {
	parameters, err := p.LoadParametersWithCustom(parametersFilePath, customsFilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var start uint16 = 1
	var count uint16 = 120
	fmt.Println(c.Blue, "type the desired start number", c.Reset)
	_, err = fmt.Scan(&start)
	if err != nil {
		fmt.Println(c.Red, "Invalid start number. default value applied ", start, c.Reset)
	}
	fmt.Println(c.Blue, "type the desired number of tickets", c.Reset)
	_, err = fmt.Scan(&count)
	if err != nil {
		fmt.Println(c.Red, "Invalid count number. default value applied ", count, c.Reset)
	}

	generator := d.Generator{}
	generator.Parameters = &parameters
	generator.NumberLength = 5
	numberer := d.CreateNumberer(10, 3, uint16(start), uint16(count))
	generator.Numberer = numberer

	first, last, count := generator.Numberer.GetFirstLastAndCount()
	fmt.Println("\nℹ️  Computed values ℹ️")
	fmt.Printf(
		"%sFirst ticket: %d\nLast ticket: %d\nNumber of tickets: %d %s\n\n",
		c.Blue,
		first,
		last,
		count,
		c.Reset,
	)

	fileInfo, _ := generator.Generate()
	if fileInfo != nil {
		fmt.Printf("Raffle book generated here: %s\"%s\"%s\n", c.Green, fileInfo.Name(), c.Reset)
		clear()
		//MainMenu()
	}

	os.Exit(1)
}

func DisplayInstructions() int {
	fmt.Println("\n\nGenerate Raffle Book")
	answers := []string{
		"Generate a new raffle book",
		"Generate a boilerplate custom file",
		"Display help",
		"Quit",
	}
	var answer string
	for {
		if answer == "" {
			fmt.Println(c.Blue, "What do you want to do?", c.Reset)
		}
		answerAsInt, err := strconv.Atoi(answer)
		if err == nil && answerAsInt >= 1 && answerAsInt <= len(answers) {
			break
		}
		for key, value := range answers {
			fmt.Printf("%d. %s\n", key+1, value)
		}
		fmt.Print(c.Purple, "Your answer: ", c.Reset)
		_, err = fmt.Scan(&answer)
		if err != nil {
			fmt.Println(c.Red, "Invalid answer", c.Reset)
		}
	}
	result, _ := strconv.Atoi(answer)
	return result
}

func DisplayHelp() {
	fmt.Println("\n\nUsage: ", c.Purple, "generate-raffle-book", c.Reset)
	fmt.Println("follow the prompted instructions")
}

func OutputCustomFile() {
	stat, _ := os.Stat(customsFilePath)
	if stat != nil {
		fmt.Println(c.Blue, "customs.json file already exists. Overwrite it ? (yes/no).", c.Reset)
		answer := ""
		_, err := fmt.Scan(&answer)
		if err != nil {
			fmt.Println(c.Red, "Invalid answer", c.Reset)
		}
		if answer != "yes" && answer != "y" {
			fmt.Println("Aborted")
			MainMenu()
		}
	}
	output, err := json.MarshalIndent(p.CreateBoilerPlateCustoms(), "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = os.WriteFile("customs.json", output, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("customs.json file created")
	fmt.Println(c.Green, "ℹ️  Edit it and rerun the program for generating the raffle book", c.Reset)
}
