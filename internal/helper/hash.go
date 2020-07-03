package helper

import "hash/fnv"

// TargetToId makes a unique id ot of the target string by hashing it
func TargetToId(target string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(target))
	return h.Sum32()
}
