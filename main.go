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
		if len(b) > 1 && b[1]&0xC0 == 0x80 {
			return 0, 0, fmt.Errorf("invalid length")
		}
		r = rune(b0)
		s = 1
	case b0&0xE0 == 0xC0: // 2 bytes charactere
		if len(b) > 2 && b[2]&0xC0 == 0x80 {
			return 0, 0, fmt.Errorf("invalid length")
		}
		if len(b) < 2 {
			return 0, 0, fmt.Errorf("Invalid length, expected 2 received %d", len(b))
		}

		b1 := b[1]		
		if b1&0xC0 != 0x80 {
			return 0, 0, fmt.Errorf("invalid length")
		}

		s = 2
		r = ((rune(b0) & 0x1f) << 6) | (rune(b1) & 0x3F)
		
		if r < 0x80 {
			return 0, 0, fmt.Errorf("overlong")
		}
	case b0&0xF0 == 0xE0: // 3 bytes charactere
		if len(b) > 3 && b[3]&0xC0 == 0x80 {
			return 0, 0, fmt.Errorf("invalid length")
		}
		if len(b) < 3 {
			return 0, 0, fmt.Errorf("Invalid length, expected 3 received %d", len(b))
		}
		
		b1 := b[1]		
		b2 := b[2]
			
				
		if b0 == 0xE0 && b1 < 0xA0 {
			return 0, 0, fmt.Errorf("overlong")
		}	


		if b1&0xC0 != 0x80 || b2&0xC0 != 0X80 {
			return 0, 0, fmt.Errorf("invalid continuation byte")
		}
		
		s = 3
		r = ((rune(b0) & 0x0F) << 12) | ((rune(b1) & 0x3F) << 6) | (rune(b2) & 0x3F) 
	case b0&0xF8 == 0xF0: // 4 bytes charactere
		if len(b) > 4 && b[4]&0xC0 == 0x80 {
			return 0, 0, fmt.Errorf("invalid length")
		}
		if len(b) < 4 {
			return 0, 0, fmt.Errorf("Invalid length, expected 4 received %d", len(b))
		}
		
		b1 := b[1]		
		b2 := b[2]
		b3 := b[3]

		if b0 == 0xF0 && b1 < 0x90 {
			return 0, 0, fmt.Errorf("overlong")
		}

		if b1&0xC0 != 0x80 || b2&0xC0 != 0X80 || b3&0xC0 != 0x80 {
			return 0, 0, fmt.Errorf("invalid continuation byte")	
		}
		
		s = 4
		r = ((rune(b0) & 0x07) << 18) | ((rune(b1) & 0x3F) << 12) | ((rune(b2) & 0x3F) << 6) | (rune(b3) & 0x3F) 
	default:
		return r, s, fmt.Errorf("invalid utf8")
	}

	
	if r >= 0xD800 && r <= 0xDFFF {
	
		return 0, 0, fmt.Errorf("surrogate pairs")
	}
	
	if r > 0x10FFFF {
		return 0, 0, fmt.Errorf("too long")
	}
	
	
	return r, s, nil
}
