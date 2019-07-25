package cmd

import (
	"fmt"
	"os"

	"github.com/identixone/identixone-go/utils"
	"github.com/spf13/cobra"
)

var (
	outputPath string
	limit      int
	offset     int
	idxid      string
	faceSize   int
	sourceName string
	token      string
	debug      bool
)
var rootCmd = &cobra.Command{
	Use:   "identixone",
	Short: "Identix.one is a real-time cloud-based facial recognition platform for businesses.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if token != "" {
			ifErrorExit(os.Setenv("IDENTIXONE_TOKEN", token))
		}
		if debug {
			ifErrorExit(os.Setenv("IDENTIXONE_DEBUG", fmt.Sprintf("%v", debug)))
			utils.Warn().Msgf("%v", os.Environ())
		}
	},
	Long: `


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

						https://identix.one
`,
}

func ifErrorExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printAndExit(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func Execute() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug cli and client")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "identix.one API token")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "path to file for writing output result")
	rootCmd.PersistentFlags().IntVar(&limit, "limit", 20, "the number of output items, maximum 1000 entries per request")
	rootCmd.PersistentFlags().IntVar(&offset, "offset", 0, "a sequential number of an output item, to return a sampling after this one")

	rootCmd.AddCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
