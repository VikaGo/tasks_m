package main

import (
	"fmt"
	"time"
)

const layout1 = "03:04:05PM"
const layout2 = "15:04:05"

func main() {
	militaryTime, err := ConvertToMilitaryTime("00:03:12")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(militaryTime)
}
func ConvertToMilitaryTime(time12format string) (string, error) {
	t, err := time.Parse(layout1, time12format)
	if err != nil {
		return "", err
	}
	militaryTime := t.Format(layout2)
	return militaryTime, nil
}
