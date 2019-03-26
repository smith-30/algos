package main

import (
	"fmt"
	"sync"
)

type ID int
type Sex int

const (
	man Sex = iota
	woman
)

func (s Sex) String() string {
	switch s {
	case man:
		return "man"
	case woman:
		return "woman"
	default:
		return "unknown"
	}
}

func IDFactory() func() ID {
	var (
		idx  int
		lock sync.Mutex
	)
	return func() ID {
		lock.Lock()
		defer lock.Unlock()
		idx++
		return ID(idx)
	}
}

type Participant interface {
	ID() ID
	Sex() Sex
	AddLikeOrder(ids []ID)
}

type Man struct {
	id ID

	Name      string
	LikeOrder []ID
}

func (a *Man) ID() ID {
	return a.id
}

func (a *Man) Sex() Sex {
	return man
}

func (a *Man) AddLikeOrder(ids []ID) {
	a.LikeOrder = ids
}

var ManFactory = func(name string) Participant {
	idf := IDFactory()
	return func(n string) Participant {
		return &Man{
			id:   idf(),
			Name: n,
		}
	}(name)
}

type Woman struct {
	id ID

	Name      string
	LikeOrder []ID
}

func (a *Woman) AddLikeOrder(ids []ID) {
	a.LikeOrder = ids
}

func (a *Woman) Sex() Sex {
	return woman
}

func (a *Woman) ID() ID {
	return a.id
}

var WomanFactory = func(name string) Participant {
	idf := IDFactory()
	return func(n string) Participant {
		return &Woman{
			id:   idf(),
			Name: n,
		}
	}(name)
}

type Event struct {
	MenList   map[ID]Participant
	WomenList map[ID]Participant
}

func NewEvent() *Event {
	return &Event{
		MenList:   map[ID]Participant{},
		WomenList: map[ID]Participant{},
	}
}

func (a *Event) AddParticipant(p Participant) {
	switch p.Sex() {
	case man:
		a.MenList[p.ID()] = p
	case woman:
		a.WomenList[p.ID()] = p
	}
}

func (a *Event) SetOrder(target ID, orderIDs []ID) error {
	if m, ok := a.MenList[target]; ok {
		m.AddLikeOrder(orderIDs)
	}

	if w, ok := a.WomenList[target]; ok {
		w.AddLikeOrder(orderIDs)
	}

	return fmt.Errorf("Not found target. id: %d", target)
}

type Matching struct {
	Man
	Woman
}

func (a *Matching) Show() string {
	return fmt.Sprintf("(man: %d, woman: %d)", a.Man.ID(), a.Woman.ID())
}

func main() {
	e := &Event{}
	// 男性受付
	e.AddParticipant(ManFactory("m_1"))
	e.AddParticipant(ManFactory("m_2"))
	e.AddParticipant(ManFactory("m_3"))

	// 女性受付
	e.AddParticipant(WomanFactory("w_1"))
	e.AddParticipant(WomanFactory("w_2"))
	e.AddParticipant(WomanFactory("w_3"))

	// Event開始
	
}
