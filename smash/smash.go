package smash

import (
	"io"
    "bufio"
    "sync"
)

type word string
    
type Counter struct {
    entries map[uint32]uint
    mux sync.Mutex
}

func (c *Counter) Increment(key uint32) {
    c.mux.Lock()
    defer c.mux.Unlock()
    c.entries[key]++
}

// {"foo", returnZero, map[uint32]uint{0: 1}},
// The call "Smash(strings.NewReader("foo"), returnZero)" should return {0:1}
func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
    c := Counter{
        entries: make(map[uint32]uint),
    }

    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    
    // For each word in r,
    for scanner.Scan() {
        str := word(scanner.Text())
        
        // Call smasher() on that word
        to_append := uint32(smasher(str))

        // Call smasher() on the word
        // Create an entry in the dict whose KEY is the return val of smasher(word)
        // and the VALUE is 1.
        // If that entry already exists, increment it
        c.Increment(to_append)
    }
	
    return c.entries
}
