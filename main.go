package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"os/exec"
	"runtime"
)

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

	fmt.Println("Press ESC to quit")
	fmt.Println("Press ← ↑ ↓ →")
	r := rocket{h: 15, w: 30, x: 5, y: 6}
	r.init()
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowLeft {
			r.left()
		}
		if key == keyboard.KeyArrowRight {
			r.right()
		}
		if key == keyboard.KeyArrowUp {
			r.up()
		}
		if key == keyboard.KeyArrowDown {
			r.down()
		}
		clearScreen()
		r.Draw()

		//fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		if key == keyboard.KeyEsc {
			break
		}
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
