package piscine

func IsUpper(s string) bool {
	if s <= "65" || s >= "90" {
		return false
	} else {
		return true
	}
}
