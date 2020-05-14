package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	diff := p[j].birthDay.Sub(p[i].birthDay).Hours() / 24
	if diff < 0 {
		return true
	}
	if diff > 0 {
		return false
	}
	if p[j].lastName < p[i].lastName {
		return false
	}
	if p[j].firstName < p[i].firstName {
		return false
	}
	return true
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SetDate(InitDate string) time.Time {
	FinalDate, err := time.Parse("2006-01-02", InitDate)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return FinalDate
}

func main() {
	p := People{
		{"Ivan", "Ivanov", SetDate("2005-08-10")},
		{"Ivan", "Ivanov", SetDate("2003-08-05")},
		{"Artiom", "Ivanov", SetDate("2005-08-10")},
	}

	sort.Sort(p)

	for _, a := range p {
		fmt.Printf("%v %v %v\n", a.firstName, a.lastName, a.birthDay.Format("2006-01-02"))
	}

	fmt.Println("\nSome test values:\n")

	t := People{
		{"A", "B", SetDate("2000-01-01")},
		{"A", "A", SetDate("2000-01-01")},
		{"B", "F", SetDate("2001-01-01")},
		{"C", "G", SetDate("2000-01-01")},
		{"C", "G", SetDate("2000-01-02")},
	}

	sort.Sort(t)

	for _, a := range t {
		fmt.Printf("%v %v %v\n", a.firstName, a.lastName, a.birthDay.Format("2006-01-02"))
	}

	fmt.Println("\nSame people but another order:\n")

	t1 := People{
		{"C", "G", SetDate("2000-01-02")},
		{"B", "F", SetDate("2001-01-01")},
		{"A", "B", SetDate("2000-01-01")},
		{"C", "G", SetDate("2000-01-01")},
		{"A", "A", SetDate("2000-01-01")},
	}

	sort.Sort(t1)

	for _, a := range t1 {
		fmt.Printf("%v %v %v\n", a.firstName, a.lastName, a.birthDay.Format("2006-01-02"))
	}
}
