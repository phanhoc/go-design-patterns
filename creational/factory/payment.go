package factory

import (
	"fmt"
)

const (
	CASH       = 1
	DEBIT_CARD = 2
)

type PaymentMethod interface {
	Pay(amount uint32) string
}

type CashPM struct {
}
type DebitCard struct {
}

func (c *CashPM) Pay(amount uint32) string {
	return fmt.Sprintf("You have just pay amount %d by cash", amount)
}

func (d *DebitCard) Pay(amount uint32) string {
	return fmt.Sprintf("You have just pay amount %d by debit card", amount)
}

func GetPaymentMethod(typePM uint32) (PaymentMethod, error) {

	switch typePM {
	case CASH:
		return new(CashPM), nil
	case DEBIT_CARD:
		return new(DebitCard), nil
	default:
		return nil, fmt.Errorf("Unsupport type payment %d", typePM)
	}
}
