/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/txrm/github-action-ide/api" // to be removed
	"github.com/txrm/github-action-ide/cmd"
)

func main() {
	go api.StartServer()
	cmd.Execute()
}
