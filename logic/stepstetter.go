package logic

import (
	"custom_step/service"
	"errors"
	"log"
	"math/rand"
	"time"

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
	if err := s.miSrv.PushData(time.Now().Format("2006-01-02"), step, tokenInfo); err != nil {
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
var hours = map[int]int{8: 1, 12: 2, 15: 3, 17: 4, 21: 5}

func getStep() string {
	index, ok := hours[time.Now().Hour()]
	if !ok {
		return ""
	}

	return cast.ToString(random(steps[index].min, steps[index].max))
}

func random(min, max int) int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rnd.Intn(max-min) + min
}
