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
	var zwht uint32
	zwht, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
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
				var zhct uint64
				zhct, err = dc.ReadUint64()
				z.Name = StringHeapID(zhct)
			}
			if err != nil {
				return
			}
		case "TypeLine":
			{
				var zcua uint64
				zcua, err = dc.ReadUint64()
				z.TypeLine = StringHeapID(zcua)
			}
			if err != nil {
				return
			}
		case "Note":
			{
				var zxhx uint64
				zxhx, err = dc.ReadUint64()
				z.Note = StringHeapID(zxhx)
			}
			if err != nil {
				return
			}
		case "RootType":
			{
				var zlqf uint64
				zlqf, err = dc.ReadUint64()
				z.RootType = StringHeapID(zlqf)
			}
			if err != nil {
				return
			}
		case "RootFlavor":
			{
				var zdaf uint64
				zdaf, err = dc.ReadUint64()
				z.RootFlavor = StringHeapID(zdaf)
			}
			if err != nil {
				return
			}
		case "League":
			{
				var zpks uint16
				zpks, err = dc.ReadUint16()
				z.League = LeagueHeapID(zpks)
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
				zcxo, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zcxo > 0 {
					zcxo--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Mod":
						{
							var zeff uint64
							zeff, err = dc.ReadUint64()
							z.Mods[zcmr].Mod = StringHeapID(zeff)
						}
						if err != nil {
							return
						}
					case "Values":
						var zrsw uint32
						zrsw, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.Mods[zcmr].Values) >= int(zrsw) {
							z.Mods[zcmr].Values = (z.Mods[zcmr].Values)[:zrsw]
						} else {
							z.Mods[zcmr].Values = make([]int, zrsw)
						}
						for zajw := range z.Mods[zcmr].Values {
							z.Mods[zcmr].Values[zajw], err = dc.ReadInt()
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
		// map header, size 2
		// write "Mod"
		err = en.Append(0x82, 0xa3, 0x4d, 0x6f, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteUint64(uint64(z.Mods[zcmr].Mod))
		if err != nil {
			return
		}
		// write "Values"
		err = en.Append(0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Mods[zcmr].Values)))
		if err != nil {
			return
		}
		for zajw := range z.Mods[zcmr].Values {
			err = en.WriteInt(z.Mods[zcmr].Values[zajw])
			if err != nil {
				return
			}
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
		// map header, size 2
		// string "Mod"
		o = append(o, 0x82, 0xa3, 0x4d, 0x6f, 0x64)
		o = msgp.AppendUint64(o, uint64(z.Mods[zcmr].Mod))
		// string "Values"
		o = append(o, 0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Mods[zcmr].Values)))
		for zajw := range z.Mods[zcmr].Values {
			o = msgp.AppendInt(o, z.Mods[zcmr].Values[zajw])
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxpk uint32
	zxpk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxpk > 0 {
		zxpk--
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
				var zdnj uint64
				zdnj, bts, err = msgp.ReadUint64Bytes(bts)
				z.Name = StringHeapID(zdnj)
			}
			if err != nil {
				return
			}
		case "TypeLine":
			{
				var zobc uint64
				zobc, bts, err = msgp.ReadUint64Bytes(bts)
				z.TypeLine = StringHeapID(zobc)
			}
			if err != nil {
				return
			}
		case "Note":
			{
				var zsnv uint64
				zsnv, bts, err = msgp.ReadUint64Bytes(bts)
				z.Note = StringHeapID(zsnv)
			}
			if err != nil {
				return
			}
		case "RootType":
			{
				var zkgt uint64
				zkgt, bts, err = msgp.ReadUint64Bytes(bts)
				z.RootType = StringHeapID(zkgt)
			}
			if err != nil {
				return
			}
		case "RootFlavor":
			{
				var zema uint64
				zema, bts, err = msgp.ReadUint64Bytes(bts)
				z.RootFlavor = StringHeapID(zema)
			}
			if err != nil {
				return
			}
		case "League":
			{
				var zpez uint16
				zpez, bts, err = msgp.ReadUint16Bytes(bts)
				z.League = LeagueHeapID(zpez)
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
							z.Mods[zcmr].Mod = StringHeapID(zyzr)
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
						if cap(z.Mods[zcmr].Values) >= int(zywj) {
							z.Mods[zcmr].Values = (z.Mods[zcmr].Values)[:zywj]
						} else {
							z.Mods[zcmr].Values = make([]int, zywj)
						}
						for zajw := range z.Mods[zcmr].Values {
							z.Mods[zcmr].Values[zajw], bts, err = msgp.ReadIntBytes(bts)
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
		s += 1 + 4 + msgp.Uint64Size + 7 + msgp.ArrayHeaderSize + (len(z.Mods[zcmr].Values) * (msgp.IntSize))
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ItemMod) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zzpf uint32
	zzpf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zzpf > 0 {
		zzpf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Mod":
			{
				var zrfe uint64
				zrfe, err = dc.ReadUint64()
				z.Mod = StringHeapID(zrfe)
			}
			if err != nil {
				return
			}
		case "Values":
			var zgmo uint32
			zgmo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Values) >= int(zgmo) {
				z.Values = (z.Values)[:zgmo]
			} else {
				z.Values = make([]int, zgmo)
			}
			for zjpj := range z.Values {
				z.Values[zjpj], err = dc.ReadInt()
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
	for zjpj := range z.Values {
		err = en.WriteInt(z.Values[zjpj])
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
	for zjpj := range z.Values {
		o = msgp.AppendInt(o, z.Values[zjpj])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ItemMod) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Mod":
			{
				var zeth uint64
				zeth, bts, err = msgp.ReadUint64Bytes(bts)
				z.Mod = StringHeapID(zeth)
			}
			if err != nil {
				return
			}
		case "Values":
			var zsbz uint32
			zsbz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Values) >= int(zsbz) {
				z.Values = (z.Values)[:zsbz]
			} else {
				z.Values = make([]int, zsbz)
			}
			for zjpj := range z.Values {
				z.Values[zjpj], bts, err = msgp.ReadIntBytes(bts)
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
		var zzdc uint64
		zzdc, err = dc.ReadUint64()
		(*z) = StringHeapID(zzdc)
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
		var zelx uint64
		zelx, bts, err = msgp.ReadUint64Bytes(bts)
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
	s = msgp.Uint64Size
	return
}
