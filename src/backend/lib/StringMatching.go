package lib

func ComputeFail(pattern string) []int {

	var (
		fail    []int
		m, i, j int
	)

	for i = 0; i < len(pattern); i++ {
		fail = append(fail, 0)
	}

	m, i, j = len(pattern), 1, 0

	for i < m {
		if pattern[j] == pattern[i] {
			fail[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			fail[i] = 0
			i++
		}
	}

	return fail[0 : len(pattern)-1]

}

func KmpMatch(text string, pattern string) int {

	var (
		fail       []int
		n, m, i, j int
	)

	n, m, i, j = len(text), len(pattern), 0, 0
	fail = ComputeFail(pattern)

	for i < n {
		if pattern[j] == text[i] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			i++
		}
	}

	return -1

}

func BuildLast(pattern string) map[string]int {

	var (
		last map[string]int
		i    int
	)

	last = make(map[string]int)

	for i = 0; i < len(pattern); i++ {
		last[string(pattern[i])] = i
	}

	return last

}

func BmMatch(text string, pattern string) int {

	var (
		last                map[string]int
		i, j, n, m, lo, min int
	)

	last = BuildLast(pattern)
	n, m = len(text), len(pattern)
	i, j = m-1, m-1

	if i > n-1 {
		return -1
	}

	for {

		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			lo = last[string(text[i])]

			if j < lo+1 {
				min = j
			} else {
				min = lo + 1
			}

			i = i + m - min
			j = m - 1
		}

		if i > n-1 {
			break
		}

	}

	return -1

}

func HammingDistance(text string, pattern string) (int, float32) {

	var (
		i, j, mismatch, min, minpos int
		percentage                  float32
	)

	min = -1
	minpos = 0
	for i = 0; i <= (len(text) - len(pattern)); i++ {
		mismatch = 0
		for j = 0; j <= len(pattern)-1; j++ {
			if text[i+j] != pattern[j] {
				mismatch += 1
			}
			if mismatch >= min && min != -1 {
				break
			}
		}
		if mismatch < min || min == -1 {
			min = mismatch
			minpos = i
		}
	}

	percentage = (float32(len(pattern) - min)) / (float32(len(pattern)))

	return minpos, percentage * 100

}
