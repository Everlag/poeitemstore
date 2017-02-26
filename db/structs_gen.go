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
	var zwht uint32
	zwht, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zwht != 12 {
		err = msgp.ArrayError{Wanted: 12, Got: zwht}
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
		var zhct uint32
		zhct, err = dc.ReadUint32()
		z.Name = StringHeapID(zhct)
	}
	if err != nil {
		return
	}
	{
		var zcua uint32
		zcua, err = dc.ReadUint32()
		z.TypeLine = StringHeapID(zcua)
	}
	if err != nil {
		return
	}
	{
		var zxhx uint32
		zxhx, err = dc.ReadUint32()
		z.Note = StringHeapID(zxhx)
	}
	if err != nil {
		return
	}
	{
		var zlqf uint32
		zlqf, err = dc.ReadUint32()
		z.RootType = StringHeapID(zlqf)
	}
	if err != nil {
		return
	}
	{
		var zdaf uint32
		zdaf, err = dc.ReadUint32()
		z.RootFlavor = StringHeapID(zdaf)
	}
	if err != nil {
		return
	}
	{
		var zpks uint16
		zpks, err = dc.ReadUint16()
		z.League = LeagueHeapID(zpks)
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
	var zjfb uint32
	zjfb, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zjfb) {
		z.Mods = (z.Mods)[:zjfb]
	} else {
		z.Mods = make([]ItemMod, zjfb)
	}
	for zcmr := range z.Mods {
		var zcxo uint32
		zcxo, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if zcxo != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zcxo}
			return
		}
		{
			var zeff uint32
			zeff, err = dc.ReadUint32()
			z.Mods[zcmr].Mod = StringHeapID(zeff)
		}
		if err != nil {
			return
		}
		var zrsw uint32
		zrsw, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if cap(z.Mods[zcmr].Values) >= int(zrsw) {
			z.Mods[zcmr].Values = (z.Mods[zcmr].Values)[:zrsw]
		} else {
			z.Mods[zcmr].Values = make([]uint16, zrsw)
		}
		for zajw := range z.Mods[zcmr].Values {
			z.Mods[zcmr].Values[zajw], err = dc.ReadUint16()
			if err != nil {
				return
			}
		}
	}
	z.UpdateSequence, err = dc.ReadUint16()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Item) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 12
	err = en.Append(0x9c)
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
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return err
		}
		err = en.WriteUint32(uint32(z.Mods[zcmr].Mod))
		if err != nil {
			return
		}
		err = en.WriteArrayHeader(uint32(len(z.Mods[zcmr].Values)))
		if err != nil {
			return
		}
		for zajw := range z.Mods[zcmr].Values {
			err = en.WriteUint16(z.Mods[zcmr].Values[zajw])
			if err != nil {
				return
			}
		}
	}
	err = en.WriteUint16(z.UpdateSequence)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Item) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 12
	o = append(o, 0x9c)
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
		// array header, size 2
		o = append(o, 0x92)
		o = msgp.AppendUint32(o, uint32(z.Mods[zcmr].Mod))
		o = msgp.AppendArrayHeader(o, uint32(len(z.Mods[zcmr].Values)))
		for zajw := range z.Mods[zcmr].Values {
			o = msgp.AppendUint16(o, z.Mods[zcmr].Values[zajw])
		}
	}
	o = msgp.AppendUint16(o, z.UpdateSequence)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zxpk uint32
	zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zxpk != 12 {
		err = msgp.ArrayError{Wanted: 12, Got: zxpk}
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
		var zdnj uint32
		zdnj, bts, err = msgp.ReadUint32Bytes(bts)
		z.Name = StringHeapID(zdnj)
	}
	if err != nil {
		return
	}
	{
		var zobc uint32
		zobc, bts, err = msgp.ReadUint32Bytes(bts)
		z.TypeLine = StringHeapID(zobc)
	}
	if err != nil {
		return
	}
	{
		var zsnv uint32
		zsnv, bts, err = msgp.ReadUint32Bytes(bts)
		z.Note = StringHeapID(zsnv)
	}
	if err != nil {
		return
	}
	{
		var zkgt uint32
		zkgt, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootType = StringHeapID(zkgt)
	}
	if err != nil {
		return
	}
	{
		var zema uint32
		zema, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootFlavor = StringHeapID(zema)
	}
	if err != nil {
		return
	}
	{
		var zpez uint16
		zpez, bts, err = msgp.ReadUint16Bytes(bts)
		z.League = LeagueHeapID(zpez)
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
	var zqke uint32
	zqke, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zqke) {
		z.Mods = (z.Mods)[:zqke]
	} else {
		z.Mods = make([]ItemMod, zqke)
	}
	for zcmr := range z.Mods {
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
			z.Mods[zcmr].Mod = StringHeapID(zyzr)
		}
		if err != nil {
			return
		}
		var zywj uint32
		zywj, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if cap(z.Mods[zcmr].Values) >= int(zywj) {
			z.Mods[zcmr].Values = (z.Mods[zcmr].Values)[:zywj]
		} else {
			z.Mods[zcmr].Values = make([]uint16, zywj)
		}
		for zajw := range z.Mods[zcmr].Values {
			z.Mods[zcmr].Values[zajw], bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		}
	}
	z.UpdateSequence, bts, err = msgp.ReadUint16Bytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Item) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint16Size + msgp.BoolSize + msgp.BoolSize + msgp.ArrayHeaderSize
	for zcmr := range z.Mods {
		s += 1 + msgp.Uint32Size + msgp.ArrayHeaderSize + (len(z.Mods[zcmr].Values) * (msgp.Uint16Size))
	}
	s += msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ItemMod) DecodeMsg(dc *msgp.Reader) (err error) {
	var zzpf uint32
	zzpf, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zzpf != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zzpf}
		return
	}
	{
		var zrfe uint32
		zrfe, err = dc.ReadUint32()
		z.Mod = StringHeapID(zrfe)
	}
	if err != nil {
		return
	}
	var zgmo uint32
	zgmo, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Values) >= int(zgmo) {
		z.Values = (z.Values)[:zgmo]
	} else {
		z.Values = make([]uint16, zgmo)
	}
	for zjpj := range z.Values {
		z.Values[zjpj], err = dc.ReadUint16()
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
	for zjpj := range z.Values {
		err = en.WriteUint16(z.Values[zjpj])
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
	for zjpj := range z.Values {
		o = msgp.AppendUint16(o, z.Values[zjpj])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ItemMod) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var ztaf uint32
	ztaf, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if ztaf != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: ztaf}
		return
	}
	{
		var zeth uint32
		zeth, bts, err = msgp.ReadUint32Bytes(bts)
		z.Mod = StringHeapID(zeth)
	}
	if err != nil {
		return
	}
	var zsbz uint32
	zsbz, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Values) >= int(zsbz) {
		z.Values = (z.Values)[:zsbz]
	} else {
		z.Values = make([]uint16, zsbz)
	}
	for zjpj := range z.Values {
		z.Values[zjpj], bts, err = msgp.ReadUint16Bytes(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ItemMod) Msgsize() (s int) {
	s = 1 + msgp.Uint32Size + msgp.ArrayHeaderSize + (len(z.Values) * (msgp.Uint16Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LeagueHeapID) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zrjx uint16
		zrjx, err = dc.ReadUint16()
		(*z) = LeagueHeapID(zrjx)
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
		var zawn uint16
		zawn, bts, err = msgp.ReadUint16Bytes(bts)
		(*z) = LeagueHeapID(zawn)
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
	var zrbe uint32
	zrbe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrbe > 0 {
		zrbe--
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
	var zmfd uint32
	zmfd, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmfd > 0 {
		zmfd--
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
		var zzdc uint32
		zzdc, err = dc.ReadUint32()
		(*z) = StringHeapID(zzdc)
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
		var zelx uint32
		zelx, bts, err = msgp.ReadUint32Bytes(bts)
		(*z) = StringHeapID(zelx)
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
