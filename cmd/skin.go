package main

import "fmt"

type Skin struct {
	Name  string  `json:"Name"`
	Wear  string  `json:"Wear"`
	Price float32 `json:"Price"`
	// Shares []*Share
}

type Share struct {
	CurrentPrice float32
	ListedPrice  float32
}

func NewSkin(name, wear string, intial float32) Skin {
	return Skin{
		Name:  name,
		Wear:  wear,
		Price: intial,
	}
}

// func (s *Skin) buyShare() {
// 	cp := s.CurrentPrice / 10
// 	lp := s.CurrentPrice / 10
// 	share := NewShare(cp, lp)
// 	s.Shares = append(s.Shares, share)
// }

// func (s *Skin) increasePrice(amount float32) {
// 	s.CurrentPrice += amount
// 	s.increaseShareValues(amount)
// }

// func (s *Skin) increaseShareValues(amount float32) {
// 	for _, s := range s.Shares {
// 		s.CurrentPrice += amount / 10
// 	}
// }

// func (s *Skin) printValues() {
// 	fmt.Println("Skin:", s.Name)
// 	fmt.Println("Skin price:", s.CurrentPrice)
// 	for _, share := range s.Shares {
// 		share.printInfo()
// 	}
// }

func NewShare(cp, lp float32) *Share {
	return &Share{
		CurrentPrice: cp,
		ListedPrice:  lp,
	}
}

func (s *Share) printInfo() {
	fmt.Printf("  Share current price: $%.2f, ", s.CurrentPrice)
	fmt.Printf("Share listed price: $%.2f\n", s.ListedPrice)
}
