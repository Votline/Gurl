package shapes

func QuadUV() []float32 {
	return []float32{
		-1, 1, 0,     0, 0,
		-1, -1, 0,    0, 1,
		1, -1, 0,     1, 1,

		-1, 1, 0,     0, 0,
		1, -1, 0,     1, 1,
		1, 1, 0,      1, 0,
	}
}
