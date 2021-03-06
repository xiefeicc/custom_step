package logic

import (
	"custom_step/service"
	"custom_step/timeutils"
	"errors"
	"log"
	"math/rand"

	"github.com/spf13/cast"
)

type StepSetter interface {
	Do(user, pwd string) error
}

type stepSetter struct {
	miSrv service.MiSrv
}

func NewStepSetter() StepSetter {
	return &stepSetter{
		miSrv: service.NewMiSrv(),
	}
}

func (s stepSetter) Do(user, pwd string) error {
	if len(user) == 0 || len(pwd) == 0 {
		return errors.New("invalid user name or password")
	}
	code, err := s.miSrv.Registrations(user, pwd)
	if err != nil {
		return err
	}

	tokenInfo, err := s.miSrv.Login(code)
	if err != nil {
		return err
	}

	step := getStep()
	log.Println("update step:", step)
	if err := s.miSrv.PushData(timeutils.GetBeijingTM().Format("2006-01-02"), step, tokenInfo); err != nil {
		return err
	}
	return nil
}

var steps = []struct {
	min int
	max int
}{
	{
		min: 0,
		max: 1000,
	},
	{
		min: 1000,
		max: 4000,
	},
	{
		min: 4000,
		max: 8000,
	},
	{
		min: 8000,
		max: 10000,
	},
	{
		min: 10000,
		max: 20000,
	},
}
var hours = map[int]int{8: 0, 12: 1, 15: 2, 17: 3, 21: 4}

func getStep() string {
	index, ok := hours[timeutils.GetBeijingTM().Hour()]
	if !ok {
		return ""
	}

	return cast.ToString(random(steps[index].min, steps[index].max))
}

func random(min, max int) int {
	rnd := rand.New(rand.NewSource(timeutils.GetBeijingTM().UnixNano()))
	return rnd.Intn(max-min) + min
}
