package pointers

import "testing"
import "fmt"

func TestWallet(t *testing.T) {
	wallet := Wallet{0}
	wallet.Deposit(10)
	got := wallet.Balance()

	fmt.Println("address of balance in test is", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
