package router

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/SteinsElite/pickGinS/service/coin"
	"github.com/SteinsElite/pickGinS/service/notification"
	"github.com/SteinsElite/pickGinS/service/transaction"
	"github.com/SteinsElite/pickGinS/service/vault"
	"github.com/SteinsElite/pickGinS/util"
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
	values, startTime := vault.PhasedVolume(phase)
	c.JSON(200, gin.H{
		"startTime": startTime,
		"volume": values,
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
	values := vault.PhasedProfit(phase)
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

// GetKeyWordHash godoc
// @summary get the keywordHash to sign
// @description the keyword hash is sign by the account to make sure that the account is accessed
// @produce json
// @param address query string true "the address of the keyword bind to"
// @success 200 "keyword hash"
// @failure 400 "invalid param"
// @failure 403 "not authorized"
// @router /auth/keyword_hash [get]
func GetKeyWordHash(c *gin.Context) {
	accountAddr := c.Query("address")
	if !util.IsValidAddress(accountAddr) {
		c.JSON(400, gin.H{
			"error":   "Invalid params",
			"message": "address is not valid ethereum address",
		})
		return
	}
	word := getAuthWord(accountAddr)
	if word == nil {
		c.JSON(403, gin.H{
			"error":   "Fail to getAuthWord",
			"message": "the address is not register as admin",
		})
		return
	}
	c.JSON(200, gin.H{
		"keyword_hash": word,
	})
}

// AddPublisher godoc
// @summary add new publisher
// @description add new publisher who is ability to publish new notification(
// only publisher could add publisher)
// @produce json
// @param address path string true "the publisher address is login now"
// @param signature formData string true "signature of the publisher address"
// @param new_publisher formData string true "the address of new publisher to add"
// @param keyword formData string true "the keyword been used to sign"
// @success 200
// @router /auth/{address}/add_publisher [post]
func AddPublisher(c *gin.Context) {
	adminAddr := c.Param("address")
	sig := c.PostForm("signature")
	newPublisher := c.PostForm("new_publisher")
	keyword := c.PostForm("keyword")

	if !IsAuth(adminAddr, sig) {
		c.JSON(403, gin.H{
			"error":   "not permission",
			"message": "current account is not the publisher",
		})
		return
	}
	SetNewPublisher(newPublisher, keyword)
}

// PublishNotification godoc
// @summary publish new notification
// @description publish new notification with title, content, category
// @produce json
// @param publisher path string true "the publisher address"
// @param signature formData string true "the signature of the publisher"
// @param title formData string true "the title of the notification"
// @param content formData string true "the content of the notification"
// @param tag formData string true "the category of the notification: { QuotaUpdate, Weekly, Activity}"
// @success 200
// @router /notification/{publisher} [post]
func PublishNotification(c *gin.Context) {
	publisher := c.Param("publisher")
	signature := c.PostForm("signature")
	if !IsAuth(publisher, signature) {
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
