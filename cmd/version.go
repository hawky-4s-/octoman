// Copyright © 2016 Christian Lipphardt <christian.lipphardt@gmail.com>
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

package cmd

import (
	"fmt"
	"github.com/hawky-4s-/octoman/helpers"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Octoman",
	RunE: func(cmd *cobra.Command, args []string) error {
		printOctomanVersion()
		return nil
	},
}

func printOctomanVersion() {
	if helpers.BuildDate == "" {
		setBuildDate() // set the build date from executable's mdate
	} else {
		formatBuildDate() // format the compile time
	}
	if helpers.CommitHash == "" {
		fmt.Printf("Octoman GitHub Manager v%s %s/%s BuildDate: %s\n", helpers.CurrentOctomanVersion, runtime.GOOS, runtime.GOARCH, helpers.BuildDate)
	} else {
		fmt.Printf("Octoman Static Site Generator v%s-%s %s/%s BuildDate: %s\n", helpers.CurrentOctomanVersion, strings.ToUpper(helpers.CommitHash), runtime.GOOS, runtime.GOARCH, helpers.BuildDate)
	}
}

// setBuildDate checks the ModTime of the Octoman executable and returns it as a
// formatted string.  This assumes that the executable name is Octoman, if it does
// not exist, an empty string will be returned.  This is only called if the
// helpers.BuildDate wasn't set during compile time.
//
// osext is used for cross-platform.
func setBuildDate() {
	fname, _ := osext.Executable()
	dir, err := filepath.Abs(filepath.Dir(fname))
	if err != nil {
		jww.ERROR.Println(err)
		return
	}
	fi, err := os.Lstat(filepath.Join(dir, filepath.Base(fname)))
	if err != nil {
		jww.ERROR.Println(err)
		return
	}
	t := fi.ModTime()
	helpers.BuildDate = t.Format(time.RFC3339)
}

// formatBuildDate formats the helpers.BuildDate according to the value in
// .Params.DateFormat, if it's set.
func formatBuildDate() {
	t, _ := time.Parse("2006-01-02T15:04:05-0700", helpers.BuildDate)
	helpers.BuildDate = t.Format(time.RFC3339)
}
