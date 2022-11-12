# Simple-Console-Menu
A Simple Console Menu for Go

How to install:
`go install github.com/TheKamboy/Simple-Console-Menu@latest`

# How to use

Creating Selection Menu
``` go
	testsel := SelectionMenu {
		title: "Test",
		description: "Test",
		options: []string{"Test 1", "Test 2", "Test 3"},
		showselectedoption: true,
		separatorstyle: ">",
	}
```

How to use the Selection Menu
``` go
i := testsel.show()
```

The options are marked by numbers. It starts at 1 then goes up the more options you add.

``` go
	if i == 1 {
		println("Test 1")
	} else if i == 2 {
		println("Test 2")
	} else if i == 3 {
		println("Test 3")
	}
```

How to use the YesNo Menu

``` go
	testyesno := YesNoMenu {
		title: "Test",
		description: "Test",
		separatorstyle: "/",
	}

	i := testyesno.show()
  
	if i == "y" {
		println("Yes")
	} else {
		println("No")
	}
```
