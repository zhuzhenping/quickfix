//Package securitylist msg type = y.
package securitylist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a SecurityList wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//SecurityReqID is a non-required field for SecurityList.
func (m Message) SecurityReqID() (*field.SecurityReqIDField, quickfix.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityList.
func (m Message) GetSecurityReqID(f *field.SecurityReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a non-required field for SecurityList.
func (m Message) SecurityResponseID() (*field.SecurityResponseIDField, quickfix.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityList.
func (m Message) GetSecurityResponseID(f *field.SecurityResponseIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a non-required field for SecurityList.
func (m Message) SecurityRequestResult() (*field.SecurityRequestResultField, quickfix.MessageRejectError) {
	f := &field.SecurityRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from SecurityList.
func (m Message) GetSecurityRequestResult(f *field.SecurityRequestResultField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for SecurityList.
func (m Message) TotNoRelatedSym() (*field.TotNoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.TotNoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from SecurityList.
func (m Message) GetTotNoRelatedSym(f *field.TotNoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityList.
func (m Message) LastFragment() (*field.LastFragmentField, quickfix.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityList.
func (m Message) GetLastFragment(f *field.LastFragmentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityList.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityList.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityReportID is a non-required field for SecurityList.
func (m Message) SecurityReportID() (*field.SecurityReportIDField, quickfix.MessageRejectError) {
	f := &field.SecurityReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReportID reads a SecurityReportID from SecurityList.
func (m Message) GetSecurityReportID(f *field.SecurityReportIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for SecurityList.
func (m Message) ClearingBusinessDate() (*field.ClearingBusinessDateField, quickfix.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SecurityList.
func (m Message) GetClearingBusinessDate(f *field.ClearingBusinessDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for SecurityList.
func (m Message) MarketID() (*field.MarketIDField, quickfix.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from SecurityList.
func (m Message) GetMarketID(f *field.MarketIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for SecurityList.
func (m Message) MarketSegmentID() (*field.MarketSegmentIDField, quickfix.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from SecurityList.
func (m Message) GetMarketSegmentID(f *field.MarketSegmentIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SecurityList.
func (m Message) ApplID() (*field.ApplIDField, quickfix.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SecurityList.
func (m Message) GetApplID(f *field.ApplIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SecurityList.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SecurityList.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SecurityList.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SecurityList.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SecurityList.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, quickfix.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SecurityList.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListID is a non-required field for SecurityList.
func (m Message) SecurityListID() (*field.SecurityListIDField, quickfix.MessageRejectError) {
	f := &field.SecurityListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListID reads a SecurityListID from SecurityList.
func (m Message) GetSecurityListID(f *field.SecurityListIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListRefID is a non-required field for SecurityList.
func (m Message) SecurityListRefID() (*field.SecurityListRefIDField, quickfix.MessageRejectError) {
	f := &field.SecurityListRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListRefID reads a SecurityListRefID from SecurityList.
func (m Message) GetSecurityListRefID(f *field.SecurityListRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListDesc is a non-required field for SecurityList.
func (m Message) SecurityListDesc() (*field.SecurityListDescField, quickfix.MessageRejectError) {
	f := &field.SecurityListDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListDesc reads a SecurityListDesc from SecurityList.
func (m Message) GetSecurityListDesc(f *field.SecurityListDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityListDescLen is a non-required field for SecurityList.
func (m Message) EncodedSecurityListDescLen() (*field.EncodedSecurityListDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityListDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityListDescLen reads a EncodedSecurityListDescLen from SecurityList.
func (m Message) GetEncodedSecurityListDescLen(f *field.EncodedSecurityListDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityListDesc is a non-required field for SecurityList.
func (m Message) EncodedSecurityListDesc() (*field.EncodedSecurityListDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityListDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityListDesc reads a EncodedSecurityListDesc from SecurityList.
func (m Message) GetEncodedSecurityListDesc(f *field.EncodedSecurityListDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListType is a non-required field for SecurityList.
func (m Message) SecurityListType() (*field.SecurityListTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityListTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListType reads a SecurityListType from SecurityList.
func (m Message) GetSecurityListType(f *field.SecurityListTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListTypeSource is a non-required field for SecurityList.
func (m Message) SecurityListTypeSource() (*field.SecurityListTypeSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityListTypeSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListTypeSource reads a SecurityListTypeSource from SecurityList.
func (m Message) GetSecurityListTypeSource(f *field.SecurityListTypeSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SecurityList.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SecurityList.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for SecurityList.
func New() Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("y"))
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "y", r
}
