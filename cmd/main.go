package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	const testUrl = "https://www.google-analytics.com/analytics.php"

	fmt.Println(color.GreenString(testUrl))
}
