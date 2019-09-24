package tool

import (
	"os"
	"io/ioutil"
	"fmt"

	"github.com/AlecAivazis/survey"
)

// Tool TODO doc
type Tool struct {
	Name string
	Values []string
}


func (t *Tool) getOption() []string {
	var options []string
	path := "../.templates/" + t.Name
	c, err := ioutil.ReadDir(path)
    if err != nil {
		fmt.Printf("impossible de trouver les templates")
		os.Exit(1)
	}

    for _, entry := range c {
		if entry.IsDir() {
			options = append(options, entry.Name())
		}
    }
	return options
}

func (t *Tool) onlyOneSelect() {
	value := ""
	prompt := &survey.Select{
		Message: "Choisir une technologie pour le " + t.Name,
		Options: t.getOption(),
	}
	survey.AskOne(prompt, &value)
	t.Values = append(t.Values, value)
}

func (t *Tool) multiSelect() {
	values := []string{}
	prompt := &survey.MultiSelect{
		Message: "Choisir une technologie pour le " + t.Name,
		Options: t.getOption(),
	}
	survey.AskOne(prompt, &values)
	t.Values = append(t.Values, values...)
}

// Select TODO doc
func (t *Tool) Select() {
	switch t.Name {
	case "BDD":
		t.multiSelect()
	case "Back":
		t.onlyOneSelect()
	case "Front":
		t.onlyOneSelect()
	default:
		fmt.Printf("%s non pris en charge", t.Name)
		os.Exit(1)
	}
}
