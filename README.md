# Simple-Console-Menu
A Simple Console Menu for Go

## Unarchived
Here's a [reason why I unarchived this.](https://github.com/TheKamboy/Simple-Console-Menu/blob/main/UNARCHIVED.md)

# Install
How to use in a go program:
`go get github.com/TheKamboy/Simple-Console-Menu`

# How to use

Creating Selection Menu
``` go
	testsel := SelectionMenu {
		Title: "Test",
		Description: "Test",
		Options: []string{"Test 1", "Test 2", "Test 3"},
		Showselectedoption: true,
		Separatorstyle: ">",
	}
```

How to use the Selection Menu
``` go
i := testsel.Show()
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
		Title: "Test",
		Description: "Test",
		Separatorstyle: "/",
	}

	i := testyesno.Show()
  
	if i == "y" {
		println("Yes")
	} else {
		println("No")
	}
```
