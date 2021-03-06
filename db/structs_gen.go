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
	var zcua uint32
	zcua, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zcua != 13 {
		err = msgp.ArrayError{Wanted: 13, Got: zcua}
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
		var zxhx uint32
		zxhx, err = dc.ReadUint32()
		z.Name = StringHeapID(zxhx)
	}
	if err != nil {
		return
	}
	{
		var zlqf uint32
		zlqf, err = dc.ReadUint32()
		z.TypeLine = StringHeapID(zlqf)
	}
	if err != nil {
		return
	}
	{
		var zdaf uint32
		zdaf, err = dc.ReadUint32()
		z.Note = StringHeapID(zdaf)
	}
	if err != nil {
		return
	}
	{
		var zpks uint32
		zpks, err = dc.ReadUint32()
		z.RootType = StringHeapID(zpks)
	}
	if err != nil {
		return
	}
	{
		var zjfb uint32
		zjfb, err = dc.ReadUint32()
		z.RootFlavor = StringHeapID(zjfb)
	}
	if err != nil {
		return
	}
	{
		var zcxo uint16
		zcxo, err = dc.ReadUint16()
		z.League = LeagueHeapID(zcxo)
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
	var zeff uint32
	zeff, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zeff) {
		z.Mods = (z.Mods)[:zeff]
	} else {
		z.Mods = make([]ItemMod, zeff)
	}
	for zwht := range z.Mods {
		var zrsw uint32
		zrsw, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if zrsw != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zrsw}
			return
		}
		{
			var zxpk uint32
			zxpk, err = dc.ReadUint32()
			z.Mods[zwht].Mod = StringHeapID(zxpk)
		}
		if err != nil {
			return
		}
		z.Mods[zwht].Value, err = dc.ReadUint16()
		if err != nil {
			return
		}
	}
	err = dc.ReadExactBytes(z.When[:])
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
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return err
		}
		err = en.WriteUint32(uint32(z.Mods[zwht].Mod))
		if err != nil {
			return
		}
		err = en.WriteUint16(z.Mods[zwht].Value)
		if err != nil {
			return
		}
	}
	err = en.WriteBytes(z.When[:])
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
		// array header, size 2
		o = append(o, 0x92)
		o = msgp.AppendUint32(o, uint32(z.Mods[zwht].Mod))
		o = msgp.AppendUint16(o, z.Mods[zwht].Value)
	}
	o = msgp.AppendBytes(o, z.When[:])
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zdnj uint32
	zdnj, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zdnj != 13 {
		err = msgp.ArrayError{Wanted: 13, Got: zdnj}
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
		var zobc uint32
		zobc, bts, err = msgp.ReadUint32Bytes(bts)
		z.Name = StringHeapID(zobc)
	}
	if err != nil {
		return
	}
	{
		var zsnv uint32
		zsnv, bts, err = msgp.ReadUint32Bytes(bts)
		z.TypeLine = StringHeapID(zsnv)
	}
	if err != nil {
		return
	}
	{
		var zkgt uint32
		zkgt, bts, err = msgp.ReadUint32Bytes(bts)
		z.Note = StringHeapID(zkgt)
	}
	if err != nil {
		return
	}
	{
		var zema uint32
		zema, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootType = StringHeapID(zema)
	}
	if err != nil {
		return
	}
	{
		var zpez uint32
		zpez, bts, err = msgp.ReadUint32Bytes(bts)
		z.RootFlavor = StringHeapID(zpez)
	}
	if err != nil {
		return
	}
	{
		var zqke uint16
		zqke, bts, err = msgp.ReadUint16Bytes(bts)
		z.League = LeagueHeapID(zqke)
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
	var zqyh uint32
	zqyh, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(zqyh) {
		z.Mods = (z.Mods)[:zqyh]
	} else {
		z.Mods = make([]ItemMod, zqyh)
	}
	for zwht := range z.Mods {
		var zyzr uint32
		zyzr, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if zyzr != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zyzr}
			return
		}
		{
			var zywj uint32
			zywj, bts, err = msgp.ReadUint32Bytes(bts)
			z.Mods[zwht].Mod = StringHeapID(zywj)
		}
		if err != nil {
			return
		}
		z.Mods[zwht].Value, bts, err = msgp.ReadUint16Bytes(bts)
		if err != nil {
			return
		}
	}
	bts, err = msgp.ReadExactBytes(bts, z.When[:])
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Item) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize + (IDSize * (msgp.ByteSize)) + msgp.ArrayHeaderSize + (GGGIDSize * (msgp.ByteSize)) + msgp.ArrayHeaderSize + (GGGIDSize * (msgp.ByteSize)) + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint32Size + msgp.Uint16Size + msgp.BoolSize + msgp.BoolSize + msgp.ArrayHeaderSize + (len(z.Mods) * (11 + msgp.Uint32Size + msgp.Uint16Size)) + msgp.ArrayHeaderSize + (TimestampSize * (msgp.ByteSize))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ItemMod) DecodeMsg(dc *msgp.Reader) (err error) {
	var zjpj uint32
	zjpj, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zjpj != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zjpj}
		return
	}
	{
		var zzpf uint32
		zzpf, err = dc.ReadUint32()
		z.Mod = StringHeapID(zzpf)
	}
	if err != nil {
		return
	}
	z.Value, err = dc.ReadUint16()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ItemMod) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return err
	}
	err = en.WriteUint32(uint32(z.Mod))
	if err != nil {
		return
	}
	err = en.WriteUint16(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ItemMod) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendUint32(o, uint32(z.Mod))
	o = msgp.AppendUint16(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ItemMod) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zrfe uint32
	zrfe, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zrfe != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zrfe}
		return
	}
	{
		var zgmo uint32
		zgmo, bts, err = msgp.ReadUint32Bytes(bts)
		z.Mod = StringHeapID(zgmo)
	}
	if err != nil {
		return
	}
	z.Value, bts, err = msgp.ReadUint16Bytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ItemMod) Msgsize() (s int) {
	s = 1 + msgp.Uint32Size + msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LeagueHeapID) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var ztaf uint16
		ztaf, err = dc.ReadUint16()
		(*z) = LeagueHeapID(ztaf)
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
		var zeth uint16
		zeth, bts, err = msgp.ReadUint16Bytes(bts)
		(*z) = LeagueHeapID(zeth)
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
	var zwel uint32
	zwel, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zwel != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zwel}
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
	var zrbe uint32
	zrbe, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Items) >= int(zrbe) {
		z.Items = (z.Items)[:zrbe]
	} else {
		z.Items = make([]GGGID, zrbe)
	}
	for zrjx := range z.Items {
		err = dc.ReadExactBytes(z.Items[zrjx][:])
		if err != nil {
			return
		}
	}
	{
		var zmfd uint16
		zmfd, err = dc.ReadUint16()
		z.League = LeagueHeapID(zmfd)
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
	for zrjx := range z.Items {
		err = en.WriteBytes(z.Items[zrjx][:])
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
	for zrjx := range z.Items {
		o = msgp.AppendBytes(o, z.Items[zrjx][:])
	}
	o = msgp.AppendUint16(o, uint16(z.League))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Stash) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zzdc uint32
	zzdc, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zzdc != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zzdc}
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
	var zelx uint32
	zelx, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Items) >= int(zelx) {
		z.Items = (z.Items)[:zelx]
	} else {
		z.Items = make([]GGGID, zelx)
	}
	for zrjx := range z.Items {
		bts, err = msgp.ReadExactBytes(bts, z.Items[zrjx][:])
		if err != nil {
			return
		}
	}
	{
		var zbal uint16
		zbal, bts, err = msgp.ReadUint16Bytes(bts)
		z.League = LeagueHeapID(zbal)
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
		var zjqz uint32
		zjqz, err = dc.ReadUint32()
		(*z) = StringHeapID(zjqz)
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
		var zkct uint32
		zkct, bts, err = msgp.ReadUint32Bytes(bts)
		(*z) = StringHeapID(zkct)
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
