package tool

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlecAivazis/survey"
)

// Tool TODO doc
type Tool struct {
	Name   string
	Values []string
	Link   []string
}

// LinkWith TODO t bdd, tool other tools
func (t *Tool) LinkWith(tool *Tool) {
	var response string
	if tool.Values != nil {
		for _, value := range t.Values {
			prompt := &survey.Select{
				Message: "Est-ce que la base de données " + value + " doit être reliée au " + tool.Name,
				Options: []string{"oui", "non"},
			}
			survey.AskOne(prompt, &response)
			if response == "oui" {
				tool.Link = append(tool.Link, value)
				t.Link = append(t.Link, tool.Name)
			}
		}
	}
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
		Message: "Choisir une technologie pour la/le " + t.Name,
		Options: t.getOption(),
	}
	survey.AskOne(prompt, &value)
	t.Values = append(t.Values, value)
}

func (t *Tool) multiSelect() {
	values := []string{}
	prompt := &survey.MultiSelect{
		Message: "Choisir une ou plusieurs technologie(s) pour la/le " + t.Name,
		Options: t.getOption(),
	}
	survey.AskOne(prompt, &values)
	t.Values = append(t.Values, values...)
}

// Select TODO doc
func (t *Tool) Select() {
	switch t.Name {
	case "bdd":
		t.multiSelect()
	case "back":
		t.onlyOneSelect()
	case "front":
		t.onlyOneSelect()
	default:
		fmt.Printf("type de données %s non pris en charge", t.Name)
		os.Exit(1)
	}
}
