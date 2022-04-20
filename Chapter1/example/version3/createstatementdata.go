package version3

import (
	"fmt"
	"math"

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
	result.Amount = result.amountFor()
	result.VolumeCredits = result.volumeCreditsFor()
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

func (p *PerformanceDetails) volumeCreditsFor() float64 {
	result := math.Max(float64(p.APerformance.Audience-30), float64(0))
	if p.Play.Type == "comedy" {
		result += math.Floor(float64(p.APerformance.Audience) / float64(5))
	}
	return result
}

func playFor(plays types.Plays, aPerformance types.Performance) types.Play {
	return plays[aPerformance.PlayID]
}

func (p *PerformanceDetails) amountFor() float64 {
	var result float64 = 0
	switch p.Play.Type {
	case "tragedy":
		{
			result = 40000
			if p.APerformance.Audience > 30 {
				result += float64(1000 * (p.APerformance.Audience - 30))
			}
			break
		}
	case "comedy":
		{
			result = 30000
			if p.APerformance.Audience > 20 {
				result += float64(10000 + 500*(p.APerformance.Audience-20))
			}
			result += float64(300 * p.APerformance.Audience)
			break
		}
	default:
		{
			panic(fmt.Sprintf("unknown tpye %s", p.Play.Type))
		}
	}
	return result
}


func createStatementdata(invoice types.Invoice, plays types.Plays) *StatementsData {
	statementsData := NewStatementsData(invoice, plays)
	statementsData.TotalAmount = statementsData.totalAmount()
	statementsData.TotalVolumeCredits = statementsData.totalVolumeCredits()
	return statementsData
}