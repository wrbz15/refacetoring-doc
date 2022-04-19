package version1

import (
	"errors"
	"fmt"
	"math"

	"github.com/wrbz15/refacetoring-doc/Chapter1/example/types"
)

func Statements(invoice types.Invoice, plays types.Plays) (string, error) {
	var totalsAmount float64 = 0
	var volumeCredits float64 = 0
	result := fmt.Sprintf("Statements for %s", invoice.Cusomer)
	result = fmt.Sprintln(result)
	for _, perf := range invoice.Performance {
		play := plays[perf.PlayID]
		var thisAmount float64 = 0
		switch play.Type {
		case "tragedy":
			{
				thisAmount = 40000
				if perf.Audience > 30 {
					thisAmount += float64(1000 * (perf.Audience - 30))
				}
				break
			}
		case "comedy":
			{
				thisAmount = 30000
				if perf.Audience > 20 {
					thisAmount += float64(10000 + 500*(perf.Audience-20))
				}
				thisAmount += float64(300 * perf.Audience)
				break
			}
		default:
			{
				return "", errors.New(fmt.Sprintf("unknown tpye %s", play.Type))
			}
		}

		// add volume credits
		volumeCredits += math.Max(float64(perf.Audience-30), float64(0))
		// add extra credits for every ten comedy attendees
		if "comedy" == play.Type {
			volumeCredits += math.Floor(float64(perf.Audience) / float64(5))
		}

		// print line for this oder
		result += fmt.Sprintf("	%v: $%v  %v seats \n", play.Name, thisAmount/100, perf.Audience)
		totalsAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is $%v \n", totalsAmount/100)
	result += fmt.Sprintf("you earned $%v credits \n", volumeCredits)
	return result, nil
}
