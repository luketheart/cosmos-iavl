package iavl

var enables []bool = make([]bool, 0)

func AddEnable() {
	enables = append(enables, false)
}

func RemoveEnable() {
	if len(enables) == 0 {
		panic("SetEnable failed")
	}
	enables = enables[:len(enables)-1]
}

func SetEnable(t bool) {
	if len(enables) == 0 {
		panic("SetEnable failed")
	}
	enables[len(enables)-1] = t
}

func GetEnable() bool {
	if len(enables) == 0 {
		panic("GetEnable failed")
	}
	return enables[len(enables)-1]
}

func Println(a ...any) {
	for _, enable := range enables {
		if !enable {
			return
		}
	}
	// fmt.Println(a...)
}

func Printf(format string, a ...any) {
	for _, enable := range enables {
		if !enable {
			return
		}
	}
	// fmt.Printf(format, a...)
}
