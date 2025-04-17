package hooka

import "github.com/irony0egoist/Hooka/evasion"

func GetSysId(funcname string) (uint16, error) {
	return evasion.GetSysId(funcname)
}

