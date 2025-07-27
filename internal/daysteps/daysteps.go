package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"errors"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// Парсим строку и проаеряем что в слайсе необходимое кол-во элементов (2)
	splittedData := strings.Split(datastring, ",") // ["Шаги", "Продолжительность"]
	if len(splittedData) != 2 {
		return fmt.Errorf("bad data: count of args %d != 2", len(splittedData))
	}
	// Приводим элементы слайса к необходимым типам данных
	// Обрабатываем ошибки
	steps, err := strconv.Atoi(splittedData[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("daysteps/Parse() - bad data: steps")
	}

	timeDur, err := time.ParseDuration(splittedData[1])
	if err != nil {
		return err
	}
	if timeDur <= 0 {
		return errors.New("daysteps/Parse() - bad data: time")
	}
	// Сохраняем расшитые данные в поля структуры
	ds.Steps = steps
	ds.Duration =  timeDur
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// вычисляем дистанцию и калории, обрабатываем ошибки
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	outStr := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)
	return outStr, nil
}
