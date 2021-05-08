package enc

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/8i8/lib/log"
)

type Encoder interface {
	Encode(string) string
	Setup(...string)
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
		b := s.caps(byt.Bytes())
		out = string(b)
	case s.SHA256:
		x := sha256.Sum256(input)
		byt.WriteString(fmt.Sprintf("%x", x))
		b := s.caps(byt.Bytes())
		out = string(b)
	case s.SHA384:
		x := sha512.Sum384(input)
		byt.WriteString(fmt.Sprintf("%x", x))
		b := s.caps(byt.Bytes())
		out = string(b)
	case s.SHA512:
		x := sha512.Sum512(input)
		byt.WriteString(fmt.Sprintf("%x", x))
		b := s.caps(byt.Bytes())
		out = string(b)
	case s.pp:
		x := md5.Sum(input)
		byt.WriteString(fmt.Sprintf("%x", x))
		b := s.caps(byt.Bytes())
		b = b[:20]
		out = string(b)
	default:
		// Default hash is md5
		x := md5.Sum(input)
		byt.WriteString(fmt.Sprintf("%x", x))
		b := s.caps(byt.Bytes())
		out = string(b)
	}
	s = settings{} // reset all
	return out
}

func (s *settings) Setup(args ...string) {
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
		default:
			log.Err(nil, "encoding", "Setup", "Unknown argument")
		}
	}
}

func isLetter(c byte) bool {
	if c > 96 && c < 123 {
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
