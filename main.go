package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"time"
)

var style string
var length int
var unique bool

var number = "0123456789"
var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lower = "abcdefghijklmnopqrstuvwxyz"
var special = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

var mask = uint8(0)

func init() {
	flag.StringVar(&style, "s", "", "密码样式")
	flag.IntVar(&length, "l", 8, "密码长度")
	flag.BoolVar(&unique, "u", false, "是否允许重复字符")
	flag.Parse()
}

func getRandomChar(s string) byte {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(s))))
	if err != nil {
		panic("生成密码失败: " + err.Error())
	}
	return s[i.Int64()]
}

func main() {
	if style == "" {
		mask = mask | 14
	} else {
		for _, b := range []byte(style) {
			if b >= 48 && b <= 57 {
				mask = mask | 8
			} else if b >= 65 && b <= 90 {
				mask = mask | 4
			} else if b >= 97 && b <= 122 {
				mask = mask | 2
			} else {
				mask = mask | 1
			}
		}
	}

	chars := ""
	result := make([]byte, 0, length)
	exist := make(map[byte]bool, length)

	if mask&8 != 0 {
		c := getRandomChar(number)
		result = append(result, c)
		length--
		chars += number
		exist[c] = true
	}

	if mask&4 != 0 {
		c := getRandomChar(upper)
		result = append(result, c)
		length--
		chars += upper
		exist[c] = true
	}

	if mask&2 != 0 {
		c := getRandomChar(lower)
		result = append(result, c)
		length--
		chars += lower
		exist[c] = true
	}

	if mask&1 != 0 {
		c := getRandomChar(special)
		result = append(result, c)
		length--
		chars += special
		exist[c] = true
	}

	for length > 0 {
		c := getRandomChar(chars)
		if unique {
			if !exist[c] {
				result = append(result, c)
				length--
				exist[c] = true
			}
		} else {
			result = append(result, c)
			length--
		}
	}

	mathRand.Seed(time.Now().UnixNano())

	mathRand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	fmt.Println(string(result))
}
