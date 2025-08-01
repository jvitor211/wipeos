package shredder

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShredder_WipeFile(t *testing.T) {
	// Create a temporary file for testing
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	
	content := "This is sensitive test data that should be wiped securely"
	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)
	
	// Test wiping the file
	shredder := New()
	options := WipeOptions{
		Recursive: false,
		Passes:    3,
		Force:     true,
		DryRun:    false,
	}
	
	result := shredder.wipeFile(testFile, options)
	
	assert.True(t, result.Success)
	assert.NoError(t, result.Error)
	assert.Equal(t, testFile, result.Path)
	assert.Equal(t, int64(len(content)), result.Size)
	
	// Verify file is deleted
	_, err = os.Stat(testFile)
	assert.True(t, os.IsNotExist(err))
}

func TestShredder_WipeFile_DryRun(t *testing.T) {
	// Create a temporary file for testing
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	
	content := "This file should not be deleted in dry run"
	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)
	
	shredder := New()
	options := WipeOptions{
		Recursive: false,
		Passes:    3,
		Force:     true,
		DryRun:    true,
	}
	
	result := shredder.wipeFile(testFile, options)
	
	assert.True(t, result.Success)
	assert.NoError(t, result.Error)
	assert.Equal(t, testFile, result.Path)
	
	// Verify file still exists
	_, err = os.Stat(testFile)
	assert.NoError(t, err)
}

func TestShredder_WipeDirectory(t *testing.T) {
	// Create a temporary directory with files
	tmpDir := t.TempDir()
	testDir := filepath.Join(tmpDir, "testdir")
	err := os.Mkdir(testDir, 0755)
	require.NoError(t, err)
	
	// Create test files
	file1 := filepath.Join(testDir, "file1.txt")
	file2 := filepath.Join(testDir, "file2.txt")
	
	err = os.WriteFile(file1, []byte("content1"), 0644)
	require.NoError(t, err)
	err = os.WriteFile(file2, []byte("content2"), 0644)
	require.NoError(t, err)
	
	shredder := New()
	options := WipeOptions{
		Recursive: true,
		Passes:    1,
		Force:     true,
		DryRun:    false,
	}
	
	results := shredder.wipeDirectory(testDir, options)
	
	assert.Len(t, results, 2)
	for _, result := range results {
		assert.True(t, result.Success)
		assert.NoError(t, result.Error)
	}
	
	// Verify directory is deleted
	_, err = os.Stat(testDir)
	assert.True(t, os.IsNotExist(err))
}

func TestShredder_WipeFiles_Multiple(t *testing.T) {
	tmpDir := t.TempDir()
	
	// Create multiple test files
	files := []string{
		filepath.Join(tmpDir, "file1.txt"),
		filepath.Join(tmpDir, "file2.txt"),
		filepath.Join(tmpDir, "file3.txt"),
	}
	
	for _, file := range files {
		err := os.WriteFile(file, []byte("test content"), 0644)
		require.NoError(t, err)
	}
	
	shredder := New()
	options := WipeOptions{
		Recursive: false,
		Passes:    1,
		Force:     true,
		DryRun:    false,
	}
	
	results := shredder.WipeFiles(files, options)
	
	assert.Len(t, results, 3)
	for _, result := range results {
		assert.True(t, result.Success)
		assert.NoError(t, result.Error)
	}
	
	// Verify all files are deleted
	for _, file := range files {
		_, err := os.Stat(file)
		assert.True(t, os.IsNotExist(err))
	}
}

func TestShredder_PerformPass(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	
	// Create a file with known content
	originalContent := "Original content that should be overwritten"
	err := os.WriteFile(testFile, []byte(originalContent), 0644)
	require.NoError(t, err)
	
	// Open file for overwriting
	file, err := os.OpenFile(testFile, os.O_WRONLY, 0)
	require.NoError(t, err)
	defer file.Close()
	
	shredder := New()
	
	// Perform an overwrite pass
	err = shredder.performPass(file, int64(len(originalContent)), 0)
	assert.NoError(t, err)
	
	// Read the file back
	newContent, err := os.ReadFile(testFile)
	require.NoError(t, err)
	
	// Content should be different (overwritten)
	assert.NotEqual(t, originalContent, string(newContent))
	assert.Equal(t, len(originalContent), len(newContent))
}

func TestGetSystemTempPaths(t *testing.T) {
	shredder := New()
	paths := shredder.getSystemTempPaths()
	
	assert.NotEmpty(t, paths)
	
	// At least one path should exist
	found := false
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			found = true
			break
		}
	}
	assert.True(t, found, "At least one temp path should exist")
}

func TestGetBrowserPaths(t *testing.T) {
	shredder := New()
	browserPaths := shredder.getBrowserPaths()
	
	assert.NotEmpty(t, browserPaths)
	
	// Should have paths for at least one browser
	totalPaths := 0
	for _, paths := range browserPaths {
		totalPaths += len(paths)
	}
	assert.Greater(t, totalPaths, 0)
} 