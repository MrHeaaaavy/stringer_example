package main

//go:generate stringer -type=People --linecomment
type People int

const (
	Bob  People = 1 // 鲍勃
	Lily People = 2 // 莉莉
)
