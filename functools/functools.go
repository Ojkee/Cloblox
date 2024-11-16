package functools

import (
	"fmt"
	"math"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

func TextWidthEx(text string) rl.Vector2 {
	retVal := rl.MeasureTextEx(
		settings.FONT,
		text,
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
	)
	return retVal
}

func GetScaledSlice(
	source []float64,
	low, high float64,
) (retVal []float64, hasPositives, havNegatives bool) {
	if len(source) == 0 {
		return source, false, false
	}
	mmin, mmax := GetMinMax(&source)
	if mmin == mmax {
		if mmin == 0 {
			return source, false, false
		} else if mmin > 0 {
			return source, true, false
		} else {
			return source, false, true
		}
	}
	maxDev := max(math.Abs(mmin), mmax)
	if mmin >= 0 {
		retVal = MulElemVec(source, (high-low)/maxDev)
		return retVal, true, false
	} else if mmax <= 0 {
		retVal = MulElemVec(source, (high-low)/maxDev)
		return retVal, false, true
	}
	retVal = MulElemVec(source, (high-low)/maxDev/2)
	return retVal, true, true
}

func GetMinMax(source *[]float64) (float64, float64) {
	if len(*source) == 0 {
		return 0, 0
	}
	mmin := (*source)[0]
	mmax := (*source)[0]
	for i := 1; i < len(*source)-1; i += 2 {
		if (*source)[i] > (*source)[i+1] {
			if (*source)[i] > mmax {
				mmax = (*source)[i]
			}
			if (*source)[i+1] < mmin {
				mmin = (*source)[i+1]
			}
		} else {
			if (*source)[i] < mmin {
				mmin = (*source)[i]
			}
			if (*source)[i+1] > mmax {
				mmax = (*source)[i+1]
			}
		}
	}
	if len(*source)%2 == 0 {
		lastElem := (*source)[len(*source)-1]
		if lastElem < mmin {
			mmin = lastElem
		} else if lastElem > mmax {
			mmax = lastElem
		}
	}
	return mmin, mmax
}

func MulElemVec(source []float64, scalar float64) []float64 {
	for i := range source {
		source[i] *= scalar
	}
	return source
}

func SplitLine(line string, maxWidth float32) []string {
	retVal := make([]string, 0)
	words := strings.Split(line, " ")
	if len(words) == 0 {
		return nil
	}
	currentLine := words[0]
	for i := 1; i < len(words); i++ {
		nextPossible := fmt.Sprintf("%s %s", currentLine, words[i])
		if rl.MeasureTextEx(
			settings.FONT,
			nextPossible,
			float32(settings.FONT_SIZE),
			settings.FONT_SPACING,
		).X <= maxWidth {
			currentLine = nextPossible
		} else {
			retVal = append(retVal, currentLine)
			currentLine = words[i]
		}
	}
	retVal = append(retVal, currentLine)

	return retVal
}
