package main

import (
	"os"
	"testing"
)

func TestWasmGo_Print(t *testing.T) {
	w := WasmGo{Message: "Test Message"}
	// Print just outputs to stdout, so we just ensure it doesn't panic
	w.Print()
}

func TestWasmGo_Write(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "wasmgo-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	testFile := tmpDir + "/test.txt"
	testContent := "This is a test"

	w := WasmGo{Message: "Test"}
	err = w.Write(testFile, testContent)
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}

	// Verify the file was created
	_, err = os.Stat(testFile)
	if err != nil {
		t.Errorf("File was not created: %v", err)
	}

	// Fix permissions to read the file (0x444 creates read-only files)
	err = os.Chmod(testFile, 0644)
	if err != nil {
		t.Fatalf("Failed to fix file permissions: %v", err)
	}

	// Verify the file has correct content
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}

	if string(data) != testContent {
		t.Errorf("Expected content %q, got %q", testContent, string(data))
	}
}

func TestWasmGo_Copy(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "wasmgo-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	sourceFile := tmpDir + "/source.txt"
	destFile := tmpDir + "/dest.txt"
	testContent := "Content to copy"

	// Create source file
	err = os.WriteFile(sourceFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	w := WasmGo{Message: "Test"}
	err = w.Copy(sourceFile, destFile)
	if err != nil {
		t.Errorf("Copy failed: %v", err)
	}

	// Verify the destination file was created
	_, err = os.Stat(destFile)
	if err != nil {
		t.Errorf("Destination file was not created: %v", err)
	}

	// Fix permissions to read the file (0x444 creates read-only files)
	err = os.Chmod(destFile, 0644)
	if err != nil {
		t.Fatalf("Failed to fix file permissions: %v", err)
	}

	// Verify the destination file has the same content
	data, err := os.ReadFile(destFile)
	if err != nil {
		t.Errorf("Failed to read destination file: %v", err)
	}

	if string(data) != testContent {
		t.Errorf("Expected content %q, got %q", testContent, string(data))
	}
}

func TestWasmGo_Copy_NonExistentSource(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "wasmgo-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	w := WasmGo{Message: "Test"}
	err = w.Copy("/non/existent/file.txt", tmpDir+"/dest.txt")
	if err == nil {
		t.Error("Expected error when copying non-existent file, got nil")
	}
}
