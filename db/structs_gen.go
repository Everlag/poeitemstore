package db

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *GGGID) DecodeMsg(dc *msgp.Reader) (err error) {
	err = dc.ReadExactBytes(z[:])
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *GGGID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes(z[:])
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GGGID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, z[:])
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GGGID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, z[:])
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *GGGID) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (GGGIDSize * (msgp.ByteSize))
	return
}

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
	var zhct uint32
	zhct, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zhct != 13 {
		err = msgp.ArrayError{Wanted: 13, Got: zhct}
		return
	}
	err = dc.ReadExactBytes(z.ID[:])
	if err != nil {
		return
	}
	err = dc.ReadExactBytes(z.GGGID[:])
	if err != nil {
		return
	}
	err = dc.ReadExactBytes(z.Stash[:])
	if err != nil {
		return
	}
	{
		var zcua uint32
		zcua, err = dc.ReadUint32()
		z.Name = StringHeapID(zcua)
	}
	if err != nil {
		return
	}
	{
		var zxhx uint32
		zxhx, err = dc.ReadUint32()
		z.TypeLine = StringHeapID(zxhx)
	}
	if err != nil {
		return
	}
	{
		var zlqf uint32
		zlqf, err = dc.ReadUint32()
		z.Note = StringHeapID(zlqf)
	}
	if err != nil {
		return
	}
	{
		var zdaf uint32
		zdaf, err = dc.ReadUint32()
		z.RootType = StringHeapID(zdaf)
	}
	if err != nil {
		return
	}
	{
		var zpks uint32
		zpks, err = dc.ReadUint32()
		z.RootFlavor = StringHeapID(zpks)
	}
	if err != nil {
		return
	}
	{
		var zjfb uint16
		zjfb, err = dc.ReadUint16()
		z.League = LeagueHeapID(zjfb)
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
	var zcxo uint32
	zcxo, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zcxo) {
		z.Mods = (z.Mods)[:zcxo]
	} else {
		z.Mods = make([]ItemMod, zcxo)
	}
	for zwht := range z.Mods {
		err = z.Mods[zwht].DecodeMsg(dc)
		if err != nil {
			return
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
	// array header, size 13
	err = en.Append(0x9d)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.ID[:])
	if err != nil {
		return
	}
	err = en.WriteBytes(z.GGGID[:])
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
	for zwht := range z.Mods {
		err = z.Mods[zwht].EncodeMsg(en)
		if err != nil {
			return
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
	// array header, size 13
	o = append(o, 0x9d)
	o = msgp.AppendBytes(o, z.ID[:])
	o = msgp.AppendBytes(o, z.GGGID[:])
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
	for zwht := range z.Mods {
		o, err = z.Mods[zwht].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o = msgp.AppendUint16(o, z.UpdateSequence)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zeff uint32
	zeff, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zeff != 13 {
		err = msgp.ArrayError{Wanted: 13, Got: zeff}
		return
	}
	bts, err = msgp.ReadExactBytes(bts, z.ID[:])
	if err != nil {
		return
	}
	bts, err = msgp.ReadExactBytes(bts, z.GGGID[:])
	if err != nil {
		return
	}
	bts, err = msgp.ReadExactBytes(bts, z.Stash[:])
	if err != nil {
		return
	}
	{
		var zrsw uint32
		zrsw, bts, err = msgp.ReadUint32Bytes(bts)
		z.Name = StringHeapID(zrsw)
	}
	if err != nil {
		return
	}
	{
		var zxpk uint32
		zxpk, bts, err = msgp.ReadUint32Bytes(bts)
		z.TypeLine = StringHeapID(zxpk)
	}
	if err != nil {
		return
	}
	{
		var zdnj uint32
		zdnj, bts, err = msgp.ReadUint32Bytes(bts)
		z.Note = StringHeapID(zdnj)
	}
	if err != nil {
		return
	}
	{
		var zobc uint32
		zobc, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootType = StringHeapID(zobc)
	}
	if err != nil {
		return
	}
	{
		var zsnv uint32
		zsnv, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootFlavor = StringHeapID(zsnv)
	}
	if err != nil {
		return
	}
	{
		var zkgt uint16
		zkgt, bts, err = msgp.ReadUint16Bytes(bts)
		z.League = LeagueHeapID(zkgt)
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
	var zema uint32
	zema, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zema) {
		z.Mods = (z.Mods)[:zema]
	} else {
		z.Mods = make([]ItemMod, zema)
	}
	for zwht := range z.Mods {
		bts, err = z.Mods[zwht].UnmarshalMsg(bts)
		if err != nil {
			return
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
	s = 1 + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + msgp.ArrayHeaderSize + (GGGIDSize * (msgp.ByteSize)) + msgp.ArrayHeaderSize + (GGGIDSize * (msgp.ByteSize)) + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint16Size + msgp.BoolSize + msgp.BoolSize + msgp.ArrayHeaderSize
	for zwht := range z.Mods {
		s += z.Mods[zwht].Msgsize()
	}
	s += msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ItemMod) DecodeMsg(dc *msgp.Reader) (err error) {
	var zqke uint32
	zqke, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zqke != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zqke}
		return
	}
	{
		var zqyh uint32
		zqyh, err = dc.ReadUint32()
		z.Mod = StringHeapID(zqyh)
	}
	if err != nil {
		return
	}
	var zyzr uint32
	zyzr, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Values) >= int(zyzr) {
		z.Values = (z.Values)[:zyzr]
	} else {
		z.Values = make([]uint16, zyzr)
	}
	for zpez := range z.Values {
		z.Values[zpez], err = dc.ReadUint16()
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
	for zpez := range z.Values {
		err = en.WriteUint16(z.Values[zpez])
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
	for zpez := range z.Values {
		o = msgp.AppendUint16(o, z.Values[zpez])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ItemMod) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zywj uint32
	zywj, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zywj != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zywj}
		return
	}
	{
		var zjpj uint32
		zjpj, bts, err = msgp.ReadUint32Bytes(bts)
		z.Mod = StringHeapID(zjpj)
	}
	if err != nil {
		return
	}
	var zzpf uint32
	zzpf, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Values) >= int(zzpf) {
		z.Values = (z.Values)[:zzpf]
	} else {
		z.Values = make([]uint16, zzpf)
	}
	for zpez := range z.Values {
		z.Values[zpez], bts, err = msgp.ReadUint16Bytes(bts)
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
		var zrfe uint16
		zrfe, err = dc.ReadUint16()
		(*z) = LeagueHeapID(zrfe)
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
		var zgmo uint16
		zgmo, bts, err = msgp.ReadUint16Bytes(bts)
		(*z) = LeagueHeapID(zgmo)
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
	var zrjx uint32
	zrjx, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zrjx != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zrjx}
		return
	}
	err = dc.ReadExactBytes(z.ID[:])
	if err != nil {
		return
	}
	z.AccountName, err = dc.ReadString()
	if err != nil {
		return
	}
	var zawn uint32
	zawn, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Items) >= int(zawn) {
		z.Items = (z.Items)[:zawn]
	} else {
		z.Items = make([]GGGID, zawn)
	}
	for zeth := range z.Items {
		err = dc.ReadExactBytes(z.Items[zeth][:])
		if err != nil {
			return
		}
	}
	{
		var zwel uint16
		zwel, err = dc.ReadUint16()
		z.League = LeagueHeapID(zwel)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Stash) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.ID[:])
	if err != nil {
		return
	}
	err = en.WriteString(z.AccountName)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Items)))
	if err != nil {
		return
	}
	for zeth := range z.Items {
		err = en.WriteBytes(z.Items[zeth][:])
		if err != nil {
			return
		}
	}
	err = en.WriteUint16(uint16(z.League))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Stash) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o = msgp.AppendBytes(o, z.ID[:])
	o = msgp.AppendString(o, z.AccountName)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Items)))
	for zeth := range z.Items {
		o = msgp.AppendBytes(o, z.Items[zeth][:])
	}
	o = msgp.AppendUint16(o, uint16(z.League))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Stash) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zrbe uint32
	zrbe, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zrbe != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zrbe}
		return
	}
	bts, err = msgp.ReadExactBytes(bts, z.ID[:])
	if err != nil {
		return
	}
	z.AccountName, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	var zmfd uint32
	zmfd, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Items) >= int(zmfd) {
		z.Items = (z.Items)[:zmfd]
	} else {
		z.Items = make([]GGGID, zmfd)
	}
	for zeth := range z.Items {
		bts, err = msgp.ReadExactBytes(bts, z.Items[zeth][:])
		if err != nil {
			return
		}
	}
	{
		var zzdc uint16
		zzdc, bts, err = msgp.ReadUint16Bytes(bts)
		z.League = LeagueHeapID(zzdc)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Stash) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize + (GGGIDSize * (msgp.ByteSize)) + msgp.StringPrefixSize + len(z.AccountName) + msgp.ArrayHeaderSize + (len(z.Items) * (GGGIDSize * (msgp.ByteSize))) + msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StringHeapID) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zelx uint32
		zelx, err = dc.ReadUint32()
		(*z) = StringHeapID(zelx)
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
		var zbal uint32
		zbal, bts, err = msgp.ReadUint32Bytes(bts)
		(*z) = StringHeapID(zbal)
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
