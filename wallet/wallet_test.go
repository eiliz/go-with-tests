package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(20)}
		err := w.Withdraw(Bitcoin(10))

		assertBalance(t, w, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		w := Wallet{startingBalance}
		err := w.Withdraw(Bitcoin(100))

		assertBalance(t, w, startingBalance)

		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("expected %s, got %s", want, got)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
