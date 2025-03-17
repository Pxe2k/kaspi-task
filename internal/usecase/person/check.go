package person

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func (u UseCase) Check(ctx context.Context, docID string) (gender string, birthDate string, err error) {
	genderDigit, err := strconv.Atoi(string(docID[6]))
	if err != nil {
		return "", "", fmt.Errorf("can't get gender digit: %w document_id is %s", err, docID)
	}

	gender, err = checkGender(genderDigit)
	if err != nil {
		return "", "", fmt.Errorf("can't get gender: %w", err)
	}

	birthDate, err = checkBirthDate(docID[0:6])
	if err != nil {
		return "", "", fmt.Errorf("can't get birth date: %w", err)
	}

	return
}

func checkGender(i int) (string, error) {
	switch i {
	case 3, 5:
		return "male", nil
	case 4, 6:
		return "female", nil
	default:
		return "", fmt.Errorf("wrong gender number: %d", i)
	}
}

func checkBirthDate(s string) (string, error) {
	currentYear := time.Now().Format(time.DateOnly)[2:4]
	currentYearInt, err := strconv.Atoi(currentYear)
	if err != nil {
		return "", fmt.Errorf("can't parse current year to int: %w", err)
	}

	yearShort, err := strconv.Atoi(s[:2])
	if err != nil {
		return "", fmt.Errorf("parsing short year: %w", err)
	}

	var fullYear int
	if currentYearInt > yearShort {
		fullYear = 2000 + yearShort
	} else {
		fullYear = 1900 + yearShort
	}

	month, err := strconv.Atoi(s[2:4])
	if err != nil {
		return "", fmt.Errorf("parse month: %w", err)
	}

	day, err := strconv.Atoi(s[4:6])
	if err != nil {
		return "", fmt.Errorf("parse day: %w", err)
	}

	birthDate := fmt.Sprintf("%02d-%02d-%04d", day, month, fullYear)

	return birthDate, nil
}
