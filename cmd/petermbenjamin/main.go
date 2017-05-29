package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
	me "github.com/petermbenjamin/go-petermbenjamin"
)

var (
	nameFlag    = flag.Bool("n", false, "Print Name")
	contactFlag = flag.Bool("c", false, "Print Contact Info")
	emailFlag   = flag.Bool("e", false, "Print Email Address")
)

func main() {

	flag.Parse()

	info := color.New(color.Bold)

	if flag.NFlag() == 0 {
		info.Println("Usage: petermbenjamin <options>")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	if *nameFlag {
		info.Println(me.FullName())
		fmt.Println(me.Title())
	}

	contacts := me.Contact()

	if *contactFlag {
		for k, v := range contacts {
			fmt.Println(k, "-", v)
		}
	}

	if *emailFlag {
		fmt.Println(contacts["email"])
	}
}
