package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// Парсим строку и проаеряем что в слайсе необходимое кол-во элементов (3)
	splittedData := strings.Split(datastring, ",") // ["Шаги", "Тип тренировки", "Продолжительность"]
	if len(splittedData) != 3 {
		return fmt.Errorf("bad data")
	}
	// Приводим элементы слайса к необходимым типам данных
	// Обрабатываем ошибки
	steps, err := strconv.Atoi(splittedData[0])
	if err != nil || steps <= 0 {
		return fmt.Errorf("bad data: steps")
	}

	timeDur, err := time.ParseDuration(splittedData[2])
	if err != nil || timeDur <= 0 {
		return fmt.Errorf("bad data: time")
	}
	// Сохраняем расшитые данные в поля структуры
	t.Steps = steps
	t.TrainingType = splittedData[1]
	t.Duration = timeDur
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Вычисляем дистанцию и среднюю скорость
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	// Расчитываем калории в зависимости от типа тренировки, обрабатываем возможные ошибки и возвращаем результирующую строку
	switch t.TrainingType{
		case "Ходьба":
			calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
			if err != nil {
				return "", err
			}
			outStr := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)
			return outStr, nil
		case "Бег":
			calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
			if err != nil {
				return "", err
			}
			outStr := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)
			return outStr, nil
		default:
			return "", fmt.Errorf("неизвестный тип тренировки")
	}
}
