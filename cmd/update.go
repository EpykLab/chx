/*
Copyright © 2025 contact@epyklab.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/go-github/v58/github"
	"github.com/spf13/cobra"
)

// set all the params for the github repo where updates will be
const (
	repoOwner = "EpykLab"
	repoName  = "chx"
	appName   = "chx"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update chx to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Check for updates...")

		err := updateApp()
		if err != nil {
			fmt.Printf("Update failed: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

// updateApp handles the update process
func updateApp() error {
	// Step 1: Get the latest release from GitHub
	client := github.NewClient(nil)
	ctx := context.Background()
	release, _, err := client.Repositories.GetLatestRelease(ctx, repoOwner, repoName)
	if err != nil {
		return fmt.Errorf("failed to fetch latest release: %v", err)
	}

	latestversion := release.GetTagName()
	if latestversion == version {
		fmt.Println("Already up to date!")
		return nil
	} else {
		fmt.Println("new version is: ", latestversion)
	}

	// Step 2: Find the correct asset (binary or archive) for the current OS/arch
	assetName := getAssetName()
	var downloadURL string
	for _, asset := range release.Assets {
		if strings.Contains(asset.GetName(), assetName) {
			downloadURL = asset.GetBrowserDownloadURL()
			log.Println(downloadURL)
			break
		}
	}
	if downloadURL == "" {
		return fmt.Errorf("no suitable asset found for %s", assetName)
	}

	// Step 3: Download the new version
	tempFile := filepath.Join(os.TempDir(), assetName)
	err = downloadFile(downloadURL, tempFile)
	if err != nil {
		return fmt.Errorf("failed to download update: %v", err)
	}
	defer os.Remove(tempFile) // Clean up temp file

	// Step 4: Replace the current binary
	currentPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to locate current executable: %v", err)
	}

	// Handle extraction if it's an archive (e.g., .zip or .tar.gz)
	if strings.HasSuffix(assetName, ".zip") || strings.HasSuffix(assetName, ".tar.gz") {
		// Extract the tar.gz and find the binary
		err = extractTarGz(tempFile, currentPath)
		if err != nil {
			return fmt.Errorf("failed to extract tar.gz: %v", err)
		}
	} else {
		// Direct binary replacement (unchanged from previous code)
		err = replaceBinary(tempFile, currentPath)
		if err != nil {
			return fmt.Errorf("failed to replace binary: %v", err)
		}
	}

	return nil
}

// getAssetName determines the expected asset name based on OS and architecture
func getAssetName() string {
	osName := runtime.GOOS
	arch := runtime.GOARCH
	// Example: "chx-linux-amd64" or "chx-windows-amd64.exe"
	if osName == "windows" {
		return fmt.Sprintf("%s_%s.exe", osName, arch)
	}
	return fmt.Sprintf("%s_%s.tar.gz", osName, arch)
}

// downloadFile downloads a file from a URL to a local path
func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// replaceBinary replaces the current binary with the downloaded one
func replaceBinary(newPath, currentPath string) error {
	// On Unix-like systems, we can't directly overwrite a running binary
	// So we move it to a temp location first, then replace
	if runtime.GOOS != "windows" {
		tempPath := currentPath + ".old"
		err := os.Rename(currentPath, tempPath)
		if err != nil {
			return err
		}
		defer os.Remove(tempPath) // Clean up old binary
	}

	err := os.Rename(newPath, currentPath)
	if err != nil {
		return err
	}

	// Ensure the new binary is executable (Unix-like systems)
	if runtime.GOOS != "windows" {
		return os.Chmod(currentPath, 0755)
	}
	return nil
}

// extractTarGz extracts a .tar.gz file and replaces the current binary
func extractTarGz(tarGzPath, destPath string) error {
	// Open the .tar.gz file
	file, err := os.Open(tarGzPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Uncompress the gzip layer
	gz, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gz.Close()

	// Create a tar reader
	tarReader := tar.NewReader(gz)

	// Look for the binary in the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}

		// We assume the binary has the same name as the app (e.g., "chx")
		if header.Typeflag == tar.TypeReg && filepath.Base(header.Name) == appName {
			// Move current binary to a temp location (Unix-like systems)
			if runtime.GOOS != "windows" {
				tempPath := destPath + ".old"
				err = os.Rename(destPath, tempPath)
				if err != nil {
					return err
				}
				defer os.Remove(tempPath)
			}

			// Create the new binary file
			outFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer outFile.Close()

			// Copy the binary content from the tar
			_, err = io.Copy(outFile, tarReader)
			if err != nil {
				return err
			}

			// Set executable permissions (Unix-like systems)
			if runtime.GOOS != "windows" {
				err = os.Chmod(destPath, 0755)
				if err != nil {
					return err
				}
			}

			return nil // Successfully replaced the binary
		}
	}

	return fmt.Errorf("binary '%s' not found in archive", appName)
}
