package main

import (
	"github.com/HectorBarrios/pos-ticket-printer-windows/ticket"
)

func main() {
	ticket, err := ticket.NewTicket(
		ticket.WithPrinterName("OKIPOS2"),
	)
	if err != nil {
		panic(err)
	}

	if err := ticket.Print(); err != nil {
		panic(err)
	}
}
