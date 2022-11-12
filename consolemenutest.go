package consolemenu

func main() {
	testsel := SelectionMenu {
		title: "Test",
		description: "Test",
		options: []string{"Test 1", "Test 2", "Test 3"},
		showselectedoption: true,
		separatorstyle: ">",
	}

	i := testsel.show()

	if i == 1 {
		println("Test 1")
	} else if i == 2 {
		println("Test 2")
	} else if i == 3 {
		println("Test 3")
	}

	testyesno := YesNoMenu {
		title: "Test",
		description: "Test",
		separatorstyle: "/",
	}

	j := testyesno.show()

	if j == "y" {
		println("Yes")
	} else {
		println("No")
	}
}