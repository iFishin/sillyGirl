package onebyone

import (
	"encoding/json"
	"strings"

	"github.com/FishEyeiii/sillyGirl/core"
	"github.com/gin-gonic/gin"
)

type AutoGenerated struct {
	PtPin   string `json:"pt_pin"`
	Message string `json:"message"`
}

func init() {

	core.Server.POST("/onebyone/push", func(c *gin.Context) {
		data, _ := c.GetRawData()
		ag := &AutoGenerated{}
		json.Unmarshal(data, ag)
		ptPin := ag.PtPin
		message := ag.Message
		for _, tp := range []string{
			"qq", "tg", "wx",
		} {
			core.Bucket("pin" + strings.ToUpper(tp)).Foreach(func(k, v []byte) error {
				translateEmoji(&message, tp == "wx")
				if string(k) == ptPin && ptPin != "" {
					if push, ok := core.Pushs[tp]; ok {
						push(string(v), message, nil, nil)
					}
				}
				return nil
			})
		}
		c.String(200, "ok")
	})
}

func translateEmoji(str *string, isWechat bool) {

	if !isWechat {
		return
	}

	*str = strings.Replace(*str, "⭕", "[emoji=\\u2b55]", -1)
	*str = strings.Replace(*str, "🧧", "[emoji=\\uD83E\\uDDE7]", -1)
	*str = strings.Replace(*str, "🥚", "[emoji=\\ud83e\\udd5a]", -1)
	*str = strings.Replace(*str, "💰", "[emoji=\\ud83d\\udcb0]", -1)
	*str = strings.Replace(*str, "⏰", "[emoji=\\u23f0]", -1)
	*str = strings.Replace(*str, "🍒", "[emoji=\\ud83c\\udf52]", -1)
	*str = strings.Replace(*str, "🐶", "[emoji=\\ud83d\\udc36]", -1)
	*str = strings.Replace(*str, "🎰", "[emoji=\\ud83c\\udfb0]", -1)
	*str = strings.Replace(*str, "🌂", "[emoji=\\ud83c\\udf02]", -1)
}
