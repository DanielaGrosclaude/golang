package main

import (
	"main.go/packageone"
)

var myVar string = "myVar"

func main() {
	//variables
	//declare a package level variable for the main
	// package named myVar

	// declar a block variable for the main function
	//called blockVar
	var blockVar string = "blockVar"

	//declare a package level varibale in the packageone
	//package named PackageVar

	//create an exported function in packageone called PrintMe

	//in the main function, prin out the values of myVar
	//blocVar, and PackageVar on one line using the PrintMe
	//function in packageOne
	packageone.PrintMe(myVar, blockVar)
}
