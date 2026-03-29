package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 2 {
		err = errors.New("invalid parameter value")
		return
	}

	ds.Steps, err = strconv.Atoi(splitData[0])
	if err != nil {
		return
	}

	ds.Duration, err = time.ParseDuration(splitData[1])
	if err != nil {
		return
	}

	return
}

func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	result += fmt.Sprintf("Дистанция составила %.2f км.\n", spentenergy.Distance(ds.Steps, ds.Personal.Height))
	result += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)
	return result, nil
}
