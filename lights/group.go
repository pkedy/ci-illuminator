package lights

import (
	"sync"
	"time"

	"github.com/2tvenom/golifx"
)

type Group struct {
	bulbs []*golifx.Bulb
}

func NewGroup(bulbs ...*golifx.Bulb) *Group {
	return &Group{
		bulbs: bulbs,
	}
}

func (g *Group) SetColor(color *golifx.HSBK, duration time.Duration) error {
	var (
		wg       sync.WaitGroup
		err      error
		errMutex sync.Mutex
	)

	millis := uint32(duration.Nanoseconds() / 1000000)
	bulbs := g.bulbs

	if len(bulbs) == 0 {
		return nil
	}

	for _, bulb := range bulbs {
		wg.Add(1)
		go func(bulb *golifx.Bulb) {
			e := bulb.SetColorState(color, millis)
			errMutex.Lock()
			if err == nil && e != nil {
				err = e
			}
			errMutex.Unlock()
			wg.Done()
		}(bulb)
	}

	wg.Wait()
	return err
}

func (g *Group) SetPower(state bool) error {
	var (
		wg       sync.WaitGroup
		err      error
		errMutex sync.Mutex
	)

	bulbs := g.bulbs

	if len(bulbs) == 0 {
		return nil
	}

	for _, bulb := range bulbs {
		wg.Add(1)
		go func(bulb *golifx.Bulb) {
			e := bulb.SetPowerState(state)
			errMutex.Lock()
			if err == nil && e != nil {
				err = e
			}
			errMutex.Unlock()
			wg.Done()
		}(bulb)
	}

	wg.Wait()
	return err
}
