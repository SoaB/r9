package es

// DNA is an slice

import "r9/sb"

type DNA struct {
	Genes []float64
}

// NewDNA :
func NewDNA(num int) *DNA {
	gens := make([]float64, 0)
	for i := 0; i < num; i++ {
		gens = append(gens, sb.RandFloat64())
	}
	return &DNA{gens}
}

//Clone : clone dna
func (d *DNA) Clone() *DNA {
	gens := make([]float64, len(d.Genes), len(d.Genes))
	copy(gens, d.Genes)
	return &DNA{gens}
}

//CrossOver : create new DNA sequence from two DNA (self & partner)
func (d *DNA) CrossOver(partner *DNA) *DNA {
	length := len(d.Genes)
	child := make([]float64, length)
	// pick a middle point
	crossOver := sb.RandIntn(length)
	// Take half from one and half from another
	for i := 0; i < length; i++ {
		if i > crossOver {
			child[i] = d.Genes[i]
		} else {
			child[i] = partner.Genes[i]
		}
	}
	return &DNA{child}
}

//Mutate : Mutation with m probability
func (d *DNA) Mutate(m float64) {
	for i := 0; i < len(d.Genes); i++ {
		if sb.RandFloat64() < m {
			d.Genes[i] = sb.RandFloat64()
		}
	}
}
