package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"
)

func main() {
	n := promptInt("enter RSA modulus")
	e := promptInt("enter public exponent")
	d := promptInt("enter private exponent")
	p, err := factor(n, e, d)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v = %v * %v\n", n, p, new(big.Int).Div(n, p))
	}
}

func promptInt(p string) *big.Int {
	fmt.Println(p)
	v := new(big.Int)
	_, err := fmt.Scan(v)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return v
}

var (
	one = big.NewInt(1)
)

func factor(n, e, d *big.Int) (*big.Int, error) {
	nm1 := new(big.Int)
	nm1.Sub(n, one)
	t := big.NewInt(0)
	t.Mul(e, d)
	t.Sub(t, one)
	s := 0
	for t.Bit(0) == 0 {
		s++
		t.Rsh(t, 1)
	}
nexta:
	for i := 0; i < 100; i++ {
		a := big.NewInt(int64(i + 2))
		fmt.Println("try with ", a)
		a.Exp(a, t, n)
		fmt.Println("a**t:", a)
		if a.Cmp(one) == 0 || a.Cmp(nm1) == 0 {
			continue
		}
		for i := s; i > 0; i-- {
			a2 := new(big.Int)
			a2.Mul(a, a)
			a2.Mod(a2, n)
			fmt.Println("a**2:", a2)
			if a2.Cmp(nm1) == 0 {
				continue nexta
			}
			if a2.Cmp(one) == 0 {
				f1 := new(big.Int)
				f1.GCD(nil, nil, n, a.Sub(a, one))
				return f1, nil
			}
			a = a2
		}
		return nil, errors.New("invalid keypair")
	}
	return nil, errors.New("no factors found after 100 iterations")
}
