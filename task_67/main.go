package main

func main() {
	//println(addBinary("1111", "1111"))
	println(addBinary("1", "111"))
}

func addBinary(a, b string) string {
	//builder := strings.Builder{}
	if len(b) > len(a) {
		a, b = b, a
	}
	res := make([]rune, len(a))
	i, k := len(a)-1, len(b)-1
	toBear := uint8(0)
	for {
		var u uint8
		if i < 0 && k < 0 {
			if toBear != 0 {
				return "1" + string(res)
				//builder.WriteString(strconv.Itoa(int(toBear)))
			}
			break
		}
		l := 0
		if i < 0 {
			l = k
			u = b[k] - '0' + toBear
		} else if k < 0 {
			l = i
			u = a[i] - '0' + toBear
		} else {
			l = i
			u = a[i] - '0' + b[k] - '0' + toBear
		}
		toBear = u / 2
		u = u % 2
		res[l] = rune('0' + u)
		//builder.WriteString(strconv.Itoa(int(u)))
		i--
		k--
	}
	return string(res)
	//return reverse(builder.String())
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
