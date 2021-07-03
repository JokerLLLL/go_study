package walk

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map [][]int

type Point struct {
	i int
	j int
}

type PointStep struct {
	Point
	s int
}

func addPoint(po1 PointStep, po2 Point, s int) PointStep {
	return PointStep{Point{po1.i + po2.i, po1.j + po2.j}, s + 1}
}

func Walk(mapInfo Map) Map {
	endI, endJ := len(mapInfo), len(mapInfo[0])
	// int steps
	var steps = make([][]int, endI)
	for i := range steps {
		steps[i] = make([]int, endJ)
	}
	fmt.Println(mapInfo)
	fmt.Println(steps)
	// position
	var position = [4]Point{
		{0, -1}, {-1, 0}, {0, 1}, {1, 0},
	}
	// 走map
	var startPoint = Point{0, 0}
	var endPoint = Point{endI - 1, endJ - 1}
	var Queue []PointStep
	Queue = append(Queue, PointStep{startPoint, 1})
	fmt.Println(endPoint)

	for {
		// 没有Queue说明走不下去了
		if len(Queue) == 0 {
			break
		}
		nowPoint := Queue[0]
		Queue = Queue[1:]
		// 记录step
		steps[nowPoint.i][nowPoint.j] = nowPoint.s
		// 判断这个点是否是重点
		if nowPoint.i == endPoint.i && nowPoint.j == endPoint.j {
			break
		}
		// 判断nowPoint四周
		for _, p := range position {
			nextPoint := addPoint(nowPoint, p, nowPoint.s)
			// 1. i j 是否越界
			if nextPoint.i < startPoint.i || nextPoint.i > endPoint.i {
				continue
			}
			if nextPoint.j < startPoint.j || nextPoint.j > endPoint.j {
				continue
			}
			fmt.Println(nextPoint)
			// 2. i j 是否是墙
			if mapInfo[nextPoint.i][nextPoint.j] == 1 {
				continue
			}
			// 3. i j 是否走过
			if steps[nextPoint.i][nextPoint.j] > 0 {
				continue
			}
			Queue = append(Queue, nextPoint)
		}
	}
	return steps
}


func GetMap(mapS string) Map {
	// first 行 * 列
	file, e := os.Open(mapS)
	if e != nil {
		panic(e)
	}
	var size string
	n, e := fmt.Fscanln(file, &size)
	if e != nil {
		panic(e.Error() + fmt.Sprintf("; get file line:%d", n))
	}
	split := strings.Split(size, ",")
	if len(split) != 2 {
		panic("not two params")
	}
	line, _ := strconv.Atoi(split[0])
	column, _ := strconv.Atoi(split[1])

	points := make([][]int, line)
	for i := range points {
		var lines string
		fmt.Fscanln(file, &lines)
		s := strings.Split(lines, "")
		points[i] = make([]int, column)
		for j :=range points[i] {
			atoi, _ := strconv.Atoi(s[j])
			points[i][j] = atoi
		}
	}
	return points
}

func Show(p Map, r Map)  {
	for i,l := range r {
		for j,d := range l {
			if p[i][j] == 1 {
				fmt.Printf("%4s","x")
			}else{
				fmt.Printf("%4d", d)
			}
		}
		fmt.Println()
	}
}