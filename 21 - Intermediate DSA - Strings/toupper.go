package dsa

func to_upper(A []byte) []byte {
	for i := 0; i < len(A); i++ {
		if A[i] >= 97 && A[i] <= 122 {
			A[i] = A[i] ^ 1<<5
		}
	}
	return A
}
