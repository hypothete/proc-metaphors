package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func pick(s []string) string {
	return s[rand.Intn(len(s))]
}

//Comparator uses templates to write comparison statements
type Comparator struct {
	templates []string
	nouns     []string
}

//PrintNew prints comparison statements based on the random templates and nouns it draws
func (c Comparator) PrintNew() {
	t := pick(c.templates)
	nouncount := strings.Count(t, "%s")
	nouns := []interface{}{}
	for i := 0; i < nouncount; i++ {
		nouns = append(nouns, pick(c.nouns))
	}
	fmt.Printf(t, nouns...)
	fmt.Println("")
}

func main() {
	comp := Comparator{}

	data, err := ioutil.ReadFile("nouns.csv")
	if err != nil {
		fmt.Print(err)
	}
	comp.nouns = strings.Split(string(data), "\r\n")

	data, err = ioutil.ReadFile("templates.txt")
	if err != nil {
		fmt.Print(err)
	}
	comp.templates = strings.Split(string(data), "\r\n")

	rand.Seed(time.Now().Unix())

	for i := 0; i < 30; i++ {
		comp.PrintNew()
	}
}
