/**
   Copyright 2025 Vassili Dzuba

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.	
**/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	
	reload "github.com/vassilidzuba/yacicctl/cmd/config/reload"

)

// configCmd represents the config command
var Cmd = &cobra.Command{
	Use:   "config",
	Short: "commands related to the configurations",
	Long: `commands related tothe various configurations (pipelines, actions,etc)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("need subcomand")
	},
}

func init() {
	Cmd.AddCommand(reload.Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
