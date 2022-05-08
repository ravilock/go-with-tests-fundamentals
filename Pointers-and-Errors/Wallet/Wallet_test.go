package Wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficiente funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrorInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but did'nt want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
