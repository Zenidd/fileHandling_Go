package main

import (
	"fmt"
	"io/ioutil" // IO Handling library
	"log"
	"strings"
	"time"
)

func main() {


	
	files, err := ioutil.ReadDir("./")
	dt := time.Now()
	header := strings.Repeat("-", 94)

	if err != nil {
		log.Fatal(err)
	}

	/* Output format */

	fmt.Println("\n Current date and time is: ", dt.String())
	fmt.Println(header)
	fmt.Printf("|%-30v|%6v|%45v|\n", "Name", "Size", "Timestamp")

	/* Go through directory */
	for _, f := range files {
		diff := dt.Sub(f.ModTime())
		check := " "
		if diff.Hours() > 1 {
			check = "X"
		}
		fmt.Printf("|%-30v|%6v|%45v|%20v|%20v|\n", f.Name(), f.Size(), f.ModTime(),diff, check)

			

	}
	fmt.Println("\n ")
}
