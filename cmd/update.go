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
	"os/exec"
	"strings"
	"sync"

	"github.com/gookit/color"
	"github.com/muultipla/kutu/internal/utils"

	"github.com/muultipla/kutu/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update K8s binaries",
	Long:  "Update K8s binaries",
	Run: func(cmd *cobra.Command, args []string) {

		// Check if binaries flag or configuration options
		// have been set so we only check those binaries
		flagBinaries := viper.GetString("binaries")
		var selectedBinaries []string
		if flagBinaries != "" {
			selectedBinaries = strings.Split(flagBinaries, ",")
		}

		var wg sync.WaitGroup

		for binary, info := range config.Binaries {
			if len(selectedBinaries) == 0 || utils.Contains(selectedBinaries, binary) {
				wg.Add(1)
				go update(binary, info, &wg)
			}
		}
		wg.Wait()
	},
}

func update(binary string, info config.BinaryConfig, wg *sync.WaitGroup) {
	defer wg.Done()

	red := color.FgRed.Render
	green := color.FgGreen.Render
	cyan := color.FgCyan.Render

	path, err := exec.LookPath(binary)
	if err == nil {
		version, err := utils.CurrentVersion(binary, info.VersionCommand)
		if err == nil {
			res, err := utils.CheckVersion(binary, info, version)
			if err == nil {
				if res.Outdated {
					updateError := utils.Update(binary, path, info.DownloadURL, res.Latest)
					if updateError == nil {
						fmt.Printf("%s has been %s. New version is %s.\n", cyan(binary), green("updated"), green(res.Latest))
					} else {
						fmt.Printf("Error: %s\n", red(updateError))
					}
				} else {
					fmt.Printf("%s is %s. No need to update it.\n", cyan(binary), green("up to date"))
				}
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
