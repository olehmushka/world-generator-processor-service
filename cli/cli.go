package cli

import (
	"errors"
	"fmt"
)

func Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("cli.Execute: args should not be zero size")
	}
	switch args[0] {
	case MigrateUpCommand:
		if err := runMigrateUpCommand(); err != nil {
			return err
		}
	case MigrateDownCommand:
		if err := runMigrateDownCommand(); err != nil {
			return err
		}
	case HTTPServerRunCommand:
		if err := runHTTPServerCommand(); err != nil {
			return err
		}
	case SyncDataCommand:
		if err := runSyncDataCommand(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cli.Execute: not found command = %s", args[0])
	}

	return nil
}
