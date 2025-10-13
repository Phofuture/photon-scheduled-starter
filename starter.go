package scheduleStarter

import (
	"github.com/Phofuture/photon-core-starter/core"
	"github.com/Phofuture/photon-scheduled-starter/schedule"
)

func init() {
	core.RegisterAddModule(schedule.Start)
	core.RegisterShutdownAddModule(schedule.Shutdown)
}
