package numberic

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	suffixes = []string{"", "k", "m", "b", "t", "p"}
)

func AbbreviateNumber(n int, x int, space string) string {
	magnitude := 0
	f := float64(n)
	for math.Abs(f) > 1000 {
		magnitude++
		f /= 1000.0
	}

	suf := ""
	if magnitude < 6 {
		suf = suffixes[magnitude]
	}

	s := fmt.Sprintf("%."+fmt.Sprintf("%d", x)+"f", f)

	s = regexp.MustCompile(`\.0+$`).ReplaceAllString(s, "")

	return s + space + suf
}

func AbbreviateNumberToInt(s string, space string) (n int) {
	s = strings.ToLower(s)
	magnitude := 0
	num := 0.0
	if space != "" {
		parts := strings.Split(s, space)
		num, _ = strconv.ParseFloat(parts[0], 64)
		if len(parts) == 2 {
			magnitude = getMagnitude(strings.TrimSpace(parts[1]))
		}

	} else {
		re := regexp.MustCompile(`[a-z]`)
		parts := re.Split(s, 2)
		num, _ = strconv.ParseFloat(parts[0], 64)
		magnitude = getMagnitude(re.FindString(s))
	}

	if magnitude > 0 {
		return int(num * math.Pow(1000, float64(magnitude)))
	}

	return int(num)
}

func getMagnitude(s string) int {
	for i, v := range suffixes {
		if s == v {
			return i
		}
	}
	return 0
}
