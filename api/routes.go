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
		portal.GET("/loanparams", s.GetLoanParams)
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

	// Wallet API Group
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
		voting.POST("/register", s.RegisterBoardCandidate)
	}
}
