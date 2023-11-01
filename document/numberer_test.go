package document

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumberer_GetNumbering(t *testing.T) {
	type test struct {
		PerBook           uint16
		PerPage           uint16
		TicketsCount      uint16
		StartAt           uint16
		ExpectedNumbering [][]uint16
	}
	tests := []test{
		{2, 2, 1, 1, [][]uint16{{1, 3}, {2, 4}}},
		{2, 2, 1, 4, [][]uint16{{5, 7}, {6, 8}}},
		{2, 2, 1, 5, [][]uint16{{5, 7}, {6, 8}}},
		{2, 2, 4, 1, [][]uint16{{1, 3}, {2, 4}}},
		{2, 2, 5, 1, [][]uint16{{1, 3}, {2, 4}, {5, 7}, {6, 8}}},
	}
	for _, test := range tests {
		numberer := CreateNumberer(test.PerBook, test.PerPage, test.StartAt, test.TicketsCount)
		numbering := numberer.GetNumbering()
		if !reflect.DeepEqual(test.ExpectedNumbering, numbering) {
			fmt.Println("expected numbering ", test.ExpectedNumbering, "actual: ", numbering)
			t.Fail()
		}
	}
}

func TestNumberer_CreateNumberer_BatchCount(t *testing.T) {
	type test struct {
		PerBook            uint16
		PerPage            uint16
		TicketsCount       uint16
		ExpectedBatchCount uint16
	}
	tests := []test{
		{2, 2, 1, 1},
		{2, 2, 4, 1},
		{2, 2, 5, 2},
		{10, 3, 10, 1},
		{10, 3, 30, 1},
		{10, 3, 31, 2},
	}
	for _, test := range tests {
		numberer := CreateNumberer(test.PerBook, test.PerPage, 1, test.TicketsCount)
		if numberer.batchCount != test.ExpectedBatchCount {
			fmt.Printf(
				"Batch count should be %d\nfor %d ticket books with %d tickets per page\nwhen asking for %d tickets\nIt is actually %d\n",
				test.ExpectedBatchCount,
				test.PerBook,
				test.PerPage,
				test.TicketsCount,
				numberer.batchCount,
			)
			t.Fail()
		}
	}
}

func TestNumberer_CreateNumberer_FirstTicket(t *testing.T) {
	type test struct {
		PerBook             uint16
		PerPage             uint16
		TicketsCount        uint16
		StartAt             uint16
		ExpectedFirstTicket uint16
	}
	tests := []test{
		{2, 2, 1, 1, 1},
		{2, 2, 1, 3, 5},
		{10, 3, 10, 0, 1},
		{10, 3, 10, 1, 1},
		{10, 3, 10, 2, 31},
		{10, 3, 10, 100, 121},
	}
	for _, test := range tests {
		numberer := CreateNumberer(test.PerBook, test.PerPage, test.StartAt, test.TicketsCount)
		if numberer.firstTicketIndex != test.ExpectedFirstTicket {
			fmt.Printf(
				"First Ticket Index should be %d\nfor %d ticket books with %d tickets per page\nwhen asking for %d tickets, starting at %d\nIt is actually %d\n",
				test.ExpectedFirstTicket,
				test.PerBook,
				test.PerPage,
				test.TicketsCount,
				test.StartAt,
				numberer.firstTicketIndex,
			)
			t.Fail()
		}
	}
}
