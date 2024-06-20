package model

import ()

type Group struct {
	Tier   string
	Active int
}

type DisplayGroup struct {
	GroupId  int
	Tier     string
	SkinId   int
	SkinName string
}

func NewGroup(t string) Group {
	return Group{
		Tier:   t,
		Active: 1,
	}
}

func (g *Group) TradeUp() Skin {
	return Skin{}
}

func (g *Group) IsReady() bool {
	return false
}
