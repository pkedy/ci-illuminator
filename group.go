package illuminator

import (
	"net"
	"os"
	"time"

	"github.com/2tvenom/golifx"
	log "github.com/Sirupsen/logrus"

	"github.com/pkedy/ci-illuminator/lights"
)

var group *lights.Group

func init() {
	bcastIP := os.Getenv("CII_BCAST_IP")
	if len(bcastIP) > 0 {
		golifx.SetBroadcastAddress(net.ParseIP(bcastIP))
	}
}

func DiscoverBulbs() {
	//Lookup all bulbs
	bulbs, _ := golifx.LookupBulbs()

	for _, bulb := range bulbs {
		label, err := bulb.GetLabel()
		if err == nil {
			log.Infof("Bulb found %s", label)
		}
	}

	group = lights.NewGroup(bulbs...)
}

func SetPowerState(on bool) error {
	return group.SetPower(on)
}

func SetColor(hsbk *golifx.HSBK, duration time.Duration) error {
	return group.SetColor(hsbk, duration)
}
