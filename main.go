package main

import (
	"fmt"
	"log"
)


const (
	DIRECTORY = "dir"
	FILE      = "file"
)

type tree struct {
	Name   string
	Type   string
	Items  []*tree
	Parent *tree
}

func New(name, Type string) *tree {
	return &tree{
		Name: name,
		Type: Type,
		Items: nil,
		Parent: nil,
	}
}

func (d *tree) Add(i *tree) {
	if d.Type == FILE {
		log.Fatalln("cannot add child to a file")
	}

	i.Parent = d
	d.Items = append(d.Items, i)
}

func (d *tree) Depth() int {
	depth := 0
	cur := d

	for cur.Parent != nil {
		cur = cur.Parent
		depth += 1
	}
	return depth
}

func (d *tree) Draw() {
	prefix := ""
	for i := 0; i < d.Depth(); i++ {
		prefix += "  "
	}

	if d.Parent != nil {
		prefix += "|__ "
	} else {
		prefix += ""
	}

	fmt.Println(prefix + d.Name)

	if d.Items != nil {
		for i := 0; i < len(d.Items); i++ {
			d.Items[i].Draw()
		}
	}

}


func main() {
	d := New("godir/", DIRECTORY)

	d.Add(New("main.go", FILE))
	utils := New("utils/", DIRECTORY)
	d.Add(utils)
	utils.Add(New("check-validity.go", FILE))
	d.Add(New("go.mod", FILE))

	d.Draw()
}