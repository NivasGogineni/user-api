package service

import "time"

// CalculateAge returns age based on date of birth
func CalculateAge(dob time.Time) int {
	now := time.Now()

	age := now.Year() - dob.Year()

	// If birthday hasn't occurred yet this year, subtract 1
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}
