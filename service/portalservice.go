package service

import (
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/dao/portal"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/ninjadotorg/constant-api-service/common"
	"github.com/ninjadotorg/constant-api-service/service/3rd/ethereum"
)

type Portal struct {
	r         *portal.Portal
	bc        *blockchain.Blockchain
	ethclient *ethereum.EthClient
}

func NewPortal(r *portal.Portal, bc *blockchain.Blockchain) *Portal {
	return &Portal{
		r:         r,
		bc:        bc,
		ethclient: ethereum.CreateEthClient(),
	}
}

func (p *Portal) CreateBorrow(u *models.User, req serializers.BorrowReq) (*serializers.BorrowResp, error) {
	startDate, err := time.Parse(common.DateTimeLayoutFormat, req.StartDate)
	if err != nil {
		return nil, errors.Wrap(err, "b.r.Create")
	}
	if req.LoanRequest.ReceiveAddress == "" {
		req.LoanRequest.ReceiveAddress = u.PaymentAddress
	}
	endDate := startDate.Add(time.Duration(req.LoanRequest.Params.Maturity) * time.Second)
	borrow, err := p.r.CreateBorrow(&models.Borrow{
		LoanAmount:       int64(req.LoanRequest.LoanAmount),
		KeyDigest:        req.LoanRequest.KeyDigest,
		LoanID:           req.LoanRequest.LoanID,
		CollateralType:   req.LoanRequest.CollateralType,
		CollateralAmount: req.LoanRequest.CollateralAmount,
		StartDate:        startDate,
		EndDate:          endDate,
		InterestRate:     int64(req.LoanRequest.Params.InterestRate),
		Maturity:         int64(req.LoanRequest.Params.Maturity),
		LiquidationStart: int64(req.LoanRequest.Params.LiquidationStart),
		PaymentAddress:   req.LoanRequest.ReceiveAddress,
		State:            models.Pending,
	})
	if err != nil {
		return nil, errors.Wrap(err, "b.r.Create")
	}

	txID, err := p.bc.CreateAndSendLoanRequest(u.PrivKey, req.LoanRequest)
	if err != nil {
		return nil, err
	}
	borrow.ConstantLoanRequestTxID = *txID
	_, err = p.r.UpdateBorrow(borrow)
	if err != nil {
		return nil, err
	}

	return AssembleBorrow(borrow), nil
}

func (p *Portal) ListBorrowsByUser(paymentAddress string, state, limit, page string) ([]*serializers.BorrowResp, error) {
	l, pg, err := p.parseQuery(limit, page)
	if err != nil {
		return nil, errors.Wrapf(err, "b.parseQuery %s %s", limit, page)
	}

	var s *models.BorrowState
	if state != "" {
		st := models.GetBorrowState(state)
		if st == models.InvalidState {
			return nil, ErrInvalidBorrowState
		}
		s = &st
	}

	borrows, err := p.r.ListBorrowByUser(paymentAddress, s, l, pg)
	if err != nil {
		return nil, errors.Wrap(err, "b.r.ListByUser")
	}

	return p.transformToResp(borrows), nil
}

func (p *Portal) ListAllBorrows(state, limit, page string) ([]*serializers.BorrowResp, error) {
	l, pg, err := p.parseQuery(limit, page)
	if err != nil {
		return nil, errors.Wrapf(err, "b.parseQuery %s %s", limit, page)
	}

	var s *models.BorrowState
	if state != "" {
		st := models.GetBorrowState(state)
		if st == models.InvalidState {
			return nil, ErrInvalidBorrowState
		}
		s = &st
	}

	borrows, err := p.r.ListAllBorrows(s, l, pg)
	if err != nil {
		return nil, errors.Wrap(err, "b.r.ListByUser")
	}

	return p.transformToResp(borrows), nil
}

func (p *Portal) FindBorrowByID(idS string) (*models.Borrow, error) {
	id, err := strconv.Atoi(idS)
	if err != nil {
		return nil, errors.Wrapf(err, "strconv.Atoi %s", idS)
	}
	borrow, err := p.r.FindBorrowByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "b.r.FindByID")
	}
	if borrow == nil {
		return nil, ErrBorrowNotFound
	}
	return borrow, nil
}

func (p *Portal) parseQuery(limitS, pageS string) (limit, page int, err error) {
	page, err = strconv.Atoi(pageS)
	if err != nil {
		return 0, 0, ErrInvalidPage
	}
	limit, err = strconv.Atoi(limitS)
	if err != nil {
		return 0, 0, ErrInvalidLimit
	}
	return
}

func (p *Portal) transformToResp(bs []*models.Borrow) []*serializers.BorrowResp {
	resp := make([]*serializers.BorrowResp, 0, len(bs))
	for _, br := range bs {
		resp = append(resp, AssembleBorrow(br))
	}
	return resp
}

func AssembleBorrow(b *models.Borrow) *serializers.BorrowResp {
	return &serializers.BorrowResp{
		ID:                       b.ID,
		LoanAmount:               b.LoanAmount,
		LoanID:                   b.LoanID,
		KeyDigest:                b.KeyDigest,
		State:                    b.State.String(),
		StartDate:                b.StartDate.Format(common.DateTimeLayoutFormat),
		EndDate:                  b.EndDate.Format(common.DateTimeLayoutFormat),
		InterestRate:             b.InterestRate,
		CollateralType:           b.CollateralType,
		CollateralAmount:         b.CollateralAmount,
		CreatedAt:                b.CreatedAt.Format(common.DateTimeLayoutFormat),
		PaymentAddress:           b.PaymentAddress,
		LiquidationStart:         b.LiquidationStart,
		Maturity:                 b.Maturity,
		ConstantLoanPaymentTxID:  b.ConstantLoanPaymentTxID,
		ConstantLoanRequestTxID:  b.ConstantLoanRequestTxID,
		ConstantLoanAcceptTxID:   b.ConstantLoanAcceptTxID,
		ConstantLoanWithdrawTxID: b.ConstantLoanWithdrawTxID,
	}
}

func (p *Portal) UpdateStatusBorrowRequest(b *models.Borrow, action string, constantLoanTxId string) (bool, error) {
	switch action {
	case "r":
		{
			// TODO call web3 to eth to check
			b.State = models.Rejected
			_, err := p.r.UpdateBorrow(b)
			if err != nil {

				return false, err
			}
			return true, nil
		}
	case "a":
		{
			// TODO check with blockchain node to get tx
			b.State = models.Approved
			b.ConstantLoanAcceptTxID = constantLoanTxId
			_, err := p.r.UpdateBorrow(b)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	default:
		return false, nil
	}
}

func (p *Portal) PaymentTxForLoanRequestByID(b *models.Borrow, constantPaymentTxId string) (*blockchain.TransactionDetail, error) {
	var txPayment *blockchain.TransactionDetail
	retry := 10
	for true {
		var err error
		txPayment, err = p.bc.GetTxByHash(constantPaymentTxId)
		if err != nil {
			return txPayment, err
		}
		// retry 10 times = 30s
		time.Sleep((3 * time.Millisecond))
		retry --
		if retry == 0 {
			break
		}
	}
	if txPayment != nil {
		b.State = models.Payment
		b.ConstantLoanPaymentTxID = txPayment.Hash
		_, err := p.r.UpdateBorrow(b)
		if err != nil {
			return txPayment, err
		}
		return txPayment, nil
	}
	return txPayment, errors.New("Not found payment tx")
}

func (p *Portal) GetLoanParams() ([]interface{}, error) {
	return p.bc.GetLoanParams()
}
