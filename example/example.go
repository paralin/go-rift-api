package main

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/paralin/go-rift-api/util/clientproc"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	le := logrus.NewEntry(log)

	if err := executeExample(le); err != nil {
		le.Fatal(err.Error())
	}
}

func executeExample(le *logrus.Entry) error {
	le.Info("attempting to find the league client")

	sproc, sprocExe, err := clientproc.LocateLeagueClient()
	if err != nil {
		return err
	}
	if sproc == nil {
		return errors.New("could not find league of legends process")
	}

	le.Infof("league client found with pid %v", sproc.Pid)
	le.Infof("path to league client: %v", sprocExe)

	return nil
}
