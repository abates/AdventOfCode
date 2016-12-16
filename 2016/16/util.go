package main

import (
	"bytes"
)

type Disk struct {
	size int
	byts []byte
}

func NewDisk(size int) *Disk {
	numByts := size / 8
	if size%8 != 0 {
		numByts++
	}

	return &Disk{
		size: size,
		byts: make([]byte, numByts),
	}
}

func (d *Disk) GetBit(i int) int {
	if i >= d.size {
		return 0
	}

	byt := d.byts[i/8]
	bit := 7 - uint(i%8)
	if byt&(1<<bit) == 0 {
		return 0
	}
	return 1
}

func (d *Disk) Fill(prefix string) {
	length := len(prefix)
	for i, c := range prefix {
		if c == '1' {
			d.SetBit(i, 1)
		}
	}

	for ; length < d.size; length += length + 1 {
		Generate(length, d)
	}
}

func (d *Disk) SetBit(i, v int) {
	if i >= d.size {
		return
	}

	byt := i / 8
	bit := 7 - uint(i%8)
	if v == 0 {
		d.byts[byt] &= ^(1 << bit)
	} else if v == 1 {
		d.byts[byt] |= (1 << bit)
	}
}

func Generate(bitLen int, disk *Disk) {
	for i := 0; i < bitLen; i++ {
		if disk.GetBit(i) == 0 {
			disk.SetBit(bitLen+bitLen-i, 1)
		}
	}
}

func (d *Disk) String() string {
	buffer := &bytes.Buffer{}

	for i := 0; i < d.size; i++ {
		if d.GetBit(i) == 1 {
			buffer.Write([]byte("1"))
		} else {
			buffer.Write([]byte("0"))
		}
	}

	return buffer.String()
}

func (d *Disk) Checksum() string {
	oldDisk := d
	var checksum *Disk

	for {
		checksum = NewDisk(oldDisk.size / 2)
		for i := 0; i < oldDisk.size; i += 2 {
			if oldDisk.GetBit(i) == oldDisk.GetBit(i+1) {
				checksum.SetBit(i/2, 1)
			}
		}
		if checksum.size%2 != 0 {
			break
		}
		oldDisk = checksum
	}
	return checksum.String()
}
