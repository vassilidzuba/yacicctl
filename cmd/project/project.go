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

package project

import (
	"fmt"

	"github.com/spf13/cobra"
	
	run "github.com/vassilidzuba/yacicctl/cmd/project/run"
	list "github.com/vassilidzuba/yacicctl/cmd/project/list"
	get "github.com/vassilidzuba/yacicctl/cmd/project/get"

)

// projectCmd represents the project command
var Cmd = &cobra.Command{
	Use:   "project",
	Short: "commands related toprojectsd",
	Long: `a project specifies agit repo and pipeline(s) that can run on it.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command project needs a subcommand: ")
	},
}

func init() {
	Cmd.AddCommand(run.Cmd, list.Cmd, get.Cmd);
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
