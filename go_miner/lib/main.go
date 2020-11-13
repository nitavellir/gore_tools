package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var (
	target_str string
)

func main() {
	//arg
	flag.StringVar(&target_str, "s", "", "Specify the phrase you want to find.")
	flag.Parse()
	if target_str == "" {
		log.Fatal("The phrase is not specified.")
	}

	//init handler
	h := &Handler{
		TargetStr: target_str,
	}

	//execute
	if status := h.execute(); status != 0 {
		log.Fatal(h.ErrorMsg)
	} else {
		fmt.Println(strings.Join(h.Outputs, "\n"))
	}
}
