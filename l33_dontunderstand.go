/*
Cさんはチワワが大好きです。そこでCさんは、文字列にもチワワを見出すことにしました。
Cさんによれば、ある文字列に 'c', 'w', 'w' がこの順で含まれるとき、その文字列を「チワワ列」であるといいます。
Cさんは小さなチワワが好きなので、できるだけ長さの小さいチワワ列を見つけたいです。
文字列 SS が与えられるので、その連続した部分文字列のうちチワワ列となるものの最小の長さを求めてください。
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	l := -1
	reg := regexp.MustCompile(`^c[^c]*?w.*?w`)
	for {
		index := strings.Index(S, "c")
		if index == -1 {
			break
		}
		S = S[index:]
		s := reg.FindString(S)
		slen := len(s)
		if slen > 0 && (slen < l || l == -1) {
			l = slen
		}
		// この部分の意味がわかんない
		S = S[1:]
	}
	fmt.Println(l)
}
