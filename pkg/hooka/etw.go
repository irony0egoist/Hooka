package hooka

import "github.com/irony0egoist/Hooka/evasion"

func PatchEtw() error {
	return evasion.PatchEtw()
}

func PatchEtw2() error {
	return evasion.PatchEtw2()
}
