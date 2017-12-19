package main

import (
	"fmt"

	"github.com/abates/AdventOfCode/2017/graph"
	"github.com/abates/AdventOfCode/2017/hash"
)

type Row struct {
	buf []byte
}

func (r *Row) Len() int {
	return 8 * len(r.buf)
}

func (r *Row) Bit(n int) int {
	bit := r.Len() - n - 1
	index := len(r.buf) - 1 - bit/8
	offset := bit % 8
	return int(0x01 & (r.buf[index] >> uint(offset)))
}

type Disk struct {
	rows  []*Row
	graph *graph.Graph
}

func NewDisk() *Disk {
	return &Disk{
		graph: graph.New(),
	}
}

func (d *Disk) Groups() int {
	count := 0
	visited := make(map[string]bool)
	for id, vertex := range d.graph.Vertices {
		if found := visited[id]; !found {
			count++
			nodes := vertex.Connected()
			for id, _ := range nodes {
				visited[id] = true
			}
		}
	}
	return count
}

func (d *Disk) Count() int {
	count := 0
	for _, row := range d.rows {
		for i := 0; i < row.Len(); i++ {
			count += int(row.Bit(i))
		}
	}
	return count
}

func (d *Disk) Append(buf []byte) {
	row := &Row{buf}
	d.rows = append(d.rows, row)
	y := len(d.rows) - 1
	for x := 0; x < row.Len(); x++ {
		id := y*row.Len() + x
		if row.Bit(x) > 0 {
			vertex := d.graph.FindOrCreateVertex(fmt.Sprintf("%d", id))
			if y > 0 {
				prevId := (y-1)*row.Len() + x
				prevRow := d.rows[y-1]
				if prevRow.Bit(x) > 0 {
					vertex.Connect(d.graph.FindOrCreateVertex(fmt.Sprintf("%d", prevId)))
				}
			}
			if x > 0 {
				if row.Bit(x-1) > 0 {
					vertex.Connect(d.graph.FindOrCreateVertex(fmt.Sprintf("%d", id-1)))
				}
			}
		}
	}
}

func main() {
	//input := "flqrgnkx"
	//input := "a0c2017..."
	input := "jxqlasbh"
	disk := NewDisk()
	for i := 0; i < 128; i++ {
		disk.Append([]byte(hash.ComputeString(fmt.Sprintf("%s-%d", input, i)).Hash()))
	}
	fmt.Printf("%d\n", disk.Count())

	fmt.Printf("Groups: %d\n", disk.Groups())
}
