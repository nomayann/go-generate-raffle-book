package document

import (
	"fmt"
	"generate-raffle-book/parameters"

	"github.com/jung-kurt/gofpdf"
)

type Generator struct {
	Numberer     Numberer
	NumberLength uint8
	Pdf          *gofpdf.Fpdf
	Parameters   *parameters.Parameters
}

func (g *Generator) Generate() {
	g.Pdf = gofpdf.New("P", "mm", "A4", "")
	g.initFonts()
	g.Parameters.GeneralDirectives.Apply(g.Pdf)
	_, pageH := g.Pdf.GetPageSize()
	_, marginT, _, _ := g.Pdf.GetMargins()
	ticketH := (pageH - (marginT * float64(g.Numberer.GetTicketsPerPage()) * 2.)) / float64(g.Numberer.GetTicketsPerPage())
	// create template
	template := g.Pdf.CreateTemplate(func(tpl *gofpdf.Tpl) {
		g.Parameters.TemplateDirectives.Apply(tpl)
	})
	_, tplSize := template.Size()
	for _, pageNumbers := range g.Numberer.GetNumbering() {
		g.Pdf.AddPage()
		_, fontH := g.Pdf.GetFontSize()
		numberFormat := fmt.Sprint("N° : %0", g.NumberLength, "d")
		for i, pageNumber := range pageNumbers {
			ticketX := 0.
			ticketY := ticketH * float64(i)
			numberY := ticketY + ticketH - fontH
			g.Pdf.UseTemplateScaled(template, gofpdf.PointType{X: ticketX, Y: ticketY}, tplSize)
			g.Pdf.Text(30, numberY, fmt.Sprintf(numberFormat, pageNumber))
			g.Pdf.Text(175, numberY, fmt.Sprintf(numberFormat, pageNumber))
		}
	}
	err := g.Pdf.OutputFileAndClose("build/tickets.pdf")
	if err != nil {
		fmt.Println("Erreur inattendue: ", err.Error())
	}
}

func (g Generator) initFonts() {
	g.Pdf.AddUTF8Font("ubuntu", "", "fonts/ubuntu/Ubuntu-R.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "B", "fonts/ubuntu/Ubuntu-B.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "BI", "fonts/ubuntu/Ubuntu-BI.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "C", "fonts/ubuntu/Ubuntu-C.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "L", "fonts/ubuntu/Ubuntu-L.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "LI", "fonts/ubuntu/Ubuntu-LI.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "M", "fonts/ubuntu/Ubuntu-M.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "MI", "fonts/ubuntu/Ubuntu-MI.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "R", "fonts/ubuntu/Ubuntu-R.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "RI", "fonts/ubuntu/Ubuntu-RI.ttf")
	g.Pdf.AddUTF8Font("ubuntu", "Th", "fonts/ubuntu/Ubuntu-Th.ttf")
	g.Pdf.AddUTF8Font("ubuntuMono", "B", "fonts/ubuntu/UbuntuMono-B.ttf")
	g.Pdf.AddUTF8Font("ubuntuMono", "BI", "fonts/ubuntu/UbuntuMono-BI.ttf")
	g.Pdf.AddUTF8Font("ubuntuMono", "R", "fonts/ubuntu/UbuntuMono-R.ttf")
	g.Pdf.AddUTF8Font("ubuntuMono", "RI", "fonts/ubuntu/UbuntuMono-RI.ttf")
	g.Pdf.SetFont("ubuntu", "R", 15)
}
