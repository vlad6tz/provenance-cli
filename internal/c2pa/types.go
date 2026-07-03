package c2pa

type ValidationStatus string

const (
	StatusValid ValidationStatus= "VALID"
	StatusInvalid ValidationStatus= "INVALID"
	StatusUnverified ValidationStatus = 
)