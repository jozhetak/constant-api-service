package api

import (
	"github.com/appleboy/gin-jwt"
)

func (s *Server) Routes(authMw *jwt.GinJWTMiddleware) {
	// auth API group
	auth := s.g.Group("/auth")
	auth.POST("/register", s.Register)
	auth.POST("/login", authMw.LoginHandler)
	auth.POST("/forgot", s.ForgotPassword)
	auth.POST("/reset", s.ResetPassword)

	// portal API group
	portal := s.g.Group("/portal")
	portal.GET("/borrows", s.ListBorrowsByUser)
	portal.GET("/borrows/:id", s.FindByID)
	portal.GET("/all_borrows", s.ListAllBorrows)
	portal.POST("/borrows/:id/process", s.ProcessStateBorrowByID)
	portal.Use(authMw.MiddlewareFunc())
	{
		portal.POST("/borrows", s.CreateNewBorrow)
		portal.POST("/borrows/:id/pay", s.PayBorrowByID)
		portal.POST("/borrows/:id/withdraw", s.WithdrawBorrowByID)
	}

	// exchange API group
	exch := s.g.Group("/exchange")
	exch.Any("/ws/trades", s.ExchangeWS)
	exch.GET("/markets", s.ListMarkets)
	exch.GET("/market_histories", s.MarketHistory)
	exch.GET("/symbol_rates", s.SymbolRates)
	exch.GET("/market_rates", s.MarketRates)
	exch.Use(authMw.MiddlewareFunc())
	{
		exch.POST("/orders", s.CreateOrder)
		exch.GET("/orders", s.UserOrderHistory)
	}

	// WalletService API Group
	wallet := s.g.Group("/wallet")
	wallet.GET("/accounts", s.ListAccounts)
	wallet.Use(authMw.MiddlewareFunc())
	{
		wallet.GET("/coinbalance", s.GetCoinBalance)
		wallet.GET("/balances", s.GetCoinAndCustomTokenBalance)
		wallet.POST("/send", s.SendCoin)
	}

	// voting API group
	voting := s.g.Group("/voting")
	voting.Use(authMw.MiddlewareFunc())
	{
		// candidate board
		voting.POST("/candidate", s.RegisterBoardCandidate)
		voting.GET("/candidates", s.GetCandidatesList)
		voting.POST("/candidate/vote", s.VoteCandidateBoard)
		// Proposal
		voting.POST("/proposal", s.CreateProposal)
		voting.GET("/proposals", s.GetProposalsList)
		voting.GET("/proposal", s.GetProposal)
		voting.POST("/proposal/vote", s.VoteProposal)
	}

	// reserve API group
	reserve := s.g.Group("/reserve")
	reserve.GET("/primetrust", s.PrimetrustWebHook)
	voting.Use(authMw.MiddlewareFunc())
	{
		reserve.GET("/getreserveparty", s.GetReserveParty)
		reserve.POST("/request/:party", s.RequestReserve)
		reserve.GET("/request/:party", s.RequestReserveHistory)
		reserve.GET("/request/:party/:requestId", s.GetRequestReserve)
		reserve.POST("/request-return/:party", s.ReturnRequestReserve)
		reserve.GET("/request-return/:party", s.ReturnRequestReserveHistory)
		reserve.GET("/request-return/:party/:requestId", s.GetReturnRequestReserve)

	}

	// common API
	common := s.g.Group("/common")
	common.GET("/loanparams", s.GetLoanParams)
	common.GET("/bondtypes", s.GetBondTypes)
}
