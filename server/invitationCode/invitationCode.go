package invitationCode

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

const maxId = 9999

var (
	ids = [maxId + 1]int{}
	end = maxId
	mu sync.Mutex
)

func init() {
	for i := range ids {
		ids[i] = i
	}
	rand.Seed(time.Now().UnixNano())
}

// Get returns a unique random 4-digit invitation code
// it gives an error if all invitation codes have been used
func Get() (int, error) {
	mu.Lock()
	defer mu.Unlock()

	if end < 0 {
		return 0, errors.New("no IDs available")
	}

	i := rand.Intn(end +1)
	id := ids[i]
	ids[i], ids[end] = ids[end], ids[i]
	end--
	return id, nil
}

// Return gives back an invitation code for future use
func Return(id int) {
	mu.Lock()
	defer mu.Unlock()

	end++
	ids[end] = id
}