package main

import (
	"flag"
	"log"
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
	}
}
