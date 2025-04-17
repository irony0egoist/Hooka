package hooka

import (
  "github.com/irony0egoist/Hooka/evasion"
)

func BlockDLLs() error {
  return evasion.BlockDLLs()
}

func CreateProcessBlockDLLs(cmd string) error {
  return evasion.CreateProcessBlockDLLs(cmd)
}

