package main

import "fmt"

func main() {
	column := []string{"id", "name", "age", "sex", "phone"}
	index := []string{"id", "name", "phone"} //first 3 fields is cluster index
	userField := []string{"id"}

	f := NewIndexBuilder(column, index)
	f.AddIndex("sex")

	fmt.Println(f.Choice(userField))
}

type IndexBuilder struct {
	columns []string
	index   []string
	size    int
}

func (ib *IndexBuilder) AddIndex(index string) {
	ib.index = append(ib.index, index)
}

func NewIndexBuilder(column, index []string) *IndexBuilder {
	var size int
	for i := 0; i < len(index); i++ {
		size += len(index[i])
	}

	return &IndexBuilder{
		columns: column,
		size:    len(index),
		index:   index,
	}
}

func (ib *IndexBuilder) Choice(userField []string) string {
	var (
		i, j  int
		index []string = ib.index
	)

	for i < ib.size && j < len(userField) {
		for j < len(userField) {
			if index[i] == userField[j] {
				userField[i], userField[j] = userField[j], userField[i]
				i++
				j = i
				break
			}
			j++
		}

		if i != j {
			break
		}
	}

	if i != 0 {
		return "cluster"
	}

	for i := ib.size; i < len(index); i++ {
		for j := 0; j < len(userField); j++ {
			if index[i] == userField[j] {
				return index[i]
			}
		}
	}

	return ""
}
