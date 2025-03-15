package dsa

func to_lower(A []byte) []byte {
	for i := 0; i < len(A); i++ {
		if A[i] >= 'A' && A[i] <= 'Z' {
			A[i] = A[i] ^ 1<<5
		}
	}
	return A
}
