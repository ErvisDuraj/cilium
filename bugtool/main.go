// Copyright 2017-2018 Authors of Cilium
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

package main

import (
	"fmt"
	"os"

	"github.com/cilium/cilium/bugtool/cmd"

	gops "github.com/google/gops/agent"
)

func main() {
	// Open socket for using gops to get stacktraces of the agent
	if err := gops.Listen(gops.Options{}); err != nil {
		fmt.Println("Unable to start gops")
		os.Exit(1)
	}
	if err := cmd.BugtoolRootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
