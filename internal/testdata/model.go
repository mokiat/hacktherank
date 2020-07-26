package testdata

type Sample struct {
	InputLines  []string
	OutputLines []string
}

func (s Sample) IsEmpty() bool {
	return (len(s.InputLines) == 0) && (len(s.OutputLines) == 0)
}
