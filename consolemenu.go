// Package
package consolemenu

// Imports
import (
	"fmt"
	"strings"
)

// The Selection Menu allows you to make a menu
// with a couple of options selectable by numbers
type SelectionMenu struct {
	Title          string
	Description    string
	Options        []string
	Separatorstyle string // Changes the separator between the number and the option title (default: >)
}

// The Multi Selection Menu is the Selection Menu
// But you can select multiple options
type MultiSelectionMenu struct {
	Title           string
	Description     string
	Options         []string
	Separatorstyle  string // Changes the separator between the number and the option title (default: >)
	ShowControlsMSG bool
}

type YesNoMenu struct {
	Title          string
	Description    string
	Separatorstyle string // Changes the separator between the number and the option title (default: >)
}

func (e SelectionMenu) Show() (Optionselected int) {
	var input int
	a := 1
	i := 0
	cs := ""
	if !(e.Title == "") {
		fmt.Println(e.Title)
		fmt.Println()
	}
	if !(e.Description == "") {
		fmt.Println(e.Description)
		fmt.Println()
	}
	if e.Separatorstyle == "" {
		cs = ">"
	} else {
		cs = e.Separatorstyle
	}
	if e.Options != nil {
		for i < len(e.Options) {
			fmt.Println(a, cs, e.Options[i])
			a++
			i++
		}
	} else {
		panic("ERROR: \"Options\" cannot be empty!")
	}
	fmt.Println()
	fmt.Print(">>")
	fmt.Scanln(&input)
	for input > a-1 || input < 1 {
		warning := fmt.Sprintf("\"%v\" is not an option.", input)
		fmt.Println()
		fmt.Println(warning)
		fmt.Println()
		fmt.Print(">>")
		fmt.Scanln(&input)
	}
	Optionselected = input
	return
}

func (e YesNoMenu) Show() (Optionselected string) {
	var input string
	cs := ""
	if !(e.Title == "") {
		fmt.Println(e.Title)
		fmt.Println()
	}
	if !(e.Description == "") {
		fmt.Println(e.Description)
		fmt.Println()
	}
	if e.Separatorstyle == "" {
		cs = "/"
	} else {
		cs = e.Separatorstyle
	}
	text := fmt.Sprintf("y%vn: ", cs)
	if e.Title == "" {
		fmt.Println()
	}
	fmt.Print(text)
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	for {
		if input == "y" || input == "n" {
			break
		}
		warning := fmt.Sprintf("\"%v\" is not an option.", input)
		fmt.Println()
		fmt.Println(warning)
		fmt.Println()
		fmt.Print(text)
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		input = strings.ReplaceAll(input, " ", "")
	}

	Optionselected = input
	return
}

func (e MultiSelectionMenu) Show() (Optionselected []string) {
	var input int = 1
	a := 1
	i := 0
	cs := ""
	if !(e.Title == "") {
		fmt.Println(e.Title)
		fmt.Println()
	}
	if !(e.Description == "") {
		fmt.Println(e.Description)
		fmt.Println()
	}
	if e.Separatorstyle == "" {
		cs = ">"
	} else {
		cs = e.Separatorstyle
	}
	if e.Options != nil {
		for i < len(e.Options) {
			fmt.Println(a, cs, e.Options[i])
			a++
			i++
		}
	} else {
		panic("ERROR: \"Options\" cannot be empty!")
	}
	if e.ShowControlsMSG {
		fmt.Println("Press Enter without any numbers to exit.")
	}
	for input != 0 {
		input = 0
		fmt.Println()
		fmt.Print(">>")
		fmt.Scanln(&input)

		if input > a-1 {
			warning := fmt.Sprintf("\"%v\" is not an option.", input)
			fmt.Println()
			fmt.Println(warning)
		} else {
			hasitem := false
			for f := 0; f < len(e.Options); f++ {
				if Optionselected != nil {
					if !(f >= len(Optionselected)) {
						if Optionselected[f] == e.Options[input-1] {
							fmt.Println("You already have that selected!")
							hasitem = true
							break
						}
					}
				} else {
					break
				}
			}

			if !(hasitem) {
				if input != 0 {
					op := Optionselected
					Optionselected = []string{}
					Optionselected = append(op, e.Options[input-1])
				}
			}
		}
	}
	return
}
