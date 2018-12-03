package service

import (
	"encoding/json"

	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/ninjadotorg/constant-api-service/service/3rd/primetrust"
	thirdpartymodels "github.com/ninjadotorg/constant-api-service/service/primetrust/models"
)

type ReserveService struct {
	r *reserve.Reserve
	b *blockchain.Blockchain
}

func NewReserveService(r *reserve.ReserveDao, b *blockchain.Blockchain) *ReserveService {
	return &ReserveService{
		r: r,
		b: b,
	}
}

func (self *ReserveService) CreateContribution(u *models.User, req *serializers.ReserveContributionRequest) (*models.ReserveContributionRequest, error) {
	// 1. Validate ReserveContributionRequest in request
	if req.PaymentForm.ContactEmail == "" || req.PaymentForm.ContactName == "" {
		return nil, ErrInvalidArgument
	}

	switch req.PaymentForm.PaymentType {
	case thirdpartymodels.PaymentTypeAch:
		if req.PaymentForm.BankAccountName == "" || req.PaymentForm.BankAccountType == "" || req.PaymentForm.BankName == "" {
			return nil, ErrInvalidArgument
		}
	case thirdpartymodels.PaymentTypeCreditCard:
		if req.PaymentForm.CreditCardCvv == "" || req.PaymentForm.CreditCardExpirationDate == "" || req.PaymentForm.CreditCardNumber == "" {
			return nil, ErrInvalidArgument
		}
	}

	if req.PaymentAddress == "" {
		req.PaymentAddress = u.PaymentAddress
	}

	// 2. insert db ReserveContributionRequest
	rcr, err := self.r.CreateReserveContributionRequest(&models.ReserveContributionRequest{
		User:           u,
		PartyID:        req.PartyID,
		PaymentAddress: req.PaymentAddress,
	})

	if err != nil {
		return nil, err
	}

	// 3. insert db ReserveContributionRequestPaymentParty
	requestData, _ := json.Marshal(req.PaymentForm)
	rcrpp, err := self.r.CreateReserveContributionRequestPaymentParty(&models.ReserveContributionRequestPaymentParty{
		ReserveContributionRequest: rcr,
		RequestData:                requestData,
		Amount:                     req.Amount,
		Status:                     models.ReserveContributionRequestPaymentPartyStatus,
	})

	if err != nil {
		self.r.DeleteReserveContributionRequest(rcr)
		return nil, err
	}

	// 4. call related party: prime trust, eth ...
	switch rcr.PartyID {
	case models.ReservePartyPrimeTrust:
		contribution := thirdpartymodels.Contribution{
			Data: thirdpartymodels.ContributionData{
				Type: thirdpartymodels.ContributionType,
			},
		}

		contributionAttributes := thirdpartymodels.ContributionAttributes{
			AccountID:     "", //todo get from config
			Amount:        rcr.Amount,
			ContractEmail: user.Email,
			ContractName:  user.Name,
			PaymentMethod: thirdpartymodels.PaymentMethodAttributes{
				PaymentType:              req.PaymentForm.PaymentType,
				RoutingNumber:            req.PaymentForm.RoutingNumber,
				Last4:                    req.PaymentForm.Last4,
				AchCheckType:             req.PaymentForm.AchCheckType,
				BankAccountName:          req.PaymentForm.BankAccountName,
				BankAccountType:          req.PaymentForm.BankAccountType,
				BankName:                 req.PaymentForm.BankName,
				RoutingNumber:            req.PaymentForm.RoutingNumber,
				CreditCardCvv:            req.PaymentForm.CreditCardCvv,
				CreditCardExpirationDate: req.PaymentForm.CreditCardExpirationDate,
				CreditCardNumber:         req.PaymentForm.CreditCardNumber,
				CreditCardName:           req.PaymentForm.CreditCardName,
				CreditCardPostalCode:     req.PaymentForm.CreditCardPostalCode,
				CreditCardType:           req.PaymentForm.CreditCardType,
			},
		}

		contribution.Attributes = contributionAttributes

		extRequestData, _ := json.Marshal(contribution)

		rcrpp.ExtRequestData = extRequestData
		response, err := primetrust.CreateContribution(contribution)

		if err != nil {
			self.r.DeleteReserveContributionRequest(rcr)
			self.r.DeleteReserveContributionRequestPaymentParty(rcrpp)
			return nil, err
		}

		extResponseData, _ := json.Marshal(response)

		rcrpp.ExtID = response.Data.ID
	}

	return rcr
}

func (self *ReserveService) GetContribution(u *models.User, id int) (*models.ReserveContributionRequest, error) {
	rcr, err := self.r.FindReserveContributionRequestByID(id)

	if err != nil {
		return nil, err
	}

	if rcr.UserID != u.ID {
		return nil, nil
	}

	return rcr, nil
}

func (self *ReserveService) GetContributions(u *models.User) ([]*models.ReserveContributionRequest, error) {
	var (
		page  = c.DefaultQuery("page", "1")
		limit = c.DefaultQuery("limit", "10")
	)
	filter := map[string]interface{}{
		UserID: u.ID,
	}

	rcrs, err := self.r.FindAllReserveContributionRequest(&filter, page, limit)
	if err != nil {
		return nil, err
	}

	return rcrs, nil
}

func (self *ReserveService) CreateDisbursement(req *serializers.ReserveDisbursementRequest) {
	// 1. Validate ReserveDisbursementRequest in request
	// 2. insert db ReserveDisbursementRequest
	// 3. insert db ReserveDisbursementRequestPaymentParty
	// 6. call blockchain network to burn constant
	// 4. call related party: prime trust, eth ... and wait for data
	// 5. update data ReserveContributionRequestPaymentParty
}

func (self *ReserveService) GetDisbursement(u *models.User, id int) (*models.ReserveDisbursementRequest, error) {
	rdr, err := self.r.FindReserveDisbursementRequestByID(id)

	if err != nil {
		return nil, err
	}

	if rdr.UserID != u.ID {
		return nil, nil
	}

	return rdr, nil
}

func (self *ReserveService) GetDisbursements(u *models.User, filter *map[string]interface{}, page, limit int) ([]*models.ReserveDisbursementRequest, error) {
	filter.UserID = u.ID

	rdrs, err := self.r.FindAllReserveDisbursementRequest(&filter, page, limit)
	if err != nil {
		return nil, err
	}

	return rdrs, nil
}

func (self *ReserveService) PrimetrustWebHook(params interface{}) {
	/*
		{
		      "id": "151f0371-d16d-49b4-bc4c-c13788432a58",
		      "account_id": "1ae0e833-b07b-4d95-a32f-16c86bed539d",
		      "action": "update",
		      "data": {
		        "attributes": {
		          "status": "settled"
		        }
		      },
		      "resource_id": "fbd8bf30-552c-4d2d-b21c-c78dbc6b05d9",
		      "resource_type": "contributions"
		}
	*/
}
