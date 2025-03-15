package dsa

func solve(A []byte) int {
	for i := 0; i < len(A); i++ {
		if (A[i] < 'a' || A[i] > 'z') && (A[i] < 'A' || A[i] > 'Z') && (A[i] < '0' || A[i] > '9') {
			return 0
		}
	}
	return 1
}
