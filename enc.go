package enc

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

type Encoder interface {
	Encode(string) string
	Setup(...string) error
}

func NewEncoder() Encoder {
	return &settings{}
}

type settings struct {
	// MD5 Input flags.
	MD5 bool
	// NEWLINE Input flags.
	NEWLINE bool
	// SHA256 Input flags.
	SHA256 bool
	// SHA384 Input flags.
	SHA384 bool
	// SHA512 Input flags.
	SHA512 bool
	// pp truncates the output to 20 characters.
	pp bool
	// CAPITAL capitalises the last letter in the sha for cases in which a capital
	// letter is a requirment.
	CAPITAL bool
	// SYMBOL repalces the last character of the string that is not
	// upper case with a symbol.
	SYMBOL bool
	// VISIBLE when set to true, displays user input whilst it is being typed.
	VISIBLE bool
}

func (s settings) Encode(str string) string {

	var out string
	input := []byte(str)

	// Set the appropriate hash for the input, the given flag or lack there
	// of, set the hashing algorithm. The resulting hash is copied to the system
	// clipboard and then written to stdout.
	byt := bytes.Buffer{}
	switch {
	case s.MD5:
		x := md5.Sum(input)
		byt.WriteString(fmt.Sprintf("%x", x))
	case s.SHA256:
		x := sha256.Sum256(input)
		byt.WriteString(fmt.Sprintf("%x", x))
	case s.SHA384:
		x := sha512.Sum384(input)
		byt.WriteString(fmt.Sprintf("%x", x))
	case s.SHA512:
		x := sha512.Sum512(input)
		byt.WriteString(fmt.Sprintf("%x", x))
	case s.pp:
		x := md5.Sum(input)
		byt.WriteString(fmt.Sprintf("%x", x))
	default:
		// Default hash is md5
		x := md5.Sum(input)
		byt.WriteString(fmt.Sprintf("%x", x))
	}
	b := byt.Bytes()
	if s.pp {
		b = b[:20]
	}
	b = s.caps(b)
	b = s.symbol(b)
	out = string(b)

	s = settings{} // reset all
	return out
}

func (s *settings) Setup(args ...string) error {
	for _, str := range args {
		switch str {
		case "md5":
			s.MD5 = true
		case "newline":
			s.NEWLINE = true
		case "sha256":
			s.SHA256 = true
		case "sha384":
			s.SHA384 = true
		case "sha512":
			s.SHA512 = true
		case "pp":
			s.pp = true
		case "capital":
			s.CAPITAL = true
		case "symbol":
			s.SYMBOL = true
		default:
			return fmt.Errorf("Unknown setup flag")
		}
	}
	return nil
}

func isLetter(c byte) bool {
	if c > 96 && c < 123 {
		return true
	}
	return false
}

func isUpperCase(c byte) bool {
	if c > 64 && c < 91 {
		return true
	}
	return false
}

func toUpper(c byte) byte {
	return c - 32
}

func (s settings) caps(b []byte) []byte {
	if s.CAPITAL {
		for c := len(b) - 1; c > 0; c-- {
			if isLetter(b[c]) {
				b[c] = toUpper(b[c])
				break
			}
		}
	}
	return b
}

func addSymbol() byte {
	return '@'
}

func (s settings) symbol(b []byte) []byte {
	if s.SYMBOL {
		for c := len(b) - 1; c > 0; c-- {
			if !isUpperCase(b[c]) {
				b[c] = addSymbol()
				break
			}
		}
	}
	return b
}
