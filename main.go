package main

import (
	"encoding/json"
	"flag"
	"fmt"
	d "generate-raffle-book/document"
	p "generate-raffle-book/parameters"
	"os"
)

var (
	start *int
	count *int
	help  *bool
	temp  *string
	force *bool
)

func init() {
	start = flag.Int("start", 1, "what is the 1st desired number ?")
	count = flag.Int("count", 10, "how many raffle tickets would you like ?")
	help = flag.Bool("help", false, "show this help")
	temp = flag.String("template", "", "path to the customs file to create")
	force = flag.Bool("force", false, "enforce overwriting customs file")
}

func main() {
	flag.Parse()

	if *help {
		DisplayHelp()
		os.Exit(0)
	}

	if *temp != "" {
		OutputCustomFile()
		os.Exit(0)
	}

	Generate()
}

func Generate() {
	parameters, err := LoadParameters()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	generator := d.Generator{}
	generator.Parameters = &parameters
	generator.NumberLength = 5
	numberer := d.CreateNumberer(10, 3, uint16(*start), uint16(*count))
	generator.Numberer = numberer

	generator.Generate()
}

func LoadParameters() (p.Parameters, error) {
	parameters, err := p.LoadParameters("assets/parameters.json")
	if err != nil {
		return parameters, err
	}
	customs, err := p.LoadCustoms("./customs.json")
	if err != nil {
		return parameters, err
	}

	parameters.Customize(customs)

	return parameters, nil
}

func DisplayHelp() {
	fmt.Println("Usage: generate-raffle-book [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func OutputCustomFile() {
	if !*force {
		stat, _ := os.Stat("customs.json")
		if stat != nil {
			fmt.Println("customs.json file already exists. Rerun the command with -force flag to enforce overwrtiting it.")
			os.Exit(1)
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
}
