package ticket

type TicketOption func(*ticket)

func WithPrinterName(n string) func(*ticket) {
	return func(t *ticket) {
		t.printerName = n
	}
}

func WithDocumentName(n string) func(*ticket) {
	return func(t *ticket) {
		t.documentName = n
	}
}

func defaultOptions() []TicketOption {
	return []TicketOption{
		WithDocumentName("document"),
	}
}
