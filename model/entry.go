package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nore-dev/fman/entry"
)

type EntryModel struct {
	entry entry.Entry
}

func NewEntryModel() EntryModel {
	return EntryModel{}
}

func (model EntryModel) Init() tea.Cmd {
	return nil
}

func (model EntryModel) Update(msg tea.Msg) (EntryModel, tea.Cmd) {

	switch msg := msg.(type) {
	case entry.EntryMsg:
		model.entry = msg.Entry
	}
	return model, nil
}

func (model EntryModel) getFileInfo() string {
	str := strings.Builder{}

	str.WriteString(model.entry.Name)

	str.WriteByte('\n')

	if model.entry.IsDir {
		str.WriteString("Folder\n")
	} else {
		str.WriteString(model.entry.Extension)
		str.WriteString(" File\n")
	}

	str.WriteString("Modified ")
	str.WriteString(model.entry.ModifyTime)

	str.WriteByte('\n')

	str.WriteString("Changed ")
	str.WriteString(model.entry.ChangeTime)

	str.WriteByte('\n')

	str.WriteString("Accessed ")
	str.WriteString(model.entry.AccessTime)

	return str.String()
}

func (model EntryModel) View() string {

	return lipgloss.JoinVertical(lipgloss.Left,
		model.entry.Preview,
		model.getFileInfo())

}