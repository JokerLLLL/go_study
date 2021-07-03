package main

import (
	"github.com/jokerl/go_study/maz/walk"
)

func main() {
	m := walk.GetMap("./maz/map2.in")
	successMap := walk.Walk(m)
	walk.Show(m, successMap)
}