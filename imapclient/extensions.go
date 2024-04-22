package imapclient

import (
	"github.com/emersion/go-imap/v2"
)

type extensions struct {
	fetchExtensions map[string]FetchExtension
}

type FetchExtension interface {
	Attr() string
	Populate(dec WireDecoder) (*FetchItemDataExtension, error)
}

type SearchExtension interface {
	Attr() string
	Value() string
}

type WireDecoder interface {
	Err() error
	EOF() bool
	Expect(ok bool, name string) bool
	SP() bool
	ExpectSP() bool
	CRLF() bool
	ExpectCRLF() bool
	Func(ptr *string, valid func(ch byte) bool) bool
	Atom(ptr *string) bool
	ExpectAtom(ptr *string) bool
	ExpectNIL() bool
	Special(b byte) bool
	ExpectSpecial(b byte) bool
	Text(ptr *string) bool
	ExpectText(ptr *string) bool
	DiscardUntilByte(untilCh byte)
	DiscardLine()
	DiscardValue() bool
	Number(ptr *uint32) bool
	ExpectNumber(ptr *uint32) bool
	ExpectBodyFldOctets(ptr *uint32) bool
	Number64(ptr *int64) bool
	ExpectNumber64(ptr *int64) bool
	ModSeq(ptr *uint64) bool
	ExpectModSeq(ptr *uint64) bool
	Quoted(ptr *string) bool
	ExpectAString(ptr *string) bool
	String(ptr *string) bool
	ExpectString(ptr *string) bool
	ExpectNString(ptr *string) bool
	List(f func() error) (isList bool, err error)
	ExpectList(f func() error) error
	ExpectNList(f func() error) error
	ExpectMailbox(ptr *string) bool
	ExpectUID(ptr *imap.UID) bool
	ExpectUIDSet(ptr *imap.UIDSet) bool
	Literal(ptr *string) bool
}
