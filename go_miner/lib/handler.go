package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Handler struct {
	TargetStr string
	ErrorMsg  string
	Outputs   []string
}

func (h *Handler) execute() int {
	wd, err := os.Getwd()
	if err != nil {
		return h.sendError("Can not get the current directory.")
	}

	file_infos, err := ioutil.ReadDir(wd)
	if err != nil {
		return h.sendError("Can not get files from the current directory.")
	}

	for _, file_info := range file_infos {
		if file_info.IsDir() {
			continue
		}

		file, err := os.Open(file_info.Name())
		if err != nil {
			return h.sendError("Can not open: " + file_info.Name())
		}

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return h.sendError("Can not read: " + file_info.Name())
		}

		if strings.Contains(string(bytes), h.TargetStr) {
			h.Outputs = append(h.Outputs, fmt.Sprintf("Found \"%s\" in %s", h.TargetStr, file_info.Name()))
		}
	}

	return 0
}

func (h *Handler) sendError(err_msg string) int {
	h.ErrorMsg = err_msg
	return 1
}

func (h *Handler) outputResponse() {
	fmt.Println(fmt.Sprintf("------ RESULT ------\n%s", strings.Join(h.Outputs, "\n")))
}
