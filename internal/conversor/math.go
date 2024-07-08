package conversor

import "math"

func sinSqrd(x float64) float64 {
	ans := math.Sin(x)
	return ans * ans
}

func sinCubed(x float64) float64 {
	ans := math.Sin(x)
	return ans * ans * ans
}

func cosCubed(x float64) float64 {
	ans := math.Cos(x)
	return ans * ans * ans
}

const (
	piRatio        = 180 / math.Pi
	inversePiRatio = math.Pi / 180
)

func RadToDeg(x float64) float64 {
	return x * piRatio
}

func DegToRad(x float64) float64 {
	return x * inversePiRatio
}
