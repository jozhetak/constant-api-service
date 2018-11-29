package models

const (
	ReservePartyInvalid    = iota
	ReservePartyPrimeTrust
	ReservePartyEth
)

type ReserseParty struct {
	ID       int
	Name     string
	IsActive bool
}
