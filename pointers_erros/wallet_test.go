package pointers_erros

import "testing"

func TestWallet(t *testing.T){
	t.Run("Deposit", func (t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t,wallet,Bitcoin(10))
	})

	t.Run("Withdraw", func (t *testing.T){
		wallet := Wallet{balance: Bitcoin(30)}
		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Withdraw insufficient funds", func (t *testing.T){
		wallet := Wallet{balance : Bitcoin(20)}
	    err := wallet.Withdraw(Bitcoin(100))

	    assertBalance(t, wallet,Bitcoin(20)	)
	    assertError(t, err,ErrInsufficientFunds)

	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error){
	t.Helper()
	if got == nil {
		t.Fatal("didnt get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q, want %q", want, got)
	}
}
