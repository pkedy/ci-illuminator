package illuminator

import (
	"time"

	"github.com/pkedy/ci-illuminator/pipeline"
)

const (
	duration = time.Millisecond * 1500
	pause    = time.Millisecond * 1500
)

var (
	StateBuilding = pipeline.New(
		true,
		&Color{
			color:    &BrightCyan,
			duration: duration,
		},
		&Wait{
			duration: pause,
		},
		&Color{
			color:    &DimCyan,
			duration: duration,
		},
		&Wait{
			duration: pause,
		},
	)

	StateInitial = pipeline.New(
		false,
		&Color{
			color:    &WarmNeutral,
			duration: duration,
		},
	)

	StateSuccess = pipeline.New(
		false,
		&Color{
			color:    &Green,
			duration: duration,
		},
		&Wait{
			duration: time.Second * 60,
		},
		&Color{
			color:    &WarmNeutral,
			duration: duration,
		},
	)

	StateFailure = pipeline.New(
		false,
		&Color{
			color:    &Red,
			duration: duration,
		},
	)
)
