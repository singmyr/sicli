package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

type Direction uint32

const (
	UNKNOWN Direction = iota
	UP
	DOWN
	RIGHT
	LEFT
)

func main() {
	options := []string{
		"Option 1",
		"Option 2",
		"Option 3",
	}
	selected := 0

	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln("failed setting stdin to raw:", err)
	}

	restore := func() {
		if err := term.Restore(0, state); err != nil {
			log.Fatalln("failed restoring terminal:", err)
		}
	}
	defer restore()

	for {
		for i, option := range options {
			if i == selected {
				fmt.Printf("\x1b[32m%s\x1b[0m\r\n", option)
			} else {
				fmt.Printf("\x1b[31m%s\x1b[0m\r\n", option)
			}
		}

		var b []byte = make([]byte, 3)
		_, err = os.Stdin.Read(b)
		if err != nil {
			log.Fatalln("failed reading from stdin:", err)
		}
		var key Direction
		switch b[2] {
		case 'A':
			// Up
			key = UP
		case 'B':
			// Down
			key = DOWN
		case 'C':
			// Right
			key = RIGHT
		case 'D':
			// Left
			key = LEFT
		default:
			key = UNKNOWN
		}

		if key == UP {
			selected--
		} else if key == DOWN {
			selected++
		}

		if b[0] == 'q' {
			break
		}

		for range options {
			fmt.Printf("%c[%dA%c[2K", 27, 1, 27)
		}

		if b[0] == '\r' {
			restore()
			fmt.Printf("Chose %d\n", selected)
			break
		}
	}
}
