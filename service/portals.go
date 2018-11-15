package service

import (
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/dao/portal"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
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

func (p *Portal) CreateBorrow(u *models.User, amount int64, hash, txID, paymentAddr string) (*serializers.BorrowResp, error) {
	if u.Type != models.Borrower {
		return nil, errors.New("user type must be borrower to create borrow")
	}
	borrow, err := p.r.CreateBorrow(&models.Borrow{
		User:           u,
		Amount:         amount,
		Hash:           hash,
		TxID:           txID,
		PaymentAddress: paymentAddr,
		State:          models.Pending,
	})
	if err != nil {
		return nil, errors.Wrap(err, "b.r.Create")
	}
	return assembleBorrow(borrow), nil
}

func (p *Portal) ListBorrowsByUser(user *models.User, state, limit, page string) ([]*serializers.BorrowResp, error) {
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

	borrows, err := p.r.ListBorrowByUser(user, s, l, pg)
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

func (p *Portal) FindBorrowByID(idS string) (*serializers.BorrowResp, error) {
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
	return assembleBorrow(borrow), nil
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
		resp = append(resp, assembleBorrow(br))
	}
	return resp
}

func assembleBorrow(b *models.Borrow) *serializers.BorrowResp {
	return &serializers.BorrowResp{
		ID:             b.ID,
		Amount:         b.Amount,
		Hash:           b.Hash,
		TxID:           b.TxID,
		PaymentAddress: b.PaymentAddress,
		State:          b.State.String(),
		CreatedAt:      b.CreatedAt.Format(time.RFC3339),
	}
}
