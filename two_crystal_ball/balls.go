package twocrystalball

import "math"

func Search(breaks []bool) int {
	jumpBy := int(math.Sqrt(float64(len(breaks))))
	jumped := jumpBy

	for jumped < len(breaks) {
		if breaks[jumped] {
			break
		}

		jumped += jumpBy
	}

	jumped -= jumpBy

	for remaining := 0; remaining < jumpBy && jumped < len(breaks); {
		if breaks[jumped] {
			return jumped
		}

		remaining += 1
		jumped += 1
	}

	return -1
}
