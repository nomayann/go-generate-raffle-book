package main

import (
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
)

func init() {
	start = flag.Int("start", 1, "what is the 1st desired number ?")
	count = flag.Int("count", 10, "how many books would you like ?")
	help = flag.Bool("help", false, "show this help")
}

func main() {
	flag.Parse()

	if *help {
		DisplayHelp()
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
	customs, err := p.LoadCustoms("assets/customs.json")
	if err != nil {
		return parameters, err
	}

	parameters.Customize(customs)

	return parameters, nil
}

func DisplayHelp() {
	fmt.Println("Usage: generate-raffle-book [options]")
	fmt.Println("Options:")
	fmt.Println("  -start=1\n What is the 1st desired start ticket number ?")
	fmt.Println("  -count=10\n How many tickets would you like ?")
}
