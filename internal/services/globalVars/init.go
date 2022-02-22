package globalVars

import "tel-note/internal/protocol"

//init global variable
var AllContact protocol.ContactStorage

func Init() {
	AllContact = protocol.ContactStorage{}
}
