package ticket

import (
	"bytes"

	"golang.org/x/text/encoding/charmap"

	"github.com/HectorBarrios/pos-ticket-printer-windows/ticket/internal/printer"
)

type ticket struct {
	printer *printer.Printer
	buff    *bytes.Buffer

	printerName  string
	documentName string
}

func (t *ticket) Print() error {
	_, err := t.printer.StartRawDocument(t.documentName, "")
	if err != nil {
		return err
	}
	defer t.printer.EndDocument()

	if err := t.printer.StartPage(); err != nil {
		return err
	}
	defer t.printer.EndPage()

	t.setText("This is an amazing text")
	t.newLine()
	t.setTextWidth1()
	t.setText("Hola soy Héctor €")
	t.newLine()
	t.setTextWidth2()
	t.setTextAndNewLine("######=ººº1!$%&$%&/$%&$%&$%")
	t.setTextAndNewLine("hector.br90@gmail.com")
	t.setTextNormal()
	t.setTextAndNewLine("hola @@ ## ## ### ####")
	t.feedAndCut()

	if _, err := t.printer.Write(t.buff.Bytes()); err != nil {
		return err
	}

	return t.printer.Close()
}

func (t *ticket) newLine() {
	t.buff.Write(COMMAND_LINE)
}

func (t *ticket) setTextAndNewLine(text string) {
	t.setText(text)
	t.newLine()
}

func (t *ticket) setText(text string) {
	result, _ := charmap.CodePage858.NewEncoder().String(text)
	t.buff.Write([]byte(result))
}

func (t *ticket) setTextNormal() {
	t.buff.Write(COMMAND_FONT_SIZE_NORMAL)
}

func (t *ticket) setTextWidth1() {
	t.buff.Write(COMMAND_FONT_SIZE_WIDTH1)
}

func (t *ticket) setTextWidth2() {
	t.buff.Write(COMMAND_FONT_SIZE_WIDTH2)
}

func (t *ticket) feedAndCut() {
	t.buff.Write(COMMAND_FEED_1_LINE)
	t.buff.Write(COMMAND_FEED_AND_CUT)
}

func NewTicket(opts ...TicketOption) (*ticket, error) {
	t := new(ticket)
	for _, o := range opts {
		o(t)
	}

	if err := validateTicket(t); err != nil {
		return nil, err
	}

	p, err := printer.Open(t.printerName)
	if err != nil {
		return nil, err
	}
	t.printer = p

	t.buff = bytes.NewBuffer([]byte{})
	if _, err := t.buff.Write(COMMAND_INIT); err != nil {
		return nil, err
	}

	return t, nil
}

func validateTicket(t *ticket) error {
	if t.printerName == "" {
		return ErrFieldPrinterNameNotProvided
	}

	return nil
}
