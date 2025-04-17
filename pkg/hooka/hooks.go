package hooka

import (
	"github.com/irony0egoist/Hooka/evasion"
)

func DetectHooks() ([]string, error) {
	return evasion.DetectHooks()
}

func IsHooked(func_name string) (bool, error) {
	return evasion.IsHooked(func_name)
}
