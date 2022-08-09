package main

import "fmt"

const emptyCel = "."

type rocket struct {
	x int
	y int

	store []string
	w     int
	h     int
}

func (m *rocket) left() {
	if m.x-1 < 0 {
		return
	}
	m.x--
}

func (m *rocket) right() {
	if m.x+1 > m.w-4 {
		return
	}
	m.x++
}

func (m *rocket) up() {
	if m.y-1 < 0 {
		return
	}
	m.y--
}

func (m *rocket) down() {
	if m.y+1 > m.h-5 {
		return
	}
	m.y++
}

func (m *rocket) init() {
	m.store = []string{
		` /\ `,
		`||||`,
		` || `,
		`/||\`,
	}
}

func (m *rocket) Draw() {
	fmt.Println(fmt.Sprintf(`┌%s┐`, m.offset(m.w, "-")))
	ol := m.offset(m.w, emptyCel)
	for i := 0; i < m.y; i++ {
		fmt.Println(fmt.Sprintf(`│%s│`, ol))
	}

	m.draw()

	for i := m.y + 4; i < m.h; i++ {
		fmt.Println(fmt.Sprintf(`│%s│`, ol))
	}
	fmt.Println(fmt.Sprintf(`└%s┘`, m.offset(m.w, "-")))
}

func (m *rocket) draw() {
	ol := m.offset(m.x, emptyCel)
	for _, v := range m.store {
		or := m.offset(m.w-len(ol)-len(v), emptyCel)
		fmt.Println(fmt.Sprintf(`│%s%s%s│`, ol, v, or))
	}
}

func (m *rocket) offset(n int, cel string) string {
	o := ""
	for i := 0; i < n; i++ {
		o += cel
	}

	return o
}
