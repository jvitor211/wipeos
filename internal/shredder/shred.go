package shredder

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// WipeOptions contains configuration for the wiping operation
type WipeOptions struct {
	Recursive bool
	Passes    int
	Force     bool
	DryRun    bool
}

// WipeResult represents the result of wiping a single file
type WipeResult struct {
	Path    string
	Success bool
	Error   error
	Size    int64
}

// Shredder handles secure file deletion
type Shredder struct {
	logger zerolog.Logger
}

// New creates a new Shredder instance
func New() *Shredder {
	return &Shredder{
		logger: log.With().Str("component", "shredder").Logger(),
	}
}

// WipeFiles securely wipes multiple files
func (s *Shredder) WipeFiles(paths []string, options WipeOptions) []WipeResult {
	var results []WipeResult
	
	for _, path := range paths {
		if options.Recursive {
			if info, err := os.Stat(path); err == nil && info.IsDir() {
				dirResults := s.wipeDirectory(path, options)
				results = append(results, dirResults...)
				continue
			}
		}
		
		result := s.wipeFile(path, options)
		results = append(results, result)
	}
	
	return results
}

// wipeFile securely wipes a single file
func (s *Shredder) wipeFile(path string, options WipeOptions) WipeResult {
	info, err := os.Stat(path)
	if err != nil {
		return WipeResult{Path: path, Success: false, Error: err}
	}
	
	if info.IsDir() {
		return WipeResult{Path: path, Success: false, Error: fmt.Errorf("is a directory, use --recursive flag")}
	}
	
	if options.DryRun {
		s.logger.Info().Str("file", path).Msg("would wipe file (dry run)")
		return WipeResult{Path: path, Success: true, Size: info.Size()}
	}
	
	// Perform overwrite passes
	if err := s.overwriteFile(path, options.Passes); err != nil {
		return WipeResult{Path: path, Success: false, Error: err}
	}
	
	// Remove the file
	if err := os.Remove(path); err != nil {
		return WipeResult{Path: path, Success: false, Error: err}
	}
	
	s.logger.Info().Str("file", path).Int64("size", info.Size()).Msg("file wiped successfully")
	return WipeResult{Path: path, Success: true, Size: info.Size()}
}

// wipeDirectory recursively wipes all files in a directory
func (s *Shredder) wipeDirectory(dirPath string, options WipeOptions) []WipeResult {
	var results []WipeResult
	
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			results = append(results, WipeResult{Path: path, Success: false, Error: err})
			return nil
		}
		
		if !info.IsDir() {
			result := s.wipeFile(path, options)
			results = append(results, result)
		}
		
		return nil
	})
	
	if err != nil {
		results = append(results, WipeResult{Path: dirPath, Success: false, Error: err})
	}
	
	// Remove empty directories
	if !options.DryRun {
		os.RemoveAll(dirPath)
	}
	
	return results
}

// overwriteFile performs multiple overwrite passes on a file
func (s *Shredder) overwriteFile(path string, passes int) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()
	
	info, err := file.Stat()
	if err != nil {
		return err
	}
	
	size := info.Size()
	
	for pass := 0; pass < passes; pass++ {
		if err := s.performPass(file, size, pass); err != nil {
			return fmt.Errorf("pass %d failed: %w", pass+1, err)
		}
		
		// Sync to ensure data is written to disk
		if err := file.Sync(); err != nil {
			return err
		}
	}
	
	return nil
}

// performPass performs a single overwrite pass with different patterns
func (s *Shredder) performPass(file *os.File, size int64, passNum int) error {
	// Seek to beginning
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	
	var pattern []byte
	
	switch passNum % 4 {
	case 0:
		// Random data
		pattern = make([]byte, 4096)
		if _, err := rand.Read(pattern); err != nil {
			return err
		}
	case 1:
		// All zeros
		pattern = make([]byte, 4096)
	case 2:
		// All ones (0xFF)
		pattern = make([]byte, 4096)
		for i := range pattern {
			pattern[i] = 0xFF
		}
	case 3:
		// Alternating pattern (0xAA)
		pattern = make([]byte, 4096)
		for i := range pattern {
			pattern[i] = 0xAA
		}
	}
	
	written := int64(0)
	for written < size {
		remaining := size - written
		writeSize := int64(len(pattern))
		if remaining < writeSize {
			writeSize = remaining
		}
		
		n, err := file.Write(pattern[:writeSize])
		if err != nil {
			return err
		}
		written += int64(n)
	}
	
	return nil
}

// WipeBrowserData wipes browser cache, history, and temporary files
func (s *Shredder) WipeBrowserData(options WipeOptions) error {
	browserPaths := s.getBrowserPaths()
	
	var allPaths []string
	for _, paths := range browserPaths {
		allPaths = append(allPaths, paths...)
	}
	
	results := s.WipeFiles(allPaths, options)
	
	failed := 0
	for _, result := range results {
		if !result.Success {
			failed++
			s.logger.Error().Str("path", result.Path).Err(result.Error).Msg("failed to wipe browser data")
		}
	}
	
	if failed > 0 {
		return fmt.Errorf("failed to wipe %d browser data files", failed)
	}
	
	return nil
}

// WipeSystemTemp wipes system temporary files
func (s *Shredder) WipeSystemTemp(options WipeOptions) error {
	tempPaths := s.getSystemTempPaths()
	
	results := s.WipeFiles(tempPaths, options)
	
	failed := 0
	for _, result := range results {
		if !result.Success {
			failed++
			s.logger.Error().Str("path", result.Path).Err(result.Error).Msg("failed to wipe temp file")
		}
	}
	
	if failed > 0 {
		return fmt.Errorf("failed to wipe %d temporary files", failed)
	}
	
	return nil
}

// getBrowserPaths returns paths to browser data based on OS
func (s *Shredder) getBrowserPaths() map[string][]string {
	browsers := make(map[string][]string)
	
	switch runtime.GOOS {
	case "windows":
		userProfile := os.Getenv("USERPROFILE")
		browsers["chrome"] = []string{
			filepath.Join(userProfile, "AppData", "Local", "Google", "Chrome", "User Data", "Default", "History"),
			filepath.Join(userProfile, "AppData", "Local", "Google", "Chrome", "User Data", "Default", "Cache"),
		}
		browsers["firefox"] = []string{
			filepath.Join(userProfile, "AppData", "Roaming", "Mozilla", "Firefox", "Profiles"),
		}
		browsers["edge"] = []string{
			filepath.Join(userProfile, "AppData", "Local", "Microsoft", "Edge", "User Data", "Default", "History"),
			filepath.Join(userProfile, "AppData", "Local", "Microsoft", "Edge", "User Data", "Default", "Cache"),
		}
	case "darwin":
		homeDir, _ := os.UserHomeDir()
		browsers["chrome"] = []string{
			filepath.Join(homeDir, "Library", "Application Support", "Google", "Chrome", "Default", "History"),
			filepath.Join(homeDir, "Library", "Caches", "Google", "Chrome"),
		}
		browsers["firefox"] = []string{
			filepath.Join(homeDir, "Library", "Application Support", "Firefox", "Profiles"),
		}
		browsers["safari"] = []string{
			filepath.Join(homeDir, "Library", "Safari", "History.db"),
			filepath.Join(homeDir, "Library", "Caches", "com.apple.Safari"),
		}
	default: // Linux
		homeDir, _ := os.UserHomeDir()
		browsers["chrome"] = []string{
			filepath.Join(homeDir, ".config", "google-chrome", "Default", "History"),
			filepath.Join(homeDir, ".cache", "google-chrome"),
		}
		browsers["firefox"] = []string{
			filepath.Join(homeDir, ".mozilla", "firefox"),
		}
	}
	
	return browsers
}

// getSystemTempPaths returns paths to system temporary directories
func (s *Shredder) getSystemTempPaths() []string {
	var paths []string
	
	switch runtime.GOOS {
	case "windows":
		paths = []string{
			os.Getenv("TEMP"),
			os.Getenv("TMP"),
			filepath.Join(os.Getenv("WINDIR"), "Temp"),
		}
	case "darwin", "linux":
		paths = []string{
			"/tmp",
			"/var/tmp",
		}
		
		if homeDir, err := os.UserHomeDir(); err == nil {
			paths = append(paths, filepath.Join(homeDir, ".cache"))
		}
	}
	
	// Filter out non-existent paths
	var validPaths []string
	for _, path := range paths {
		if path != "" {
			if _, err := os.Stat(path); err == nil {
				validPaths = append(validPaths, path)
			}
		}
	}
	
	return validPaths
} 