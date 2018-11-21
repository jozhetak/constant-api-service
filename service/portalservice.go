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
)

type Portal struct {
	r  *portal.Portal
	bc *blockchain.Blockchain
}

func NewPortal(r *portal.Portal, bc *blockchain.Blockchain) *Portal {
	return &Portal{
		r:  r,
		bc: bc,
	}
}

func (p *Portal) CreateBorrow(req serializers.BorrowReq) (*serializers.BorrowResp, error) {
	// if u.Type != models.Borrower {
	//         return nil, errors.New("user type must be borrower to create borrow")
	// }
	endDate, err := time.Parse(common.DateTimeLayoutFormat, req.EndDate)
	if err != nil {
		return nil, errors.Wrap(err, "b.r.Create")
	}
	startDate, err := time.Parse(common.DateTimeLayoutFormat, req.StartDate)
	if err != nil {
		return nil, errors.Wrap(err, "b.r.Create")
	}
	borrow, err := p.r.CreateBorrow(&models.Borrow{
		Amount:         req.Amount,
		Hash:           req.HashKey,
		CollateralTxID: req.CollateralTxID,
		State:          models.Pending,
		Collateral:     req.Collateral,
		StartDate:      startDate,
		EndDate:        endDate,
		Rate:           req.Rate,
		PaymentAddress: req.PaymentAddress,
	})
	if err != nil {
		return nil, errors.Wrap(err, "b.r.Create")
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

func (p *Portal) Approve(idS string) {
	// call to bccore
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
		ID:             b.ID,
		Amount:         b.Amount,
		Hash:           b.Hash,
		CollateralTxID: b.CollateralTxID,
		State:          b.State.String(),
		StartDate:      b.StartDate.Format(common.DateTimeLayoutFormat),
		EndDate:        b.EndDate.Format(common.DateTimeLayoutFormat),
		Rate:           b.Rate,
		Collateral:     b.Collateral,
		CreatedAt:      b.CreatedAt.Format(time.RFC3339),
		PaymentAdrress: b.PaymentAddress,
	}
}

func (p *Portal) UpdateStatusBorrowRequest(b *models.Borrow, action string, constantTxId string) (bool, error) {
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
			b.ConstantTxID = constantTxId
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
