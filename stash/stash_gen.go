package stash

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

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
		case "Verified":
			z.Verified, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "League":
			z.League, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ID":
			z.ID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "TypeLine":
			z.TypeLine, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Identified":
			z.Identified, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Corrupted":
			z.Corrupted, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "ImplicitMods":
			var zhct uint32
			zhct, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.ImplicitMods) >= int(zhct) {
				z.ImplicitMods = (z.ImplicitMods)[:zhct]
			} else {
				z.ImplicitMods = make([]ItemMod, zhct)
			}
			for zxvk := range z.ImplicitMods {
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
					case "Template":
						z.ImplicitMods[zxvk].Template, err = dc.ReadBytes(z.ImplicitMods[zxvk].Template)
						if err != nil {
							return
						}
					case "Values":
						var zxhx uint32
						zxhx, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.ImplicitMods[zxvk].Values) >= int(zxhx) {
							z.ImplicitMods[zxvk].Values = (z.ImplicitMods[zxvk].Values)[:zxhx]
						} else {
							z.ImplicitMods[zxvk].Values = make([]uint16, zxhx)
						}
						for zbzg := range z.ImplicitMods[zxvk].Values {
							z.ImplicitMods[zxvk].Values[zbzg], err = dc.ReadUint16()
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
		case "ExplicitMods":
			var zlqf uint32
			zlqf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.ExplicitMods) >= int(zlqf) {
				z.ExplicitMods = (z.ExplicitMods)[:zlqf]
			} else {
				z.ExplicitMods = make([]ItemMod, zlqf)
			}
			for zbai := range z.ExplicitMods {
				var zdaf uint32
				zdaf, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zdaf > 0 {
					zdaf--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Template":
						z.ExplicitMods[zbai].Template, err = dc.ReadBytes(z.ExplicitMods[zbai].Template)
						if err != nil {
							return
						}
					case "Values":
						var zpks uint32
						zpks, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.ExplicitMods[zbai].Values) >= int(zpks) {
							z.ExplicitMods[zbai].Values = (z.ExplicitMods[zbai].Values)[:zpks]
						} else {
							z.ExplicitMods[zbai].Values = make([]uint16, zpks)
						}
						for zcmr := range z.ExplicitMods[zbai].Values {
							z.ExplicitMods[zbai].Values[zcmr], err = dc.ReadUint16()
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
		case "Note":
			z.Note, err = dc.ReadString()
			if err != nil {
				return
			}
		case "UtilityMods":
			var zjfb uint32
			zjfb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.UtilityMods) >= int(zjfb) {
				z.UtilityMods = (z.UtilityMods)[:zjfb]
			} else {
				z.UtilityMods = make([]string, zjfb)
			}
			for zajw := range z.UtilityMods {
				z.UtilityMods[zajw], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "DescrText":
			z.DescrText, err = dc.ReadString()
			if err != nil {
				return
			}
		case "StashID":
			z.StashID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "RootType":
			z.RootType, err = dc.ReadString()
			if err != nil {
				return
			}
		case "RootFlavor":
			z.RootFlavor, err = dc.ReadString()
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
func (z *Item) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 15
	// write "Verified"
	err = en.Append(0x8f, 0xa8, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Verified)
	if err != nil {
		return
	}
	// write "League"
	err = en.Append(0xa6, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.League)
	if err != nil {
		return
	}
	// write "ID"
	err = en.Append(0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ID)
	if err != nil {
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "TypeLine"
	err = en.Append(0xa8, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TypeLine)
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
	// write "Corrupted"
	err = en.Append(0xa9, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Corrupted)
	if err != nil {
		return
	}
	// write "ImplicitMods"
	err = en.Append(0xac, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.ImplicitMods)))
	if err != nil {
		return
	}
	for zxvk := range z.ImplicitMods {
		// map header, size 2
		// write "Template"
		err = en.Append(0x82, 0xa8, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteBytes(z.ImplicitMods[zxvk].Template)
		if err != nil {
			return
		}
		// write "Values"
		err = en.Append(0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.ImplicitMods[zxvk].Values)))
		if err != nil {
			return
		}
		for zbzg := range z.ImplicitMods[zxvk].Values {
			err = en.WriteUint16(z.ImplicitMods[zxvk].Values[zbzg])
			if err != nil {
				return
			}
		}
	}
	// write "ExplicitMods"
	err = en.Append(0xac, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.ExplicitMods)))
	if err != nil {
		return
	}
	for zbai := range z.ExplicitMods {
		// map header, size 2
		// write "Template"
		err = en.Append(0x82, 0xa8, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteBytes(z.ExplicitMods[zbai].Template)
		if err != nil {
			return
		}
		// write "Values"
		err = en.Append(0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.ExplicitMods[zbai].Values)))
		if err != nil {
			return
		}
		for zcmr := range z.ExplicitMods[zbai].Values {
			err = en.WriteUint16(z.ExplicitMods[zbai].Values[zcmr])
			if err != nil {
				return
			}
		}
	}
	// write "Note"
	err = en.Append(0xa4, 0x4e, 0x6f, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Note)
	if err != nil {
		return
	}
	// write "UtilityMods"
	err = en.Append(0xab, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.UtilityMods)))
	if err != nil {
		return
	}
	for zajw := range z.UtilityMods {
		err = en.WriteString(z.UtilityMods[zajw])
		if err != nil {
			return
		}
	}
	// write "DescrText"
	err = en.Append(0xa9, 0x44, 0x65, 0x73, 0x63, 0x72, 0x54, 0x65, 0x78, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.DescrText)
	if err != nil {
		return
	}
	// write "StashID"
	err = en.Append(0xa7, 0x53, 0x74, 0x61, 0x73, 0x68, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.StashID)
	if err != nil {
		return
	}
	// write "RootType"
	err = en.Append(0xa8, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.RootType)
	if err != nil {
		return
	}
	// write "RootFlavor"
	err = en.Append(0xaa, 0x52, 0x6f, 0x6f, 0x74, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.RootFlavor)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Item) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 15
	// string "Verified"
	o = append(o, 0x8f, 0xa8, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Verified)
	// string "League"
	o = append(o, 0xa6, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65)
	o = msgp.AppendString(o, z.League)
	// string "ID"
	o = append(o, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, z.ID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "TypeLine"
	o = append(o, 0xa8, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65)
	o = msgp.AppendString(o, z.TypeLine)
	// string "Identified"
	o = append(o, 0xaa, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Identified)
	// string "Corrupted"
	o = append(o, 0xa9, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Corrupted)
	// string "ImplicitMods"
	o = append(o, 0xac, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.ImplicitMods)))
	for zxvk := range z.ImplicitMods {
		// map header, size 2
		// string "Template"
		o = append(o, 0x82, 0xa8, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65)
		o = msgp.AppendBytes(o, z.ImplicitMods[zxvk].Template)
		// string "Values"
		o = append(o, 0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.ImplicitMods[zxvk].Values)))
		for zbzg := range z.ImplicitMods[zxvk].Values {
			o = msgp.AppendUint16(o, z.ImplicitMods[zxvk].Values[zbzg])
		}
	}
	// string "ExplicitMods"
	o = append(o, 0xac, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.ExplicitMods)))
	for zbai := range z.ExplicitMods {
		// map header, size 2
		// string "Template"
		o = append(o, 0x82, 0xa8, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65)
		o = msgp.AppendBytes(o, z.ExplicitMods[zbai].Template)
		// string "Values"
		o = append(o, 0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.ExplicitMods[zbai].Values)))
		for zcmr := range z.ExplicitMods[zbai].Values {
			o = msgp.AppendUint16(o, z.ExplicitMods[zbai].Values[zcmr])
		}
	}
	// string "Note"
	o = append(o, 0xa4, 0x4e, 0x6f, 0x74, 0x65)
	o = msgp.AppendString(o, z.Note)
	// string "UtilityMods"
	o = append(o, 0xab, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.UtilityMods)))
	for zajw := range z.UtilityMods {
		o = msgp.AppendString(o, z.UtilityMods[zajw])
	}
	// string "DescrText"
	o = append(o, 0xa9, 0x44, 0x65, 0x73, 0x63, 0x72, 0x54, 0x65, 0x78, 0x74)
	o = msgp.AppendString(o, z.DescrText)
	// string "StashID"
	o = append(o, 0xa7, 0x53, 0x74, 0x61, 0x73, 0x68, 0x49, 0x44)
	o = msgp.AppendString(o, z.StashID)
	// string "RootType"
	o = append(o, 0xa8, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.RootType)
	// string "RootFlavor"
	o = append(o, 0xaa, 0x52, 0x6f, 0x6f, 0x74, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x72)
	o = msgp.AppendString(o, z.RootFlavor)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcxo uint32
	zcxo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Verified":
			z.Verified, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "League":
			z.League, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ID":
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TypeLine":
			z.TypeLine, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Identified":
			z.Identified, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Corrupted":
			z.Corrupted, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "ImplicitMods":
			var zeff uint32
			zeff, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.ImplicitMods) >= int(zeff) {
				z.ImplicitMods = (z.ImplicitMods)[:zeff]
			} else {
				z.ImplicitMods = make([]ItemMod, zeff)
			}
			for zxvk := range z.ImplicitMods {
				var zrsw uint32
				zrsw, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zrsw > 0 {
					zrsw--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Template":
						z.ImplicitMods[zxvk].Template, bts, err = msgp.ReadBytesBytes(bts, z.ImplicitMods[zxvk].Template)
						if err != nil {
							return
						}
					case "Values":
						var zxpk uint32
						zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.ImplicitMods[zxvk].Values) >= int(zxpk) {
							z.ImplicitMods[zxvk].Values = (z.ImplicitMods[zxvk].Values)[:zxpk]
						} else {
							z.ImplicitMods[zxvk].Values = make([]uint16, zxpk)
						}
						for zbzg := range z.ImplicitMods[zxvk].Values {
							z.ImplicitMods[zxvk].Values[zbzg], bts, err = msgp.ReadUint16Bytes(bts)
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
		case "ExplicitMods":
			var zdnj uint32
			zdnj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.ExplicitMods) >= int(zdnj) {
				z.ExplicitMods = (z.ExplicitMods)[:zdnj]
			} else {
				z.ExplicitMods = make([]ItemMod, zdnj)
			}
			for zbai := range z.ExplicitMods {
				var zobc uint32
				zobc, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zobc > 0 {
					zobc--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Template":
						z.ExplicitMods[zbai].Template, bts, err = msgp.ReadBytesBytes(bts, z.ExplicitMods[zbai].Template)
						if err != nil {
							return
						}
					case "Values":
						var zsnv uint32
						zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.ExplicitMods[zbai].Values) >= int(zsnv) {
							z.ExplicitMods[zbai].Values = (z.ExplicitMods[zbai].Values)[:zsnv]
						} else {
							z.ExplicitMods[zbai].Values = make([]uint16, zsnv)
						}
						for zcmr := range z.ExplicitMods[zbai].Values {
							z.ExplicitMods[zbai].Values[zcmr], bts, err = msgp.ReadUint16Bytes(bts)
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
		case "Note":
			z.Note, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "UtilityMods":
			var zkgt uint32
			zkgt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.UtilityMods) >= int(zkgt) {
				z.UtilityMods = (z.UtilityMods)[:zkgt]
			} else {
				z.UtilityMods = make([]string, zkgt)
			}
			for zajw := range z.UtilityMods {
				z.UtilityMods[zajw], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "DescrText":
			z.DescrText, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "StashID":
			z.StashID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RootType":
			z.RootType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RootFlavor":
			z.RootFlavor, bts, err = msgp.ReadStringBytes(bts)
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
func (z *Item) Msgsize() (s int) {
	s = 1 + 9 + msgp.BoolSize + 7 + msgp.StringPrefixSize + len(z.League) + 3 + msgp.StringPrefixSize + len(z.ID) + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.StringPrefixSize + len(z.TypeLine) + 11 + msgp.BoolSize + 10 + msgp.BoolSize + 13 + msgp.ArrayHeaderSize
	for zxvk := range z.ImplicitMods {
		s += 1 + 9 + msgp.BytesPrefixSize + len(z.ImplicitMods[zxvk].Template) + 7 + msgp.ArrayHeaderSize + (len(z.ImplicitMods[zxvk].Values) * (msgp.Uint16Size))
	}
	s += 13 + msgp.ArrayHeaderSize
	for zbai := range z.ExplicitMods {
		s += 1 + 9 + msgp.BytesPrefixSize + len(z.ExplicitMods[zbai].Template) + 7 + msgp.ArrayHeaderSize + (len(z.ExplicitMods[zbai].Values) * (msgp.Uint16Size))
	}
	s += 5 + msgp.StringPrefixSize + len(z.Note) + 12 + msgp.ArrayHeaderSize
	for zajw := range z.UtilityMods {
		s += msgp.StringPrefixSize + len(z.UtilityMods[zajw])
	}
	s += 10 + msgp.StringPrefixSize + len(z.DescrText) + 8 + msgp.StringPrefixSize + len(z.StashID) + 9 + msgp.StringPrefixSize + len(z.RootType) + 11 + msgp.StringPrefixSize + len(z.RootFlavor)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ItemMod) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zpez uint32
	zpez, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpez > 0 {
		zpez--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Template":
			z.Template, err = dc.ReadBytes(z.Template)
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
				z.Values = make([]uint16, zqke)
			}
			for zema := range z.Values {
				z.Values[zema], err = dc.ReadUint16()
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
	// write "Template"
	err = en.Append(0x82, 0xa8, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Template)
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
	for zema := range z.Values {
		err = en.WriteUint16(z.Values[zema])
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
	// string "Template"
	o = append(o, 0x82, 0xa8, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65)
	o = msgp.AppendBytes(o, z.Template)
	// string "Values"
	o = append(o, 0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Values)))
	for zema := range z.Values {
		o = msgp.AppendUint16(o, z.Values[zema])
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
		case "Template":
			z.Template, bts, err = msgp.ReadBytesBytes(bts, z.Template)
			if err != nil {
				return
			}
		case "Values":
			var zyzr uint32
			zyzr, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Values) >= int(zyzr) {
				z.Values = (z.Values)[:zyzr]
			} else {
				z.Values = make([]uint16, zyzr)
			}
			for zema := range z.Values {
				z.Values[zema], bts, err = msgp.ReadUint16Bytes(bts)
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
	s = 1 + 9 + msgp.BytesPrefixSize + len(z.Template) + 7 + msgp.ArrayHeaderSize + (len(z.Values) * (msgp.Uint16Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PropertyValue) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zywj uint32
	zywj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zywj > 0 {
		zywj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Value":
			z.Value, err = dc.ReadString()
			if err != nil {
				return
			}
		case "PrintKey":
			z.PrintKey, err = dc.ReadInt()
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
func (z PropertyValue) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Value"
	err = en.Append(0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Value)
	if err != nil {
		return
	}
	// write "PrintKey"
	err = en.Append(0xa8, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x4b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PrintKey)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z PropertyValue) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Value"
	o = append(o, 0x82, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.Value)
	// string "PrintKey"
	o = append(o, 0xa8, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x4b, 0x65, 0x79)
	o = msgp.AppendInt(o, z.PrintKey)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PropertyValue) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjpj uint32
	zjpj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Value":
			z.Value, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "PrintKey":
			z.PrintKey, bts, err = msgp.ReadIntBytes(bts)
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
func (z PropertyValue) Msgsize() (s int) {
	s = 1 + 6 + msgp.StringPrefixSize + len(z.Value) + 9 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Response) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrfe uint32
	zrfe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrfe > 0 {
		zrfe--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "NextChangeID":
			z.NextChangeID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Stashes":
			var zgmo uint32
			zgmo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Stashes) >= int(zgmo) {
				z.Stashes = (z.Stashes)[:zgmo]
			} else {
				z.Stashes = make([]Stash, zgmo)
			}
			for zzpf := range z.Stashes {
				err = z.Stashes[zzpf].DecodeMsg(dc)
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
func (z *Response) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "NextChangeID"
	err = en.Append(0x82, 0xac, 0x4e, 0x65, 0x78, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.NextChangeID)
	if err != nil {
		return
	}
	// write "Stashes"
	err = en.Append(0xa7, 0x53, 0x74, 0x61, 0x73, 0x68, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Stashes)))
	if err != nil {
		return
	}
	for zzpf := range z.Stashes {
		err = z.Stashes[zzpf].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Response) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "NextChangeID"
	o = append(o, 0x82, 0xac, 0x4e, 0x65, 0x78, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x44)
	o = msgp.AppendString(o, z.NextChangeID)
	// string "Stashes"
	o = append(o, 0xa7, 0x53, 0x74, 0x61, 0x73, 0x68, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Stashes)))
	for zzpf := range z.Stashes {
		o, err = z.Stashes[zzpf].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Response) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "NextChangeID":
			z.NextChangeID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Stashes":
			var zeth uint32
			zeth, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Stashes) >= int(zeth) {
				z.Stashes = (z.Stashes)[:zeth]
			} else {
				z.Stashes = make([]Stash, zeth)
			}
			for zzpf := range z.Stashes {
				bts, err = z.Stashes[zzpf].UnmarshalMsg(bts)
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
func (z *Response) Msgsize() (s int) {
	s = 1 + 13 + msgp.StringPrefixSize + len(z.NextChangeID) + 8 + msgp.ArrayHeaderSize
	for zzpf := range z.Stashes {
		s += z.Stashes[zzpf].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Stash) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrjx uint32
	zrjx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrjx > 0 {
		zrjx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "AccountName":
			z.AccountName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "LastCharacterName":
			z.LastCharacterName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ID":
			z.ID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Stash":
			z.Stash, err = dc.ReadString()
			if err != nil {
				return
			}
		case "StashType":
			z.StashType, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Items":
			var zawn uint32
			zawn, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Items) >= int(zawn) {
				z.Items = (z.Items)[:zawn]
			} else {
				z.Items = make([]Item, zawn)
			}
			for zsbz := range z.Items {
				err = z.Items[zsbz].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Public":
			z.Public, err = dc.ReadBool()
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
	// map header, size 7
	// write "AccountName"
	err = en.Append(0x87, 0xab, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AccountName)
	if err != nil {
		return
	}
	// write "LastCharacterName"
	err = en.Append(0xb1, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.LastCharacterName)
	if err != nil {
		return
	}
	// write "ID"
	err = en.Append(0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ID)
	if err != nil {
		return
	}
	// write "Stash"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x73, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Stash)
	if err != nil {
		return
	}
	// write "StashType"
	err = en.Append(0xa9, 0x53, 0x74, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.StashType)
	if err != nil {
		return
	}
	// write "Items"
	err = en.Append(0xa5, 0x49, 0x74, 0x65, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Items)))
	if err != nil {
		return
	}
	for zsbz := range z.Items {
		err = z.Items[zsbz].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Public"
	err = en.Append(0xa6, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Public)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Stash) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "AccountName"
	o = append(o, 0x87, 0xab, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.AccountName)
	// string "LastCharacterName"
	o = append(o, 0xb1, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.LastCharacterName)
	// string "ID"
	o = append(o, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, z.ID)
	// string "Stash"
	o = append(o, 0xa5, 0x53, 0x74, 0x61, 0x73, 0x68)
	o = msgp.AppendString(o, z.Stash)
	// string "StashType"
	o = append(o, 0xa9, 0x53, 0x74, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.StashType)
	// string "Items"
	o = append(o, 0xa5, 0x49, 0x74, 0x65, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Items)))
	for zsbz := range z.Items {
		o, err = z.Items[zsbz].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Public"
	o = append(o, 0xa6, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63)
	o = msgp.AppendBool(o, z.Public)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Stash) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwel uint32
	zwel, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwel > 0 {
		zwel--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "AccountName":
			z.AccountName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "LastCharacterName":
			z.LastCharacterName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ID":
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Stash":
			z.Stash, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "StashType":
			z.StashType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Items":
			var zrbe uint32
			zrbe, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Items) >= int(zrbe) {
				z.Items = (z.Items)[:zrbe]
			} else {
				z.Items = make([]Item, zrbe)
			}
			for zsbz := range z.Items {
				bts, err = z.Items[zsbz].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Public":
			z.Public, bts, err = msgp.ReadBoolBytes(bts)
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
	s = 1 + 12 + msgp.StringPrefixSize + len(z.AccountName) + 18 + msgp.StringPrefixSize + len(z.LastCharacterName) + 3 + msgp.StringPrefixSize + len(z.ID) + 6 + msgp.StringPrefixSize + len(z.Stash) + 10 + msgp.StringPrefixSize + len(z.StashType) + 6 + msgp.ArrayHeaderSize
	for zsbz := range z.Items {
		s += z.Items[zsbz].Msgsize()
	}
	s += 7 + msgp.BoolSize
	return
}
