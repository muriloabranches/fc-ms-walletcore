package gateway

import "github.com/muriloabranches/fc-ms-walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
