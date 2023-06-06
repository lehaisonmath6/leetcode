package main

import "fmt"

// y = ax + b
func getLine(pointA, pointB []int) (a, b float64, isHoziontal bool, k int) {

	if pointA[0] == pointB[0] {
		return 1, 0, true, pointA[0]
	}
	// if pointA[1] == pointB[1] {
	// 	return 0, float64(pointB[1])
	// }
	a = float64(pointA[1]-pointB[1]) / float64(pointA[0]-pointB[0])
	b = float64(pointB[1]) - a*float64(pointB[0])
	return
}

func checkPointInLine(a, b float64, point []int) bool {
	c := float64(point[0])*a + b
	if int(c) == point[1] {
		return true
	}
	return false
}

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}
	a, b, isHoziontal, k := getLine(coordinates[0], coordinates[1])
	if isHoziontal {
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][0] != k {
				return false
			}
		}
		return true
	}
	for i := 2; i < len(coordinates); i++ {
		if !checkPointInLine(a, b, coordinates[i]) {
			return false
		}
	}
	return true
}

func main() {
	coordinates := [][]int{
		{0, 0}, {0, 1}, {0, -1},
	}
	fmt.Println(checkStraightLine(coordinates))
}
