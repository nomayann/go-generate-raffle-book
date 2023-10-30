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
)

func init() {
	start = flag.Int("start", 1, "what is the 1st desired number ?")
	count = flag.Int("count", 10, "how many books would you like ?")
}

func main() {
	flag.Parse()

	parameters, err := p.LoadParameters("assets/parameters.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	customs, err := p.LoadCustoms("assets/customs.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	parameters.Customize(customs)

	generator := d.Generator{}
	generator.Parameters = &parameters
	generator.NumberLength = 5
	numberer := d.CreateNumberer(10, 3, 1)
	param := d.BookParam{Start: uint16(*start), Count: uint16(*count)}

	numberer.SetBookParam(param)
	generator.Numberer = numberer

	generator.Generate()
}
