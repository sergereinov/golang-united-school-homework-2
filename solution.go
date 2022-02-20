package square

import "math"

type figure int

const (
	SidesCircle   figure = 0
	SidesTriangle figure = 3
	SidesSquare   figure = 4
)

func CalcSquare(sideLen float64, sidesNum figure) float64 {
	switch sidesNum {
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	case SidesTriangle:
		return sideLen * sideLen * math.Sqrt(3) / 4
	case SidesSquare:
		return sideLen * sideLen
	}
	return 0
}
