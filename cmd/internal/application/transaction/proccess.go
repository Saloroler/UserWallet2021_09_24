package transaction

import "UserWallet2021_09_24/cmd/internal/db/mysql/repo"

type Process struct {
	transactionRepo repo.TransactionRepo
}

func NewProcess(transactionRepo repo.TransactionRepo) Process {
	return Process{transactionRepo}
}
