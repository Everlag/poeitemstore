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
	var zajw uint32
	zajw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
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
				var zwht uint64
				zwht, err = dc.ReadUint64()
				z.Name = StringHeapID(zwht)
			}
			if err != nil {
				return
			}
		case "TypeLine":
			{
				var zhct uint64
				zhct, err = dc.ReadUint64()
				z.TypeLine = StringHeapID(zhct)
			}
			if err != nil {
				return
			}
		case "Note":
			{
				var zcua uint64
				zcua, err = dc.ReadUint64()
				z.Note = StringHeapID(zcua)
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
		case "League":
			{
				var zdaf uint16
				zdaf, err = dc.ReadUint16()
				z.League = LeagueHeapID(zdaf)
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
		case "Mods":
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
	// map header, size 11
	// write "ID"
	err = en.Append(0x8b, 0xa2, 0x49, 0x44)
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
	// write "League"
	err = en.Append(0xa6, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint16(uint16(z.League))
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
	// write "Mods"
	err = en.Append(0xa4, 0x4d, 0x6f, 0x64, 0x73)
	if err != nil {
		return err
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
	// map header, size 11
	// string "ID"
	o = append(o, 0x8b, 0xa2, 0x49, 0x44)
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
	// string "RootType"
	o = append(o, 0xa8, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendUint64(o, uint64(z.RootType))
	// string "RootFlavor"
	o = append(o, 0xaa, 0x52, 0x6f, 0x6f, 0x74, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x72)
	o = msgp.AppendUint64(o, uint64(z.RootFlavor))
	// string "League"
	o = append(o, 0xa6, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65)
	o = msgp.AppendUint16(o, uint16(z.League))
	// string "Corrupted"
	o = append(o, 0xa9, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Corrupted)
	// string "Identified"
	o = append(o, 0xaa, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Identified)
	// string "Mods"
	o = append(o, 0xa4, 0x4d, 0x6f, 0x64, 0x73)
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
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
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
				var zcxo uint64
				zcxo, bts, err = msgp.ReadUint64Bytes(bts)
				z.Name = StringHeapID(zcxo)
			}
			if err != nil {
				return
			}
		case "TypeLine":
			{
				var zeff uint64
				zeff, bts, err = msgp.ReadUint64Bytes(bts)
				z.TypeLine = StringHeapID(zeff)
			}
			if err != nil {
				return
			}
		case "Note":
			{
				var zrsw uint64
				zrsw, bts, err = msgp.ReadUint64Bytes(bts)
				z.Note = StringHeapID(zrsw)
			}
			if err != nil {
				return
			}
		case "RootType":
			{
				var zxpk uint64
				zxpk, bts, err = msgp.ReadUint64Bytes(bts)
				z.RootType = StringHeapID(zxpk)
			}
			if err != nil {
				return
			}
		case "RootFlavor":
			{
				var zdnj uint64
				zdnj, bts, err = msgp.ReadUint64Bytes(bts)
				z.RootFlavor = StringHeapID(zdnj)
			}
			if err != nil {
				return
			}
		case "League":
			{
				var zobc uint16
				zobc, bts, err = msgp.ReadUint16Bytes(bts)
				z.League = LeagueHeapID(zobc)
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
		case "Mods":
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
	s = 1 + 3 + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + 6 + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + 5 + msgp.Uint64Size + 9 + msgp.Uint64Size + 5 + msgp.Uint64Size + 9 + msgp.Uint64Size + 11 + msgp.Uint64Size + 7 + msgp.Uint16Size + 10 + msgp.BoolSize + 11 + msgp.BoolSize + 5 + msgp.ArrayHeaderSize
	for zcmr := range z.Mods {
		s += z.Mods[zcmr].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ItemMod) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zema uint32
	zema, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zema > 0 {
		zema--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Mod":
			{
				var zpez uint64
				zpez, err = dc.ReadUint64()
				z.Mod = StringHeapID(zpez)
			}
			if err != nil {
				return
			}
		case "Values":
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
func (z *ItemMod) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Mod"
	err = en.Append(0x82, 0xa3, 0x4d, 0x6f, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Mod))
	if err != nil {
		return
	}
	// write "Values"
	err = en.Append(0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
	if err != nil {
		return err
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
	// map header, size 2
	// string "Mod"
	o = append(o, 0x82, 0xa3, 0x4d, 0x6f, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Mod))
	// string "Values"
	o = append(o, 0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Values)))
	for zkgt := range z.Values {
		o = msgp.AppendInt(o, z.Values[zkgt])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ItemMod) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zqyh uint32
	zqyh, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqyh > 0 {
		zqyh--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Mod":
			{
				var zyzr uint64
				zyzr, bts, err = msgp.ReadUint64Bytes(bts)
				z.Mod = StringHeapID(zyzr)
			}
			if err != nil {
				return
			}
		case "Values":
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
func (z *ItemMod) Msgsize() (s int) {
	s = 1 + 4 + msgp.Uint64Size + 7 + msgp.ArrayHeaderSize + (len(z.Values) * (msgp.IntSize))
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
		var zeth uint64
		zeth, err = dc.ReadUint64()
		(*z) = StringHeapID(zeth)
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
		var zsbz uint64
		zsbz, bts, err = msgp.ReadUint64Bytes(bts)
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
	s = msgp.Uint64Size
	return
}
