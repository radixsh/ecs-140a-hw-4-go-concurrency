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
    c.entries[key]++
    c.mux.Unlock()
}

func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
    // From https://go.dev/play/p/Z_Td6Kn_hMT
    var wg sync.WaitGroup
    
    c := Counter{
        entries: make(map[uint32]uint),
    }

    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    
    // For each word in r,
    for scanner.Scan() {
        // Call smasher() on that word
        to_append := uint32(smasher(word(scanner.Text())))
        
        wg.Add(1)

        // Create an entry in the dict whose KEY is the return val of smasher(word)
        // and the VALUE is 1.
        // If that entry already exists, increment it
        go func() {
            c.Increment(to_append)
            wg.Done()
        }()
    }

    wg.Wait()

    return c.entries
}
