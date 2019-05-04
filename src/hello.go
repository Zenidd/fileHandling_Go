package main

import (
	"fmt"
	"io/ioutil" // IO Handling library
	"strings"
	"time"
	"os"
	"strconv"
	"github.com/gookit/color" //output color library
	"bufio"
)

	func check(e error) {
		if e != nil {
			panic(e)
		}
	}

 func showFiles(fi []os.FileInfo, limit float64, log string) {

	dt := time.Now()

	file, err := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	fmt.Fprintf(w, "Log file from fileHandler.\n")


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
		/*Printing each file properties*/
		color.Cyan.Printf("|%-30v|%12v|%25v|%20v|%10v|\n", file_name, f.Size(), f.ModTime().Format("15:04:05 2006.01.02 "), diff, check)
		/*Printing to log file */
		fmt.Fprintf(w,"|%-30v|%12v|%25v|%20v|%10v|\n", file_name, f.Size(), f.ModTime().Format("15:04:05 2006.01.02 "), diff, check)

	}
		w.Flush()
 }


func main() {


	dirToRead:="./"
	hours2Old := ""
	log := ""

	/*Getting args*/
	if len(os.Args) > 1 {
		hours2Old = os.Args[1]
		/* Without a dir target input, current directory is target*/
		if len(os.Args) == 3 {
			dirToRead = os.Args[2]
		} else if len(os.Args) == 3 {
			log = os.Args[0]
		}

	} else {
		fmt.Println("\n please input needed Hours to make a file erasable")
		os.Exit(3)
	}

	files, err := ioutil.ReadDir(dirToRead)
	check(err)

	user_limit, err := strconv.ParseFloat(hours2Old, 64)
	check(err)

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
	showFiles(files, user_limit, log)

	d.Println(header)
}
