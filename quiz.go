package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var file = flag.String("filename", "problems.csv", "name of file")
var time_var = flag.Int("time", 30, "timelimit")

//	func Input(q, a chan string) {
//		var input string
//		<-q
//		fmt.Scan(&input)
//		a <- input
//	}
func Do(reader *csv.Reader, correct, count *int) {
	rec, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error")
	}
	*count = len(rec)
	for _, x := range rec {
		fmt.Printf("%s = ", x[0])
		var input string
		fmt.Scan(&input)
		if input == x[1] {
			*correct++
		}

	}
}
func main() {
	flag.Parse()
	countdown := time.NewTimer(time.Duration(*time_var) * time.Second)
	f, err := os.Open(*file)

	if err != nil {
		fmt.Println("Oops")
		return
	}
	defer f.Close()
	csvreader := csv.NewReader(f)
	count, correct := 0, 0
	go Do(csvreader, &correct, &count)
	<-countdown.C
	fmt.Println()
	fmt.Println("You scored", correct, "out of", count)

}
