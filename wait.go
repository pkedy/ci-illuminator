package illuminator

import (
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/pkedy/ci-illuminator/pipeline"
)

type Wait struct {
	duration time.Duration
}

func (p *Wait) Process(done chan pipeline.Signal) {
	log.Debug("Waiting")
	time.Sleep(p.duration)
	done <- pipeline.Next
}
