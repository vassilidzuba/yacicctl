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
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var Cmd = &cobra.Command{
	Use:   "get <project> [<branch>] file",
	Short: "get a filee",
	Long:  `get a file fom a projects`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			execute(cmd, args[0], "main", args[1])
		} else if len(args) == 3 {
			execute(cmd, args[0], args[1], args[2])
		} else {
			_ = cmd.Usage()
		}
	},
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func execute(cmd *cobra.Command, project string, branch string, file string) {

	username := viper.GetString("username")
	password := viper.GetString("password")
	host := viper.GetString("host")

	fmt.Println("project run called on", project, branch, "!")

	url := "http://" + host + "/yacic/project/get?project=" + project + "&file=" + file
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

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if resp.StatusCode != http.StatusOK {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		data, _ := io.ReadAll(resp.Body)
		e := ErrorResult{}

		err := json.Unmarshal(data, &e)
		if err != nil {
			log.Fatal(err)
		}

		log.Fatalf("status code error: %d %s\nmessage: %s", resp.StatusCode, resp.Status, e.Message)

		os.Exit(1)
	}
	
	show, _ := cmd.Flags().GetBool("show")

	// print the response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if show {
		f, err := os.Create(file)
		check(err)
		defer func() {
		        if err := f.Close(); err != nil {
		            panic(err)
		        }
		    }()
					
		_, err = f.Write(data)
		check(err)
		
		err = f.Close()
		check(err)
		
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", file).Run()
		check(err)
	} else {
		fmt.Println(string(data))
	}
}

func closeBody(body *io.ReadCloser) {
	err := (*body).Close()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	Cmd.Flags().StringP("format", "f", "nice", "format, can be 'raw' or 'nice' (default)")
	Cmd.Flags().BoolP("show", "s", false, "show in browser")
}
