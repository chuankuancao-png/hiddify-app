package hcore

import "github.com/sagernet/sing-box/common/trafficlimit"

func applyTrafficRateLimit() {
	if static.HiddifyOptions == nil {
		trafficlimit.Configure(0, 0)
		return
	}
	trafficlimit.Configure(
		static.HiddifyOptions.RateLimit.UploadMbps,
		static.HiddifyOptions.RateLimit.DownloadMbps,
	)
}
