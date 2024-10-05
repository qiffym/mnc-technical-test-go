package transaction

import "time"

type (
	SendTrx struct {
		Title       string    `json:"title"`
		Amount      int       `json:"amount"`
		Description string    `json:"desc"`
		Date        time.Time `json:"date"`
	}

	UseCaseSendResult struct {
		Trx SendTrx `json:"transaction"`
	}
)
