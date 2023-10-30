package parameters

import (
	"github.com/jung-kurt/gofpdf"
)

type Directive interface {
	Apply(p *gofpdf.Fpdf)
}

type GeneralDirectives struct {
	SetCreator       string  `json:"SetCreator,omitempty"`
	SetAuthor        string  `json:"SetAuthor,omitempty"`
	SetTitle         string  `json:"SetTitle,omitempty"`
	SetSubject       string  `json:"SetSubject,omitempty"`
	SetAutoPageBreak bool    `json:"SetAutoPageBreak,omitempty"`
	SetPrintHeader   bool    `json:"setPrintHeader,omitempty"`
	SetPrintFooter   bool    `json:"setPrintFooter,omitempty"`
	Margins          Margins `json:"margins,omitempty"`
}

func (g GeneralDirectives) Apply(p *gofpdf.Fpdf) {
	if g.SetCreator != "" {
		p.SetCreator(g.SetCreator, true)
	}
	if g.SetAuthor != "" {
		p.SetAuthor(g.SetAuthor, true)
	}
	if g.SetTitle != "" {
		p.SetTitle(g.SetTitle, true)
	}
	p.SetMargins(g.Margins.Left, g.Margins.Top, g.Margins.Right)
}

type TemplateDirectives struct {
	LeftTitle        CellFormats `json:"leftTitle"`
	RightTitle       CellFormats `json:"rightTitle"`
	EventDescription CellFormats `json:"eventDescription"`
	PrizeTitle       CellFormats `json:"prizeTitle"`
	PrizeList        CellFormats `json:"prizeList"`
	Price            Text        `json:"price"`
	Buyer            Text        `json:"buyer"`
	Seller           Text        `json:"seller"`
	VerticalLine     Line        `json:"verticalLine"`
	CutLine          Line        `json:"cutLine"`
	BuyerLine        Line        `json:"buyerLine"`
	SellerLine       Line        `json:"sellerLine"`
	RectEvent        Rect        `json:"rectEvent"`
	RectPrize        Rect        `json:"rectPrize"`
}

func (t TemplateDirectives) Apply(tpl gofpdf.Pdf) {
	t.CutLine.Apply(tpl)
	t.VerticalLine.Apply(tpl)
	t.BuyerLine.Apply(tpl)
	t.SellerLine.Apply(tpl)
	t.RectEvent.Apply(tpl)
	t.RectPrize.Apply(tpl)
	t.LeftTitle.Apply(tpl)
	t.RightTitle.Apply(tpl)
	t.EventDescription.Apply(tpl)
	t.PrizeTitle.Apply(tpl)
	t.PrizeList.Apply(tpl)
	t.Price.Apply(tpl)
	t.Buyer.Apply(tpl)
	t.Seller.Apply(tpl)
}

type CellFormats struct {
	Texts    []string  `json:"texts"`
	Width    float64   `json:"width"`
	Align    string    `json:"align"`
	Font     *Font     `json:"font,omitempty"`
	Position *Position `json:"position,omitempty"`
}

type Margins struct {
	Left  float64 `json:"left"`
	Top   float64 `json:"top"`
	Right float64 `json:"right"`
}

type Font struct {
	Family string  `json:"family"`
	Style  string  `json:"style"`
	Size   float64 `json:"size"`
}

type Text struct {
	Text     string    `json:"text"`
	Position *Position `json:"position"`
	Font     Font      `json:"font"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Line struct {
	DashPattern DashPattern `json:"dashPattern"`
	X1          float64     `json:"x1"`
	X2          float64     `json:"x2"`
	Y1          float64     `json:"y1"`
	Y2          float64     `json:"y2"`
}

type Rect struct {
	DashPattern DashPattern `json:"dashPattern"`
	X           float64     `json:"X"`
	Y           float64     `json:"Y"`
	W           float64     `json:"w"`
	H           float64     `json:"h"`
	Style       string      `json:"style"`
}

type DashPattern []float64

func (c CellFormats) Apply(pdf gofpdf.Pdf) {
	if c.Font != nil {
		pdf.SetFont(c.Font.Family, c.Font.Style, float64(c.Font.Size))
	}
	if c.Position != nil {
		pdf.SetXY(c.Position.X, c.Position.Y)
	}
	_, lineHt := pdf.GetFontSize()
	for _, text := range c.Texts {
		pdf.CellFormat(c.Width, lineHt, text, "0", 2, c.Align, false, 0, "")
	}
}

func (t Text) Apply(pdf gofpdf.Pdf) {
	pdf.SetFont(t.Font.Family, t.Font.Style, t.Font.Size)
	pdf.Text(t.Position.X, t.Position.Y, t.Text)
}

func (l Line) Apply(pdf gofpdf.Pdf) {
	pdf.SetDashPattern(l.DashPattern, 0)
	pdf.Line(l.X1, l.Y1, l.X2, l.Y2)
}

func (r Rect) Apply(pdf gofpdf.Pdf) {
	pdf.SetDashPattern(r.DashPattern, 0)
	pdf.Rect(r.X, r.Y, r.W, r.H, r.Style)
}
