package stash

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *ChangeSet) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Changes":
			var zbai uint32
			zbai, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Changes) >= int(zbai) {
				z.Changes = (z.Changes)[:zbai]
			} else {
				z.Changes = make([]CompressedResponse, zbai)
			}
			for zxvk := range z.Changes {
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
					case "Content":
						z.Changes[zxvk].Content, err = dc.ReadBytes(z.Changes[zxvk].Content)
						if err != nil {
							return
						}
					case "Size":
						z.Changes[zxvk].Size, err = dc.ReadInt()
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
			}
		case "Size":
			z.Size, err = dc.ReadInt()
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
func (z *ChangeSet) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Changes"
	err = en.Append(0x82, 0xa7, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Changes)))
	if err != nil {
		return
	}
	for zxvk := range z.Changes {
		// map header, size 2
		// write "Content"
		err = en.Append(0x82, 0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
		if err != nil {
			return err
		}
		err = en.WriteBytes(z.Changes[zxvk].Content)
		if err != nil {
			return
		}
		// write "Size"
		err = en.Append(0xa4, 0x53, 0x69, 0x7a, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Changes[zxvk].Size)
		if err != nil {
			return
		}
	}
	// write "Size"
	err = en.Append(0xa4, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Size)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ChangeSet) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Changes"
	o = append(o, 0x82, 0xa7, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Changes)))
	for zxvk := range z.Changes {
		// map header, size 2
		// string "Content"
		o = append(o, 0x82, 0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
		o = msgp.AppendBytes(o, z.Changes[zxvk].Content)
		// string "Size"
		o = append(o, 0xa4, 0x53, 0x69, 0x7a, 0x65)
		o = msgp.AppendInt(o, z.Changes[zxvk].Size)
	}
	// string "Size"
	o = append(o, 0xa4, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt(o, z.Size)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ChangeSet) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Changes":
			var zwht uint32
			zwht, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Changes) >= int(zwht) {
				z.Changes = (z.Changes)[:zwht]
			} else {
				z.Changes = make([]CompressedResponse, zwht)
			}
			for zxvk := range z.Changes {
				var zhct uint32
				zhct, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zhct > 0 {
					zhct--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Content":
						z.Changes[zxvk].Content, bts, err = msgp.ReadBytesBytes(bts, z.Changes[zxvk].Content)
						if err != nil {
							return
						}
					case "Size":
						z.Changes[zxvk].Size, bts, err = msgp.ReadIntBytes(bts)
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
			}
		case "Size":
			z.Size, bts, err = msgp.ReadIntBytes(bts)
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
func (z *ChangeSet) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zxvk := range z.Changes {
		s += 1 + 8 + msgp.BytesPrefixSize + len(z.Changes[zxvk].Content) + 5 + msgp.IntSize
	}
	s += 5 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CompressedResponse) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Content":
			z.Content, err = dc.ReadBytes(z.Content)
			if err != nil {
				return
			}
		case "Size":
			z.Size, err = dc.ReadInt()
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
func (z *CompressedResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Content"
	err = en.Append(0x82, 0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Content)
	if err != nil {
		return
	}
	// write "Size"
	err = en.Append(0xa4, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Size)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CompressedResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Content"
	o = append(o, 0x82, 0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	o = msgp.AppendBytes(o, z.Content)
	// string "Size"
	o = append(o, 0xa4, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt(o, z.Size)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CompressedResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Content":
			z.Content, bts, err = msgp.ReadBytesBytes(bts, z.Content)
			if err != nil {
				return
			}
		case "Size":
			z.Size, bts, err = msgp.ReadIntBytes(bts)
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
func (z *CompressedResponse) Msgsize() (s int) {
	s = 1 + 8 + msgp.BytesPrefixSize + len(z.Content) + 5 + msgp.IntSize
	return
}
