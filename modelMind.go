package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type MIND struct {
	Cells map[string]*Cell `json:"cells"`
}

// Initialize a new MIND
func NewMIND() *MIND {
	return &MIND{
		Cells: make(map[string]*Cell),
	}
}

func (m *MIND) AddCell(point [2]int) error {
	return m.editContent(time.Now().Format(time.RFC3339), "", &point)
}

func (m *MIND) UpdateCell(key string) error {
	if text, exists := m.Cells[key]; exists {
		return m.editContent(key, text.Content, nil)
	}
	return fmt.Errorf("text with key '%s' not found", key)
}

// editText handles the editing of a text by key
func (m *MIND) editContent(key string, existingContent string, point *[2]int) error {
	// Create a temporary file to store the input text
	tmpfile, err := os.CreateTemp("", "temp*.txt")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up after use

	// Write existing content to the temporary file if available
	if existingContent != "" {
		if err := os.WriteFile(tmpfile.Name(), []byte(existingContent), 0644); err != nil {
			return fmt.Errorf("failed to write existing content to temporary file: %v", err)
		}
	}

	// Detect the terminal type using $TERM
	term := os.Getenv("TERM")
	cmd := prepareEditorCommand(term, tmpfile.Name())

	// Start the editor process
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to open editor in new terminal: %v", err)
	}

	// Wait for the Vim process to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("editor did not close properly: %v", err)
	}

	// Read the content from the temporary file after editing
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		return fmt.Errorf("failed to read from temporary file: %v", err)
	}

	if existingContent == "" {
		m.Cells[key] = &Cell{
			Content:  strings.TrimSpace(string(content)),
			Position: *point,
			Status:   CellStatusActive,
		}
		fmt.Println("Text added successfully!")
	} else {
        m.Cells[key].Content = strings.TrimSpace(string(content))
		fmt.Println("Text updated successfully!")
	}
	saveData()
	return nil
}

// prepareEditorCommand prepares the command to open the editor based on $TERM
func prepareEditorCommand(term string, filePath string) *exec.Cmd {
	var cmd *exec.Cmd
	switch term {
	case "xterm", "xterm-256color", "screen", "st", "konsole":
		cmd = exec.Command(term, "-e", "vim", filePath)
	case "gnome-terminal":
		cmd = exec.Command("gnome-terminal", "--", "vim", filePath)
	default:
		cmd = exec.Command("xterm", "-e", "vim", filePath)
	}
	return cmd
}
