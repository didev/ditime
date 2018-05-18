package ditime

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

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

//이 함수는 날짜를 입력받아 CSI의 시간형식으로 변경하고 시간이 잘못 입력 되었다면 err를 반환한다.
func CsiTime(t string) (string, error) {
	MatchShortTime := regexp.MustCompile(`^\d{4}$`)                                                                            // 1019
	MatchNormalTime := regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])$`)                                // 2016-10-19
	MatchFullTime := regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$`) // 2016-10-19T16:41:24+09:00
	ts := ""
	// 입력받은 날짜를 "2018-05-02T19:00:00+09:00" 와 같은 시간형식으로 변경
	if MatchFullTime.MatchString(t) {
		ts = t
	} else if MatchNormalTime.MatchString(t) {
		ts = fmt.Sprintf("%sT%02d:%02d:%02d%s", t, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Format(time.RFC3339)[19:])
	} else if MatchShortTime.MatchString(t) {
		t1 := fmt.Sprintf("%04d-%s-%s", time.Now().Year(), t[0:2], t[2:])
		if !MatchNormalTime.MatchString(t1) {
			return "", errors.New("올바른 날짜가 아닙니다!!")
		}
		ts = fmt.Sprintf("%sT%02d:%02d:%02d%s", t1, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Format(time.RFC3339)[19:])
	} else {
		return "", errors.New("입력한 날짜형식이 잘못 되었습니다.")

	}
	// 날짜형식이 올바른지 한번더 체크하기 위해 사용.
	// 아래의 time.Parse는 "ex) 2018-05-02 19:00:00 +0900 KST" 형식을 반환하지만
	// CSI에서 ex)"2018-05-02T19:00:00+09:00" 형식을 사용하기 때문에 값은 _ 처리하고 err만 체크한다.
	_, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return "", err
	}
	return ts, nil
}
