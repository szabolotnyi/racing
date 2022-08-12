package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var rocketT = `
 /\
||||
 ||
/||\`

var rocketT1 = ` 
  ^
 ^ ^
 |#|
/└ ┘\`

var rocketT2 = `
  /\
 /__\
|    |
 \--/
 /\/\
 ||||`

func main() {
	// GOOS=linux GOARCH=amd64 go build -o rocket_linux
	// GOOS=windows GOARCH=amd64 go build -o rocket.exe
	// GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o mac_intel
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	width := 18
	height := 50

	r1 := NewEngine(width, height, rocketT, 2, 3)
	r2 := NewEngine(width, height, rocketT1, 5, 8)
	r3 := NewEngine(width, height, rocketT2, 9, 0)

	for {
		clearScreen()
		draw(merge(
			addBorderAndGrid(matrix(width, height)),
			r1.draw(),
			r2.draw(),
			r3.draw(),
		))

		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}

		switch key {
		case keyboard.KeyArrowLeft:
			r1.left()
		case keyboard.KeyArrowRight:
			r1.right()
		case keyboard.KeyArrowUp:
			r1.up()
		case keyboard.KeyArrowDown:
			r1.down()
		}
	}
}

func merge(data ...[][]string) [][]string {
	res := data[0]
	for n := 1; n < len(data); n++ {
		for i := 1; i < len(data[n]); i++ {
			for j := 1; j < len(data[n][i]); j++ {
				if data[n][i][j] != "" {
					res[i][j] = data[n][i][j]
				}
			}
		}
	}

	return res
}

func addBorderAndGrid(m [][]string) [][]string {
	max := len(m) - 1
	for i := 0; i <= max; i++ {
		ll := len(m[i]) - 1

		switch i {
		case 0:
			fillLine(m[i], "-")
			m[i][0] = "┌"
			m[i][ll] = "┐"
		case max:
			fillLine(m[i], "-")
			m[max][0] = "└"
			m[max][ll] = "┘"
		default:
			if i%2 == 0 {
				fillLine(m[i], []string{".", ".", " ", " "}...)
			} else {
				fillLine(m[i], []string{" ", " ", ".", "."}...)
			}
			m[i][0] = "│"
			m[i][ll] = "│"
		}
	}

	return m
}

func fillLine(m []string, c ...string) {
	t := 0
	for i := 0; i < len(m); i++ {
		m[i] = c[t]
		t++
		if t == len(c) {
			t = 0
		}
	}
}

func matrix(n, m int) [][]string {
	res := make([][]string, n)
	for i := 0; i < n; i++ {
		res[i] = make([]string, m)
	}

	return res
}

func draw(m [][]string) {
	for i := 0; i < len(m); i++ {
		fmt.Println(strings.Join(m[i], ""))
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}
