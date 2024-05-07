package model

import (
	"fmt"
	"slices"
)

type Group struct {
	tier  string
	skins map[string][]Skin
	total int
}

func NewGroup(t string) *Group {
	return &Group{
		tier:  t,
		skins: make(map[string][]Skin, 10),
	}
}

func (g *Group) AddSkin(s Skin) {
	g.total = g.total + 1
	key := s.generateKey()
	g.skins[key] = append(g.skins[key], s)
}

func (g *Group) RemoveSkin(s Skin) {
	g.total = g.total - 1
	key := s.generateKey()
	i := slices.Index(g.skins[key], s)
	g.skins[key] = slices.Delete(g.skins[key], i, i+1)
}

func (g *Group) TradeUp() Skin {
	for k, v := range g.skins {
		fmt.Printf("%s\n", k)
		for _, s := range v {
			fmt.Printf("%v ", s)
		}
		fmt.Println()
	}
	return Skin{}
}

func (g *Group) IsReady() bool {
	return g.total == 10
}
