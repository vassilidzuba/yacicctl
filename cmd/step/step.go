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

package step

import (
	"fmt"

	"github.com/spf13/cobra"
	
	list "github.com/vassilidzuba/yacicctl/cmd/step/list"
)

// stepCmd represents the step command
var Cmd = &cobra.Command{
	Use:   "step <subcommand>",
	Short: "commands related to steps",
	Long: `commands related to steps`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("step needs a subcommand")
		cmd.Usage()
	},
}

func init() {
	Cmd.AddCommand(list.Cmd)
}
