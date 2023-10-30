package document

import (
	"math"
)

type Numberer struct {
	batchCount       uint16
	startBatchIndex  uint16
	firstTicketIndex uint16
	ticketsPerBatch  uint16
	numberOffset     uint16
	ticketsPerBook   uint16
	ticketsPerPage   uint16
}

func (n *Numberer) SetNumberOffset(numberOffset uint16) {
	n.numberOffset = numberOffset
}

func (n *Numberer) SetTicketsPerBatch(ticketPerBatch uint16) {
	n.ticketsPerBatch = ticketPerBatch
}

func (n Numberer) GetTicketsPerPage() uint16 {
	return n.ticketsPerPage
}

func (n Numberer) GetStartBatchIndex() uint16 {
	return n.startBatchIndex
}

func (n Numberer) GetFirstTicketNumber() uint16 {
	return n.firstTicketIndex + n.numberOffset
}

func (n Numberer) GetNumbering() [][]uint16 {
	res := make([][]uint16, int(n.batchCount)*int(n.ticketsPerBook))

	for batchI := 0; batchI < int(n.batchCount); batchI++ {
		// Loop on the pages inside one batch
		for pageI := 0; pageI < int(n.ticketsPerBook); pageI++ {
			// Loop on the tickets inside one page
			newPage := make([]uint16, n.ticketsPerPage)
			for ticketI := 0; ticketI < int(n.ticketsPerPage); ticketI++ {
				var ticketNumber uint16
				ticketNumber = n.GetFirstTicketNumber()
				ticketNumber += uint16(batchI) * uint16(n.ticketsPerBatch)
				ticketNumber += uint16(pageI)
				ticketNumber += uint16(ticketI * int(n.ticketsPerBook))
				newPage[ticketI] = ticketNumber
			}
			res[batchI*int(n.ticketsPerBook)+pageI] = newPage
		}
	}

	return res
}

func (n *Numberer) SetBookParam(bookParam BookParam) {
	// What will be the 1st ticket index ?
	firstTicketIndex := math.Ceil(float64(bookParam.Start-uint16(1)) / float64(n.ticketsPerBatch))
	firstTicketIndex *= float64(n.ticketsPerBatch)
	firstTicketIndex += float64(1)
	n.firstTicketIndex = uint16(firstTicketIndex)
	// How many books shall we generate ?
	bookCount := math.Ceil(float64(bookParam.Count)/float64(n.ticketsPerPage)) * float64(n.ticketsPerPage)
	// Now we can know how many batches to print
	n.batchCount = uint16(bookCount / float64(n.ticketsPerPage))
	// Let's consider the shift param to get the actual 1st index
	startBatchIndex := float64(n.firstTicketIndex)/float64(n.ticketsPerBatch) - float64(1)
	n.startBatchIndex = uint16(startBatchIndex)
}

func CreateNumberer(ticketPerBook uint16, ticketPerPage uint16, startAt uint16) Numberer {
	numberer := Numberer{}
	numberer.ticketsPerBook = ticketPerBook
	numberer.ticketsPerPage = ticketPerPage
	numberer.SetNumberOffset(startAt - uint16(1))
	numberer.SetTicketsPerBatch(ticketPerPage * ticketPerBook)

	return numberer
}
