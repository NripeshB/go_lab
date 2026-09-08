package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Withdraw Bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: 100}
		err := wallet.Withdraw(67)
		assertNoError(t, err)
		assertBalance(t, wallet, 33)

	})
	t.Run("Deposit Bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: 100}
		wallet.Deposit(Bitcoin(67))
		assertBalance(t, wallet, 167)

	})
	t.Run("withdraw more than Balance", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(101))
		assertError(t, err, ErrorInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})

}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error("Wanted no errors but still go some")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}
func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected an error and didn't get one")
	}
	if got.Error() != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}
