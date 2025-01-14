package main

import (
	"fmt"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func otus_print() {

	line := "Привет,"
	fmt.Println(line, "это мой первый проект")
}

func otus_branch_feature00() {

	line := "Привет,"
	fmt.Println(line, "это новая фича 00")
}

func otus_branch_feature01() {

	line := "Привет,"
	fmt.Println(line, "это новая фича 01")
}

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}

	otus_branch_feature00()
	otus_branch_feature01()
}
