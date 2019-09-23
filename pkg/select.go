package slct

import (
	"fmt"

	"github.com/AlecAivazis/survey"
)

// Tool TODO doc
type Tool struct {
	Name string
	Values []string
}

func (t Tool) getOption() []string {
	return 
}

func (t Tool) onlyOneSelect() {
	prompt := &survey.Select{
		Message: "Choisir une technologie pour le " + t.Name ,
		Options: t.getOption(), // []string
	}
	survey.AskOne(prompt, &t.Values)
}

func (t Tool) multiSelect() {
	prompt := &survey.MultiSelect{
		Message: "What days do you prefer:",
		Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	}
	survey.AskOne(prompt, &t.Values)
}

// Select TODO doc
func (t Tool) Select() {
	switch t.Name {
	case "bdd":
		t.multiSelect()
	case "backend":
		t.onlyOneSelect()
	case "frontend":
		t.onlyOneSelect()
	default:
		fmt.Errorf("%s non pris en charge", t.Name)
	}
}
