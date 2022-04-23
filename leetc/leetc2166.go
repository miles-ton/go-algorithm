package leetc

import "strings"

type Bitset struct {
	bits     []byte
	oneCnt   int
	size     int
	reversed byte
}

func Constructor2166(size int) Bitset {
	return Bitset{bits: make([]byte, size), size: size}
}

func (this *Bitset) Fix(idx int) {
	if this.bits[idx]&1^this.reversed == 0 {
		this.oneCnt++
		this.bits[idx] ^= 1
	}
}

func (this *Bitset) Unfix(idx int) {
	if this.bits[idx]&1^this.reversed == 1 {
		this.oneCnt--
		this.bits[idx] ^= 1
	}
}

func (this *Bitset) Flip() {
	this.reversed ^= 1
	this.oneCnt = this.size - this.oneCnt
}

func (this *Bitset) All() bool {
	return this.oneCnt == this.size
}

func (this *Bitset) One() bool {
	return this.oneCnt > 0
}

func (this *Bitset) Count() int {
	return this.oneCnt
}

func (this *Bitset) ToString() string {
	ret := strings.Builder{}
	for _, v := range this.bits {
		if v^this.reversed == 1 {
			ret.WriteString("1")
		} else {
			ret.WriteString("0")
		}
	}
	return ret.String()
}
