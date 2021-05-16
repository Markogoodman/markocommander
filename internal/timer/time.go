package timer

import "time"

func GetNowTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Taipei")
	return time.Now().In(loc)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	dur, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(dur), nil
}
