package consolemenu

import (
	"fmt"
	"os"
	"strings"
)

type SelectionMenu struct {
	Title              string
	Description        string
	Options            []string
	Showselectedoption bool
	Separatorstyle     string
}

type YesNoMenu struct {
	Title          string
	Description    string
	Separatorstyle string
}

func (e SelectionMenu) show() (Optionselected int) {
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
		fmt.Println(fmt.Errorf("nothing found in Options variable"))
		os.Exit(1)
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
	if e.Showselectedoption {
		fmt.Println()
		fmt.Println("You selected Option", input, "which has the name:", e.Options[input-1])
	}
	Optionselected = input
	return
}

func (e YesNoMenu) show() (Optionselected string) {
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
