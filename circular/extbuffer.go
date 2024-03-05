package circular

type ExtendedCircularBuffer struct {
	CircularBuffer
}

func (cb *ExtendedCircularBuffer) AddValues(vals ...float64) {
	for _, val := range vals {
		cb.AddValue(val)
	}
}

func NewExtCircularBuffer(size int) ExtendedCircularBuffer {
	return ExtendedCircularBuffer{
		CircularBuffer: NewCircularBuffer(size),
	}
}
