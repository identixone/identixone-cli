// Copyright © 2019 Maks Balashov <maksbalashov@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/identixone/identixone-go/api/client"
	"github.com/identixone/identixone-go/api/common"
	"github.com/identixone/identixone-go/utils"
	"github.com/spf13/cobra"
)

const (
	sourcesList   = "list"
	sourcesDelete = "delete"
	sourcersGet   = "get"
	sourcesUpdate = "update"
	sourcesCreate = "create"
)

var (
	q        string
	sourceId int
)

//func get

// serveCmd represents the serve command
var sourcesCmd = &cobra.Command{
	Use:   "sources",
	Short: "sources API",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("only one argument supported at time")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(args)

		switch args[0] {
		case sourcesList:
			query := common.NewSearchPaginationQuery(q, limit, offset)
			c, err := client.NewClient()
			ifErrorExit(err)
			sources, err := c.Sources().List(query)
			ifErrorExit(err)
			if outputPath != "" {
				preOut, err := utils.GetPretty(sources)
				ifErrorExit(err)
				err = utils.WriteToFile(outputPath, preOut)
				ifErrorExit(err)
			} else {
				ifErrorExit(utils.PrettyPrint(sources))
			}
		case sourcesDelete:
			if sourceId == 0 {
				printAndExit("source id is required")
			}
			c, err := client.NewClient()
			ifErrorExit(err)
			ifErrorExit(c.Sources().Delete(sourceId))
			fmt.Printf("source %d successfuly deleted", sourceId)
			fmt.Println()
		case sourcersGet:
			if sourceId == 0 {
				printAndExit("source id is required")
			}
			c, err := client.NewClient()
			ifErrorExit(err)
			source, err := c.Sources().Get(sourceId)
			ifErrorExit(err)
			ifErrorExit(utils.PrettyPrint(source))
		}
	},
}

func init() {
	sourcesCmd.Flags().StringVarP(&q, "search", "s", "", "filtering of a source sourcesList by partly or fully specified name")
	sourcesCmd.Flags().IntVar(&sourceId, "id", 0, "source id")
	rootCmd.AddCommand(sourcesCmd)
}
