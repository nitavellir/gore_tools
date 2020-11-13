package main

import "os"

type Handler struct {
	TargetStr string
	ErrorMsg  string
}

func (h *Handler) execute() int {
	_, err := os.Getwd()
	if err != nil {
		return h.sendError("Can not get the current directory.")
	}

	return 0
}

func (h *Handler) sendError(err_msg string) int {
	h.ErrorMsg = err_msg
	return 1
}
