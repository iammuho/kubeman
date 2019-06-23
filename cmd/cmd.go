package cmd

import (

	"fmt"
	"os"
	"log"
	"strconv"

	"../pkg/aws"

	"github.com/spf13/cobra"
)

// Variables
var (
	generatorVersion string

	debug bool

	provider     string

	masterCount int
	workerCount int
)

// Execute will run main logic
func Execute(version string) {
	generatorVersion = version

	cmd := &cobra.Command{
		Use:     "generator",
		Short:   "CLI Tool for generating terraform templates for Kubernetes",
		Example: "kubeman --provider aws --master 3 --worker 3",
		Version: generatorVersion,
		Run:     runGenerator,
	}

	cmd.PersistentFlags().StringVar(&provider, "provider", "aws", "Terraform provider")
	cmd.PersistentFlags().IntVar(&masterCount, "master", 3, "Master Count")
	cmd.PersistentFlags().IntVar(&workerCount, "worker", 3, "Worker Count")
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "Debug Mode")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func runGenerator(cmd *cobra.Command, args []string) {
	log.Print("Kubeman started")

	if debug {
		log.Print("Provider : "  + string(provider))
		log.Print("Master Count : "  + strconv.Itoa(masterCount))
		log.Print("Worker Count : "  + strconv.Itoa(workerCount))
	}

	aws.Build(debug, masterCount, workerCount)
}
