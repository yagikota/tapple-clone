package model

import "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"

type HobbyID int

type HobbySlice []*Hobby

type Hobby struct {
	ID  HobbyID `json:"id"`
	Tag string  `json:"tag"`
}

func HobbyFromDomainModel(m *model.Hobby) *Hobby {
	return &Hobby{
		ID:  HobbyID(m.ID),
		Tag: m.Tag,
	}
}
