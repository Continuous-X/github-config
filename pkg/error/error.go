package error

import "github-config/pkg/output"

func FailHandleCommand(err error)  {
	if err != nil {
		output.PrintCliError(err)
	}
}