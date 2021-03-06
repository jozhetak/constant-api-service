package service

import (
	"encoding/json"
	"fmt"

	"github.com/ninjadotorg/constant-api-service/dao/reserve"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/ninjadotorg/constant-api-service/service/3rd/primetrust"
	thirdpartymodels "github.com/ninjadotorg/constant-api-service/service/3rd/primetrust/models"
)

type ReserveService struct {
	r *reserve.Reserve
	b *blockchain.Blockchain
	p *primetrust.Primetrust
}

func NewReserveService(r *reserve.Reserve, b *blockchain.Blockchain, p *primetrust.Primetrust) *ReserveService {
	return &ReserveService{
		r: r,
		b: b,
		p: p,
	}
}

func (self *ReserveService) CreateContribution(u *models.User, req *serializers.ReserveContributionRequest) (*models.ReserveContributionRequest, error) {
	// 1. Validate ReserveContributionRequest in request
	if req.PaymentForm.ContactEmail == "" || req.PaymentForm.ContactName == "" {
		return nil, ErrInvalidArgument
	}

	if req.PaymentForm.BankAccountName == "" || req.PaymentForm.BankAccountType == "" || req.PaymentForm.BankName == "" {
		return nil, ErrInvalidArgument
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
	requestData, _ := json.Marshal(req)
	rcrpp, err := self.r.CreateReserveContributionRequestPaymentParty(&models.ReserveContributionRequestPaymentParty{
		ReserveContributionRequest: rcr,
		RequestData:                string(requestData),
		Amount:                     req.Amount,
		Status:                     models.ReserveContributionRequestPaymentPartyStatusPending,
	})

	if err != nil {
		delErr := self.r.DeleteReserveContributionRequest(rcr)
		if delErr != nil {
			fmt.Println("ReserveService Delete Error", delErr.Error())
		}

		return nil, err
	}

	// 4. call related party: prime trust, eth ...
	switch rcr.PartyID {
	case models.ReservePartyPrimeTrust:
		contributionData := thirdpartymodels.ContributionData{
			Type: thirdpartymodels.ContributionType,
		}

		contributionAttributes := thirdpartymodels.ContributionAttributes{
			AccountID: self.p.AccountID,
			Amount:    rcrpp.Amount,
		}

		fundsTransferMethod := thirdpartymodels.FundsTransferMethodAttributes{
			FundsTransferType: thirdpartymodels.FundsTransferMethodFundsTransferTypeAch,
			ContactEmail:      u.Email,
			ContactName:       fmt.Sprintf("%s %s", u.FirstName, u.LastName),
			AchCheckType:      string(req.PaymentForm.AchCheckType),
			BankAccountName:   req.PaymentForm.BankAccountName,
			BankAccountType:   req.PaymentForm.BankAccountType,
			BankAccountNumber: req.PaymentForm.BankAccountNumber,
			BankName:          req.PaymentForm.BankName,
			RoutingNumber:     req.PaymentForm.RoutingNumber,
		}

		contributionAttributes.FundsTransferMethod = &fundsTransferMethod

		contributionData.Attributes = &contributionAttributes

		contribution := thirdpartymodels.Contribution{
			Data: &contributionData,
		}

		extRequestData, _ := json.Marshal(contribution)

		rcrpp.ExtRequestData = string(extRequestData)
		response, err := self.p.CreateContribution(&contribution)

		if err != nil {
			contributionBytes, _ := json.Marshal(contribution)
			fmt.Println("Call API error", string(contributionBytes))
			delErr := self.r.DeleteReserveContributionRequestPaymentParty(rcrpp)
			if delErr != nil {
				fmt.Println("ReserveService Delete Error", delErr.Error())
			}
			delErr2 := self.r.DeleteReserveContributionRequest(rcr)
			if delErr2 != nil {
				fmt.Println("ReserveService Delete Error", delErr2.Error())
			}
			return nil, err
		}

		extResponseData, _ := json.Marshal(response)
		rcrpp.ExtResponseData = string(extResponseData)
		fmt.Println("dump response", string(extResponseData))

		rcrpp.ExtID = response.Data.ID

		_, err = self.r.UpdateReserveContributionRequestPaymentParty(rcrpp)
		if err != nil {
			delErr := self.r.DeleteReserveContributionRequestPaymentParty(rcrpp)
			if delErr != nil {
				fmt.Println("ReserveService Delete Error", delErr.Error())
			}
			delErr2 := self.r.DeleteReserveContributionRequest(rcr)
			if delErr2 != nil {
				fmt.Println("ReserveService Delete Error", delErr2.Error())
			}
			return nil, err
		}
	}

	return rcr, nil
}

func (self *ReserveService) GetContribution(id int) (*models.ReserveContributionRequest, error) {
	rcr, err := self.r.FindReserveContributionRequestByID(id)

	if err != nil {
		return nil, err
	}

	return rcr, nil
}

func (self *ReserveService) GetContributions(filter *map[string]interface{}, page, limit int) ([]*models.ReserveContributionRequest, error) {
	rcrs, err := self.r.FindAllReserveContributionRequest(filter, page, limit)
	if err != nil {
		return nil, err
	}

	return rcrs, nil
}

func (self *ReserveService) CreateDisbursement(u *models.User, req *serializers.ReserveDisbursementRequest) (*models.ReserveDisbursementRequest, error) {
	// 1. Validate ReserveDisbursementRequest in request
	if req.PaymentForm.ContactEmail == "" || req.PaymentForm.ContactName == "" {
		return nil, ErrInvalidArgument
	}

	if req.PaymentForm.BankAccountName == "" || req.PaymentForm.BankAccountType == "" || req.PaymentForm.BankName == "" {
		return nil, ErrInvalidArgument
	}

	// 2. insert db ReserveDisbursementRequest
	rdr, err := self.r.CreateReserveDisbursementRequest(&models.ReserveDisbursementRequest{
		User:    u,
		PartyID: req.PartyID,
		Status:  models.ReserveDisbursementRequestStatusPending,
	})

	if err != nil {
		return nil, err
	}

	// 3. insert db ReserveDisbursementRequestPaymentParty
	requestData, _ := json.Marshal(req.PaymentForm)
	rdrpp, err := self.r.CreateReserveDisbursementRequestPaymentParty(&models.ReserveDisbursementRequestPaymentParty{
		ReserveDisbursementRequest: rdr,
		RequestData:                string(requestData),
		Amount:                     req.Amount,
		Status:                     models.ReserveDisbursementRequestPaymentPartyStatusPending,
	})

	if err != nil {
		delErr := self.r.DeleteReserveDisbursementRequest(rdr)
		if delErr != nil {
			fmt.Println("ReserveService Delete Error", delErr.Error())
		}
		return nil, err
	}

	// 3. TODO call blockchain network to burn constant & update tx id

	// 4. call related party: prime trust, eth ... and wait for data
	switch rdr.PartyID {
	case models.ReservePartyPrimeTrust:
		disbursementData := thirdpartymodels.DisbursementData{
			Type: thirdpartymodels.DisbursementType,
		}

		disbursementAttributes := thirdpartymodels.DisbursementAttributes{
			AccountID: self.p.AccountID,
			Amount:    rdrpp.Amount,
		}

		fundsTransferMethod := thirdpartymodels.FundsTransferMethodAttributes{
			FundsTransferType: thirdpartymodels.FundsTransferMethodFundsTransferTypeAch,
			ContactEmail:      u.Email,
			ContactName:       fmt.Sprintf("%s %s", u.FirstName, u.LastName),
			AchCheckType:      string(req.PaymentForm.AchCheckType),
			BankAccountName:   req.PaymentForm.BankAccountName,
			BankAccountType:   req.PaymentForm.BankAccountType,
			BankAccountNumber: req.PaymentForm.BankAccountNumber,
			BankName:          req.PaymentForm.BankName,
			RoutingNumber:     req.PaymentForm.RoutingNumber,
		}

		disbursementAttributes.FundsTransferMethod = &fundsTransferMethod

		disbursementData.Attributes = &disbursementAttributes

		disbursement := thirdpartymodels.Disbursement{
			Data: &disbursementData,
		}

		extRequestData, _ := json.Marshal(disbursement)

		rdrpp.ExtRequestData = string(extRequestData)
		response, err := self.p.CreateDisbursement(&disbursement)

		if err != nil {
			delErr := self.r.DeleteReserveDisbursementRequestPaymentParty(rdrpp)
			if delErr != nil {
				fmt.Println("ReserveService Delete Error", delErr.Error())
			}
			delErr2 := self.r.DeleteReserveDisbursementRequest(rdr)
			if delErr2 != nil {
				fmt.Println("ReserveService Delete Error", delErr2.Error())
			}
			return nil, err
		}

		extResponseData, _ := json.Marshal(response)
		rdrpp.ExtResponseData = string(extResponseData)

		rdrpp.ExtID = response.Data.ID

		_, err = self.r.UpdateReserveDisbursementRequestPaymentParty(rdrpp)
		if err != nil {
			delErr := self.r.DeleteReserveDisbursementRequestPaymentParty(rdrpp)
			if delErr != nil {
				fmt.Println("ReserveService Delete Error", delErr.Error())
			}
			delErr2 := self.r.DeleteReserveDisbursementRequest(rdr)
			if delErr2 != nil {
				fmt.Println("ReserveService Delete Error", delErr2.Error())
			}
			return nil, err
		}
	}

	return rdr, nil
}

func (self *ReserveService) GetDisbursement(id int) (*models.ReserveDisbursementRequest, error) {
	rdr, err := self.r.FindReserveDisbursementRequestByID(id)

	if err != nil {
		return nil, err
	}

	return rdr, nil
}

func (self *ReserveService) GetDisbursements(filter *map[string]interface{}, page, limit int) ([]*models.ReserveDisbursementRequest, error) {
	rdrs, err := self.r.FindAllReserveDisbursementRequest(filter, page, limit)
	if err != nil {
		return nil, err
	}

	return rdrs, nil
}

func (self *ReserveService) PrimetrustWebHook(req *serializers.PrimetrustChangedRequest) error {
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
	if data, ok := req.Data["attributes"]; ok {
		status, exist := data.(map[string]interface{})["status"]
		if exist {
			if req.ResourceType == thirdpartymodels.ContributionType {
				status := models.GetContributionPaymentPartyState(status.(string))
				if status != models.ReserveContributionRequestPaymentPartyStatusInvalid {
					rcrpp, err := self.r.FindReserveContributionRequestPaymentPartyByExtID(req.ResourceID)
					if err != nil {
						return err
					}

					rcrpp.Status = status
					_, err = self.r.UpdateReserveContributionRequestPaymentParty(rcrpp)

					if err != nil {
						return err
					}

					switch rcrpp.Status {
					case models.ReserveContributionRequestPaymentPartyStatusSettled:
						// TODO call blockchain send coin & update tx id

						// update reserve contribution status
						rcr, err := self.r.FindReserveContributionRequestByID(rcrpp.ReserveContributionRequestID)
						if err != nil {
							return err
						}

						rcr.Status = models.ReserveContributionRequestStatusFilled
						_, err = self.r.UpdateReserveContributionRequest(rcr)
						if err != nil {
							return err
						}
					}
				}
			}
			/*
				else if req.ResourceType == thirdpartymodels.DisbursementType {
					status := models.GetDisbursementPaymentPartyState(status.(string))
					if status != models.ReserveDisbursementRequestPaymentPartyStatusInvalid {
						rdrpp, err := self.r.FindReserveDisbursementRequestPaymentPartyByExtID(req.ResourceID)
						if err != nil {
							return err
						}

						rdrpp.Status = status
						_, err = self.r.UpdateReserveDisbursementRequestPaymentParty(rdrpp)

						if err != nil {
							return err
						}

						switch rdrpp.Status {
						case models.ReserveDisbursementRequestPaymentPartyStatusSettled:
							// todo call blockchain burn coin

							// update reserve contribution status
							rdr, err := self.r.FindReserveDisbursementRequestByID(rdrpp.ReserveDisbursementRequestID)
							if err != nil {
								return err
							}

							rdr.Status = models.ReserveDisbursementRequestStatusFilled
							_, err = self.r.UpdateReserveDisbursementRequest(rdr)
							if err != nil {
								return err
							}
						}
					}
				}
			*/
		}
	}
	return nil
}
