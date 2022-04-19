package types

import (
	"encoding/json"
	"io/ioutil"
)

type Invoices []Invoice

type Invoice struct {
	Cusomer     string        `json:"cusomer"`
	Performance []Performance `json:"Performance"`
}

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int64  `json:"audience"`
}

func LoadInvoice(path string) (Invoices, error) {
	invoiceData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	invoices := make([]Invoice, 0, 0)
	err = json.Unmarshal(invoiceData, &invoices)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}
