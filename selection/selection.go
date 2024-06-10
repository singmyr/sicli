package selection

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

type direction int

const (
	unknown direction = iota
	up
	down
	right
	left
)

func GetKey(options []string) int {
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

	for {
		for i, option := range options {
			if i == selected {
				fmt.Printf("\x1b[32m%s\x1b[0m\r\n", option)
			} else {
				fmt.Printf("\x1b[31m%s\x1b[0m\r\n", option)
			}
		}

		b := make([]byte, 3)
		_, err = os.Stdin.Read(b)
		if err != nil {
			log.Fatalln("failed reading from stdin:", err)
		}
		var key direction
		switch b[2] {
		case 'A':
			key = up
		case 'B':
			key = down
		case 'C':
			key = right
		case 'D':
			key = left
		default:
			key = unknown
		}

		if key == up {
			if selected > 0 {
				selected--
			}
		} else if key == down {
			if selected < len(options)-1 {
				selected++
			}
		}

		for range options {
			fmt.Printf("%c[%dA%c[2K", 27, 1, 27)
		}

		if b[0] == '\r' {
			break
		}

		if b[0] == '\x03' && b[1] == '\x00' && b[2] == '\x00' {
			restore()
			os.Exit(1)
		}
	}

	restore()

	return selected
}

func GetValue(options []string) string {
	key := GetKey(options)

	return options[key]
}
