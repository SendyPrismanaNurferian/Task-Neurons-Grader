package main

import (
	"fmt"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	switch t := time.(type) {
	case string:
		hour, minute, err := parseTimeString(t)
		if err != nil {
			return "Invalid input"
		}
		return convertToStandardTime(hour, minute)
	case []int:
		if len(t) != 2 {
			return "Invalid input"
		}
		return convertToStandardTime(t[0], t[1])
	case map[string]int:
		hour, ok := t["hour"]
		if !ok {
			return "Invalid input"
		}
		minute, ok := t["minute"]
		if !ok {
			return "Invalid input"
		}
		return convertToStandardTime(hour, minute)
	case Time:
		return convertToStandardTime(t.Hour, t.Minute)
	default:
		return "Invalid input"
	}
}

func parseTimeString(timeStr string) (int, int, error) {
	var hour, minute int
	_, err := fmt.Sscanf(timeStr, "%d:%d", &hour, &minute)
	if err != nil {
		return 0, 0, err
	}
	return hour, minute, nil
}

func convertToStandardTime(hour, minute int) string {
	ampm := "AM"
	if hour >= 12 {
		ampm = "PM"
	}
	if hour > 12 {
		hour -= 12
	}
	return fmt.Sprintf("%02d:%02d %s", hour, minute, ampm)
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
