package vault

import (
	"github.com/SteinsElite/pickGinS/internal/coin"
	"github.com/SteinsElite/pickGinS/internal/token"
)

// get the ratio of each asset in the vault
func AssetRatio() map[string]float64 {
	ratio := make(map[string]float64)

	for k, v := range vaultWatcher.stats.CoinAmount {
		ids, _ := token.TokenIds(k)
		ratio[k] = v * coin.GetCurrentCoinPrice(ids)
	}
	return ratio
}

// get the [(timestamp, volumevalue)] in a specific time range
func PhasedVolume(phase string) []ValuePair {
	qualifiedStats := getQualifiedStatsFromDb(phase)
	// when query the volume, we shoule include the query time now
	qualifiedStats = append(qualifiedStats, vaultWatcher.stats)

	volumeUsd := make([]ValuePair, len(qualifiedStats))

	for i := range volumeUsd {
		volumeUsd[i] = ValuePair{
			TimeStamp: qualifiedStats[i].TimeStamp,
			Value:     volumeValue(qualifiedStats[i]),
		}
	}
	return volumeUsd
}

// get the [(timestamp, profit)] in the period of time
func PhasedProfit(phase string) []ValuePair {
	profitPair := getQulifiedProfitFromDb(phase)
	profitPair = append(profitPair, ValuePair{
		TimeStamp: vaultWatcher.stats.TimeStamp,
		Value:     vaultWatcher.stats.Profit,
	})
	profitUsd := []ValuePair{}
	if len(profitPair) == 1 {
		profitUsd = append(profitUsd, ValuePair{
			TimeStamp: profitPair[0].TimeStamp,
			Value:     profitValue(profitValue(profitPair[0].Value)),
		})
	}
	for i := 0; i < len(profitPair)-1; i++ {
		deltaProfit := profitPair[i+1].Value - profitPair[i].Value
		profitUsd = append(profitUsd, ValuePair{
			TimeStamp: profitPair[i+1].TimeStamp,
			Value:     profitValue(deltaProfit),
		})
	}
	return profitUsd

}
