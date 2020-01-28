package clientproc

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
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

// Lockfile is information determined from the process "lockfile"
type Lockfile struct {
	Port  int
	Token string
}

// LockfileFromClient determines api information from a process path.
func LockfileFromClient(processPath string) (*Lockfile, error) {
	dirPath := filepath.Dir(processPath)
	lockfilePath := filepath.Join(dirPath, "lockfile")
	lockfileContents, err := ioutil.ReadFile(lockfilePath)
	if err != nil {
		return nil, err
	}
	lockfilePts := strings.Split(string(lockfileContents), ":")
	if len(lockfilePts) < 5 {
		return nil, errors.Errorf(
			"expected at least 5 lockfile parts, %d found",
			len(lockfilePts),
		)
	}
	var l Lockfile
	l.Port, err = strconv.Atoi(lockfilePts[2])
	if err != nil {
		return nil, errors.Wrap(err, "parse port from lockfile")
	}
	l.Token = lockfilePts[3]
	return &l, nil
}
