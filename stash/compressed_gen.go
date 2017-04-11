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
		case "ChangeIDToIndex":
			var zajw uint32
			zajw, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.ChangeIDToIndex == nil && zajw > 0 {
				z.ChangeIDToIndex = make(map[string]int, zajw)
			} else if len(z.ChangeIDToIndex) > 0 {
				for key, _ := range z.ChangeIDToIndex {
					delete(z.ChangeIDToIndex, key)
				}
			}
			for zajw > 0 {
				zajw--
				var zxvk string
				var zbzg int
				zxvk, err = dc.ReadString()
				if err != nil {
					return
				}
				zbzg, err = dc.ReadInt()
				if err != nil {
					return
				}
				z.ChangeIDToIndex[zxvk] = zbzg
			}
		case "Changes":
			var zwht uint32
			zwht, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Changes) >= int(zwht) {
				z.Changes = (z.Changes)[:zwht]
			} else {
				z.Changes = make([]CompressedResponse, zwht)
			}
			for zbai := range z.Changes {
				var zhct uint32
				zhct, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zhct > 0 {
					zhct--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Content":
						z.Changes[zbai].Content, err = dc.ReadBytes(z.Changes[zbai].Content)
						if err != nil {
							return
						}
					case "Size":
						z.Changes[zbai].Size, err = dc.ReadInt()
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
	// map header, size 3
	// write "ChangeIDToIndex"
	err = en.Append(0x83, 0xaf, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x44, 0x54, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.ChangeIDToIndex)))
	if err != nil {
		return
	}
	for zxvk, zbzg := range z.ChangeIDToIndex {
		err = en.WriteString(zxvk)
		if err != nil {
			return
		}
		err = en.WriteInt(zbzg)
		if err != nil {
			return
		}
	}
	// write "Changes"
	err = en.Append(0xa7, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Changes)))
	if err != nil {
		return
	}
	for zbai := range z.Changes {
		// map header, size 2
		// write "Content"
		err = en.Append(0x82, 0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
		if err != nil {
			return err
		}
		err = en.WriteBytes(z.Changes[zbai].Content)
		if err != nil {
			return
		}
		// write "Size"
		err = en.Append(0xa4, 0x53, 0x69, 0x7a, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Changes[zbai].Size)
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
	// map header, size 3
	// string "ChangeIDToIndex"
	o = append(o, 0x83, 0xaf, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x44, 0x54, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78)
	o = msgp.AppendMapHeader(o, uint32(len(z.ChangeIDToIndex)))
	for zxvk, zbzg := range z.ChangeIDToIndex {
		o = msgp.AppendString(o, zxvk)
		o = msgp.AppendInt(o, zbzg)
	}
	// string "Changes"
	o = append(o, 0xa7, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Changes)))
	for zbai := range z.Changes {
		// map header, size 2
		// string "Content"
		o = append(o, 0x82, 0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
		o = msgp.AppendBytes(o, z.Changes[zbai].Content)
		// string "Size"
		o = append(o, 0xa4, 0x53, 0x69, 0x7a, 0x65)
		o = msgp.AppendInt(o, z.Changes[zbai].Size)
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
	var zcua uint32
	zcua, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ChangeIDToIndex":
			var zxhx uint32
			zxhx, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.ChangeIDToIndex == nil && zxhx > 0 {
				z.ChangeIDToIndex = make(map[string]int, zxhx)
			} else if len(z.ChangeIDToIndex) > 0 {
				for key, _ := range z.ChangeIDToIndex {
					delete(z.ChangeIDToIndex, key)
				}
			}
			for zxhx > 0 {
				var zxvk string
				var zbzg int
				zxhx--
				zxvk, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zbzg, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					return
				}
				z.ChangeIDToIndex[zxvk] = zbzg
			}
		case "Changes":
			var zlqf uint32
			zlqf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Changes) >= int(zlqf) {
				z.Changes = (z.Changes)[:zlqf]
			} else {
				z.Changes = make([]CompressedResponse, zlqf)
			}
			for zbai := range z.Changes {
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
					case "Content":
						z.Changes[zbai].Content, bts, err = msgp.ReadBytesBytes(bts, z.Changes[zbai].Content)
						if err != nil {
							return
						}
					case "Size":
						z.Changes[zbai].Size, bts, err = msgp.ReadIntBytes(bts)
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
	s = 1 + 16 + msgp.MapHeaderSize
	if z.ChangeIDToIndex != nil {
		for zxvk, zbzg := range z.ChangeIDToIndex {
			_ = zbzg
			s += msgp.StringPrefixSize + len(zxvk) + msgp.IntSize
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zbai := range z.Changes {
		s += 1 + 8 + msgp.BytesPrefixSize + len(z.Changes[zbai].Content) + 5 + msgp.IntSize
	}
	s += 5 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CompressedResponse) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
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
