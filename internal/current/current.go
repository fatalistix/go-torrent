package current

import (
	"crypto/sha1"
	"time"
)

func generateId() [20]byte {
	result := [20]byte{}
	timeSum := sha1.Sum([]byte(time.Now().String()))
	tmpStr := "-GT0001-" + string(timeSum[:12])
	copy(result[:], tmpStr)
	return result
}
