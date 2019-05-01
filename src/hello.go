package main

import (
	"fmt"
	"io/ioutil" // IO Handling library
	"log"
	"strings"
	"time"
	"os"
	"strconv"
	"github.com/gookit/color" //output color library
)

 func showFiles(fi []os.FileInfo, limit float64) {
	dt := time.Now()

	/*Iterating over files*/
	for _, f := range fi {

		/*Adjusting filename lenght*/
		file_name := f.Name()
		if len(file_name) > 30 {
			file_name = file_name[:30]
		}



		/*Checking hours since last modification*/
		diff := dt.Sub(f.ModTime())
		check := " "
		n:= diff.Hours()
		if n > limit {
			check = "X"
		}


		color.Cyan.Printf("|%-30v|%12v|%25v|%20v|%10v|\n", file_name, f.Size(), f.ModTime().Format("15:04:05 2006.01.02 "), diff, check)
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

	/* Output format and styles */
	dt := time.Now()
	header := strings.Repeat("-", 103)

	c := color.C256(132) // fg color
	d := color.C256(51) // fg color
	c.Println("\n Current date and time is: ", dt.Format("15:04:05 2006.01.02 "))
	d.Println(header)
	color.Style{color.FgCyan, color.OpBold}.Printf("|%-30v|%12v|%25v|%20v|%10v|\n", "Name", "Size", "Timestamp","","Erasable")
	d.Println(header)

	/* Showing and looking for files */
	showFiles(files, user_limit)

	d.Println(header)
}
