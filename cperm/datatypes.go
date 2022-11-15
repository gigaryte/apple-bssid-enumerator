package cperm

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lcperm -Wl,-rpath,$ORIGIN
#include "cperm.h"
*/
import "C"
import (
	"fmt"
	"github.com/gigaryte/apple-bssid-enumerator/constants"
	"math/rand"
)

var (
	OUIInfos []*OUIInfo
)

type OUIInfo struct {
	OUI   string
	CPerm CPerm
}

func (ou *OUIInfo) Next() (uint32, int) {
	var Cint C.int
	var Cval C.uint32_t

	//Cint is the return code, 0 tells us no error (more permutations to come),
	//and PERM_END (-5) means stop and is used in loop iteration
	Cint = C.cperm_next(ou.CPerm.CPerm, &Cval)

	//(next value, return code)
	return uint32(Cval), int(Cint)
}

func (ou *OUIInfo) InitCPerm() {
	randBytes := make([]byte, constants.KEYLEN)
	rand.Read(randBytes)
	key := (*[constants.KEYLEN]byte)(randBytes)
	ou.CPerm.CreateCPerm(constants.NPerOUI, *key)
}

func (ou *OUIInfo) NextMAC() (string, int) {
	permVal, ret := ou.Next()
	if ret == C.PERM_END {
		return "", ret
	}

	upper := permVal >> 16
	middle := (permVal & 0x00ff00) >> 8
	lower := permVal & 0x0000ff

	return ou.OUI + fmt.Sprintf(":%02x:%02x:%02x", upper, middle, lower), ret
}
func (ou *OUIInfo) CPermDestroy() {
	C.cperm_destroy(ou.CPerm.CPerm)
}
