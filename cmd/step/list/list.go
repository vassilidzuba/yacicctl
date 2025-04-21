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

package list

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var Cmd = &cobra.Command{
	Use:   "list <project> [<branch>] <timestamp>",
	Short: "list the steps",
	Long: `list the steps for a given build, identified by its project, branch and timestamp`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			execute(cmd, args[0], "", args[1])
		} else if len(args) == 3 {
			execute(cmd, args[0], args[1], args[2])
		} else {
			log.Println("Help")
		}
	},
}


func execute(cmd *cobra.Command, project string, branch string, timestamp string) {
	
	username := viper.GetString("username")
	password := viper.GetString("password")

	url := "http://localhost:8080/yacic/build/list?project=" + project
	if branch != "" {
		url = url + "&branch=" + branch
	}
	url = url + "&timestamp=" + timestamp
	
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
	defer resp.Body.Close()

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if resp.StatusCode != http.StatusOK {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		//panic("bad")
	}

	// print the response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}


func init() {
}
