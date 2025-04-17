package hooka

import "github.com/irony0egoist/Hooka/utils"

func GetShellcodeFromUrl(url string) ([]byte, error) {
	return utils.GetShellcodeFromUrl(url)
}

func GetShellcodeFromFile(file string) ([]byte, error) {
	return utils.GetShellcodeFromFile(file)
}

func CalcShellcode() []byte {
	return utils.CalcShellcode()
}

func CheckHighPrivs() (bool, error) {
  return utils.CheckHighPrivs()
}


