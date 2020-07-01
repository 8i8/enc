// enc requires the package github.com/atotto/clipboard
package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

var (
	// MD5 Input flags.
	MD5 = flag.Bool("md5", false, "Output a md5 hash.")
	// NEWLINE Input flags.
	NEWLINE = flag.Bool("n", false, "Output the hash generated when a newline char is included in the string to be hashed.")
	// SHA256 Input flags.
	SHA256 = flag.Bool("sha256", false, "Output a sha512 hash.")
	// SHA384 Input flags.
	SHA384 = flag.Bool("sha384", false, "Output a sha384 hash.")
	// SHA512 Input flags.
	SHA512 = flag.Bool("sha512", false, "Output a sha512 hash.")
	pp     = flag.Bool("pp", false, "")
)

// CAPITAL capitalises the last letter in the sha for cases in which a capital
// letter is a requirment.
var CAPITAL = flag.Bool("c", false, "Capitalise the last letter in the hash.")

func main() {

	// Get flags from the os on program start.
	flag.Parse()

	for {
		// Read stdin.
		scanner := bufio.NewScanner(os.Stdin)

		if !scanner.Scan() {
			return
		}

		input := scanner.Bytes()
		if len(input) == 0 {
			return
		}
		if *NEWLINE {
			input = append(input, '\n')
		}

		// Set the appropriate hash for the input, the given flag or lack there
		// of, set the hashing algorithm. The resulting hash is copied to the system
		// clipboard and then written to stdout.
		byt := bytes.Buffer{}
		switch {
		case *MD5:
			x := md5.Sum(input)
			byt.WriteString(fmt.Sprintf("%x", x))
			b := caps(byt.Bytes())
			clipboard.WriteAll(string(b))
			fmt.Println(string(b))
		case *SHA256:
			x := sha256.Sum256(input)
			byt.WriteString(fmt.Sprintf("%x", x))
			b := caps(byt.Bytes())
			clipboard.WriteAll(string(b))
			fmt.Println(string(b))
		case *SHA384:
			x := sha512.Sum384(input)
			byt.WriteString(fmt.Sprintf("%x", x))
			b := caps(byt.Bytes())
			clipboard.WriteAll(string(b))
			fmt.Println(string(b))
		case *SHA512:
			x := sha512.Sum512(input)
			byt.WriteString(fmt.Sprintf("%x", x))
			b := caps(byt.Bytes())
			clipboard.WriteAll(string(b))
			fmt.Println(string(b))
		case *pp:
			x := md5.Sum(input)
			byt.WriteString(fmt.Sprintf("%x", x))
			b := caps(byt.Bytes())
			var b2 [20]byte
			copy(b2[:], b)
			clipboard.WriteAll(string(b2[0:20]))
			fmt.Println(string(b2[0:20]))
		default:
			// Default hash is md5
			x := md5.Sum(input)
			byt.WriteString(fmt.Sprintf("%x", x))
			b := caps(byt.Bytes())
			clipboard.WriteAll(string(b))
			fmt.Println(string(b))
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

func caps(b []byte) []byte {
	if *CAPITAL {
		for c := len(b) - 1; c > 0; c-- {
			if isLetter(b[c]) {
				b[c] = toUpper(b[c])
				break
			}
		}
	}
	return b
}
