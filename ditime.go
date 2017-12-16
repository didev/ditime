package ditime

import (
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
