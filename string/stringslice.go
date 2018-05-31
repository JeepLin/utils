package string

import "strings"

type StringSlice []string

func New(s []string) StringSlice {
	return StringSlice(s)
}

//Contains reports whether StringSlice is within s.
func (self StringSlice) Contains(s string) bool {
	return self.Index(s) >= 0
}

//Index returns the index of the first instance of StringSlice in s, or -1 if StringSlice is not present in s.
func (self StringSlice) Index(s string) int {
	n := len(self)
	if n == 0 {
		return -1
	}

	i := 0
	for i < n {
		if self[i] == s {
			return i
		}
		i ++
	}
	return -1
}

//Join concatenates the elements of StringSlice to create a single string.
func (self StringSlice) Join(sep string) string {
	return strings.Join(self, sep)
}

//ToUpper upper the elements of StringSlice
func (self StringSlice) ToUpper() []string {
	ret := make([]string, 0)
	for _, v := range self {
		ret = append(ret, strings.ToUpper(v))
	}
	return ret
}

//ToLower lower the elements of StringSlice
func (self StringSlice) ToLower() []string {
	ret := make([]string, 0)
	for _, v := range self {
		ret = append(ret, strings.ToLower(v))
	}
	return ret
}



