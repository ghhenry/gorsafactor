package main

import (
	"fmt"
	"math/big"
	"testing"
)

func sv(s string) *big.Int {
	v, _ := new(big.Int).SetString(s, 10)
	return v
}

var factorCases = []struct {
	n *big.Int
	d *big.Int
	e *big.Int
}{
	{big.NewInt(527), big.NewInt(13), big.NewInt(37)},
	{sv("52137619201621371893143"),
		big.NewInt(7541),
		sv("6913886646486391469")},
	{sv("52137619201621371893143"),
		big.NewInt(7),
		sv("14896462628901108019351")},
	{sv("52137619201621371893143"),
		big.NewInt(3715),
		sv("42103057228388057659")},
	{sv("52137619201621371893143"),
		big.NewInt(65537),
		sv("374701598238300144497")},
	{sv("52137619201621371893143"),
		sv("1862559646135"),
		sv("357865878507063099223")},
}

func TestFactor(t *testing.T) {
	for _, c := range factorCases {
		p, err := factor(c.n, c.e, c.d)
		if err != nil {
			t.Fatalf("factor gave an error: %v", err)
		}
		fmt.Printf("%v has factor %v\n", c.n, p)
		if new(big.Int).Mod(c.n, p).Cmp(big.NewInt(0)) != 0 {
			t.Error("but it is not a factor")
		}
	}
}
