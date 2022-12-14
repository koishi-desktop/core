package koicmd

import (
	"github.com/samber/do"
	"gopkg.ilharper.com/koi/core/god/daemonproc"
	"gopkg.ilharper.com/koi/core/god/proto"
	"gopkg.ilharper.com/koi/core/koierr"
	"gopkg.ilharper.com/koi/core/logger"
)

func koiCmdStart(i *do.Injector) *proto.Response {
	var err error

	l := do.MustInvoke[*logger.Logger](i)
	command := do.MustInvoke[*proto.CommandRequest](i)
	daemonProc := do.MustInvoke[*daemonproc.DaemonProcess](i)

	l.Debug("Trigger KoiCmd start")

	// Parse command
	instances, ok := command.Flags["instances"].([]any)
	if !ok {
		return proto.NewErrorResult(koierr.ErrBadRequest)
	}

	for _, instanceAny := range instances {
		instance := instanceAny.(string)

		l.Infof("Starting instance %s...", instance)

		err = daemonProc.Start(instance)
		if err != nil {
			return proto.NewErrorResult(koierr.NewErrInternalError(err))
		}
	}

	return proto.NewSuccessResult(nil)
}
