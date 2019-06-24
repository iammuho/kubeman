package cmd

import (

	"fmt"
	"os"
	"log"
	"strconv"

	Utils "../utils"

	"github.com/spf13/cobra"
)

// Execute will run main logic
func Execute(generatorVersion string) {

	cmd := &cobra.Command{
		Use:     "generator",
		Short:   "CLI Tool for generating terraform templates for Kubernetes",
		Example: "kubeman --provider aws --master 3 --worker 3",
		Version: generatorVersion,
		Run:     runGenerator,
	}

	cmd.PersistentFlags().StringVar(&Utils.createOptions.provider, "provider", "aws", "Terraform provider")
	cmd.PersistentFlags().IntVar(&Utils.createOptions.masterCount, "master", 3, "Master Count")
	cmd.PersistentFlags().IntVar(&Utils.createOptions.workerCount, "worker", 3, "Worker Count")
	cmd.PersistentFlags().BoolVar(&Utils.createOptions.debug, "debug", false, "Debug Mode")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func runGenerator(cmd *cobra.Command, args []string) {
	log.Print("Kubeman started")

	if Utils.createOptions.debug {
		log.Print("Provider : "  + string(Utils.createOptions.provider))
		log.Print("Master Count : "  + strconv.Itoa(Utils.createOptions.masterCount))
		log.Print("Worker Count : "  + strconv.Itoa(Utils.createOptions.workerCount))
	}

}
