package main

import "fmt"

func main() {}

func decodeRune(b []byte) (r rune, s int, err error) {
	if len(b) == 0 {
		return 0, 0, fmt.Errorf("empty input")
	}
	b0 := b[0]
	
	switch {
	case b0 < 0x80: // 128 - ASCII
		r = rune(b0)
		s = 1
	case b0&0xE0 == 0xC0: // 2 bytes charactere
		if len(b) < 2 {
			return 0, 0, fmt.Errorf("Invalid length, expected 2 received %d", len(b))
		}
		s = 2
		r = ((rune(b0) & 0x1f) << 6) | (rune(b[1]) & 0x3F)
	case b0&0xF0 == 0xE0: // 3 bytes charactere
		if len(b) < 3 {
			return 0, 0, fmt.Errorf("Invalid length, expected 3 received %d", len(b))
		}
		s = 3
		r = ((rune(b0) & 0x0F) << 12) | ((rune(b[1]) & 0x3F) << 6) | (rune(b[2]) & 0x3F) 
	case b0&0xF8 == 0xF0: // 4 bytes charactere
		if len(b) < 4 {
			return 0, 0, fmt.Errorf("Invalid length, expected 4 received %d", len(b))
		}
		s = 4
		r = ((rune(b0) & 0x07) << 18) | ((rune(b[1]) & 0x3F) << 12) | ((rune(b[2]) & 0x3F) << 6) | (rune(b[3]) & 0x3F) 
	}

	return r, s, nil
}
