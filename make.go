// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

//go:build tools

package main

import (
	"os"

	"github.com/magefile/mage/mage"
)

// Allows to run the project tasks without installing the mage binary.
// See https://magefile.org/zeroinstall for more details.
func main() { os.Exit(mage.Main()) }
