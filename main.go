package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	licenseCheck()
	fmt.Println("license check passed")
}

func licenseCheck() {
	if time.Now().After(time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)) {
		fmt.Println("license expired")
		os.Exit(1)
	}
}
