package aws

import (

	"html/template"
)

// Variables
var (
	debug bool

	masterCount int
	workerCount int

	templates []
)


// Build will run main logic
func Build(debug bool, masterCount int, workerCount int) {

	debug = debug
	masterCount = masterCount
	workerCount = workerCount

	getTemplates()
}

func getTemplates() {

	templates["main"], templates["error"] = template.ParseFiles("../template/aws/main.tf.kubeman")

}



