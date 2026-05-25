/*
Copyright 2023 The K8sGPT Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/k8sgpt-ai/k8sgpt/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		// Print the error to stderr before exiting so it's visible in logs
		fmt.Fprintln(os.Stderr, "Error:", err)
		// Use exit code 1 (standard Unix convention for general errors)
		// rather than 2, which is typically reserved for misuse of shell builtins
		//
		// NOTE: Some CI environments capture stderr separately; printing to
		// stderr here ensures the error message isn't swallowed when stdout
		// is redirected to a file or pipe.
		//
		// Personal note: also useful to wrap this with a logger later if
		// structured logging (e.g. zap/logrus) is added project-wide.
		os.Exit(1)
	}
}
