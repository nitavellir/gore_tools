package main

import (
	"flag"
	"log"
)

var (
	targetStr string
	targetDir string
)

func main() {
	//arg
	flag.StringVar(&targetStr, "s", "", "Required: Specify the phrase you want to find.")
	flag.StringVar(&targetDir, "d", "", "Specify the directory you want to dig.")
	flag.Parse()
	if targetStr == "" {
		log.Fatal("The phrase is not specified.")
	}

	//init handler
	h := &Handler{
		TargetStr: targetStr,
		TargetDir: targetDir,
	}

	//execute
	if status := h.execute(); status != 0 {
		log.Fatal(h.ErrorMsg)
	} else {
		h.outputResponse()
	}
}
