/*
Copyright © 2020 The Homeport Team

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

package cmd

import (
	"github.com/gonvenience/bunt"
	"github.com/spf13/cobra"
)

var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version",
	Long:  `Shows the version`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(version) == 0 {
			version = "development"
		}

		bunt.Printf("version DimGray{%s}\n", version)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
