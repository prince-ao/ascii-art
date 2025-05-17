package main

import (
	"github.com/alecthomas/kong"
	"asciify/internal/commands"
	"log"
)

var CLI struct {
	Info struct {
		Path string `arg:"" name:"path" help:"Path to image."`
	} `cmd:"" help:"Get info on file."`

	Convert struct {
		Width *int `optional:"" help:"Desired width of image."`

		Path string `arg:"" name:"path" help:"Path to image."`
	} `cmd:"" help:"Convert image to ascii."`
}

func main() {
	ctx := kong.Parse(&CLI)

	switch ctx.Command() {
	case "info <path>":
		commands.HandleInfo(CLI.Info.Path)
	case "convert <path>":
		commands.HandleConvert(CLI.Convert.Path, CLI.Convert.Width)
	default:
		log.Fatal("command not found.")
	}
}
