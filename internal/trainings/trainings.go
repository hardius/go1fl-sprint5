package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 3 {
		err = errors.New("invalid parameter value")
		return
	}

	t.Steps, err = strconv.Atoi(splitData[0])
	if err != nil {
		return
	}
	if t.Steps <= 0 {
		err = errors.New("invalid parameter value")
		return
	}

	t.TrainingType = splitData[1]

	t.Duration, err = time.ParseDuration(splitData[2])
	if err != nil {
		return
	}
	if t.Duration <= 0 {
		err = errors.New("invalid parameter value")
		return
	}

	return
}

func (t Training) ActionInfo() (string, error) {
	if !(t.TrainingType == "Бег" || t.TrainingType == "Ходьба") {
		return "", errors.New("неизвестный тип тренировки")
	}

	var calories float64
	var err error
	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	} else {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	}

	result := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	result += fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	result += fmt.Sprintf("Дистанция: %.2f км.\n", spentenergy.Distance(t.Steps, t.Personal.Height))
	result += fmt.Sprintf("Скорость: %.2f км/ч\n", spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration))
	result += fmt.Sprintf("Сожгли калорий: %.2f\n", calories)
	return result, nil
}
