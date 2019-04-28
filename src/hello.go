package main

import (
	"fmt"
	"io/ioutil" // IO Handling library
	"log"
	"strings"
	"time"
	"os"
	"strconv"
)

 func showFiles(fi []os.FileInfo, limit float64) {
	dt := time.Now()

	for _, f := range fi {
		diff := dt.Sub(f.ModTime())
		check := " "
		n:= diff.Hours()
		if n > limit {
			check = "X"
		}
		fmt.Printf("|%-30v|%6v|%40v|%20v|%10v|\n", f.Name(), f.Size(), f.ModTime(),diff, check)
	}
 }


func main() {

	dirToRead:="./"
	hours2Old := ""

	/*Getting args*/
	if len(os.Args) > 1 {
		hours2Old = os.Args[1]
		/* Without a dir target input, current directory is target*/
		if len(os.Args) == 3 {
			dirToRead = os.Args[2]
		}

	} else {
		fmt.Println("\n please input needed Hours to make a file erasable")
		os.Exit(3)
	}


	files, err := ioutil.ReadDir(dirToRead)

	if err != nil {
		log.Fatal(err)
	}

	user_limit, err := strconv.ParseFloat(hours2Old, 64)

	if err == nil {
        fmt.Println(user_limit)
    }

	/* Output format */
	dt := time.Now()
	header := strings.Repeat("-", 112)
	fmt.Println("\n Current date and time is: ", dt.String())
	fmt.Println(header)
	fmt.Printf("|%-30v|%6v|%40v|%20v|%10v|\n", "Name", "Size", "Timestamp","","Erasable")
	fmt.Println(header)

	/* Showing and looking for files */
	showFiles(files, user_limit)

	fmt.Println(header)
}
