package db

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *ID) DecodeMsg(dc *msgp.Reader) (err error) {
	err = dc.ReadExactBytes(z[:])
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes(z[:])
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, z[:])
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, z[:])
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ID) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Item) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			err = dc.ReadExactBytes(z.ID[:])
			if err != nil {
				return
			}
		case "Stash":
			err = dc.ReadExactBytes(z.Stash[:])
			if err != nil {
				return
			}
		case "Name":
			{
				var zajw uint64
				zajw, err = dc.ReadUint64()
				z.Name = StringHeapID(zajw)
			}
			if err != nil {
				return
			}
		case "TypeLine":
			{
				var zwht uint64
				zwht, err = dc.ReadUint64()
				z.TypeLine = StringHeapID(zwht)
			}
			if err != nil {
				return
			}
		case "Note":
			{
				var zhct uint64
				zhct, err = dc.ReadUint64()
				z.Note = StringHeapID(zhct)
			}
			if err != nil {
				return
			}
		case "League":
			{
				var zcua uint64
				zcua, err = dc.ReadUint64()
				z.League = StringHeapID(zcua)
			}
			if err != nil {
				return
			}
		case "RootType":
			{
				var zxhx uint64
				zxhx, err = dc.ReadUint64()
				z.RootType = StringHeapID(zxhx)
			}
			if err != nil {
				return
			}
		case "RootFlavor":
			{
				var zlqf uint64
				zlqf, err = dc.ReadUint64()
				z.RootFlavor = StringHeapID(zlqf)
			}
			if err != nil {
				return
			}
		case "Corrupted":
			z.Corrupted, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Identified":
			z.Identified, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Item) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "ID"
	err = en.Append(0x8a, 0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.ID[:])
	if err != nil {
		return
	}
	// write "Stash"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x73, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Stash[:])
	if err != nil {
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Name))
	if err != nil {
		return
	}
	// write "TypeLine"
	err = en.Append(0xa8, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.TypeLine))
	if err != nil {
		return
	}
	// write "Note"
	err = en.Append(0xa4, 0x4e, 0x6f, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Note))
	if err != nil {
		return
	}
	// write "League"
	err = en.Append(0xa6, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.League))
	if err != nil {
		return
	}
	// write "RootType"
	err = en.Append(0xa8, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.RootType))
	if err != nil {
		return
	}
	// write "RootFlavor"
	err = en.Append(0xaa, 0x52, 0x6f, 0x6f, 0x74, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.RootFlavor))
	if err != nil {
		return
	}
	// write "Corrupted"
	err = en.Append(0xa9, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Corrupted)
	if err != nil {
		return
	}
	// write "Identified"
	err = en.Append(0xaa, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Identified)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Item) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "ID"
	o = append(o, 0x8a, 0xa2, 0x49, 0x44)
	o = msgp.AppendBytes(o, z.ID[:])
	// string "Stash"
	o = append(o, 0xa5, 0x53, 0x74, 0x61, 0x73, 0x68)
	o = msgp.AppendBytes(o, z.Stash[:])
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendUint64(o, uint64(z.Name))
	// string "TypeLine"
	o = append(o, 0xa8, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65)
	o = msgp.AppendUint64(o, uint64(z.TypeLine))
	// string "Note"
	o = append(o, 0xa4, 0x4e, 0x6f, 0x74, 0x65)
	o = msgp.AppendUint64(o, uint64(z.Note))
	// string "League"
	o = append(o, 0xa6, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65)
	o = msgp.AppendUint64(o, uint64(z.League))
	// string "RootType"
	o = append(o, 0xa8, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendUint64(o, uint64(z.RootType))
	// string "RootFlavor"
	o = append(o, 0xaa, 0x52, 0x6f, 0x6f, 0x74, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x72)
	o = msgp.AppendUint64(o, uint64(z.RootFlavor))
	// string "Corrupted"
	o = append(o, 0xa9, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Corrupted)
	// string "Identified"
	o = append(o, 0xaa, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Identified)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdaf uint32
	zdaf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdaf > 0 {
		zdaf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			bts, err = msgp.ReadExactBytes(bts, z.ID[:])
			if err != nil {
				return
			}
		case "Stash":
			bts, err = msgp.ReadExactBytes(bts, z.Stash[:])
			if err != nil {
				return
			}
		case "Name":
			{
				var zpks uint64
				zpks, bts, err = msgp.ReadUint64Bytes(bts)
				z.Name = StringHeapID(zpks)
			}
			if err != nil {
				return
			}
		case "TypeLine":
			{
				var zjfb uint64
				zjfb, bts, err = msgp.ReadUint64Bytes(bts)
				z.TypeLine = StringHeapID(zjfb)
			}
			if err != nil {
				return
			}
		case "Note":
			{
				var zcxo uint64
				zcxo, bts, err = msgp.ReadUint64Bytes(bts)
				z.Note = StringHeapID(zcxo)
			}
			if err != nil {
				return
			}
		case "League":
			{
				var zeff uint64
				zeff, bts, err = msgp.ReadUint64Bytes(bts)
				z.League = StringHeapID(zeff)
			}
			if err != nil {
				return
			}
		case "RootType":
			{
				var zrsw uint64
				zrsw, bts, err = msgp.ReadUint64Bytes(bts)
				z.RootType = StringHeapID(zrsw)
			}
			if err != nil {
				return
			}
		case "RootFlavor":
			{
				var zxpk uint64
				zxpk, bts, err = msgp.ReadUint64Bytes(bts)
				z.RootFlavor = StringHeapID(zxpk)
			}
			if err != nil {
				return
			}
		case "Corrupted":
			z.Corrupted, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Identified":
			z.Identified, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Item) Msgsize() (s int) {
	s = 1 + 3 + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + 6 + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + 5 + msgp.Uint64Size + 9 + msgp.Uint64Size + 5 + msgp.Uint64Size + 7 + msgp.Uint64Size + 9 + msgp.Uint64Size + 11 + msgp.Uint64Size + 10 + msgp.BoolSize + 11 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Stash) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zobc uint32
	zobc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zobc > 0 {
		zobc--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			err = dc.ReadExactBytes(z.ID[:])
			if err != nil {
				return
			}
		case "AccountName":
			z.AccountName, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Stash) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "ID"
	err = en.Append(0x82, 0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.ID[:])
	if err != nil {
		return
	}
	// write "AccountName"
	err = en.Append(0xab, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AccountName)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Stash) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "ID"
	o = append(o, 0x82, 0xa2, 0x49, 0x44)
	o = msgp.AppendBytes(o, z.ID[:])
	// string "AccountName"
	o = append(o, 0xab, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.AccountName)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Stash) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsnv uint32
	zsnv, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsnv > 0 {
		zsnv--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			bts, err = msgp.ReadExactBytes(bts, z.ID[:])
			if err != nil {
				return
			}
		case "AccountName":
			z.AccountName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Stash) Msgsize() (s int) {
	s = 1 + 3 + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + 12 + msgp.StringPrefixSize + len(z.AccountName)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StringHeapID) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zkgt uint64
		zkgt, err = dc.ReadUint64()
		(*z) = StringHeapID(zkgt)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z StringHeapID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint64(uint64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z StringHeapID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint64(o, uint64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StringHeapID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zema uint64
		zema, bts, err = msgp.ReadUint64Bytes(bts)
		(*z) = StringHeapID(zema)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z StringHeapID) Msgsize() (s int) {
	s = msgp.Uint64Size
	return
}
