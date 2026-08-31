package engine

import "hash/fnv"

var hash = fnv.New32a()

func getHash(s string) uint32 {
	hash.Reset()
	hash.Write([]byte(s))
	return hash.Sum32()
}
