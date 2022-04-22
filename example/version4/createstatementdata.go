package version4

import (
	"github.com/wrbz15/refacetoring-doc/Chapter1/example/types"
)

type PerformanceDetails struct {
	APerformance  types.Performance
	Play          types.Play
	Amount        float64
	VolumeCredits float64
}

func NewPerformanceDetails(plays types.Plays, aPerformance types.Performance) *PerformanceDetails {
	result := &PerformanceDetails{
		APerformance: aPerformance,
		Play:         playFor(plays, aPerformance),
	}
	PerformanceCalculator := createPerformanceCalculator(result.Play, aPerformance)
	result.Amount = PerformanceCalculator.GetAmount()
	result.VolumeCredits = PerformanceCalculator.GetVolumeCredits()
	return result
}

type StatementsData struct {
	Consumer           string                `json:"consumer"`
	Performances       []*PerformanceDetails `json:"performances"`
	TotalAmount        float64               `json:"totalAmount"`
	TotalVolumeCredits float64               `json:"totalvolumeCredits"`
}

func NewStatementsData(invoice types.Invoice, plays types.Plays) *StatementsData {
	result := &StatementsData{}
	result.Consumer = invoice.Cusomer
	result.Performances = make([]*PerformanceDetails, 0)
	for i, _ := range invoice.Performance {
		result.Performances = append(result.Performances, NewPerformanceDetails(plays, invoice.Performance[i]))
	}
	return result
}


func (p *StatementsData) totalAmount() float64 {
	var result float64 = 0
	for _, perf := range p.Performances {
		result += perf.Amount
	}
	return result
}

func (p *StatementsData) totalVolumeCredits() float64 {
	var result float64 = 0
	for _, perf := range p.Performances {
		result = result + perf.VolumeCredits
	}
	return result
}

func playFor(plays types.Plays, aPerformance types.Performance) types.Play {
	return plays[aPerformance.PlayID]
}

func createStatementdata(invoice types.Invoice, plays types.Plays) *StatementsData {
	statementsData := NewStatementsData(invoice, plays)
	statementsData.TotalAmount = statementsData.totalAmount()
	statementsData.TotalVolumeCredits = statementsData.totalVolumeCredits()
	return statementsData
}