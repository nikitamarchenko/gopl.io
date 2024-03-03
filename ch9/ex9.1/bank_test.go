// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	bank "gopl.io/ch9/ex9.1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBankWithdraw(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	var op1, op2 bool
	// Alice
	go func() {
		op1 = bank.Withdraw(300)
		done <- struct{}{}
	}()

	// Bob
	go func() {
		op2 = bank.Withdraw(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if op1 && op2 {
		t.Errorf("error both operations success Balance %d", bank.Balance())
	}

	if b := bank.Balance(); b != 0 && b != 200{
		t.Errorf("error operation invalid Balance %d must be 0 or 200", bank.Balance())
	}

}
