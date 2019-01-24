package main

import "github.com/fatih/color"

func CheckErr(err error, level int) {

	if err != nil {
		if level == errorWarning {
			color.Red("Error occured %v", err)
		} else if level == errorPanic {
			panic(err)
		}
	}
}
