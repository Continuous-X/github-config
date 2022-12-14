/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"github-config/cmd"
	"github-config/pkg/versions"
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

var (
	sha1ver   string
	buildTime string
	version   string
)

func main() {

	// create logging handler
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger.Enabled(slog.LevelDebug)
	// set logger as default
	slog.SetDefault(logger)

	versions.CommitFromGit = sha1ver
	versions.VersionFromGit = version
	versionArray := strings.Split(version, ".")
	if len(version) > 0 {
		versions.MajorFromGit = versionArray[0]
	}
	if len(version) > 1 {
		versions.MinorFromGit = versionArray[1]
	}
	versions.BuildDate = buildTime

	slog.Info("start app", "name", "ghc", "version", version)

	cmd.Execute()
}
