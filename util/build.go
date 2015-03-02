// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Peter Mattis (petermattis@gmail.com)

package util

var (
	// These variables are initialized via the linker -X flag in the
	// top-level Makefile when compiling release binaries.
	buildSHA  string // Git SHA
	buildTag  string // Tag of this build (git describe)
	buildTime string // Build time in UTC (year/month/day hour:min:sec)
	buildDeps string // Git SHAs of dependencies
)

// BuildInfo ...
type BuildInfo struct {
	SHA  string
	Tag  string
	Time string
	Deps string
}

// GetBuildInfo ...
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		SHA:  buildSHA,
		Tag:  buildTag,
		Time: buildTime,
		Deps: buildDeps,
	}
}