//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidt() int {
	summ := 0
	for i := 0; i < len(b.amount); i++ {
		summ += int(b.amount[i])
	}

	return summ / len(b.amount)
}

func (t *CpuTemp) AverageCpuTemp() int {
	summ := 0
	for i := 0; i < len(t.temp); i++ {
		summ += int(t.temp[i])
	}
	return summ / len(t.temp)
}

// func (t *CpuTemp) AverageCpuTemp() int {
// 	summ := 0
// 	for i := 0; i < len(t.temp); i++ {
// 		summ += int(t.temp[i])
// 	}

// 	return summ / len(t.temp)
// }

func (m *MemoryUsage) AverageMemoryUse() int {
	summ := 0
	for _, val := range m.amount {
		summ += int(val)
	}

	return summ / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	fmt.Printf("Average bandwidth usage: %v\n", dash.AverageBandwidt())
	fmt.Printf("Average temp of CPU: %v\n", dash.AverageCpuTemp())
	fmt.Printf("Average memory usage: %v\n", dash.AverageMemoryUse())

	// b := bandwidth.AverageBandwidt()
	// t := temp.AverageTemp()
	// m := memory.AverageMemory()

	// fmt.Println(b, t, m)
}
