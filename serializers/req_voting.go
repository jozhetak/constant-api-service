package serializers

// Candidate
type VotingBoardCandidateRequest struct {
	BoardType      int    `json:"BoardType"`
	PaymentAddress string `json:"PaymentAddress"`
}

// end Candidate

// Proposal
type VotingProposalRequest struct {
	Type int `json:"Type"` // 1: DCB, 2 GOV

	DCB VotingProposalDCBRequest `json:"DCB"`
	GOV VotingProposalGOVRequest `json:"DCB"`
}

type VotingProposalDCBRequest struct {
	SaleData
}
type VotingProposalGOVRequest struct {
	SalaryPerTx  uint64 // salary for each tx in block(mili constant)
	BasicSalary  uint64 // basic salary per block(mili constant)
	TxFee        uint64
	SellingBonds SellingBonds
	RefundInfo   RefundInfo
}

type SaleData struct {
	SaleID string // Unique id of the crowdsale to store in db
	BondID string // in case either base or quote asset is bond

	BuyingAsset  string
	SellingAsset string
	Price        uint64
	EndBlock     int32
}

type SellingBonds struct {
	BondsToSell    uint64
	BondPrice      uint64 // in Constant unit
	Maturity       uint32
	BuyBackPrice   uint64 // in Constant unit
	StartSellingAt uint32 // start selling bonds at block height
	SellingWithin  uint32 // selling bonds within n blocks
}

type RefundInfo struct {
	ThresholdToLargeTx uint64
	RefundAmount       uint64
}

// end Proposal
