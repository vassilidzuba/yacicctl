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
	"log"

	//	"strconv"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var Cmd = &cobra.Command{
	Use:   "list",
	Short: "list the projects",
	Long:  `list the projects referenbces in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

type Branch struct {
	Branch string `json:"branchId"`
	Dir    string `json:"branchDir"`
}

type Result struct {
	Project  string   `json:"projectId"`
	Repo     string   `json:"repo"`
	Branches []Branch `json:"branches"`
}

func execute(cmd *cobra.Command) {

	username := viper.GetString("username")
	password := viper.GetString("password")
	host := viper.GetString("host")

	url := "http://" + host + "/yacic/project/list"

	log.Println("url:", url)

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer closeBody(&resp.Body)

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if resp.StatusCode != http.StatusOK {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		//panic("bad")
	}

	// print the response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	format, _ := cmd.Flags().GetString("format")

	switch format {

	case "raw":
		fmt.Println(string(data))

	case "nice":
		var r []Result

		err = json.Unmarshal(data, &r)
		if err != nil {
			log.Fatal(err)
		}

		//		if len(r) > 0 {
		//			pterm.DefaultBasicText.Println(
		//				pterm.LightCyan("project   ") + ": " + r[0].Project)
		//		}

		tab := [][]string{{"Project", "Repo", "Branch", "Branch dir"}}

		for _, e := range r {
			for _, e2 := range e.Branches {
				tab = append(tab, []string{e.Project, e.Repo, e2.Branch, e2.Dir})
				e.Project = ""
				e.Repo = ""
			}
		}

		err = pterm.DefaultTable.WithHasHeader().WithData(tab).Render()
		if err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatal("-format can be 'raw' or 'nice'")
	}

}

func closeBody(body *io.ReadCloser) {
	err := (*body).Close()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	Cmd.Flags().StringP("format", "f", "nice", "format, can be 'raw' or 'nice' (default)")
}
