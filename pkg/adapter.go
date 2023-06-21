package pkg

import (
	"fmt"
	"github.com/guneyin/gobist-importer/pkg/broker"
	"github.com/guneyin/gobist-importer/pkg/broker/ncm"

	"github.com/guneyin/gobist-importer/pkg/broker/garanti"
	"github.com/guneyin/gobist-importer/pkg/entity"
)

var (
	_ IBroker = (*garanti.Garanti)(nil)
	_ IBroker = (*ncm.NCM)(nil)
)

type IBroker interface {
	Get() broker.Model
	Parse(content []byte) (*entity.Transactions, error)
}

type BrokerAdapter struct {
	broker IBroker
}

func NewBrokerAdapter(b broker.Broker) (*BrokerAdapter, error) {
	var v IBroker

	switch b {
	case broker.Garanti:
		v = garanti.Garanti{}
	case broker.NCM:
		v = ncm.NCM{}
	default:
		return nil, fmt.Errorf("unspported broker %s", b)
	}

	return &BrokerAdapter{broker: v}, nil
}

func (va *BrokerAdapter) Parse(content []byte) (*entity.Transactions, error) {
	fmt.Println(fmt.Sprintf("reading %v transactions..", va.broker.Get().Name))

	return va.broker.Parse(content)
}
