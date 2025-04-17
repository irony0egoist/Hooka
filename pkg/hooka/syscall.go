package hooka

import "github.com/irony0egoist/Hooka/evasion"

func Syscall(callid uint16, argh ...uintptr) (uint32, error) {
	return evasion.Syscall(callid, argh...)
}


