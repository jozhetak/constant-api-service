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
	portal.Use(authMw.MiddlewareFunc())
	{
		portal.POST("/borrows", s.CreateNewBorrow)
		portal.GET("/borrows/:id", s.FindByID)
		portal.GET("/borrows", s.ListBorrowsByUser)
		portal.GET("/all_borrows", s.ListAllBorrows)
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
}
