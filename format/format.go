package format

import (
	"github.com/fanap-infra/rtsp/av/avutil"
	"github.com/fanap-infra/rtsp/format/aac"
	"github.com/fanap-infra/rtsp/format/flv"
	"github.com/fanap-infra/rtsp/format/mp4"
	"github.com/fanap-infra/rtsp/format/rtmp"
	"github.com/fanap-infra/rtsp/format/rtsp"
	"github.com/fanap-infra/rtsp/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
