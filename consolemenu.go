package consolemenu

import (
	"fmt"
	"os"
	"strings"
)

type SelectionMenu struct {
	title              string
	description        string
	options            []string
	showselectedoption bool
	separatorstyle     string
}

type YesNoMenu struct {
	title          string
	description    string
	separatorstyle string
}

func (e SelectionMenu) show() (optionselected int) {
	var input int
	a := 1
	i := 0
	cs := ""
	if !(e.title == "") {
		fmt.Println(e.title)
		fmt.Println()
	}
	if !(e.description == "") {
		fmt.Println(e.description)
		fmt.Println()
	}
	if e.separatorstyle == "" {
		cs = ">"
	} else {
		cs = e.separatorstyle
	}
	if e.options != nil {
		for i < len(e.options) {
			fmt.Println(a, cs, e.options[i])
			a++
			i++
		}
	} else {
		fmt.Println(fmt.Errorf("nothing found in options variable"))
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
	if e.showselectedoption {
		fmt.Println()
		fmt.Println("You selected Option", input, "which has the name:", e.options[input-1])
	}
	optionselected = input
	return
}

func (e YesNoMenu) show() (optionselected string) {
	var input string
	cs := ""
	if !(e.title == "") {
		fmt.Println(e.title)
		fmt.Println()
	}
	if !(e.description == "") {
		fmt.Println(e.description)
		fmt.Println()
	}
	if e.separatorstyle == "" {
		cs = "/"
	} else {
		cs = e.separatorstyle
	}
	text := fmt.Sprintf("y%vn: ", cs)
	if e.title == "" {
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

	optionselected = input
	return
}
