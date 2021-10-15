package router

import (
	"strconv"
	"time"

	"github.com/SteinsElite/pickGinS/internal/auth"
	"github.com/gin-gonic/gin"

	"github.com/SteinsElite/pickGinS/service/coin"
	"github.com/SteinsElite/pickGinS/service/notification"
	"github.com/SteinsElite/pickGinS/service/transaction"
	"github.com/SteinsElite/pickGinS/service/vault"
)

// GetTransaction godoc
// @summary get the transaction info
// @description gets the history transaction of specific account
// @produce  json
// @success 200 "a page of transaction and the total transaction amount-{"transaction":...,
// "count": ...}"
// @failure 400 "invalid params"
// @failure 500 "server error"
// @param address path string true "user account address"
// @param tag query string false "tag of the transaction-{deposit,withdraw,claimProfit}, if not specify, get all the category"
// @param page query int true "index of page"
// @param page_size query int true "size of each page"
// @router /transaction/{address} [get]
func GetTransaction(c *gin.Context) {
	userAddr := c.Param("address")
	tag := c.Query("tag")
	if !validTxTag(tag) {
		c.JSON(400, gin.H{
			"error":   "invalid params",
			"message": "tags should be one of {deposit, withdraw, claimProfit}",
		})
		return
	}
	if c.Query("page") == "" || c.Query("page_size") == "" {
		c.JSON(400, gin.H{
			"error":   "missing params",
			"message": "specify the page & page_size",
		})
		return
	}
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Query("page_size"), 10, 64)

	res, count, err := transaction.LoadTxFromDb(page, pageSize, tag, userAddr)
	if err != nil {
		c.JSON(500, gin.H{
			"err": "internal error in server",
			"msg": err,
		})
	} else {
		c.JSON(200, gin.H{
			"transaction": res,
			"count":       count,
		})
	}

}

// GetVolume godoc
// @Summary get the total volume info
// @Description gets the volume of each asset in usd, and the start of the query timestamp
// @Produce  json
// @Success 200 "the {startTime: ..., volume: ...} in the time range"
// @Param range query string true "the duration to query-{7D,1M,1Y}"
// @Router /chart/volume [get]
func GetVolume(c *gin.Context) {
	phase := c.Query("range")
	if !validQueryPhase(phase) {
		c.JSON(400, gin.H{
			"error":   "invalid params",
			"message": "range should be one of {7D,1M,1Y}",
		})
	}
	values, startTime, err := vault.PhasedVolume(phase)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"startTime": startTime,
		"volume":    values,
	})
}

// GetProfit godoc
// @Summary get the phased profit info
// @Description gets the profit in a time range
// @Produce  json
// @Success 200 "the [(timestamp,profit)] in the time range"
// @Param range query string true "the duration to query-{7D,1M,1Y}"
// @Router /chart/profit [get]
func GetProfit(c *gin.Context) {
	phase := c.Query("range")
	if !validQueryPhase(phase) {
		c.JSON(400, gin.H{
			"error":   "invalid params",
			"message": "range should be one of {7D,1M,1Y}",
		})
	}
	values, err := vault.PhasedProfit(phase)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"points": values,
	})
}

// GetRatio godoc
// @Summary get the ratio info
// @Description gets the ratio of each asset info
// @Produce  json
// @Success 200 "amount of each asset in usd"
// @Router /chart/ratio [get]
func GetRatio(c *gin.Context) {
	values := vault.AssetRatio()
	c.JSON(200, gin.H{
		"ratio": values,
	})
}

// GetCoinPriceInfo godoc
// @Summary get the Coin Price info and trend
// @Description gets the Coin Price info and trend
// @Produce json
// @Param coin path string true "{BTC,ETH,USDT,HT,MDX}"
// @Success 200 "the price trend of coin, {trend_info:{"rate": ..., "trend": ...}}"
// @Router /price_info/{coin} [get]
func GetCoinPriceInfo(c *gin.Context) {
	coinSymbol := c.Param("coin")
	if !validCoinSymbol(coinSymbol) {
		c.JSON(400, gin.H{
			"error":   "invalid coin symbol",
			"message": "should be one of {BTC,ETH,USDT,HT,MDX}",
		})
		return
	}
	trend := coin.GetCoinTrend(coinSymbol)
	c.JSON(200, gin.H{
		"trend_info": trend,
	})
}

// GetNotification godoc
// @summary get notification info
// @description obtains the specific notification by the tag
// @produce json
// @param tag query string false "tag of the notification-{QuotaUpdate,Activity,Weekly}, if not specify, get all the category"
// @param page query int true "index of page"
// @param page_size query int true "size of each page"
// @success 200 "array of notification"
// @router /notification [get]
func GetNotification(c *gin.Context) {
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

	res, count := notification.GetNotification(tag, page, pageSize)
	c.JSON(200, gin.H{
		"notification": res,
		"count":        count,
	})
}

// PublishNotification godoc
// @summary publish new notification
// @description publish new notification with title, content, category
// @produce json
// @param raw_data formData string true "the hash of specific information to sign"
// @param signature formData string true "the signature of the raw_data by the publisher"
// @param title formData string true "the title of the notification"
// @param content formData string true "the content of the notification"
// @param tag formData string true "the category of the notification: { QuotaUpdate, Weekly, Activity}"
// @success 200
// @router /notification [post]
func PublishNotification(c *gin.Context) {
	rawData := c.Param("raw_data")
	signature := c.PostForm("signature")
	if !auth.IsPublisher(rawData, signature) {
		c.JSON(403, gin.H{
			"error":   "invalid publisher",
			"message": "current account is not the valid publisher",
		})
		return
	}
	tag := c.PostForm("tag")
	if validNotificationTag(tag) {
		c.JSON(400, gin.H{
			"error":   "invalid tag",
			"message": "tag should be one of {QuotaUpdate, Weekly, Activity}",
		})
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	announcement := notification.Notification{
		Title:     title,
		Category:  tag,
		Content:   content,
		TimeStamp: time.Now().Unix(),
	}
	if err := notification.PublishNotification(announcement); err != nil {
		c.JSON(500, gin.H{
			"error":   "fail to publish notification",
			"message": err,
		})
		return
	}
}
