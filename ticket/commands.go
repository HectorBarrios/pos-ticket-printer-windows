package ticket

var (
	COMMAND_INIT             = []byte{27, 82, 7}
	COMMAND_FEED_1_LINE      = []byte{27, 100, 1}
	COMMAND_FEED_AND_CUT     = []byte{29, 'V', 66, 0}
	COMMAND_OPEN_CASH_DRAWER = []byte{27, 112, 0, 60, 120}
	COMMAND_LINE             = []byte{10}

	COMMAND_FONT_SIZE_NORMAL = []byte{29, 33, 0}
	COMMAND_FONT_SIZE_WIDTH1 = []byte{29, 33, 16}
	COMMAND_FONT_SIZE_WIDTH2 = []byte{29, 33, 32}
)
