package aws

// Variables
var (
	debug bool

	masterCount int
	workerCount int
)

// Execute will run main logic
func Build(debug bool, masterCount int, workerCount int) {

	debug = debug
	masterCount = masterCount
	workerCount = workerCount
}