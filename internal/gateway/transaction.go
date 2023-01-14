package gateway

import "github.com/MuriloAbranches/fc-ms-walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
