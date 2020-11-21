package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//Handler main
type Handler struct {
	TargetStr string
	TargetDir string
	ErrorMsg  string
	FileInfos []string
	Warnings  []string
	Outputs   []string
}

func (h *Handler) recursiveReadDir(dir string, infos []os.FileInfo) {
	for _, info := range infos {
		if info.IsDir() {
			childInfos, err := ioutil.ReadDir(filepath.Join(dir, info.Name()))
			if err != nil {
				continue
			}
			h.recursiveReadDir(filepath.Join(dir, info.Name()), childInfos)
		} else {
			h.FileInfos = append(h.FileInfos, filepath.Join(dir, info.Name()))
		}
	}
}

func (h *Handler) execute() int {
	if h.TargetDir != "" {
		infos, err := ioutil.ReadDir(h.TargetDir)
		if err != nil {
			return h.sendError("Can not get the specified directory.")
		}
		h.recursiveReadDir(h.TargetDir, infos)
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return h.sendError("Can not get the current directory.")
		}

		infos, err := ioutil.ReadDir(wd)
		if err != nil {
			return h.sendError("Can not get files from the current directory.")
		}
		h.recursiveReadDir(wd, infos)
	}

	for _, fileInfo := range h.FileInfos {
		file, err := os.Open(fileInfo)
		if err != nil {
			h.Warnings = append(h.Warnings, "Can not open: "+fileInfo)
			continue
		}

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			h.Warnings = append(h.Warnings, "Can not read: "+fileInfo)
			continue
		}

		if strings.Contains(string(bytes), h.TargetStr) {
			h.Outputs = append(h.Outputs, fmt.Sprintf("Found \"%s\" in %s", h.TargetStr, fileInfo))
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
