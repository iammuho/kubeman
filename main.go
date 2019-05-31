/*
Copyright 2019 - Muhammet Arslan
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"log"
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