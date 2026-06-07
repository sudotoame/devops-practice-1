package logic

import "sync/atomic"

type Logic struct {
	Number atomic.Int64
}

func NewLogic() *Logic {
	return &Logic{
		Number: atomic.Int64{},
	}
}

func (l *Logic) Incriment() {
	l.Number.Add(1)
}

func (l *Logic) NumberOutput() int64 {
	return l.Number.Load()
}

func (l *Logic) Healthcheck() bool {
	_ = l.Number.Load()
	return true
}
