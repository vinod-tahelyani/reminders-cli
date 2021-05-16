package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type idsFlag []string

type Person struct {
	name string
	born time.Time
}

func (ids idsFlag) String() string {
	return strings.Join(ids, ", ")
}

func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %s, ad I was born on %s", p.name, p.born.String())
}

func (p *Person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}

func main() {
	// var ids idsFlag

	// flag.Var(&ids, "id", "ids to be appendsd")
	// flag.Parse()
	// fmt.Println(ids)

	var p Person
	flag.Var(&p, "name", "name of the person")
	flag.Parse()
	fmt.Println(p)
}
