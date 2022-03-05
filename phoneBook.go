package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}
var MIN = 0
var MAX = 26

func search(Key string) *Entry {
	for i, v := range data {
		if v.Surname == Key {
			return &data[i]
		}
	}
	return nil
}
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}

	return temp
}

func popoulate(n int, s []Entry) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(10, 199))
		data = append(data, Entry{name, surname, n})
	}
}

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		// exe var holds the path to the executable file
		exe := path.Base(arguments[0])
		fmt.Printf("Usage %s searh|list <arguments> \n", exe)
		return
	}
	SEED := time.Now().Unix()
	rand.Seed(SEED)

	// Records to populate
	n := 100
	popoulate(n, data)
	fmt.Printf("Data has %d entries.\n", len(data))

	// Differentiate between the commands
	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search <surname>")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not fount", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}
