// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package filepath_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	glFilepath "github.com/svengreb/golib/pkg/io/fs/filepath"
)

func TestIsSubDir(t *testing.T) {
	dir := t.TempDir()
	testCases := []struct {
		parentPath string
		subPath    string
		isSubDir   bool
		isSymlink  bool
		isInvalid  bool
	}{
		{dir, filepath.Join(dir, "validSubDir"), true, false, false},
		{dir, ".", true, false, false},
		{dir, t.TempDir(), false, false, false},
		{dir, filepath.Join(dir, ".."), false, false, false},
		{dir, filepath.Join(t.TempDir(), "symlink"), false, true, false},
		{dir, filepath.Join(dir, "..", filepath.Base(dir), filepath.Base(t.TempDir())), true, true, false},
	}

	for _, tc := range testCases {
		if tc.isSymlink && tc.isSubDir {
			if err := os.Symlink(tc.parentPath, tc.subPath); err != nil {
				assert.Fail(t, "failed to create symlink", "source: %q\ndest: %q\nerror: %v", tc.parentPath, tc.subPath, err)
			}
		}

		isSubDir, err := glFilepath.IsSubDir(tc.parentPath, tc.subPath, tc.isSymlink)

		assert.Equal(t, tc.isSubDir, isSubDir)
		assert.NoError(t, err)
	}
}

func TestIsSubDir_FailWithRegularFileAsDirs(t *testing.T) {
	dir := t.TempDir()
	file, err := ioutil.TempFile(dir, "*")
	if err != nil {
		assert.Fail(t, "failed to create file", "error: %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			assert.Fail(t, "failed to close file", "file: %q\nerror: %v", file.Name(), closeErr)
		}
	}()

	isSubDirParentDirFile, err := glFilepath.IsSubDir(file.Name(), dir, true)
	assert.False(t, isSubDirParentDirFile)
	assert.Error(t, err)

	isSubDirSubDirFile, err := glFilepath.IsSubDir(dir, file.Name(), true)
	assert.False(t, isSubDirSubDirFile)
	assert.Error(t, err)
}

func TestIsSubDir_FailWithRelativeParentDir(t *testing.T) {
	dir := t.TempDir()
	isSubDir, err := glFilepath.IsSubDir(filepath.Base(dir), dir, false)

	assert.False(t, isSubDir)
	assert.Error(t, err)
}
