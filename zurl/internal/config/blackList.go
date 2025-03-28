package config

type BlackListConfig struct {
	BlackList []string
}

var BlackLC BlackListConfig

var BlackLCMap map[string]struct{}

func (blc *BlackListConfig) ToMap() {
	BlackLCMap = make(map[string]struct{}, len(blc.BlackList))
	for _, v := range blc.BlackList {
		BlackLCMap[v] = struct{}{}
	}
}
