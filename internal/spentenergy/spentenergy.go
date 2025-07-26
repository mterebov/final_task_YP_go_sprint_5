package spentenergy

import (
	"time"
	"fmt"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных данных
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("bad data")
	}
	// Подсчет средней скорости и  затраченных калорий
	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / float64(minInH) * float64(walkingCaloriesCoefficient)
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных данных
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("bad data")
	}
	// Подсчет средней скорости и  затраченных калорий
	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / float64(minInH)
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверяем входные данные
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	// Вычисляем и возвращаем сренюю скорость
	meanSpeed := Distance(steps, height) / duration.Hours()
	return meanSpeed 
}

func Distance(steps int, height float64) float64 {
	// Проверяем входные данные
	if steps <= 0 || height <= 0 {
		return 0
	}
	// Вычисляем длину шага и дистанцию
	stepLength := height * stepLengthCoefficient
	distance := stepLength * float64(steps) / mInKm
	return distance
}
