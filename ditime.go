package ditime

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var regexpShortTime = regexp.MustCompile(`^\d{4}$`)                                                                            // 1019
var regexpNormalTime = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])$`)                                // 2016-10-19
var regexpFullTime = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$`) // 2016-10-19T16:41:24+09:00

// RegexpMMDD 레귤러 익스프레션은 MM*DD 패턴의 날짜를 처리한다.
var RegexpMMDD = regexp.MustCompile(`^(0?[1-9]|1[012])[-.,/]\s?(0?[1-9]|[12][0-9]|3[01])[,.]*$`) // 10/19, 10-19, 10.10. 1.1 1.1. "1. 1."
// RegexpYYYYMMDD 레귤러 익스프레션은 YYYY*MM*DD* 패턴의 날짜를 처리한다.
var RegexpYYYYMMDD = regexp.MustCompile(`^\d{4}[-,./]\s?(0?[1-9]|1[012])[-,./]\s?(0?[1-9]|[12][0-9]|3[01])[,.]*$`) // 2016-10-19, 2016/10/19, 2016.10.19, 2016.10.19. 2016. 10. 10.

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

// Worktime 함수는 현재 업무시간인지 체크하는 함수이다.
func Worktime(t time.Time) bool {
	// 토요일, 일요일이면 업무시간이 아니다.
	if t.Weekday() == 0 || t.Weekday() == 6 {
		return false
	}
	h := t.Hour()
	// 회사 업무시간은 10:00~19:00 입니다.
	switch h {
	case 10, 11, 12: // 오전업무시간
		return true
	case 13: // 점심시간
		return false
	case 14, 15, 16, 17, 18: // 오후업무시간
		return true
	default:
		return false
	}
}

// ToFullTime 함수는 모드와 시간문자를 입력받아서 FullTime(RFC3339)으로 변환한다.
// 모드는 start, end, current 값을 설정할 수 있다. 각각 출근,퇴근,현재시간으로 FullTime값을 바꿀 수 있다.
func ToFullTime(mode, t string) (string, error) {
	if t == "" {
		return "", nil // 빈문자열이 들어는 것은 시간을 제거하는 것과 같다.
	}
	var hour, min, sec, nsec int
	switch mode {
	case "start":
		hour = 10
	case "end":
		hour = 19
	case "current":
		currentTime := time.Now()
		hour = currentTime.Hour()
		min = currentTime.Minute()
		sec = currentTime.Second()
		nsec = currentTime.Nanosecond()
	default:
		return t, errors.New("mode에는 start,end,current 값만 입력할 수 있습니다")
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
	} else if RegexpMMDD.MatchString(t) {
		m, err := strconv.Atoi(t[0:2])
		if err != nil {
			return t, err
		}
		d, err := strconv.Atoi(t[3:])
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
		switch mode {
		case "start", "end":
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
		default:
			_, err := time.Parse(time.RFC3339, t)
			if err != nil {
				return t, err
			}
			return t, nil
		}
	} else if RegexpYYYYMMDD.MatchString(t) {
		re, err := regexp.Compile(`^(\d+).(\d+).(\d+)`)
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
	} else if RegexpMMDD.MatchString(t) {
		re, err := regexp.Compile(`(\d+)[.,-/]\s?(\d+)`)
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
	} else {
		return t, errors.New(`입력한 날짜형식이 "0113","1982-01-13","1982-01-13T10:38:37+09:00" 형태가 아닙니다`)
	}
}
