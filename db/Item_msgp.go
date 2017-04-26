package db

import "github.com/tinylib/msgp/msgp"

// DecodeMsg implements msgp.Decodable
func (z *Item) DecodeMsg(dc *msgp.Reader) (err error) {
	var tmpUint32 uint32
	if tmpUint32, err = dc.ReadArrayHeader(); err != nil {
		return
	}
	if tmpUint32 != 13 {
		err = msgp.ArrayError{Wanted: 13, Got: tmpUint32}
		return
	}
	if err = dc.ReadExactBytes(z.ID[:]); err != nil {
		return
	}
	if err = dc.ReadExactBytes(z.GGGID[:]); err != nil {
		return
	}
	if err = dc.ReadExactBytes(z.Stash[:]); err != nil {
		return
	}

	if tmpUint32, err = dc.ReadUint32(); err != nil {
		return
	}
	z.Name = StringHeapID(tmpUint32)

	if tmpUint32, err = dc.ReadUint32(); err != nil {
		return
	}
	z.TypeLine = StringHeapID(tmpUint32)

	if tmpUint32, err = dc.ReadUint32(); err != nil {
		return
	}
	z.Note = StringHeapID(tmpUint32)

	if tmpUint32, err = dc.ReadUint32(); err != nil {
		return
	}
	z.RootType = StringHeapID(tmpUint32)

	if tmpUint32, err = dc.ReadUint32(); err != nil {
		return
	}
	z.RootFlavor = StringHeapID(tmpUint32)

	var tmpUint16 uint16
	if tmpUint16, err = dc.ReadUint16(); err != nil {
		return
	}
	z.League = LeagueHeapID(tmpUint16)

	if z.Corrupted, err = dc.ReadBool(); err != nil {
		return
	}
	if z.Identified, err = dc.ReadBool(); err != nil {
		return
	}

	tmpUint32, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Mods) >= int(tmpUint32) {
		z.Mods = (z.Mods)[:tmpUint32]
	} else {
		z.Mods = make([]ItemMod, tmpUint32)
	}
	for i := range z.Mods {
		if tmpUint32, err = dc.ReadArrayHeader(); err != nil {
			return
		}
		if tmpUint32 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: tmpUint32}
			return
		}

		if tmpUint32, err = dc.ReadUint32(); err != nil {
			return
		}
		z.Mods[i].Mod = StringHeapID(tmpUint32)

		if z.Mods[i].Value, err = dc.ReadUint16(); err != nil {
			return
		}
	}

	if err = dc.ReadExactBytes(z.When[:]); err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Item) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 13
	if err = en.Append(0x9d); err != nil {
		return err
	}
	if err = en.WriteBytes(z.ID[:]); err != nil {
		return
	}
	if err = en.WriteBytes(z.GGGID[:]); err != nil {
		return
	}
	if err = en.WriteBytes(z.Stash[:]); err != nil {
		return
	}

	if err = en.WriteUint32(uint32(z.Name)); err != nil {
		return
	}
	if err = en.WriteUint32(uint32(z.TypeLine)); err != nil {
		return
	}
	if err = en.WriteUint32(uint32(z.Note)); err != nil {
		return
	}
	if err = en.WriteUint32(uint32(z.RootType)); err != nil {
		return
	}
	if err = en.WriteUint32(uint32(z.RootFlavor)); err != nil {
		return
	}

	if err = en.WriteUint16(uint16(z.League)); err != nil {
		return
	}

	if err = en.WriteBool(z.Corrupted); err != nil {
		return
	}
	if err = en.WriteBool(z.Identified); err != nil {
		return
	}

	if err = en.WriteArrayHeader(uint32(len(z.Mods))); err != nil {
		return
	}
	for i := range z.Mods {
		// array header, size 2
		if err = en.Append(0x92); err != nil {
			return err
		}
		if err = en.WriteUint32(uint32(z.Mods[i].Mod)); err != nil {
			return
		}
		if err = en.WriteUint16(z.Mods[i].Value); err != nil {
			return
		}
	}
	if err = en.WriteBytes(z.When[:]); err != nil {
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
	for i := range z.Mods {
		// array header, size 2
		o = append(o, 0x92)
		o = msgp.AppendUint32(o, uint32(z.Mods[i].Mod))
		o = msgp.AppendUint16(o, z.Mods[i].Value)
	}
	o = msgp.AppendBytes(o, z.When[:])
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var tmpUint32 uint32
	if tmpUint32, bts, err = msgp.ReadArrayHeaderBytes(bts); err != nil {
		return
	}
	if tmpUint32 != 13 {
		err = msgp.ArrayError{Wanted: 13, Got: tmpUint32}
		return
	}
	if bts, err = msgp.ReadExactBytes(bts, z.ID[:]); err != nil {
		return
	}
	if bts, err = msgp.ReadExactBytes(bts, z.GGGID[:]); err != nil {
		return
	}
	if bts, err = msgp.ReadExactBytes(bts, z.Stash[:]); err != nil {
		return
	}

	if tmpUint32, bts, err = msgp.ReadUint32Bytes(bts); err != nil {
		return
	}
	z.Name = StringHeapID(tmpUint32)

	if tmpUint32, bts, err = msgp.ReadUint32Bytes(bts); err != nil {
		return
	}
	z.TypeLine = StringHeapID(tmpUint32)

	if tmpUint32, bts, err = msgp.ReadUint32Bytes(bts); err != nil {
		return
	}
	z.Note = StringHeapID(tmpUint32)

	if tmpUint32, bts, err = msgp.ReadUint32Bytes(bts); err != nil {
		return
	}
	z.RootType = StringHeapID(tmpUint32)

	if tmpUint32, bts, err = msgp.ReadUint32Bytes(bts); err != nil {
		return
	}
	z.RootFlavor = StringHeapID(tmpUint32)

	var tmpUint16 uint16
	if tmpUint16, bts, err = msgp.ReadUint16Bytes(bts); err != nil {
		return
	}
	z.League = LeagueHeapID(tmpUint16)

	if z.Corrupted, bts, err = msgp.ReadBoolBytes(bts); err != nil {
		return
	}

	if z.Identified, bts, err = msgp.ReadBoolBytes(bts); err != nil {
		return
	}

	if tmpUint32, bts, err = msgp.ReadArrayHeaderBytes(bts); err != nil {
		return
	}
	if cap(z.Mods) >= int(tmpUint32) {
		z.Mods = (z.Mods)[:tmpUint32]
	} else {
		z.Mods = make([]ItemMod, tmpUint32)
	}
	for i := range z.Mods {
		tmpUint32, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if tmpUint32 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: tmpUint32}
			return
		}

		if tmpUint32, bts, err = msgp.ReadUint32Bytes(bts); err != nil {
			return
		}
		z.Mods[i].Mod = StringHeapID(tmpUint32)

		if z.Mods[i].Value, bts, err = msgp.ReadUint16Bytes(bts); err != nil {
			return
		}
	}

	if bts, err = msgp.ReadExactBytes(bts, z.When[:]); err != nil {
		return
	}

	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Item) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize +
		(IDSize * (msgp.ByteSize)) +
		msgp.ArrayHeaderSize + (GGGIDSize * (msgp.ByteSize)) +
		msgp.ArrayHeaderSize + (GGGIDSize * (msgp.ByteSize)) +
		msgp.Uint32Size +
		msgp.Uint32Size +
		msgp.Uint32Size +
		msgp.Uint32Size +
		msgp.Uint32Size +
		msgp.Uint16Size +
		msgp.BoolSize +
		msgp.BoolSize +
		msgp.ArrayHeaderSize + (len(z.Mods) * (11 + msgp.Uint32Size + msgp.Uint16Size)) +
		msgp.ArrayHeaderSize + (TimestampSize * (msgp.ByteSize))
	return
}
