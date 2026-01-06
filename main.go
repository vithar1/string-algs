package main

import (
	"fmt"
	"math/rand"
	"strings"

	"string-algs/algs"
)

func main() {
	// fmt.Println(algs.LCP())
	// s := "abcabbx"
	// tree := algs.BuildUkkonen(s)
	// algs.PrintSTree(tree, "", true, s)
	// res := algs.Match("babbababaababbaa", "bbaa")
	// res := algs.Match("abcabd", "b")
	// res := algs.Match("babaababbaa", "bbaa")
	reversTest()
	// res := algs.Match("mississippi", "sipp")
	// fmt.Println(res)
}

func reversTest() {
	sl := 2000
	for range 2000 {
		s := randStr(sl)
		from := rand.Intn(sl)
		offset := rand.Intn(sl - from)
		pattern := s[from:from+offset]
		res := algs.Match(s, pattern)
		if s[res:res+offset] != s[from:from+offset] {
			fmt.Println(s, s[from:from+offset])
			fmt.Println("error")
			break
		}
	}
	if algs.Match("babbababaababbaa", "bbaa") != 12 {
		fmt.Println("error")
	}
	if algs.Match("abcabd", "b") != 1 {
		fmt.Println("error")
	}
	if algs.Match("babaababbaa", "bbaa") != 7 {
		fmt.Println("error")
	}
	if algs.Match("mississippi", "sipp") != 6 {
		fmt.Println("error")
	}
}

func randStr(length int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var b strings.Builder
	b.Grow(length)
	for range length{
		b.WriteByte(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
