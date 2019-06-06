package ditime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/didev/ditime"
)

func Test_ToFullTime(t *testing.T) {
	n := time.Now()
	timeOffset := n.Format(time.RFC3339)[19:len(n.Format(time.RFC3339))]
	cases := []struct {
		in   string
		mode string
		want string
		err  error
	}{{
		in:   "0618",
		mode: "start",
		want: fmt.Sprintf("%04d-06-18T10:00:00%s", n.Year(), timeOffset),
		err:  nil,
	}, {
		in:   "0618",
		mode: "end",
		want: fmt.Sprintf("%04d-06-18T19:00:00%s", n.Year(), timeOffset),
		err:  nil,
	}, {
		in:   "06/18",
		mode: "start",
		want: fmt.Sprintf("%04d-06-18T10:00:00%s", n.Year(), timeOffset),
		err:  nil,
	}, {
		in:   "2018-06-18",
		mode: "start",
		want: fmt.Sprintf("2018-06-18T10:00:00%s", timeOffset),
		err:  nil,
	}, {
		in:   "2018-06-18T18:45:23+09:00",
		mode: "current",
		want: "2018-06-18T18:45:23+09:00",
		err:  nil,
	}, {
		in:   "2018-06-18T18:45:23+09:00",
		mode: "start",
		want: fmt.Sprintf("2018-06-18T10:00:00%s", timeOffset),
		err:  nil,
	}, {
		in:   "2018-06-18T18:45:23+09:00",
		mode: "end",
		want: fmt.Sprintf("2018-06-18T19:00:00%s", timeOffset),
		err:  nil,
	}}
	for _, c := range cases {
		result, err := ditime.ToFullTime(c.mode, c.in)
		if result != c.want {
			t.Fatalf("Test_ToFullTime(%v,%v): 얻은 값 %v, 원하는 값 %v, 에러 %v", c.mode, c.in, result, c.want, err)
		}
	}
}
