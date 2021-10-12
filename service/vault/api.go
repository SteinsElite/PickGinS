package vault

import (
	"fmt"
	"time"

	"github.com/SteinsElite/pickGinS/service/coin"
)

// AssetRatio get the ratio of each asset in the vault
func AssetRatio() map[string]float64 {
	ratio := make(map[string]float64)
	for k, v := range vaultWatcher.stats.CoinAmount {
		ratio[k] = v * coin.GetCurrentCoinPrice(k)
	}
	return ratio
}

// PhasedVolume get the [(timestamp, volumeValue)] in a specific time range
func PhasedVolume(phase string) []ValuePair {
	startTime, _ := queryStartTimeForVolume(phase)
	qualifiedStats := getQualifiedStatsFromDb(startTime)
	// when query the volume, we should include the query time now
	qualifiedStats = append(qualifiedStats, vaultWatcher.stats)

	volumeUsd := make([]ValuePair, len(qualifiedStats))

	for i := 0; i < len(volumeUsd); i++ {
		volumeUsd[i] = ValuePair{
			TimeStamp: qualifiedStats[i].TimeStamp,
			Value:     volumeValue(qualifiedStats[i]),
		}
	}
	return volumeUsd
}

// PhasedProfit get the [(timestamp, profit)] in the period of time
func PhasedProfit(phase string) []ValuePair {
	ticks, _ := queryTimeTickForProfit(phase)
	profitPair := getQualifiedProfitFromDb(ticks)

	profitPair = append(profitPair, ValuePair{
		TimeStamp: vaultWatcher.stats.TimeStamp,
		Value:     vaultWatcher.stats.Profit,
	})
	var profitUsd []ValuePair

	for i := 0; i < len(profitPair)-1; i++ {
		deltaProfit := profitPair[i+1].Value - profitPair[i].Value
		profitUsd = append(profitUsd, ValuePair{
			TimeStamp: profitPair[i+1].TimeStamp,
			Value:     profitValue(deltaProfit),
		})
	}
	return profitUsd
}

// get the start timestamp of volume status for query the database
// when we deal with the volume, we just need get the unix time of the start day,
func queryStartTimeForVolume(phase string) (it int64, err error) {
	midnight := midnightOfDay()
	var t time.Time
	switch phase {
	case Week:
		t = midnight.AddDate(0, 0, -6)
	case Month:
		t = midnight.AddDate(0, 0, -29)
	case Year:
		t = midnight.AddDate(0, 0, -364)
	default:
		err = fmt.Errorf("wrong time range")
	}
	it = t.Unix()
	return
}

// get the collection of the timestamp that we should get Form the database and then calculate the
// profit we need
func queryTimeTickForProfit(phase string) (timetick []int64, err error) {
	midnight := midnightOfDay()
	switch phase {
	case Week:
		// get 7 timetick
		for i := 7; i > 0; i-- {
			tick := midnight.AddDate(0, 0, -1*i)
			timetick = append(timetick, tick.Unix())
		}
	case Month:
		// we suspend get the latest 30 data
			for i := 30; i > 0; i-- {
			tick := midnight.AddDate(0, 0, -1*i)
			timetick = append(timetick, tick.Unix())
		}
	case Year:
		currentMonth := startOfMonth()
		for i := 12; i > 0; i-- {
			tick := currentMonth.AddDate(0, -1*i, 0)
			timetick = append(timetick, tick.Unix())
		}
	default:
		err = fmt.Errorf("get the wrong time range")
	}
	return

}

// get the midnight time of the current day
func midnightOfDay() time.Time {
	current := time.Now()
	return time.Date(
		current.Year(),
		current.Month(),
		current.Day(),
		0, 0, 0, 0,
		time.UTC,
	)
}

// ge the midnight of the first day of the current month
func startOfMonth() time.Time {
	current := time.Now()
	return time.Date(
		current.Year(),
		current.Month(),
		0, 0, 0, 0, 0,
		time.UTC,
	)
}
