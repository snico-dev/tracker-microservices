package models

//TLV data type
type TLV struct {
	Tag        string
	Length     int
	ValueArray []byte
}
