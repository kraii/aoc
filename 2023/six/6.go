package six

import "math"

/*
Time:        40     70     98     79
Distance:   215   1051   2147   1005
*/
func solveA() int {
	times := []int{40, 70, 98, 79}
	distances := []int{215, 1051, 2147, 1005}
	total := 1
	for i := 0; i < len(times); i++ {
		t, d := times[i], distances[i]
		total *= countWins(t, d)
	}
	return total
}

func countWins(t int, d int) int {
	wins := 0
	for acc := 1; acc < t; acc++ {
		distance := (t - acc) * acc
		if distance > d {
			wins++
		}
	}
	return wins
}

func solveBruteForce(t int, d int) int {
	return countWins(t, d)
}

/*
p: press time
t: total time
d: distance
p(t-p) > d
p^2 + tp > d
p^2 + tp + d < 0

It's a quadratic inequality
*/
func solveB(t int, d int) int {
	determinant := t*t - 4*d
	if determinant <= 0 {
		return 0
	}
	determinantRoot := math.Sqrt(float64(determinant))
	tf := float64(t)
	a := (tf - determinantRoot) / 2
	b := (tf + determinantRoot) / 2

	minRange := int(math.Ceil(a))
	maxRange := int(math.Floor(b))

	return maxRange - minRange + 1
}
