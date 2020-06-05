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
	"strings"

	"github.com/gookit/color"
	"github.com/muultipla/kutu/internal/config"
	"github.com/muultipla/kutu/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if new versions are available",
	Long:  "Check if new versions are available",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if binaries flag or configuration options
		// have been set so we only check those binaries
		flagBinaries := viper.GetString("binaries")
		var selectedBinaries []string
		if flagBinaries != "" {
			selectedBinaries = strings.Split(flagBinaries, ",")
		}

		red := color.FgRed.Render
		green := color.FgGreen.Render
		cyan := color.FgCyan.Render

		for binary, info := range config.Binaries {

			if len(selectedBinaries) == 0 || utils.Contains(selectedBinaries, binary) {
				if version, err := utils.CurrentVersion(binary, info.VersionCommand); err == nil {
					if res, err := utils.CheckVersion(binary, info, version); err == nil {
						if res.Outdated {
							fmt.Printf("%s is %s. Current version is %s. Latest version is %s\n", cyan(binary), red("outdated"), version, res.Current)
						} else {
							fmt.Printf("%s is %s. Latest version is %s\n", cyan(binary), green("up to date"), res.Current)
						}
					} else {
						fmt.Printf("%s.\n", red(err))
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
