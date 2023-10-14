package builder

import (
	"belanjabackend/entity"
	"belanjabackend/repository"
	"context"
	"database/sql"
	"fmt"
)

type walletImpl struct {
	DB *sql.DB
}

func NewWalletRepo(db *sql.DB) repository.WalletRepository {
	return &walletImpl{DB: db}
}

func (walletRepoImpl *walletImpl) Insert(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error) {
	script := "INSERT INTO wallet(username, balance, card_name) VALUES($1, $2, $3)"

	_, err := walletRepoImpl.DB.ExecContext(ctx, script, wallet.User_name, wallet.Balance, wallet.Card_name)
	defer walletRepoImpl.DB.Close()

	if err != nil {
		panic(err)
	}

	return wallet, nil
}

func (walletRepoImpl *walletImpl) FindById(ctx context.Context, id string) (entity.Wallet, error) {
	wallet := entity.Wallet{}
	defer walletRepoImpl.DB.Close()

	script := "SELECT username, balance, card_name FROM wallet WHERE id = $1"
	rows, err := walletRepoImpl.DB.QueryContext(ctx, script, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&wallet.User_name, &wallet.Balance, &wallet.Card_name)

		fmt.Println("Name: ", wallet.User_name, "\n Balance", wallet.Balance, "\n Card Name", wallet.Card_name)
	}
	return wallet, nil
}

func (walletRepoImpl *walletImpl) FindAll(ctx context.Context) ([]entity.Wallet, error) {
	defer walletRepoImpl.DB.Close()

	script := "SELECT username, balance, card_name FROM wallet"
	rows, err := walletRepoImpl.DB.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var wallets []entity.Wallet
	for rows.Next() {
		wallet := entity.Wallet{}
		rows.Scan(&wallet.User_name, &wallet.Balance, &wallet.Card_name)
		wallets = append(wallets, wallet)
		fmt.Println("Name: ", wallet.User_name, "\n Balance", wallet.Balance, "\n Card Name", wallet.Card_name)
	}

	return wallets, nil
}

func (walletRepoImpl *walletImpl) UpdateWallet(ctx context.Context, id string, wallet entity.Wallet) (entity.Wallet, error) {
	script := "UPDATE wallet SET username = $1, balance = $2, card_name = $3 WHERE id = $4"

	_, err := walletRepoImpl.DB.ExecContext(ctx, script, wallet.User_name, wallet.Balance, wallet.Card_name, wallet.Id)
	defer walletRepoImpl.DB.Close()
	if err != nil {
		panic(err)
	}

	return wallet, nil
}
