package parsing

import "strconv"

func StrSliceToIntSlice(ss []string) []int {
	r := []int{}

	for _, s := range ss {

		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}


		r = append(r, n)
	}

	return r
}
