package router

import (
	"github.com/SteinsElite/pickGinS/util"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/SteinsElite/pickGinS/service/coin"
	"github.com/SteinsElite/pickGinS/service/transaction"
	"github.com/SteinsElite/pickGinS/service/vault"
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

// GetTransaction godoc
// @Summary get the transaction info
// @Produce  json
// @Success 200 "the transaction of the page"
// @Failure 400 "Invalid params"
// @Failure 500 "Server error"
// @Param address path string true "user account address"
// @Param tag query string false "tag of the transaction-{deposit,withdraw,claimProfit}, if not specify, get all the category"
// @Param page query int true "index of page"
// @Param page_size query int true "size of each page"
// @Router /api/v1/transaction/{address} [get]
func GetTransaction(c *gin.Context) {
	userAddr := c.Param("address")
	tag := c.Query("tag")
	if !validTxTag(tag) {
		c.JSON(400, gin.H{
			"err": "Invalid params",
			"msg": "tags should be one of {deposit, withdraw, claimProfit}",
		})
		return
	}
	if c.Query("page") == "" || c.Query("page_size") == "" {
		c.JSON(400, gin.H{
			"err": "Missing params",
			"msg": "specify the page & page_size",
		})
		return
	}
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Query("page_size"), 10, 64)

	res, err := transaction.LoadTxFromDb(page, pageSize, tag, userAddr)
	if err != nil {
		c.JSON(500, gin.H{
			"err": "internal error in server",
			"msg": err,
		})
	} else {
		c.JSON(200, gin.H{
			"transaction": res,
		})
	}

}

//func GetAnn(c *gin.Context) {
//
//}

//func GetSpecificAnn(c *gin.Context) {
//
//}

// GetVolume godoc
// @Summary get the total volume info
// @Produce  json
// @Success 200 "the [(timestamp,volume)] in the time range"
// @Param range query string true "the duration to query-{7D,1M,1Y}"
// @Router /api/v1/chart/volume [get]
func GetVolume(c *gin.Context) {
	phase := c.Query("range")
	if !validQueryPhase(phase) {
		c.JSON(400, gin.H{
			"error":   "Invalid params",
			"message": "range should be one of {7D,1M,1Y}",
		})
	}
	values := vault.PhasedVolume(phase)
	c.JSON(200, gin.H{
		"points": values,
	})
}

// GetProfit godoc
// @Summary get the phased profit info
// @Produce  json
// @Success 200 "the [(timestamp,profit)] in the time range"
// @Param range query string true "the duration to query-{7D,1M,1Y}"
// @Router /api/v1/chart/profit [get]
func GetProfit(c *gin.Context) {
	phase := c.Query("range")
	if !validQueryPhase(phase) {
		c.JSON(400, gin.H{
			"error":   "Invalid params",
			"message": "range should be one of {7D,1M,1Y}",
		})
	}
	values := vault.PhasedProfit(phase)
	c.JSON(200, gin.H{
		"points": values,
	})
}

// GetRatio godoc
// @Summary get the ratio info
// @Produce  json
// @Success 200 "amount of each asset in usd"
// @Router /api/v1/chart/ratio [get]
func GetRatio(c *gin.Context) {
	values := vault.AssetRatio()
	c.JSON(200, gin.H{
		"ratio": values,
	})
}

// GetCoinPriceInfo godoc
//@Summary get the Coin Price info and trend
//@Produce json
//@Param coin path string true "{BTC,ETH,USDT,HT,MDX}"
//@Success 200 "the price trend of coin, {"rate": ..., "trend": ...}"
//@Router /api/v1/price_info/{coin} [get]
func GetCoinPriceInfo(c *gin.Context) {
	coinSymbol := c.Param("coin")
	if !validCoinSymbol(coinSymbol){
		c.JSON(400, gin.H{
			"error": "invalid coin symbol",
			"message": "should be one of {BTC,ETH,USDT,HT,MDX}",
		})
		return
	}
	c.JSON(200, coin.GetCoinTrend(coinSymbol))
}
