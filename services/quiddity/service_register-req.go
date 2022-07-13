/* For license and copyright information please see LEGAL file in repository */

package quiddity

import (
	"../../libgo/json"
	"../../libgo/protocol"
	"../../libgo/syllab"
)

/*
	Service response structure
*/

type RegisterRequest struct {
	quiddityID [16]byte
	domainID   [16]byte
}

func (req *RegisterRequest) QuiddityID() [16]byte       { return req.quiddityID }
func (req *RegisterRequest) DomainID() [16]byte         { return req.domainID }
func (req *RegisterRequest) SetQuiddityID(qid [16]byte) { req.quiddityID = qid }
func (req *RegisterRequest) SetDomainID(did [16]byte)   { req.domainID = did }

// Syallb codec
func (req *RegisterRequest) FromSyllab(payload []byte, stackIndex uint32) {
	if uint32(len(payload)) < req.LenOfSyllabStack() {
		err = syllab.ErrShortArrayDecode
		return
	}

	req.Language = protocol.Language(syllab.GetUInt32(payload, 0))
	req.URI = syllab.UnsafeGetString(payload, 4)
	req.Title = syllab.UnsafeGetString(payload, 12)
	return
}
func (req *RegisterRequest) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	syllab.SetUInt32(payload, 0, uint32(req.Language))
	heapIndex = syllab.SetString(payload, req.URI, 4, heapIndex)
	heapIndex = syllab.SetString(payload, req.Title, 12, heapIndex)
	return
}
func (req *RegisterRequest) LenOfSyllabStack() uint32 {
	return 20
}
func (req *RegisterRequest) LenOfSyllabHeap() (ln uint32) {
	ln += uint32(len(req.URI))
	ln += uint32(len(req.Title))
	return
}
func (req *RegisterRequest) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}

// JSON codec
func (req *RegisterRequest) FromJSON(payload []byte) (err protocol.Error) {
	var decoder = json.DecoderUnsafeMinified{Buf: payload}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "Language":
			var num uint32
			num, err = decoder.DecodeUInt32()
			req.Language = protocol.Language(num)
		case "URI":
			req.URI, err = decoder.DecodeString()
		case "Title":
			req.Title, err = decoder.DecodeString()
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if decoder.End() {
			return
		}
	}
	return
}
func (req *RegisterRequest) ToJSON(payload []byte) []byte {
	var encoder = json.Encoder{Buf: payload}

	encoder.EncodeString(`{"Language":`)
	encoder.EncodeUInt32(uint32(req.Language))

	encoder.EncodeString(`,"URI":"`)
	encoder.EncodeString(req.URI)

	encoder.EncodeString(`","Title":"`)
	encoder.EncodeString(req.Title)

	encoder.EncodeString(`"}`)
	return encoder.Buf
}
func (req *RegisterRequest) LenAsJSON() (ln int) {
	ln = len(req.URI) + len(req.Title)
	ln += 43
	return
}

/*
	Service response as Syllab
*/

type registerRequest_Syllab []byte

/*
	Service response as FlatBuffer
*/

type registerRequest_FlatBuffer []byte
