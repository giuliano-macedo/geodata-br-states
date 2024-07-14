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

// uses the Chebyshev's method to compute:
//
//	xcos := [4]float64{
//		math.Cos(2*x),
//		math.Cos(4*x),
//		math.Cos(6*x),
//		math.Cos(8*x),
//	}
//	 xsin:= [4]float64{
//		math.Sin(2*x),
//		math.Sin(4*x),
//		math.Sin(6*x),
//		math.Sin(8*x),
//	}
//
// such that only 4 trig functions are computed, instead of 8.
// reference: https://trans4mind.com/personal_development/mathematics/trigonometry/multipleAnglesRecursiveFormula.htm#Recursive_Formula
func fastCosSinDoubles(x float64) (xcos, xsin [4]float64) {
	var (
		sinx  = math.Sin(x)
		sin2x = math.Sin(2 * x)
		cosx  = math.Cos(x)
		cos2x = math.Cos(2 * x)
	)

	//2x
	xsin[0] = sin2x
	xcos[0] = cos2x

	aSin, bSin := sinx, sin2x
	aCos, bCos := cosx, cos2x

	//4x
	aSin, bSin = bSin, 2*cosx*bSin-aSin
	aCos, bCos = bCos, 2*cosx*bCos-aCos
	aSin, bSin = bSin, 2*cosx*bSin-aSin
	aCos, bCos = bCos, 2*cosx*bCos-aCos

	xsin[1] = aSin
	xcos[1] = aCos

	//6x
	aSin, bSin = bSin, 2*cosx*bSin-aSin
	aCos, bCos = bCos, 2*cosx*bCos-aCos
	aSin, bSin = bSin, 2*cosx*bSin-aSin
	aCos, bCos = bCos, 2*cosx*bCos-aCos

	xsin[2] = aSin
	xcos[2] = aCos

	//8x
	aSin, bSin = bSin, 2*cosx*bSin-aSin
	aCos, bCos = bCos, 2*cosx*bCos-aCos
	aSin = 2*cosx*bSin - aSin
	aCos = 2*cosx*bCos - aCos

	xsin[3] = aSin
	xcos[3] = aCos

	return
}

// computes:
//
//	xcosh := [4]float64{
//			math.Cosh(2*x),
//			math.Cosh(4*x),
//			math.Cosh(6*x),
//			math.Cosh(8*x),
//		}
//	xsinh:= [4]float64{
//			math.Sinh(2*x),
//			math.Sinh(4*x),
//			math.Sinh(6*x),
//			math.Sinh(8*x),
//		}
//
// such that only 1 exp functions is computed, instead of 8 trig functions
func fastCoshSinhDoubles(x float64) (xcosh, xsinh [4]float64) {
	ex2 := math.Exp(2 * x)

	xsinh[0] = (ex2 - (1 / ex2)) / 2
	xcosh[0] = (ex2 + (1 / ex2)) / 2

	ex4 := ex2 * ex2

	xsinh[1] = (ex4 - (1 / ex4)) / 2
	xcosh[1] = (ex4 + (1 / ex4)) / 2

	ex6 := ex2 * ex2 * ex2

	xsinh[2] = (ex6 - (1 / ex6)) / 2
	xcosh[2] = (ex6 + (1 / ex6)) / 2

	ex8 := ex4 * ex4

	xsinh[3] = (ex8 - (1 / ex8)) / 2
	xcosh[3] = (ex8 + (1 / ex8)) / 2

	return
}
