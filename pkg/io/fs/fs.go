// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package fs provides I/O utility functions for filesystem interactions.
// Please note that it makes use of the underlying filesystem and only serves as additional utility for the "os" Go
// standard library package.
// For more advanced and extended features see packages like https://github.com/spf13/afero instead.
package fs

import (
	"fmt"
	"os"
)

// DirExists checks if a directory exists.
// If an error occurs, "false" is returned along with the error.
//
// See
//
//   (1) https://en.wikipedia.org/wiki/Directory_(computing)
//   (2) https://en.wikipedia.org/wiki/Unix_file_types#Directory
//nolint:wrapcheck // Returning standard library errors is perfectly fine.
func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	if !info.IsDir() {
		return false, fmt.Errorf("%q is not a directory", path)
	}

	return false, err
}

// FileExists checks if a regular file or directory exists.
// If an error occurs, "false" is returned along with the error.
//
// See
//
//   (1) https://en.wikipedia.org/wiki/Computer_file
//   (2) https://en.wikipedia.org/wiki/Unix_file_types
//nolint:wrapcheck // Returning standard library errors is perfectly fine.
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// IsFileWritable checks if a file is writable.
// If an error occurs, "false" is returned along with the error.
//nolint:wrapcheck // Returning standard library errors is perfectly fine.
func IsFileWritable(path string) (bool, error) {
	_, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return false, err
	}

	return true, nil
}

// IsSymlink checks if a file is a symbolic link.
// If an error occurs, "false" is returned along with the error.
//
// See
//
//   (1) https://en.wikipedia.org/wiki/Symbolic_link
//   (2) https://en.wikipedia.org/wiki/Unix_file_types#Symbolic_link
//nolint:wrapcheck // Returning standard library errors is perfectly fine.
func IsSymlink(path string) (bool, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	return fi.Mode()&os.ModeSymlink == os.ModeSymlink, nil
}

// RegularFileExists checks if a regular file exists.
// If an error occurs, "false" is returned along with the error.
//
// See
//
//   (1) https://en.wikipedia.org/wiki/Unix_file_types#Regular_file
//   (2) https://en.wikipedia.org/wiki/Computer_file
//nolint:wrapcheck // Returning standard library errors is perfectly fine.
func RegularFileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil && info.Mode().IsRegular() {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
