package StringMatching

func computeFail(pattern string) []int {

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

func kmpMatch(text string, pattern string) int {

	var (
		fail       []int
		n, m, i, j int
	)

	n, m, i, j = len(text), len(pattern), 0, 0
	fail = computeFail(pattern)

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

func buildLast(pattern string) map[string]int {

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

func bmMatch(text string, pattern string) int {

	var (
		last                map[string]int
		i, j, n, m, lo, min int
	)

	last = buildLast(pattern)
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
