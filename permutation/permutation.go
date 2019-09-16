package permutation

//策略排列
//输入p，返回全排列
//如输入[1 2 3]，则返回[123 132 213 231 312 321]
func permutation(p []string) [][]string {
	if len(p) == 1 {
		return [][]string{p}
	}
	pa := p[:len(p)-1]
	pl := p[len(p)-1]
	return permutationLoop(permutation(pa), pl)
}

//返回集合全排列
//输入p，返回全排列
//如输入[1 2 3]，则返回[123 132 213 231 312 321]
func All(p []string) [][]string {
	l := len(p)
	if l < 3 || l > 5 {
		panic("Illegal argument. array length must > 3 and <= 5.")
	}
	return permutation(p)
}

//返回集合首尾相同全排列
//输入p，返回全排列
//如输入[1 2 3]，则返回[1231 1321 2132 2312 3123 3213]
func EndEqBegin(p []string) [][]string {
	r := All(p)
	s := make([][]string, len(r))
	for i, v := range r {
		s[i] = append(v, v[0])
	}
	return s
}

//返回集合全排列（指定flag开头，且首尾相同）
//输入p，返回全排列
//如输入[1 2 3]，1，则返回[1231 1321]
func StartAsFlagEndEqBegin(p []string, f string) [][]string {
	r := All(p)
	m := make([][]string, 0)
	for _, v := range r {
		if v[0] == f {
			n := append(v, v[0])
			m = append(m, n)
		}
	}
	return m
}

func permutationLoop(p [][]string, l string) [][]string {
	//保存结果的slice
	result := make([][]string, 0)
	for _, v := range p {
		result = append(result, permutationItem(v, l)...)
	}
	return result
}

func permutationItem(p []string, last string) [][]string {
	l := len(p)
	result := make([][]string, l+1)
	index := 0

	for i := 0; i < len(p); i++ {
		newEle := make([]string, len(p)+1)
		newEle[i] = last
		if i == 0 {
			copy(newEle[i+1:], p)
		} else {
			copy(newEle[:i], p[:i])
			copy(newEle[i+1:], p[i:])
		}
		result[index] = newEle
		index++
	}
	result[index] = append(p, last)
	index++
	return result
}
