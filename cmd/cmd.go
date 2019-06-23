package cmd

import (
	"log"

	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Variables
var (
	generatorVersion string
)

// Execute will run main logic
func Execute(version string) {
	generatorVersion = version

	log.Print("WKM app started")

	cmd := &cobra.Command{
		Use:     "generator",
		Short:   "CLI for generating terraform variables",
		Example: "  terraform-variable-generator",
		Version: generatorVersion,
		Run:     runGenerator,
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func runGenerator(cmd *cobra.Command, args []string) {
	fmt.Println("Muho")
}
