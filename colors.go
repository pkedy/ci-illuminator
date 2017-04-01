package illuminator

import (
	"github.com/2tvenom/golifx"
)

var (
	BrightCyan = golifx.HSBK{
		Hue:        32511,
		Saturation: 65535,
		Brightness: 65535,
		Kelvin:     3500,
	}

	DimCyan = golifx.HSBK{
		Hue:        32511,
		Saturation: 65535,
		Brightness: 65535 / 3,
		Kelvin:     3500,
	}

	Green = golifx.HSBK{
		Hue:        21504,
		Saturation: 65535,
		Brightness: 65535,
		Kelvin:     3500,
	}

	Red = golifx.HSBK{
		Hue:        0,
		Saturation: 65535,
		Brightness: 65535,
		Kelvin:     3500,
	}

	WarmNeutral = golifx.HSBK{
		Hue:        0,
		Saturation: 0,
		Brightness: 65535,
		Kelvin:     3200,
	}
)
