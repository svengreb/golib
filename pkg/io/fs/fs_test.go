// Copyright (c) 2020-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package fs_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/svengreb/golib/pkg/io/fs"
)

func TestDirExists(t *testing.T) {
	dir := t.TempDir()
	exists, err := fs.DirExists(dir)

	assert.True(t, exists)
	assert.NoError(t, err)
}

func TestDirExists_FailWithNonExistingDir(t *testing.T) {
	dir := t.TempDir()
	nonExistingDir := filepath.Join(dir, "non-existing-path")
	exists, err := fs.DirExists(nonExistingDir)

	assert.False(t, exists)
	assert.NoError(t, err)
}

func TestDirExists_FailWithRegularFile(t *testing.T) {
	dir := t.TempDir()
	file, err := ioutil.TempFile(dir, "golib")
	if err != nil {
		assert.Fail(t, "failed to create file", "error: %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			assert.Fail(t, "failed to close file", "file: %q\nerror: %v", file.Name(), closeErr)
		}
	}()

	exists, err := fs.DirExists(file.Name())

	assert.False(t, exists)
	assert.Error(t, err)
}

func TestFileExists(t *testing.T) {
	dir := t.TempDir()
	file, err := ioutil.TempFile(dir, "golib")
	if err != nil {
		assert.Fail(t, "failed to create file", "error: %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			assert.Fail(t, "failed to close file", "file: %q\nerror: %v", file.Name(), closeErr)
		}
	}()

	exists, err := fs.FileExists(file.Name())

	assert.True(t, exists)
	assert.NoError(t, err)
}

func TestFileExists_FailWithNonExistingFile(t *testing.T) {
	dir := t.TempDir()
	nonExistingFile := filepath.Join(dir, "non-existing-path")
	exists, err := fs.FileExists(nonExistingFile)

	assert.False(t, exists)
	assert.NoError(t, err)
}

func TestIsFileWritable(t *testing.T) {
	dir := t.TempDir()
	file, err := ioutil.TempFile(dir, "golib")
	if err != nil {
		assert.Fail(t, "failed to create file", "error: %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			assert.Fail(t, "failed to close file", "file: %q\nerror: %v", file.Name(), closeErr)
		}
	}()

	writable, err := fs.IsFileWritable(file.Name())

	assert.True(t, writable)
	assert.NoError(t, err)
}

func TestIsFileWritable_FailWithReadOnlyFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "read-only_file")
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0)
	if err != nil {
		assert.Fail(t, "failed to create file in read-only mode", "file: %q\nerror: %v", path, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			assert.Fail(t, "failed to close read-only file", "file: %q\nerror: %v", file.Name(), closeErr)
		}
	}()

	writable, err := fs.IsFileWritable(file.Name())

	assert.False(t, writable)
	assert.Error(t, err)
}

func TestIsSymlink(t *testing.T) {
	dir := t.TempDir()
	t.TempDir()
	src, err := ioutil.TempFile(dir, "source")
	if err != nil {
		assert.Fail(t, "failed to create source file", "file: %q\nerror:%v", src, err)
	}
	defer func() {
		if closeErr := src.Close(); closeErr != nil {
			assert.Fail(t, "failed to close file", "file: %q\nerror: %v", src.Name(), closeErr)
		}
	}()

	dest := filepath.Join(dir, "symlink")
	if err = os.Symlink(src.Name(), dest); err != nil {
		assert.Fail(t, "failed to create symbolic link", "src: %q\ndest: %q\nerror: %v", src, dest, err)
	}

	isSymlink, err := fs.IsSymlink(dest)

	assert.True(t, isSymlink)
	assert.NoError(t, err)
}

func TestIsSymlink_FailWithInvalidPath(t *testing.T) {
	isSymlink, err := fs.IsSymlink("")

	assert.False(t, isSymlink)
	assert.Error(t, err)
}

func TestIsSymlink_FailWithNonSymlink(t *testing.T) {
	dir := t.TempDir()
	src, err := ioutil.TempFile(dir, "no_symlink")
	if err != nil {
		assert.Fail(t, "failed to create source file", "file: %q\nerror: %v", src, err)
	}
	defer func() {
		if closeErr := src.Close(); closeErr != nil {
			assert.Fail(t, "failed to close file", "file: %q\nerror: %v", src.Name(), closeErr)
		}
	}()

	isSymlink, err := fs.IsSymlink(src.Name())

	assert.False(t, isSymlink)
	assert.NoError(t, err)
}

func TestRegularFileExists(t *testing.T) {
	dir := t.TempDir()
	file, err := ioutil.TempFile(dir, "file")
	if err != nil {
		assert.Fail(t, "failed to create file", "error: %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			assert.Fail(t, "failed to close file", "file: %q\nerror: %v", file.Name(), closeErr)
		}
	}()

	exists, err := fs.RegularFileExists(file.Name())

	assert.True(t, exists)
	assert.NoError(t, err)
}

func TestRegularFileExists_FailWithDir(t *testing.T) {
	dir := t.TempDir()

	exists, err := fs.RegularFileExists(dir)

	assert.False(t, exists)
	assert.NoError(t, err)
}

func TestRegularFileExists_FailWithInvalidPath(t *testing.T) {
	exists, err := fs.RegularFileExists("")

	assert.False(t, exists)
	assert.NoError(t, err)
}
