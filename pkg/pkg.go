package pkg

import "../utils"

// Build will run main logic
func Build() {

	if utils.createOptions.provider == "aws" {
		aws.Build(params)
	}
}