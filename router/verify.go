package router

import (
	"github.com/SteinsElite/pickGinS/service/notification"
	"github.com/SteinsElite/pickGinS/service/vault"
	"github.com/SteinsElite/pickGinS/util"
)

func validQueryPhase(phase string) bool {
	if phase == vault.Week || phase == vault.Month || phase == vault.Year {
		return true
	}
	return false
}

func validTxTag(tag string) bool {
	if tag == "" ||
		tag == "deposit" ||
		tag == "profit" ||
		tag == "withdraw" {
		return true
	}
	return false
}

func validCoinSymbol(coin string) bool {
	if coin == util.MDX ||
		coin == util.BTC ||
		coin == util.ETH ||
		coin == util.USDT ||
		coin == util.HT {
		return true
	}
	return false
}

func validNotificationTag(tag string) bool {
	if tag == "" ||
		tag == notification.QuotaUpdate ||
		tag == notification.Activity ||
		tag == notification.Weekly {
		return true
	}
	return false
}