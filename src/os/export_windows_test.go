// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Export for testing.

var (
	AddExtendedPrefix  = addExtendedPrefix
	NewConsoleFile     = newConsoleFile
	CommandLineToArgv  = commandLineToArgv
	AllowReadDirFileID = &allowReadDirFileID
	SplitPath          = splitPath
)
