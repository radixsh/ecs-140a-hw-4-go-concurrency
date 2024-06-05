package smash

import (
	"io"
    "bufio"
)

type word string

// {"foo", returnZero, map[uint32]uint{0: 1}},
// The call "Smash(strings.NewReader("foo"), returnZero)" should return {0:1}
func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
	m := make(map[uint32]uint)

    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    
    // For each word in r,
    for scanner.Scan() {
        // Call smasher() on that word
        thing := strings.Fields(scanner.Scan()
        to_append := smasher(word)
        fmt.Println(to_append)

        // Append the value that smasher() returns to the map m
        m = append(m, to_append)
    }


	return m
}
