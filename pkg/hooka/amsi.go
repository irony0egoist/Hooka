package hooka

import "github.com/irony0egoist/Hooka/evasion"

func PatchAmsi() error {
	return evasion.PatchAmsi()
}

func PatchAmsi2() error {
	return evasion.PatchAmsi2()
}
