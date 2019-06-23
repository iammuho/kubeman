package cmd

import (
	"flag"
	"log"

	"github.com/spf13/cobra"
)

// Variables
var (
	DefaultConfigPath = "terraform"
)

// Flag options from CLI
type Options struct {
    Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
    Output string `short:"o" long:"output" description:"What is the output type?"`
}

// Controller struct
type kubeManController struct {
	Options Options
}

// MAIN Function that to be run on init
func main() {

	verbose := flag.Bool("v", true, "Show verbose debug information")
	output := flag.String("o", "terraform", "Generate output as")

	flag.Parse()

	options := Options {
		Verbose: *verbose,
		Output: *output,
	}

	kubeMan := kubeManController {
		Options: options,
	}

    kubeMan.Run()
}

// RUN function that to be triggered by main function for initial startup
func (k *kubeManController) Run() {

	log.Print("WKM app started")
}