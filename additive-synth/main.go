package main
// As is, this is populating basic files. must be expanded on
import (
	"additive-synth/gui"
	"fmt"
)

// Add load config logging e.g. pass/fail
func main() {
	config, err := util.LoadConfig('config.json')
	if err != nil {
		fmt.Println("Something wrong loading config: ", err)
		// Continue or not based on the error. figure it out later.
		os.Exit(1)
	}
	// Proceed with 'config' shit in main if necessary
	fmt.Println("Starting Additive Synth...")
	gui.StartApp() // Placeholder for when we make a launcher for Fyne
}
