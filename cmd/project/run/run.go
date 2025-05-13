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

package run

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var Cmd = &cobra.Command{
	Use:   "run <project> [<branch>]",
	Short: "run a pipeline",
	Long:  `run a pipeline for a projects`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			execute(cmd, args[0], "main")
		} else if len(args) == 2 {
			execute(cmd, args[0], args[1])
		} else {
			_ = cmd.Usage()
		}
	},
}

type Result struct {
	Project   string `json:"projectId"`
	Branch    string `json:"branchId"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
	Duration  int    `json:"duration"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func execute(cmd *cobra.Command, project string, branch string) {

	username := viper.GetString("username")
	password := viper.GetString("password")
	host := viper.GetString("host")

	fmt.Println("project run called on", project, branch, "!")

	url := "http://" + host + "/yacic/project/run?project=" + project
	if branch != "" {
		url = url + "&branch=" + branch
	}

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

	format, _ := cmd.Flags().GetString("format")

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if resp.StatusCode != http.StatusOK {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		data1, _ := io.ReadAll(resp.Body)
		e := ErrorResult{}

		err := json.Unmarshal(data1, &e)
		if err != nil {
			log.Fatal(err)
		}

		if format == "raw" {
			log.Fatalf("status code error: %d %s\nmessage: %s", resp.StatusCode, resp.Status, e.Message)
		} else if format == "nice" {
			pterm.DefaultBasicText.Println(
				pterm.LightCyan("HTTP status") + ": " + resp.Status + "\n" +
					pterm.LightCyan("project    ") + ": " + project + "\n" +
					pterm.LightCyan("branch     ") + ": " + branch + "\n" +
					pterm.LightCyan("message    ") + ": " + e.Message + "\n")
		}

		os.Exit(1)
	}

	// print the response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	switch format {

	case "raw":
		fmt.Println(string(data))

	case "nice":
		var r Result

		err := json.Unmarshal(data, &r)
		if err != nil {
			log.Fatal(err)
		}

		pterm.DefaultBasicText.Println(
			pterm.LightCyan("project   ") + ": " + r.Project + "\n" +
				pterm.LightCyan("branch    ") + ": " + r.Branch + "\n" +
				pterm.LightCyan("timestamp ") + ": " + r.Timestamp + "\n" +
				pterm.LightCyan("duration  ") + ": " + strconv.Itoa(r.Duration/1000) + "s\n" +
				pterm.LightCyan("status    ") + ": " + r.Status)

	default:
		log.Fatal("--format can be 'raw' or 'nice'")
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
