package gateway

import "github.com/MuriloAbranches/fc-ms-walletcore/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
