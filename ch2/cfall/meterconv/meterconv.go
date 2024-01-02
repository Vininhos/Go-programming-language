package meterconv

type Feet float64
type Meter float64

func CFtM(f Feet) Meter {
	return Meter(f / 30.48)
}

func CMtF(f Meter) Feet {
	return Feet(f * 30.48)
}
