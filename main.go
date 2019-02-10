package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	n := promptInt(in, "enter RSA modulus")
	e := promptInt(in, "enter public exponent")
	d := promptInt(in, "enter private exponent")
	p, q := factor(n, e, d)
	fmt.Printf("%v = %v * %v\n", n, p, q)
}

func promptInt(in *bufio.Reader, p string) *big.Int {
	fmt.Println(p)
	ns, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	var n big.Int
	_, ok := n.SetString(ns[:len(ns)-1], 10)
	if !ok {
		fmt.Println("Could not convert number", ns)
		os.Exit(1)
	}
	return &n
}

var (
	one = big.NewInt(1)
)

func factor(n, e, d *big.Int) (p, q *big.Int) {
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
				return f1, new(big.Int).Div(n, f1)
			}
			a = a2
		}
		fmt.Println("invalid keypair", a)
		os.Exit(1)
	}
	fmt.Println("factors not found after 100 iterations")
	os.Exit(1)
	return // not reached
}
