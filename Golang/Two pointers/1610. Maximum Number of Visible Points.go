package main

import "math"

func visiblePoints(points [][]int, angle int, location []int) int {


	max:=0
	for _,i := range points{
		cur :=0

		for _, j := range points {
			if calangle(i, j, location) <= float64(angle) {
				cur += 1
			}
		}

		if cur>max{
			max = cur
		}
	}
	return max
}

func calangle(pointA []int,pointB []int,loc []int) float64{

	if pointA[0]==pointB[0]&&pointA[1]==pointB[1]{
		return 0.0
	}
	dx1 := pointA[0] - loc[0];
	dy1 :=  pointA[1] - loc[1];
	if dx1==0&&dy1==0{
		dx1 =1
		dy1 =1
	}
	dx2 := pointB[0]- loc[0];

	dy2 := pointB[1] - loc[1];

	c :=math.Sqrt(float64(dx1 * dx1 + dy1 * dy1)) * math.Sqrt(float64(dx2 * dx2 + dy2 * dy2))

	if c == 0 {
		return 0.0
	}

	ang:= math.Acos(float64(dx1 * dx2 + dy1 * dy2) / c)*180/math.Pi;
	return ang
}