package main

import (
	"strings"
)

type engine struct {
	x  int
	y  int
	xx int

	store [][]string
	sw    int
	sh    int
}

func NewEngine(w, h int, s string, x, y int) *engine {
	var data [][]string
	var xx int
	for i, v := range strings.Split(s, "\n") {
		data = append(data, []string{})
		data[i] = append(data[i], strings.Split(v, "")...)
		if xx < len(data[i]) {
			xx = len(data[i])
		}
	}

	return &engine{
		x:     x,
		y:     y,
		xx:    xx,
		sw:    h,
		sh:    w,
		store: data,
	}
}

func (m *engine) left() {
	if m.x-2 < 0 {
		return
	}
	m.x--
}

func (m *engine) right() {
	if m.x+1 > m.sw-m.xx-1 {
		return
	}
	m.x++
}

func (m *engine) up() {
	if m.y-1 < 0 {
		return
	}
	m.y--
}

func (m *engine) down() {
	if m.y+1 > m.sh-len(m.store)-1 {
		return
	}
	m.y++
}

func (m *engine) draw() [][]string {
	data := matrix(m.sh, m.sw)

	for i := 0; i < len(m.store); i++ {
		for j := 0; j < len(m.store[i]); j++ {
			data[m.y+i][m.x+j] = m.store[i][j]
		}
	}
	return data
}
