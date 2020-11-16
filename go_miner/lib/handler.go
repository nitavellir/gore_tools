package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Handler main
type Handler struct {
	TargetStr string
	TargetDir string
	ErrorMsg  string
	Warnings  []string
	Outputs   []string
}

func (h *Handler) execute() int {
	fileInfos := []os.FileInfo{}
	if h.TargetDir != "" {
		infos, err := ioutil.ReadDir(h.TargetDir)
		if err != nil {
			return h.sendError("Can not get the specified directory.")
		}
		fileInfos = infos
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return h.sendError("Can not get the current directory.")
		}

		infos, err := ioutil.ReadDir(wd)
		if err != nil {
			return h.sendError("Can not get files from the current directory.")
		}
		fileInfos = infos
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		file, err := os.Open(fileInfo.Name())
		if err != nil {
			h.Warnings = append(h.Warnings, "Can not open: "+fileInfo.Name())
			continue
		}

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			h.Warnings = append(h.Warnings, "Can not read: "+fileInfo.Name())
			continue
		}

		if strings.Contains(string(bytes), h.TargetStr) {
			h.Outputs = append(h.Outputs, fmt.Sprintf("Found \"%s\" in %s", h.TargetStr, fileInfo.Name()))
		}
	}

	return 0
}

func (h *Handler) sendError(errMsg string) int {
	h.ErrorMsg = errMsg
	return 1
}

func (h *Handler) outputResponse() {
	if len(h.Warnings) > 0 {
		fmt.Println(fmt.Sprintf("------ WARNING ------\n%s", strings.Join(h.Warnings, "\n")))
	}
	fmt.Println(fmt.Sprintf("------ RESULT ------\n%s", strings.Join(h.Outputs, "\n")))
}
