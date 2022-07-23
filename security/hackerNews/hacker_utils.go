package hackerNews

import (
	"math"
	"security/proto/hot"
	"strconv"
	"time"
)

func MakeHot(hot *hot.HotRequest) float64 {
	t, err := strconv.ParseInt(hot.Time, 10, 64)
	if err != nil {
	}
	creteTime := time.Unix(t, 0)
	nowTime := time.Now()
	dif := nowTime.Sub(creteTime).Hours() + 2
	number := hot.Number + 100 - 1
	newhot := float64(number) / math.Pow(dif, 1.5)
	return newhot
}
