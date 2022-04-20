package version3

import (
	"fmt"

	"github.com/wrbz15/refacetoring-doc/Chapter1/example/types"
)

func Statements(invoice types.Invoice, plays types.Plays) string {
	return renderPlainText(createStatementdata(invoice, plays))
}

func renderPlainText(data *StatementsData) string {
	result := fmt.Sprintf("Statements for %s \n", data.Consumer)
	for _, perf := range data.Performances {
		result += fmt.Sprintf("	%v: $%v  %v seats \n", perf.Play.Name, perf.Amount/100, perf.APerformance.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%v \n", data.TotalAmount/100)
	result += fmt.Sprintf("you earned $%v credits \n", data.TotalVolumeCredits)
	return result
}
func HtmlStatement(invoice types.Invoice, plays types.Plays) string {
	return renderHtml(createStatementdata(invoice, plays))
}

func renderHtml(data *StatementsData) string {
	result := fmt.Sprintf("<h1>Statement for %s</h1>\n", data.Consumer)
	result += "<table>\n"
	result += "<tr><th>play</th><th>seats</th><th>cost</th></tr>"
	for _, perf := range data.Performances {
		result += fmt.Sprintf(" <tr><td>%s</td><td>%v</td>", perf.Play.Name, perf.APerformance.Audience)
		result += fmt.Sprintf("<td>%v</td></tr>\n", perf.Amount)
	}
	result += "</table>\n"
	result += fmt.Sprintf("<p>Amount owed is <em>%v</em></p>\n", data.TotalAmount/100)
	result += fmt.Sprintf("<p>You earned <em>%v</em> credits</p>\n", data.TotalVolumeCredits)
	return result
}
