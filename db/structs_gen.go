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
	var zxhx uint32
	zxhx, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zxhx != 14 {
		err = msgp.ArrayError{Wanted: 14, Got: zxhx}
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
		var zlqf uint32
		zlqf, err = dc.ReadUint32()
		z.Name = StringHeapID(zlqf)
	}
	if err != nil {
		return
	}
	{
		var zdaf uint32
		zdaf, err = dc.ReadUint32()
		z.TypeLine = StringHeapID(zdaf)
	}
	if err != nil {
		return
	}
	{
		var zpks uint32
		zpks, err = dc.ReadUint32()
		z.Note = StringHeapID(zpks)
	}
	if err != nil {
		return
	}
	{
		var zjfb uint32
		zjfb, err = dc.ReadUint32()
		z.RootType = StringHeapID(zjfb)
	}
	if err != nil {
		return
	}
	{
		var zcxo uint32
		zcxo, err = dc.ReadUint32()
		z.RootFlavor = StringHeapID(zcxo)
	}
	if err != nil {
		return
	}
	{
		var zeff uint16
		zeff, err = dc.ReadUint16()
		z.League = LeagueHeapID(zeff)
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
	var zrsw uint32
	zrsw, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zrsw) {
		z.Mods = (z.Mods)[:zrsw]
	} else {
		z.Mods = make([]ItemMod, zrsw)
	}
	for zwht := range z.Mods {
		var zxpk uint32
		zxpk, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if zxpk != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zxpk}
			return
		}
		{
			var zdnj uint32
			zdnj, err = dc.ReadUint32()
			z.Mods[zwht].Mod = StringHeapID(zdnj)
		}
		if err != nil {
			return
		}
		var zobc uint32
		zobc, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if cap(z.Mods[zwht].Values) >= int(zobc) {
			z.Mods[zwht].Values = (z.Mods[zwht].Values)[:zobc]
		} else {
			z.Mods[zwht].Values = make([]uint16, zobc)
		}
		for zhct := range z.Mods[zwht].Values {
			z.Mods[zwht].Values[zhct], err = dc.ReadUint16()
			if err != nil {
				return
			}
		}
	}
	err = dc.ReadExactBytes(z.When[:])
	if err != nil {
		return
	}
	z.UpdateSequence, err = dc.ReadUint16()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Item) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 14
	err = en.Append(0x9e)
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
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return err
		}
		err = en.WriteUint32(uint32(z.Mods[zwht].Mod))
		if err != nil {
			return
		}
		err = en.WriteArrayHeader(uint32(len(z.Mods[zwht].Values)))
		if err != nil {
			return
		}
		for zhct := range z.Mods[zwht].Values {
			err = en.WriteUint16(z.Mods[zwht].Values[zhct])
			if err != nil {
				return
			}
		}
	}
	err = en.WriteBytes(z.When[:])
	if err != nil {
		return
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
	// array header, size 14
	o = append(o, 0x9e)
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
		// array header, size 2
		o = append(o, 0x92)
		o = msgp.AppendUint32(o, uint32(z.Mods[zwht].Mod))
		o = msgp.AppendArrayHeader(o, uint32(len(z.Mods[zwht].Values)))
		for zhct := range z.Mods[zwht].Values {
			o = msgp.AppendUint16(o, z.Mods[zwht].Values[zhct])
		}
	}
	o = msgp.AppendBytes(o, z.When[:])
	o = msgp.AppendUint16(o, z.UpdateSequence)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zsnv uint32
	zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zsnv != 14 {
		err = msgp.ArrayError{Wanted: 14, Got: zsnv}
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
		var zkgt uint32
		zkgt, bts, err = msgp.ReadUint32Bytes(bts)
		z.Name = StringHeapID(zkgt)
	}
	if err != nil {
		return
	}
	{
		var zema uint32
		zema, bts, err = msgp.ReadUint32Bytes(bts)
		z.TypeLine = StringHeapID(zema)
	}
	if err != nil {
		return
	}
	{
		var zpez uint32
		zpez, bts, err = msgp.ReadUint32Bytes(bts)
		z.Note = StringHeapID(zpez)
	}
	if err != nil {
		return
	}
	{
		var zqke uint32
		zqke, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootType = StringHeapID(zqke)
	}
	if err != nil {
		return
	}
	{
		var zqyh uint32
		zqyh, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootFlavor = StringHeapID(zqyh)
	}
	if err != nil {
		return
	}
	{
		var zyzr uint16
		zyzr, bts, err = msgp.ReadUint16Bytes(bts)
		z.League = LeagueHeapID(zyzr)
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
	var zywj uint32
	zywj, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zywj) {
		z.Mods = (z.Mods)[:zywj]
	} else {
		z.Mods = make([]ItemMod, zywj)
	}
	for zwht := range z.Mods {
		var zjpj uint32
		zjpj, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if zjpj != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zjpj}
			return
		}
		{
			var zzpf uint32
			zzpf, bts, err = msgp.ReadUint32Bytes(bts)
			z.Mods[zwht].Mod = StringHeapID(zzpf)
		}
		if err != nil {
			return
		}
		var zrfe uint32
		zrfe, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if cap(z.Mods[zwht].Values) >= int(zrfe) {
			z.Mods[zwht].Values = (z.Mods[zwht].Values)[:zrfe]
		} else {
			z.Mods[zwht].Values = make([]uint16, zrfe)
		}
		for zhct := range z.Mods[zwht].Values {
			z.Mods[zwht].Values[zhct], bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		}
	}
	bts, err = msgp.ReadExactBytes(bts, z.When[:])
	if err != nil {
		return
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
		s += 1 + msgp.Uint32Size + msgp.ArrayHeaderSize + (len(z.Mods[zwht].Values) * (msgp.Uint16Size))
	}
	s += msgp.ArrayHeaderSize + (TimestampSize * (msgp.ByteSize)) + msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ItemMod) DecodeMsg(dc *msgp.Reader) (err error) {
	var ztaf uint32
	ztaf, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if ztaf != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: ztaf}
		return
	}
	{
		var zeth uint32
		zeth, err = dc.ReadUint32()
		z.Mod = StringHeapID(zeth)
	}
	if err != nil {
		return
	}
	var zsbz uint32
	zsbz, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Values) >= int(zsbz) {
		z.Values = (z.Values)[:zsbz]
	} else {
		z.Values = make([]uint16, zsbz)
	}
	for zgmo := range z.Values {
		z.Values[zgmo], err = dc.ReadUint16()
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
	for zgmo := range z.Values {
		err = en.WriteUint16(z.Values[zgmo])
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
	for zgmo := range z.Values {
		o = msgp.AppendUint16(o, z.Values[zgmo])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ItemMod) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zrjx uint32
	zrjx, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zrjx != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zrjx}
		return
	}
	{
		var zawn uint32
		zawn, bts, err = msgp.ReadUint32Bytes(bts)
		z.Mod = StringHeapID(zawn)
	}
	if err != nil {
		return
	}
	var zwel uint32
	zwel, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Values) >= int(zwel) {
		z.Values = (z.Values)[:zwel]
	} else {
		z.Values = make([]uint16, zwel)
	}
	for zgmo := range z.Values {
		z.Values[zgmo], bts, err = msgp.ReadUint16Bytes(bts)
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
		var zrbe uint16
		zrbe, err = dc.ReadUint16()
		(*z) = LeagueHeapID(zrbe)
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
		var zmfd uint16
		zmfd, bts, err = msgp.ReadUint16Bytes(bts)
		(*z) = LeagueHeapID(zmfd)
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
	var zjqz uint32
	zjqz, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zjqz != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zjqz}
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
	var zkct uint32
	zkct, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Items) >= int(zkct) {
		z.Items = (z.Items)[:zkct]
	} else {
		z.Items = make([]GGGID, zkct)
	}
	for zelx := range z.Items {
		err = dc.ReadExactBytes(z.Items[zelx][:])
		if err != nil {
			return
		}
	}
	{
		var ztmt uint16
		ztmt, err = dc.ReadUint16()
		z.League = LeagueHeapID(ztmt)
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
	for zelx := range z.Items {
		err = en.WriteBytes(z.Items[zelx][:])
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
	for zelx := range z.Items {
		o = msgp.AppendBytes(o, z.Items[zelx][:])
	}
	o = msgp.AppendUint16(o, uint16(z.League))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Stash) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var ztco uint32
	ztco, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if ztco != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: ztco}
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
	var zana uint32
	zana, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Items) >= int(zana) {
		z.Items = (z.Items)[:zana]
	} else {
		z.Items = make([]GGGID, zana)
	}
	for zelx := range z.Items {
		bts, err = msgp.ReadExactBytes(bts, z.Items[zelx][:])
		if err != nil {
			return
		}
	}
	{
		var ztyy uint16
		ztyy, bts, err = msgp.ReadUint16Bytes(bts)
		z.League = LeagueHeapID(ztyy)
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
		var zinl uint32
		zinl, err = dc.ReadUint32()
		(*z) = StringHeapID(zinl)
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
		var zare uint32
		zare, bts, err = msgp.ReadUint32Bytes(bts)
		(*z) = StringHeapID(zare)
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
