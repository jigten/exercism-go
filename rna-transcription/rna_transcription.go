package strand

func ToRNA(dna string) string {
	rna := ""

	for _, c := range dna {
		if c == 'G' {
			rna += "C"
		} else if c == 'C' {
			rna += "G"
		} else if c == 'T' {
			rna += "A"
		} else {
			rna += "U"
		}
	}
	return rna
}
