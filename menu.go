package main

import (
	"fmt"
	"io"
)

type MenuStream interface {
	nextLine() (string, error)
}

type MenuItem struct {
	Id         string
	Type       string
	Name       string
	Price      string
	References []string
}

type Menu struct{}

func main() {
	ms := newMenuStream()

	//   phases := ["ID", "TYPE", "NAME", "PRICE", "REFERENCE"]
	phaseIndex := 0

	items := []MenuItem{}
	var mi MenuItem
	for {
		value, err := ms.nextLine()
		if err != nil {
			break
		}
		if value == "" {
			items = append(items, mi)
			phaseIndex = 0
			continue
		}
		if phaseIndex == 0 {
			mi = MenuItem{Id: value}
			phaseIndex++
		} else if phaseIndex == 1 {
			mi.Type = value
			phaseIndex++
		} else if phaseIndex == 2 {
			mi.Name = value
			phaseIndex++
		} else if phaseIndex == 3 {
			if mi.Type == "CATEGORY" { // references
				mi.References = append(mi.References, value)
			} else {
				mi.Price = value
			}
			phaseIndex++
		} else {
			mi.References = append(mi.References, value)
		}
	}

	for _, item := range items {
		fmt.Printf("%+v\n", item)
	}
}

func newMenuStream() MenuStream {
	return &menuStreamImpl{
		lines: []string{"4", "ENTREE", "Spaghetti", "10.95", "2", "3", "", "1", "CATEGORY", "Pasta", "4", "5", "", "2", "OPTION", "Meatballs", "1.00", "", "3", "OPTION", "Chicken", "2.00", "", "5", "ENTREE", "Lasagna", "12.00", "", "6", "ENTREE", "Caesar Salad", "9.75", "3", ""},
	}
}

type menuStreamImpl struct {
	lines []string
}

func (m *menuStreamImpl) nextLine() (string, error) {
	if len(m.lines) == 0 {
		return "", io.EOF
	}
	var result string
	result, m.lines = m.lines[0], m.lines[1:]
	return result, nil
}
