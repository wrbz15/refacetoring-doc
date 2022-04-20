package version2

import (
	"fmt"
	"math"

	"github.com/wrbz15/refacetoring-doc/Chapter1/example/types"
)

func Statements(invoice types.Invoice, plays types.Plays) string {
	result := fmt.Sprintf("Statements for %s \n", invoice.Cusomer)
	for _, perf := range invoice.Performance {
		result += fmt.Sprintf("	%v: $%v  %v seats \n", playFor(plays, perf).Name, amountFor(plays, perf)/100, perf.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%v \n", totalAmount(invoice, plays)/100)
	result += fmt.Sprintf("you earned $%v credits \n", totalVolumeCredits(invoice, plays))
	return result
}

func totalAmount(invoice types.Invoice, plays types.Plays) float64 {
	var result float64 = 0
	for _, perf := range invoice.Performance {
		result += amountFor(plays, perf)
	}
	return result
}

func totalVolumeCredits(invoice types.Invoice, plays types.Plays) float64 {
	var result float64 = 0
	for _, perf := range invoice.Performance {
		result = result + volumeCreditsFor(plays, perf)
	}
	return result
}

func volumeCreditsFor(plays types.Plays, aPerformance types.Performance) float64 {
	result := math.Max(float64(aPerformance.Audience-30), float64(0))
	if playFor(plays, aPerformance).Type == "comedy" {
		result += math.Floor(float64(aPerformance.Audience) / float64(5))
	}
	return result
}

func playFor(plays types.Plays, aPerformance types.Performance) types.Play {
	return plays[aPerformance.PlayID]
}

func amountFor(plays types.Plays, aPerformance types.Performance) float64 {
	var result float64 = 0
	switch playFor(plays, aPerformance).Type {
	case "tragedy":
		{
			result = 40000
			if aPerformance.Audience > 30 {
				result += float64(1000 * (aPerformance.Audience - 30))
			}
			break
		}
	case "comedy":
		{
			result = 30000
			if aPerformance.Audience > 20 {
				result += float64(10000 + 500*(aPerformance.Audience-20))
			}
			result += float64(300 * aPerformance.Audience)
			break
		}
	default:
		{
			panic(fmt.Sprintf("unknown tpye %s", playFor(plays, aPerformance).Type))
		}
	}
	return result
}
