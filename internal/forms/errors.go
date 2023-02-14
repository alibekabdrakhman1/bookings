package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
func gcdOfStrings(str1 string, str2 string) string {
	out := ""
	if str1+str2 != str2+str1 {
		return ""
	} else {
		l := 0
		if len(str1) > len(str2) {
			l = len(str2)
		} else {
			l = len(str1)
		}
		for i := l - 1; i < l; i++ {
			if str1[i] == str2[i] {
				out = out + string(str1[i])
			} else {
				return out
			}
		}
	}
	return out

}
