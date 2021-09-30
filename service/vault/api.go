package vault

import (
	"github.com/SteinsElite/pickGinS/internal/coin"
	"github.com/SteinsElite/pickGinS/internal/vault"
	"github.com/SteinsElite/pickGinS/types"
)

// get the ratio of each asset in the vault
func AssetRatio() map[string]float64 {
	ratio := make(map[string]float64)

	for k, v := range vault.vaultWatcher.stats.CoinAmount {
		ids, _ := _type.TokenIds(k)
		ratio[k] = v * coin.GetCurrentCoinPrice(ids)
	}
	return ratio
}

// get the [(timestamp, volumevalue)] in a specific time range
func PhasedVolume(phase string) []vault.ValuePair {
	qualifiedStats := vault.getQualifiedStatsFromDb(phase)
	// when query the volume, we shoule include the query time now
	qualifiedStats = append(qualifiedStats, vault.vaultWatcher.stats)

	volumeUsd := make([]vault.ValuePair, len(qualifiedStats))

	for i := range volumeUsd {
		volumeUsd[i] = vault.ValuePair{
			TimeStamp: qualifiedStats[i].TimeStamp,
			Value:     vault.volumeValue(qualifiedStats[i]),
		}
	}
	return volumeUsd
}

// get the [(timestamp, profit)] in the period of time
func PhasedProfit(phase string) []vault.ValuePair {
	profitPair := vault.getQulifiedProfitFromDb(phase)
	profitPair = append(profitPair, vault.ValuePair{
		TimeStamp: vault.vaultWatcher.stats.TimeStamp,
		Value:     vault.vaultWatcher.stats.Profit,
	})
	profitUsd := []vault.ValuePair{}
	if len(profitPair) == 1 {
		profitUsd = append(profitUsd, vault.ValuePair{
			TimeStamp: profitPair[0].TimeStamp,
			Value:     vault.profitValue(vault.profitValue(profitPair[0].Value)),
		})
	}
	for i := 0; i < len(profitPair)-1; i++ {
		deltaProfit := profitPair[i+1].Value - profitPair[i].Value
		profitUsd = append(profitUsd, vault.ValuePair{
			TimeStamp: profitPair[i+1].TimeStamp,
			Value:     vault.profitValue(deltaProfit),
		})
	}
	return profitUsd

}
