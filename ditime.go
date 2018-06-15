package ditime

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

var regexpShortTime = regexp.MustCompile(`^\d{4}$`)                                                                            // 1019
var regexpNormalTime = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])$`)                                // 2016-10-19
var regexpFullTime = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$`) // 2016-10-19T16:41:24+09:00

// csi, dilog처럼 국제서비스를 기준으로 제작되는 툴에서 사용하는 시간포멧
func Now() string {
	return time.Now().Format(time.RFC3339)
}

// Str2time함수는 RFC3339 시간문자열을 Time 자료구조로 바꾼다.
func Str2time(str string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return t, err // 이때 t 값은 1970년 초기값을 반환한다.
	}
	return t, nil
}

// IsWorktime 함수는 현재 업무시간인지 체크하는 함수이다.
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

// ToFullTime함수는 시간을 받아서 FullTime(RFC3339)으로 변환한다.
func ToFullTime(t string) (string, error) {
	// 10,19,현재시를 상황에 맞게 추가하는 기능 넣기
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return t, err
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
		t := time.Date(time.Now.Year(), m, d, 19, 0, 0, 0, location)
		return t.Format(time.RFC3339)
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
		t := time.Date(y, m, d, 19, 0, 0, 0, location)
		return t.Format(time.RFC3339)
	} else if regexpFullTime.MatchString(t) {
		return t, nil
	} else {
		return "", errors.New("입력한 날짜형식이 잘못 되었습니다.")
	}
}
