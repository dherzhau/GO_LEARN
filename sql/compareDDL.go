//
package main

import (
	//  "bufio"
	"fmt"
	"io/ioutil"

	//   "io/ioutil"
	//	"os"
	"flag"
)


// Parser

var reservedWords = []string{
	"(", ")", ">=", "<=", "!=", ",", "=", ">", "<",
	"SELECT", "INSERT INTO", "VALUES", "UPDATE",
	"DELETE FROM", "WHERE", "FROM", "SET",
	"CREATE TABLE"
  }

type parser struct {
	sql             string        // The query to parse
	i               int           // Where we are in the query
	query           query.Query   // The "query struct" we'll build
	step            step          // What's this? Read on...
  }

  // Main function that returns the "query struct" or an error
func (p *parser) Parse() (query.Query, error) {}

// A "look-ahead" function that returns the next token to parse
func (p *parser) peek() (string) {}

// same as peek(), but advancing our "i" index
func (p *parser) pop() (string) {}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// parsing Inputparameter
	sourceFilePtr := flag.String("source", "", "Source File")
	targetFilePtr := flag.String("target", "", "Target File")
	diffFilePtr := flag.String("diff", "", "Diff File")
	debugPtr := flag.Bool("debug", false, "Debug ")

	flag.Parse()

	if *debugPtr == true {
		fmt.Println("source", *sourceFilePtr)

		fmt.Println("target", *targetFilePtr)
		fmt.Println("diff", *diffFilePtr)
	}

	// slurping sourceFile
	dat, err := ioutil.ReadFile(*sourceFilePtr)
	check(err)
	fmt.Print(string(dat))
}
