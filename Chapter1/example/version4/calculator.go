package version4

import (
	"math"

	"github.com/wrbz15/refacetoring-doc/Chapter1/example/types"
)
type PerformanceCalculatorInterface interface {
	GetAmount() float64
	GetVolumeCredits() float64
}

type PerformanceCalculator struct {
	APerformance types.Performance
	Play      types.Play
}

type TragedyCalculator struct {
	PerformanceCalculator
}

func (c *TragedyCalculator) GetAmount() float64 {
	var result float64 = 40000
	if c.APerformance.Audience > 30 {
		result += float64(1000 * (c.APerformance.Audience - 30))
	}
	return result
}

func (c *TragedyCalculator) GetVolumeCredits() float64 {
	return math.Max(float64(c.APerformance.Audience-30), float64(0))
}

type ComedyCalculator struct {
	PerformanceCalculator
}

func (c *ComedyCalculator) GetAmount() float64 {
	var result float64 = 30000
	if c.APerformance.Audience > 20 {
		result += float64(10000 + 500*(c.APerformance.Audience-20))
	}
	result += float64(300 * c.APerformance.Audience)
	return result
}

func (c *ComedyCalculator) GetVolumeCredits() float64 {
	result := math.Max(float64(c.APerformance.Audience-30), float64(0))
	result += math.Floor(float64(c.APerformance.Audience) / float64(5))
	return result
}

func createPerformanceCalculator(play types.Play, aPerformance types.Performance) PerformanceCalculatorInterface {
	switch play.Type {
	case "tragedy":
		return &TragedyCalculator{
			PerformanceCalculator:  PerformanceCalculator{
				APerformance: aPerformance,
				Play: play,
			},
		}
		case "comedy":
			return &ComedyCalculator{
				PerformanceCalculator:  PerformanceCalculator{
					APerformance: aPerformance,
					Play: play,
				},
			}
		default:
			panic("unknown play type")
	}
	
}