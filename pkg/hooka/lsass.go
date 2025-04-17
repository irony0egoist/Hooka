package hooka

import "github.com/irony0egoist/Hooka/evasion"

func DumpLsass(output_file string) error {
	return evasion.DumpLsass(output_file)
}
