package main

type Cell struct {
	column uint
	row    uint
	vacant bool
	valid  bool
}

func NewCell(column uint, row uint, vacant bool, valid bool) *Cell {
	return &Cell{
		column: column,
		row:    row,
		vacant: vacant,
		valid:  valid}
}
