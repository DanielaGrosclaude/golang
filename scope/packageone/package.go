package packageone

import "fmt"

var PackageVar string = "PackageVar"

func PrintMe(df1, s2 string) {
	fmt.Println(df1, s2, PackageVar)

}
