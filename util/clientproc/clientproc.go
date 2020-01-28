package clientproc

import (
	"strings"
	"github.com/pkg/errors"
		ps "github.com/shirou/gopsutil/process"
)

// CheckIsLeagueExecutable returns if the string might be the league executable.
func CheckIsLeagueExecutable(name string) bool {
	return strings.HasSuffix(strings.TrimSpace(strings.ToLower(name)), "leagueclient.exe")
}

// LocateLeagueClient attempts to locate the league of legends client returning the exepath and proc.
// May return nil and empty string if not found.
func LocateLeagueClient() (*ps.Process, string, error) {
	processes, err := ps.Processes()
	if err != nil {
		return nil, "", errors.Wrap(err, "cannot list processes")
	}
	var sproc *ps.Process
	var sprocExe string
	for _, proc := range processes {
		exePath, err := proc.Exe()
		if err != nil {
			// Happens frequently
			// le.WithError(err).Warnf("cannot determine exe path for pid %v", proc.Pid)
			continue
		}
		if CheckIsLeagueExecutable(exePath) {
			sproc = proc
			sprocExe = exePath
		}
	}
	return sproc, sprocExe, nil
}