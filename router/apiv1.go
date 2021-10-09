package router

import (
	"github.com/SteinsElite/pickGinS/service/notification"
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

func validNotificationTag(tag string) bool {
	if tag == "" ||
		tag == notification.QuotaUpdate||
		tag == notification.Activity ||
		tag == notification.Weekly {
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
	if !validCoinSymbol(coinSymbol) {
		c.JSON(400, gin.H{
			"error":   "invalid coin symbol",
			"message": "should be one of {BTC,ETH,USDT,HT,MDX}",
		})
	}
	c.JSON(200, coin.GetCoinTrend(coinSymbol))
}

// GetNotification godoc
// @summary get notification info
// @description obtains the specific notification by the tag
// @produce json
// @param tag query string false "tag of the notification-{QuotaUpdate,Activity,Weekly}, if not specify, get all the category"
// @param page query int true "index of page"
// @param page_size query int true "size of each page"
// @success 200 {array} notification.Notification
// @failure 400 {json}
// router /api/v1/notification [get]
func GetNotification(c *gin.Context){
	tag := c.Query("tag")
	if !validNotificationTag(tag) {
		c.JSON(400, gin.H{
			"err": "Invalid params",
			"msg": "tags should be one of {QuotaUpdate,Activity,Weekly}",
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

	res := notification.GetNotification(tag,page, pageSize)
	c.JSON(200, res)
}

func GetWordHash(c *gin.Context) {
	addr := c.Query("address")
	if !util.IsValidAddress(addr){
		c.JSON(400, gin.H{
			"error": "Invalid params",
			"message": "address is not valid ethereum address",
		})
		return
	}
	word := getAuthWord(addr)
	if word == nil {
		c.JSON(3001, gin.H{
			"error": "Fail to getAuthWord",
			"message": "the address is not register as admin",
		})
		return
	}
	c.JSON(200, word)
}

func RegisterPublisher(c *gin.Context){
	adminAddr := c.Param("address")
	sig := c.Query("sig")
	addr := c.Query("addr")

	_ = addr
	if !IsAuth(adminAddr, sig){
		c.JSON(3001, gin.H{
			"error":   "not permission",
			"message": "current address is not able to register new publisher",
		})
	}



}