package document

import (
	"math"
)

type Numberer struct {
	batchCount       uint16
	startBatchIndex  uint16
	firstTicketIndex uint16
	ticketsPerBatch  uint16
	ticketsPerBook   uint16
	ticketsPerPage   uint16
}

func (n Numberer) GetTicketsPerPage() uint16 {
	return n.ticketsPerPage
}

func (n Numberer) GetFirstLastAndCount() (uint16, uint16, uint16) {
	count := n.batchCount * n.ticketsPerBatch
	last := n.firstTicketIndex + count - 1
	return n.firstTicketIndex, last, count
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
				ticketNumber = n.firstTicketIndex
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

func CreateNumberer(ticketPerBook uint16, ticketPerPage uint16, StartAt uint16, TicketsCount uint16) Numberer {
	numberer := Numberer{}
	numberer.ticketsPerBook = ticketPerBook
	numberer.ticketsPerPage = ticketPerPage
	numberer.ticketsPerBatch = ticketPerPage * ticketPerBook

	numberer.handleStartAndCount(StartAt, TicketsCount)

	return numberer
}

func (n *Numberer) handleStartAndCount(StartAt uint16, TicketsCount uint16) {
	// What will be the 1st ticket index ?
	// Start is minimum 0 OR startAtIndex - 1
	startAtIndex := math.Max(float64(StartAt)-1, float64(0))
	firstTicketIndex := math.Ceil(startAtIndex/float64(n.ticketsPerBatch))*float64(n.ticketsPerBatch) + 1
	n.firstTicketIndex = uint16(firstTicketIndex)
	// Now we can know how many batches to print
	n.batchCount = uint16(math.Ceil(float64(TicketsCount) / float64(n.ticketsPerBatch)))
	// Let's consider the shift param to get the actual 1st index
	startBatchIndex := math.Floor(float64(n.firstTicketIndex) / float64(n.ticketsPerBatch))
	n.startBatchIndex = uint16(startBatchIndex)
}
