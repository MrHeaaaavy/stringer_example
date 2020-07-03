package main

//go:generate stringer -type=Sport --linecomment
type Sport int

const (
	Football    Sport = 1 // 足球
	Basketball  Sport = 2 // 篮球
	TableTennis Sport = 3 // 乒乓球
)
