/*

Exercise 9.1: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1
program. The result should indicate whether the transaction succeeded or failed
due to insufficient funds. The message sent to the monitor goroutine must
contain both the amount to withdraw and a new channel over which the monitor
goroutine can send the boolean result back to Withdraw.

*/

//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan struct {
	a int
	r chan<- bool
})

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdraw:
			if (balance - w.a) < 0 {
				w.r <- false
			} else {
				balance -= w.a
				w.r <- true
			}
		}
	}
}

func Withdraw(amount int) bool {
	res := make(chan bool)
	withdraw <- struct {
		a int
		r chan<- bool
	}{amount, res}
	return <-res
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
