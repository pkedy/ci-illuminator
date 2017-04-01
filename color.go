package illuminator

import (
	"time"

	"github.com/2tvenom/golifx"
	log "github.com/Sirupsen/logrus"

	"github.com/pkedy/ci-illuminator/pipeline"
)

type Color struct {
	color    *golifx.HSBK
	duration time.Duration
}

func (p *Color) Process(done chan pipeline.Signal) {
	log.Debug("Setting color")
	group.SetColor(p.color, p.duration)
	done <- pipeline.Next
}
