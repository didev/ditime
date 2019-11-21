package ditime

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var regexpShortTime = regexp.MustCompile(`^(0[1-9]|1[012])(0?[1-9]|[12][0-9]|3[01])$`)                                         // 1019
var regexpNormalTime = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])$`)                                // 2016-10-19
var regexpFullTime = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$`) // 2016-10-19T16:41:24+09:00

// RegexpExcelDaynum 레귤러 익스프레션은 1899-12-31 부터 지난날수에 대한 레귤러 익스프레션이다.
var RegexpExcelDaynum = regexp.MustCompile(`^\d{5}$`)

// RegexpMMDD 레귤러 익스프레션은 MM*DD 패턴의 날짜를 처리한다.
var RegexpMMDD = regexp.MustCompile(`^(0?[1-9]|1[012])\D\s?(0?[1-9]|[12][0-9]|3[01])\D?\s?$`) // 10/19, 10-19, 10.10. 1.1 1.1. "1. 1."
// RegexpMDD 레귤러 익스프레션은 MDD 패턴의 날짜를 처리한다.
var RegexpMDD = regexp.MustCompile(`^([1-9])(0?[1-9]|[12][0-9]|3[01])$`) // 101, 110"
// RegexpYYYYMMDD 레귤러 익스프레션은 YYYY*MM*DD* 패턴의 날짜를 처리한다.
var RegexpYYYYMMDD = regexp.MustCompile(`^\d{4}\D\s?(0?[1-9]|1[012])\D\s?(0?[1-9]|[12][0-9]|3[01])\D?\s?$`) // 2016-10-19, 2016/10/19, 2016.10.19, 2016.10.19. 2016. 10. 10.
// RegexpMMDDYYYY 레귤러 익스프레션은 MM*DD*YYYY 패턴의 날짜를 처리한다.
var RegexpMMDDYYYY = regexp.MustCompile(`^(0?[1-9]|1[012])\D\s?(0?[1-9]|[12][0-9]|3[01])\D\s?\d{4}\D?\s?$`) // 10-19-2016, 10/19/2016, 10.19.2016, 10. 19. 2016.
// RegexpSix 레귤러 익스프레션은 YY*MM*DD* 패턴의 날짜를 처리한다.
var RegexpSix = regexp.MustCompile(`^\d{2}\D\s?(0[1-9]|[12][0-9]|3[01])\D\s?\d{2}\D?\s?$`) // 16-10-19, 16/10/19, 16.10.19, 16.10.19. 16. 10. 10.

// Now 함수는 디지털아이디어에서 사용하는 서비스의 현재 시간을 RFC3339 포멧으로 반환한다.
func Now() string {
	return time.Now().Format(time.RFC3339)
}

// Str2time 함수는 RFC3339 시간문자열을 Time 자료구조로 바꾼다.
func Str2time(str string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return t, err // 이때 t 값은 1970년 초기값을 반환한다.
	}
	return t, nil
}

// ToFullTime 함수는 시간(출근시작시간, 퇴근시간등)값과 시간문자를 입력받아서 FullTime(RFC3339)으로 변환한다.
func ToFullTime(hourNum int, t string) (string, error) {
	if t == "" {
		return "", nil // 빈문자열이 들어는 것은 시간을 제거하는 것과 같다.
	}
	var hour, min, sec, nsec int
	if 0 < hourNum && hourNum < 23 {
		hour = hourNum
	}
	if regexpShortTime.MatchString(t) {
		m, err := strconv.Atoi(t[0:2])
		if err != nil {
			return t, err
		}
		d, err := strconv.Atoi(t[2:])
		if err != nil {
			return t, err
		}
		t := time.Date(time.Now().Year(), time.Month(m), d, hour, min, sec, nsec, time.Local)
		return t.Format(time.RFC3339), nil
	} else if regexpNormalTime.MatchString(t) {
		y, err := strconv.Atoi(t[0:4])
		if err != nil {
			return t, err
		}
		m, err := strconv.Atoi(t[5:7])
		if err != nil {
			return t, err
		}
		d, err := strconv.Atoi(t[8:])
		if err != nil {
			return t, err
		}
		t := time.Date(y, time.Month(m), d, hour, min, sec, nsec, time.Local)
		return t.Format(time.RFC3339), nil
	} else if regexpFullTime.MatchString(t) {
		y, err := strconv.Atoi(t[0:4])
		if err != nil {
			return t, err
		}
		m, err := strconv.Atoi(t[5:7])
		if err != nil {
			return t, err
		}
		d, err := strconv.Atoi(t[8:10])
		if err != nil {
			return t, err
		}
		t := time.Date(y, time.Month(m), d, hour, min, sec, nsec, time.Local)
		return t.Format(time.RFC3339), nil
	} else if RegexpMDD.MatchString(t) {
		m, err := strconv.Atoi(string(t[0]))
		if err != nil {
			return t, err
		}
		d, err := strconv.Atoi(t[1:])
		if err != nil {
			return t, err
		}
		t := time.Date(time.Now().Year(), time.Month(m), d, hour, min, sec, nsec, time.Local)
		return t.Format(time.RFC3339), nil
	} else if RegexpYYYYMMDD.MatchString(t) {
		re, err := regexp.Compile(`^(\d+)\D\s?(\d+)\D\s?(\d+)\D?\s?`)
		if err != nil {
			return t, err
		}
		result := re.FindStringSubmatch(t)
		y, err := strconv.Atoi(result[1])
		if err != nil {
			return t, err
		}
		m, err := strconv.Atoi(result[2])
		if err != nil {
			return t, err
		}
		d, err := strconv.Atoi(result[3])
		if err != nil {
			return t, err
		}
		t := time.Date(y, time.Month(m), d, hour, min, sec, nsec, time.Local)
		return t.Format(time.RFC3339), nil
	} else if RegexpMMDDYYYY.MatchString(t) {
		re, err := regexp.Compile(`^(\d+)\D\s?(\d+)\D\s?(\d+)\D*`)
		if err != nil {
			return t, err
		}
		result := re.FindStringSubmatch(t)
		m, err := strconv.Atoi(result[1])
		if err != nil {
			return t, err
		}
		d, err := strconv.Atoi(result[2])
		if err != nil {
			return t, err
		}
		y, err := strconv.Atoi(result[3])
		if err != nil {
			return t, err
		}
		t := time.Date(y, time.Month(m), d, hour, min, sec, nsec, time.Local)
		return t.Format(time.RFC3339), nil
	} else if RegexpMMDD.MatchString(t) {
		re, err := regexp.Compile(`^(\d+)\D\s?(\d+)\D*`)
		if err != nil {
			return t, err
		}
		result := re.FindStringSubmatch(t)
		m, err := strconv.Atoi(result[1])
		if err != nil {
			return t, err
		}
		d, err := strconv.Atoi(result[2])
		if err != nil {
			return t, err
		}
		t := time.Date(time.Now().Year(), time.Month(m), d, hour, min, sec, nsec, time.Local)
		return t.Format(time.RFC3339), nil
	} else if RegexpExcelDaynum.MatchString(t) {
		num, err := strconv.Atoi(t)
		if err != nil {
			return t, err
		}
		s := time.Date(1899, time.December, 30, hour, min, sec, nsec, time.Local)
		return s.AddDate(0, 0, num).Format(time.RFC3339), nil
	} else if RegexpSix.MatchString(t) {
		re, err := regexp.Compile(`^(\d+)\D\s?(\d+)\D\s?(\d+)\D*`)
		if err != nil {
			return t, err
		}
		result := re.FindStringSubmatch(t)
		head, err := strconv.Atoi(result[1])
		if err != nil {
			return t, err
		}
		body, err := strconv.Atoi(result[2])
		if err != nil {
			return t, err
		}
		tail, err := strconv.Atoi(result[3])
		if err != nil {
			return t, err
		}
		var y, m, d int
		if head > 18 {
			y = 2000 + head
			m = body
			d = tail
		} else {
			m = head
			d = body
			y = 2000 + tail
		}
		t := time.Date(y, time.Month(m), d, hour, min, sec, nsec, time.Local)
		return t.Format(time.RFC3339), nil
	} else {
		return t, errors.New(`입력한 날짜형식이 "0113","1982-01-13","1982-01-13T10:38:37+09:00" 형태가 아닙니다`)
	}
}
