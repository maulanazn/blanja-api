package builder

import (
	"belanjabackend/config"
	"belanjabackend/entity"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	_ "github.com/lib/pq"
)

func TestInsertData(t *testing.T) {
	walletConnect := NewWalletRepo(config.GetConnection())
	WalletImpl := entity.Wallet{
		User_name: faker.New().Person().Name(),
		Balance:   faker.New().RandomNumber(10),
		Card_name: faker.New().Payment().CreditCardType(),
	}

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 2*time.Second)

	result, err := walletConnect.Insert(ctx, WalletImpl)
	defer cancel()

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestWalletById(t *testing.T) {
	walletConnect := NewWalletRepo(config.GetConnection())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 2*time.Second)

	result, err := walletConnect.FindById(ctx, "7a86a689-ab9d-4f96-a2b6-247e21af3a88")
	defer cancel()

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestWalletFindAll(t *testing.T) {
	walletConnect := NewWalletRepo(config.GetConnection())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 2*time.Second)

	wallets, err := walletConnect.FindAll(ctx)
	defer cancel()
	if err != nil {
		panic(err)
	}

	for result := range wallets {
		fmt.Println(result)
	}
}

func TestWalletUpdate(t *testing.T) {
	walletConnect := NewWalletRepo(config.GetConnection())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 2*time.Second)

	WalletImpl := entity.Wallet{
		Balance: faker.New().RandomNumber(8),
	}

	result, err := walletConnect.UpdateWallet(ctx, "7a86a689-ab9d-4f96-a2b6-247e21af3a88", WalletImpl)
	defer cancel()

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
