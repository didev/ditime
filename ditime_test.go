package ditime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/digital-idea/ditime"
)

func TestToFullTime(t *testing.T) {
	n := time.Now()
	// travisCI 에서는 UTC형식인 "2018-06-18T19:00:00Z" 라고 RFC3339 형식의 시간이 표기된다.
	timeZone := n.Format(time.RFC3339)[19:len(n.Format(time.RFC3339))]
	cases := []struct {
		in   string
		mode string
		want string
		err  error
	}{{
		in:   "0618",
		mode: "start",
		want: fmt.Sprintf("%04d-06-18T10:00:00%s", n.Year(), timeZone),
		err:  nil,
	}, {
		in:   "0618",
		mode: "end",
		want: fmt.Sprintf("%04d-06-18T19:00:00%s", n.Year(), timeZone),
		err:  nil,
	}, {
		in:   "06.18",
		mode: "end",
		want: fmt.Sprintf("%04d-06-18T19:00:00%s", n.Year(), timeZone),
		err:  nil,
	}, {
		in:   "06.19.",
		mode: "end",
		want: fmt.Sprintf("%04d-06-19T19:00:00%s", n.Year(), timeZone),
		err:  nil,
	}, {
		in:   "06/18",
		mode: "start",
		want: fmt.Sprintf("%04d-06-18T10:00:00%s", n.Year(), timeZone),
		err:  nil,
	}, {
		in:   "2018-06-18",
		mode: "start",
		want: fmt.Sprintf("2018-06-18T10:00:00%s", timeZone),
		err:  nil,
	}, {
		in:   "2018-06-18T18:45:23+09:00",
		mode: "current",
		want: "2018-06-18T18:45:23+09:00",
		err:  nil,
	}, {
		in:   "2018-06-18T18:45:23+09:00",
		mode: "start",
		want: fmt.Sprintf("2018-06-18T10:00:00%s", timeZone),
		err:  nil,
	}, {
		in:   "2018-06-18T18:45:23+09:00",
		mode: "end",
		want: fmt.Sprintf("2018-06-18T19:00:00%s", timeZone),
		err:  nil,
	}}
	for _, c := range cases {
		result, err := ditime.ToFullTime(c.mode, c.in)
		if result != c.want {
			t.Fatalf("TestToFullTime(%v,%v): 얻은 값 %v, 원하는 값 %v, 에러 %v", c.mode, c.in, result, c.want, err)
		}
	}
}

func TestRegexYYYYMMDD(t *testing.T) {
	cases := []struct {
		time string
		want bool
	}{{
		time: "2019-11-19",
		want: true,
	}, {
		time: "2019/11/19",
		want: true,
	}, {
		time: "2019/1/13",
		want: true,
	}, {
		time: "2019/1/1",
		want: true,
	}, {
		time: "2019.11.19.",
		want: true,
	}, {
		time: "2019,11,19",
		want: true,
	}, {
		time: "2019,11,19,",
		want: true,
	}, {
		time: "2019-1-19",
		want: true,
	}, {
		time: "2019-1-1",
		want: true,
	}, {
		time: "2019. 1. 1.",
		want: true,
	}, {
		time: "2019, 1, 1,",
		want: true,
	}, {
		time: "2019, 1, 1",
		want: true,
	}, {
		time: "2019. 1. 1",
		want: true,
	},
	}
	for _, c := range cases {
		reg := ditime.RegexpYYYYMMDD
		if reg.MatchString(c.time) != c.want {
			t.Fatalf("Test_regexYYYYMMDD: 입력 값 %v, 원하는 값 %v, 결과 %v", c.time, c.want, reg.MatchString(c.time))
		}
	}
}

func TestRegexMMDD(t *testing.T) {
	cases := []struct {
		time string
		want bool
	}{{
		time: "11-19",
		want: true,
	}, {
		time: "11/19",
		want: true,
	}, {
		time: "1/13",
		want: true,
	}, {
		time: "1/1",
		want: true,
	}, {
		time: "06.19.",
		want: true,
	}, {
		time: "11,19",
		want: true,
	}, {
		time: "11,19,",
		want: true,
	}, {
		time: "1-19",
		want: true,
	}, {
		time: "1-1",
		want: true,
	}, {
		time: "1. 1.",
		want: true,
	}, {
		time: "1, 1,",
		want: true,
	}, {
		time: "1, 1",
		want: true,
	}, {
		time: "1. 1",
		want: true,
	},
	}
	for _, c := range cases {
		reg := ditime.RegexpMMDD
		if reg.MatchString(c.time) != c.want {
			t.Fatalf("Test_regexMMDD: 입력 값 %v, 원하는 값 %v, 결과 %v", c.time, c.want, reg.MatchString(c.time))
		}
	}
}
