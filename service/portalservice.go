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
	//_ "github.com/ninjadotorg/constant-api-service/service/3rd/ethereum"
)

type PortalService struct {
	portalDao         *portal.PortalDao
	blockchainService *blockchain.Blockchain
}

func NewPortal(r *portal.PortalDao, bc *blockchain.Blockchain) *PortalService {
	return &PortalService{
		portalDao:         r,
		blockchainService: bc,
	}
}

func (p *PortalService) CreateBorrow(u *models.User, req serializers.BorrowReq) (*serializers.BorrowResp, error) {
	startDate, err := time.Parse(common.DateTimeLayoutFormatIn, req.StartDate)
	if err != nil {
		return nil, errors.Wrap(err, "b.portalDao.Create")
	}
	if req.LoanRequest.ReceiveAddress == "" {
		req.LoanRequest.ReceiveAddress = u.PaymentAddress
	}
	endDate := startDate.Add(time.Duration(req.LoanRequest.Params.Maturity) * time.Second)
	if req.LoanRequest.ReceiveAddress == "" {
		req.LoanRequest.ReceiveAddress = u.PaymentAddress
	}
	borrow, err := p.portalDao.CreateBorrow(&models.Borrow{
		User:             u,
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
		return nil, errors.Wrap(err, "b.portalDao.Create")
	}

	txID, err := p.blockchainService.CreateAndSendLoanRequest(u.PrivKey, req.LoanRequest)
	if err != nil && false { // TODO
		err1 := p.portalDao.DeleteBorrow(borrow)
		if err1 != nil {
			return nil, err1
		}
		return nil, err
	}
	if txID != nil {
		borrow.ConstantLoanRequestTxID = *txID
		_, err = p.portalDao.UpdateBorrow(borrow)
		if err != nil {
			return nil, err
		}
	}

	return AssembleBorrow(borrow), nil
}

func (p *PortalService) ListBorrowsByUser(paymentAddress string, state, limit, page string) ([]*serializers.BorrowResp, error) {
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

	borrows, err := p.portalDao.ListBorrowByUser(paymentAddress, s, l, pg)
	if err != nil {
		return nil, errors.Wrap(err, "b.portalDao.ListByUser")
	}

	return p.transformToResp(borrows), nil
}

func (p *PortalService) ListAllBorrows(state, limit, page string) ([]*serializers.BorrowResp, error) {
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

	borrows, err := p.portalDao.ListAllBorrows(s, l, pg)
	if err != nil {
		return nil, errors.Wrap(err, "b.portalDao.ListByUser")
	}

	return p.transformToResp(borrows), nil
}

func (p *PortalService) FindBorrowByID(idS string) (*models.Borrow, error) {
	id, err := strconv.Atoi(idS)
	if err != nil {
		return nil, errors.Wrapf(err, "strconv.Atoi %s", idS)
	}
	borrow, err := p.portalDao.FindBorrowByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "b.portalDao.FindByID")
	}
	if borrow == nil {
		return nil, ErrBorrowNotFound
	}
	return borrow, nil
}

func (p *PortalService) parseQuery(limitS, pageS string) (limit, page int, err error) {
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

func (p *PortalService) transformToResp(bs []*models.Borrow) []*serializers.BorrowResp {
	resp := make([]*serializers.BorrowResp, 0, len(bs))
	for _, br := range bs {
		resp = append(resp, AssembleBorrow(br))
	}
	return resp
}

func AssembleBorrow(b *models.Borrow) *serializers.BorrowResp {
	result := &serializers.BorrowResp{
		ID:                       b.ID,
		LoanAmount:               b.LoanAmount,
		LoanID:                   b.LoanID,
		KeyDigest:                b.KeyDigest,
		State:                    b.State.String(),
		StartDate:                b.StartDate.Format(common.DateTimeLayoutFormatOut),
		EndDate:                  b.EndDate.Format(common.DateTimeLayoutFormatOut),
		InterestRate:             b.InterestRate,
		CollateralType:           b.CollateralType,
		CollateralAmount:         b.CollateralAmount,
		CreatedAt:                b.CreatedAt.Format(common.DateTimeLayoutFormatOut),
		PaymentAddress:           b.PaymentAddress,
		LiquidationStart:         b.LiquidationStart,
		Maturity:                 b.Maturity,
		ConstantLoanPaymentTxID:  b.ConstantLoanPaymentTxID,
		ConstantLoanRequestTxID:  b.ConstantLoanRequestTxID,
		ConstantLoanWithdrawTxID: b.ConstantLoanWithdrawTxID,
		ConstantLoanAcceptTxID:   []string{},
	}

	for _, temp := range b.BorrowResponses {
		result.ConstantLoanAcceptTxID = append(result.ConstantLoanAcceptTxID, temp.ConstantLoanResponseTxID)
	}

	return result
}

func (p *PortalService) UpdateStatusBorrowRequest(b *models.Borrow, action string, constantLoanTxId string) (bool, error) {
	switch action {
	case "portalDao": // reject
		{
			b.State = models.Rejected
			_, err := p.portalDao.UpdateBorrow(b)
			if err != nil {

				return false, err
			}
			// TODO call web3 to eth to check
			// reject loan
			//
			return true, nil
		}
	case "a": // accept
		{
			// call to check tx in constant network
			tx, err := GetBlockchainTxByHash(constantLoanTxId, 10, p.blockchainService)
			if err != nil {
				return false, err
			}
			if tx == nil {
				return false, err
			}

			enoughAccept := false
			// TODO 0xsirrush call block chain to check enough accept
			borrowResponse := models.BorrowResponse{
				ConstantLoanResponseTxID: constantLoanTxId,
				Borrow:                   *b,
			}
			_, err = p.portalDao.CreateBorrowResponse(&borrowResponse)
			if err != nil {
				return false, err
			}
			if enoughAccept {
				b.State = models.Approved
			}
			_, err = p.portalDao.UpdateBorrow(b)
			if err != nil {
				return false, err
			}

			return true, nil
		}
	default:
		return false, nil
	}
}

func (p *PortalService) WithdrawTxForLoanRequest(u *models.User, b *models.Borrow, key string) (*blockchain.TransactionDetail, error) {
	request := serializers.LoanWithdraw{
		LoanID: b.LoanID,
		Key:    key,
	}
	// call constant network to create loan withdraw
	txId, err := p.blockchainService.CreateAndSendLoanWithdraw(u.PrivKey, request)
	if err != nil {
		return nil, err
	}
	if txId != nil {
		tx, err := GetBlockchainTxByHash(*txId, 10, p.blockchainService)
		if err != nil {
			return nil, err
		}
		// update db
		b.ConstantLoanWithdrawTxID = tx.Hash
		_, err = p.portalDao.UpdateBorrow(b)
		if err != nil {
			return nil, err
		}

		if b.State == models.Approved {
			switch b.CollateralType {
			case "ETH":
				// TODO call web3 to process collateral
				// accept loan
				//
			}
		}
		return tx, nil
	} else {
		return nil, errors.New("Fail")
	}
}

func (p *PortalService) PaymentTxForLoanRequest(u *models.User, b *models.Borrow, constantPaymentTxId string) (*blockchain.TransactionDetail, error) {
	request := serializers.LoanPayment{
		LoanID: b.LoanID,
	}
	// call constant network to create loan withdraw
	txId, err := p.blockchainService.CreateAndSendLoanPayment(u.PrivKey, request)
	if err != nil {
		return nil, err
	}
	if txId != nil {
		tx, err := GetBlockchainTxByHash(*txId, 10, p.blockchainService)
		if err != nil {
			return nil, err
		}

		// update db
		b.ConstantLoanPaymentTxID = tx.Hash
		_, err = p.portalDao.UpdateBorrow(b)
		if err != nil {
			return nil, err
		}

		if true {
			switch b.CollateralType {
			case "ETH":
				// TODO call web3 to process collateral
				// call addpayment
				//

			}
		}

		return tx, nil
	} else {
		return nil, errors.New("Fail")
	}
}

func (p *PortalService) GetLoanParams() ([]interface{}, error) {
	return p.blockchainService.GetLoanParams()
}
