// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package filepath provides utility functions for manipulating filename paths for the target operating system-defined
// file paths, using either forward slashes or backslashes.
// It extends the "path/filepath" Go standard library package with more utilities. Please note that some functions
// interact with the underlying filesystem through on-disk operations.
package filepath

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/svengreb/golib/pkg/io/fs"
)

// IsSubDir checks if a path is a subdirectory of another path.
// The parent path must be an absolute path while the sub-path can either be relative to the parent path or an absolute
// path.
//
// Note that this function interacts with the underlying filesystem through on-disk operations when the "evalSymlinks"
// parameter is "true" in order to achieve required file information!
//nolint:wrapcheck // Returning standard library errors is perfectly fine.
func IsSubDir(parentPath, subPath string, evalSymlinks bool) (bool, error) {
	if !filepath.IsAbs(parentPath) {
		return false, errors.New("parent path is not an absolute path")
	}

	pp := parentPath
	sp := subPath
	if evalSymlinks {
		ppExists, fsErr := fs.DirExists(pp)
		if fsErr != nil {
			return false, fsErr
		}
		if ppExists {
			pp, fsErr = filepath.EvalSymlinks(pp)
			if fsErr != nil {
				return false, fsErr
			}
		}

		spExists, fsErr := fs.DirExists(sp)
		if fsErr != nil {
			return false, fsErr
		}
		if spExists {
			sp, fsErr = filepath.EvalSymlinks(sp)
			if fsErr != nil {
				return false, fsErr
			}
		}
	}

	if !filepath.IsAbs(sp) {
		sp = filepath.Join(pp, sp)
	}

	rel, relErr := filepath.Rel(pp, sp)
	if relErr != nil {
		return false, relErr
	}

	if !strings.HasPrefix(rel, fmt.Sprintf("..%s", string(filepath.Separator))) && rel != ".." {
		return true, nil
	}

	return false, nil
}
