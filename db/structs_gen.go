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
	var zajw uint32
	zajw, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zajw != 11 {
		err = msgp.ArrayError{Wanted: 11, Got: zajw}
		return
	}
	err = dc.ReadExactBytes(z.ID[:])
	if err != nil {
		return
	}
	err = dc.ReadExactBytes(z.Stash[:])
	if err != nil {
		return
	}
	{
		var zwht uint32
		zwht, err = dc.ReadUint32()
		z.Name = StringHeapID(zwht)
	}
	if err != nil {
		return
	}
	{
		var zhct uint32
		zhct, err = dc.ReadUint32()
		z.TypeLine = StringHeapID(zhct)
	}
	if err != nil {
		return
	}
	{
		var zcua uint32
		zcua, err = dc.ReadUint32()
		z.Note = StringHeapID(zcua)
	}
	if err != nil {
		return
	}
	{
		var zxhx uint32
		zxhx, err = dc.ReadUint32()
		z.RootType = StringHeapID(zxhx)
	}
	if err != nil {
		return
	}
	{
		var zlqf uint32
		zlqf, err = dc.ReadUint32()
		z.RootFlavor = StringHeapID(zlqf)
	}
	if err != nil {
		return
	}
	{
		var zdaf uint16
		zdaf, err = dc.ReadUint16()
		z.League = LeagueHeapID(zdaf)
	}
	if err != nil {
		return
	}
	z.Corrupted, err = dc.ReadBool()
	if err != nil {
		return
	}
	z.Identified, err = dc.ReadBool()
	if err != nil {
		return
	}
	var zpks uint32
	zpks, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zpks) {
		z.Mods = (z.Mods)[:zpks]
	} else {
		z.Mods = make([]ItemMod, zpks)
	}
	for zcmr := range z.Mods {
		err = z.Mods[zcmr].DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Item) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 11
	err = en.Append(0x9b)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.ID[:])
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Stash[:])
	if err != nil {
		return
	}
	err = en.WriteUint32(uint32(z.Name))
	if err != nil {
		return
	}
	err = en.WriteUint32(uint32(z.TypeLine))
	if err != nil {
		return
	}
	err = en.WriteUint32(uint32(z.Note))
	if err != nil {
		return
	}
	err = en.WriteUint32(uint32(z.RootType))
	if err != nil {
		return
	}
	err = en.WriteUint32(uint32(z.RootFlavor))
	if err != nil {
		return
	}
	err = en.WriteUint16(uint16(z.League))
	if err != nil {
		return
	}
	err = en.WriteBool(z.Corrupted)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Identified)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Mods)))
	if err != nil {
		return
	}
	for zcmr := range z.Mods {
		err = z.Mods[zcmr].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Item) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 11
	o = append(o, 0x9b)
	o = msgp.AppendBytes(o, z.ID[:])
	o = msgp.AppendBytes(o, z.Stash[:])
	o = msgp.AppendUint32(o, uint32(z.Name))
	o = msgp.AppendUint32(o, uint32(z.TypeLine))
	o = msgp.AppendUint32(o, uint32(z.Note))
	o = msgp.AppendUint32(o, uint32(z.RootType))
	o = msgp.AppendUint32(o, uint32(z.RootFlavor))
	o = msgp.AppendUint16(o, uint16(z.League))
	o = msgp.AppendBool(o, z.Corrupted)
	o = msgp.AppendBool(o, z.Identified)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Mods)))
	for zcmr := range z.Mods {
		o, err = z.Mods[zcmr].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zjfb uint32
	zjfb, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zjfb != 11 {
		err = msgp.ArrayError{Wanted: 11, Got: zjfb}
		return
	}
	bts, err = msgp.ReadExactBytes(bts, z.ID[:])
	if err != nil {
		return
	}
	bts, err = msgp.ReadExactBytes(bts, z.Stash[:])
	if err != nil {
		return
	}
	{
		var zcxo uint32
		zcxo, bts, err = msgp.ReadUint32Bytes(bts)
		z.Name = StringHeapID(zcxo)
	}
	if err != nil {
		return
	}
	{
		var zeff uint32
		zeff, bts, err = msgp.ReadUint32Bytes(bts)
		z.TypeLine = StringHeapID(zeff)
	}
	if err != nil {
		return
	}
	{
		var zrsw uint32
		zrsw, bts, err = msgp.ReadUint32Bytes(bts)
		z.Note = StringHeapID(zrsw)
	}
	if err != nil {
		return
	}
	{
		var zxpk uint32
		zxpk, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootType = StringHeapID(zxpk)
	}
	if err != nil {
		return
	}
	{
		var zdnj uint32
		zdnj, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootFlavor = StringHeapID(zdnj)
	}
	if err != nil {
		return
	}
	{
		var zobc uint16
		zobc, bts, err = msgp.ReadUint16Bytes(bts)
		z.League = LeagueHeapID(zobc)
	}
	if err != nil {
		return
	}
	z.Corrupted, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		return
	}
	z.Identified, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		return
	}
	var zsnv uint32
	zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zsnv) {
		z.Mods = (z.Mods)[:zsnv]
	} else {
		z.Mods = make([]ItemMod, zsnv)
	}
	for zcmr := range z.Mods {
		bts, err = z.Mods[zcmr].UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Item) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint16Size + msgp.BoolSize + msgp.BoolSize + msgp.ArrayHeaderSize
	for zcmr := range z.Mods {
		s += z.Mods[zcmr].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ItemMod) DecodeMsg(dc *msgp.Reader) (err error) {
	var zema uint32
	zema, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zema != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zema}
		return
	}
	{
		var zpez uint32
		zpez, err = dc.ReadUint32()
		z.Mod = StringHeapID(zpez)
	}
	if err != nil {
		return
	}
	var zqke uint32
	zqke, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Values) >= int(zqke) {
		z.Values = (z.Values)[:zqke]
	} else {
		z.Values = make([]int, zqke)
	}
	for zkgt := range z.Values {
		z.Values[zkgt], err = dc.ReadInt()
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ItemMod) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return err
	}
	err = en.WriteUint32(uint32(z.Mod))
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Values)))
	if err != nil {
		return
	}
	for zkgt := range z.Values {
		err = en.WriteInt(z.Values[zkgt])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ItemMod) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendUint32(o, uint32(z.Mod))
	o = msgp.AppendArrayHeader(o, uint32(len(z.Values)))
	for zkgt := range z.Values {
		o = msgp.AppendInt(o, z.Values[zkgt])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ItemMod) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zqyh uint32
	zqyh, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zqyh != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zqyh}
		return
	}
	{
		var zyzr uint32
		zyzr, bts, err = msgp.ReadUint32Bytes(bts)
		z.Mod = StringHeapID(zyzr)
	}
	if err != nil {
		return
	}
	var zywj uint32
	zywj, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Values) >= int(zywj) {
		z.Values = (z.Values)[:zywj]
	} else {
		z.Values = make([]int, zywj)
	}
	for zkgt := range z.Values {
		z.Values[zkgt], bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ItemMod) Msgsize() (s int) {
	s = 1 + msgp.Uint32Size + msgp.ArrayHeaderSize + (len(z.Values) * (msgp.IntSize))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LeagueHeapID) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zjpj uint16
		zjpj, err = dc.ReadUint16()
		(*z) = LeagueHeapID(zjpj)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z LeagueHeapID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint16(uint16(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LeagueHeapID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint16(o, uint16(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LeagueHeapID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zzpf uint16
		zzpf, bts, err = msgp.ReadUint16Bytes(bts)
		(*z) = LeagueHeapID(zzpf)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z LeagueHeapID) Msgsize() (s int) {
	s = msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Stash) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zgmo uint32
	zgmo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zgmo > 0 {
		zgmo--
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
	var ztaf uint32
	ztaf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztaf > 0 {
		ztaf--
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
		var zeth uint32
		zeth, err = dc.ReadUint32()
		(*z) = StringHeapID(zeth)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z StringHeapID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint32(uint32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z StringHeapID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint32(o, uint32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StringHeapID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zsbz uint32
		zsbz, bts, err = msgp.ReadUint32Bytes(bts)
		(*z) = StringHeapID(zsbz)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z StringHeapID) Msgsize() (s int) {
	s = msgp.Uint32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Timestamp) DecodeMsg(dc *msgp.Reader) (err error) {
	err = dc.ReadExactBytes(z[:])
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Timestamp) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes(z[:])
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Timestamp) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, z[:])
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Timestamp) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, z[:])
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Timestamp) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (TimestampSize * (msgp.ByteSize))
	return
}
