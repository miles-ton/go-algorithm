package struc

type SdsHdr struct {
	len int
	free int
	buf []byte
}

func (this *SdsHdr) SdsCreate(s []byte) *SdsHdr {
	return nil
}
// SdsNew create an empty sds
func (this *SdsHdr) SdsNew() *SdsHdr {
	return nil
}
func (this *SdsHdr) SdsAvailLen() *SdsHdr {
	return nil
}

// SdsDuplicate Copy new duplicate
func (this *SdsHdr) SdsDuplicate() *SdsHdr {
	return nil
}
func (this *SdsHdr) SdsClear() *SdsHdr {
	return nil
}
func (this *SdsHdr) SdsCatB(s []byte) *SdsHdr {
	return nil
}
func (this *SdsHdr) SdsCatSds(tail *SdsHdr) *SdsHdr {
	return nil
}
func (this *SdsHdr) SdsReplace(s []byte) *SdsHdr {
	return nil
}
func (this *SdsHdr) SdsCmp() *SdsHdr {
	return nil
}
func (this *SdsHdr) SdsTrim() *SdsHdr {
	return nil
}
