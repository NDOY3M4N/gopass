package main

import (
	"fmt"
	"os"

	"github.com/NDOY3M4N/gopass/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(tui.NewModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
