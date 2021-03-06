package eop

import (
	"github.com/jhillyerd/enmime"
	"github.com/olekukonko/tablewriter"
	"os"
)

type Parser struct {
	envelope *enmime.Envelope
	table    *tablewriter.Table

	Fields []*FilteringField
}

func NewParserWithEnvelop(env *enmime.Envelope) *Parser {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Header", "Field", "Field Meaning", "Value", "Value Meaning"})
	table.SetRowSeparator("-")
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	return &Parser{
		envelope: env,
		table:    table,
	}
}

func (parser *Parser) ParseAndRender(tableWidth int) {
	parser.ParseAntiSpamReport()
	parser.ParseMicrosoftAntiSpam()
	parser.ParseAuthenticationResult()

	parser.table.SetColWidth(tableWidth)
	parser.Render()
}

func (parser *Parser) Render() {
	for _, field := range parser.Fields {
		parser.table.Append(field.TableRow())
	}
	parser.table.Render()
}
