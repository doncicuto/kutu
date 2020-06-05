// Copyright © 2020 Muultipla Devops <devops@muultipla.com>
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/muultipla/kutu/internal/config"

	"github.com/spf13/cobra"
)

var shortList bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Kutu's supported K8s tools",
	Long:  "List Kutu's supported K8s tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kutu can check and update the following binaries:")
		for key, info := range config.Binaries {
			if shortList {
				fmt.Println("- ", key)
			} else {
				fmt.Println("- ", key, info.Description)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&shortList, "short", "s", false, "Get shorter version for list values")
}
