package main

import (
	"fmt"
	"os"

	"github.com/identixone/identixone-cli/cmd"
)

func init() {
	if os.Getenv("IDENTIXONE_TOKEN") == "" {
		fmt.Print(`

8888888     888                888                                          
  888       888                888                                          
  888       888                888                                          
  888   .d88888 .d88b. 88888b. 888888888888  888    .d88b. 88888b.  .d88b.  
  888  d88" 888d8P  Y8b888 "88b888   888 Y8bd8P'   d88""88b888 "88bd8P  Y8b 
  888  888  88888888888888  888888   888  X88K     888  888888  88888888888 
  888  Y88b 888Y8b.    888  888Y88b. 888.d8""8b.d8bY88..88P888  888Y8b.     
8888888 "Y88888 "Y8888 888  888 "Y888888888  888Y8P "Y88P" 888  888 "Y8888  
																			
									 d8b									
									 Y8P									
									 

IDENTIXONE_TOKEN environment not found!


Set it:
	export IDENTIXONE_TOKEN=<token>
Or run programm:
	env IDENTIXONE_TOKEN=<token> identixone ....

Do you not have Token? Get your free API token for development at https://identix.one
`)
		os.Exit(1)
	}
}

func main() {
	cmd.Execute()
}
