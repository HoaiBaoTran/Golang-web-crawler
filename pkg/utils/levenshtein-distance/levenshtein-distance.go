package levenshteindistance

func LevenshteinDistance(str1, str2 string) int {
	len1 := len(str1)
	len2 := len(str2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			cost := 0
			if str1[i-1] != str2[j-1] {
				cost = 1
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}

	return dp[len1][len2]
}
