package spentenergy

import (
	"errors"
	"time"
)

var (
	ErrInvalidDuration = errors.New("invalid duration")
	ErrInvalidSteps    = errors.New("invalid steps")
	ErrInvalidWeight   = errors.New("invalid weight")
	ErrInvalidHeight   = errors.New("invalid height")
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 {
		return 0, ErrInvalidDuration
	}
	if steps <= 0 {
		return 0, ErrInvalidSteps
	}
	if weight <= 0 {
		return 0, ErrInvalidWeight
	}
	if height <= 0 {
		return 0, ErrInvalidHeight
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * meanSpeed * durationInMinutes * walkingCaloriesCoefficient / minInH
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 {
		return 0, ErrInvalidDuration
	}
	if steps <= 0 {
		return 0, ErrInvalidSteps
	}
	if weight <= 0 {
		return 0, ErrInvalidWeight
	}
	if height <= 0 {
		return 0, ErrInvalidHeight
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * meanSpeed * durationInMinutes / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps < 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength / mInKm
	return distance
}
