package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "context"
	//"time"
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)


	http.HandleFunc("/callback", callbackHandler)


	// fmt.Println(prof)
	binary, _ := Redis_Get(mac)
	user := new(USER_MAC)
	json.Unmarshal(binary,&user)
	onlyonecontent := user.USER[1].CONTENT
	// var allcontent string
	 // for i:=0;i < len(user.USER) ; i++{
	 //  allcontent = allcontent+user.USER[i].CONTENT
	 // }

	template := linebot.NewButtonsTemplate(
			"", "請告訴我們，您對於以上文章的評價" , "" ,
			linebot.NewPostbackTemplateAction("滿意 :)", "Y" , ""),
			linebot.NewPostbackTemplateAction("不滿意 :(", "N" , ""),
	)

	bot.PushMessage(
		"Uecc089487f1487a78637be4e2fe3dca9",
		linebot.NewTextMessage(onlyonecontent)).Do()

	bot.PushMessage(
		"Uecc089487f1487a78637be4e2fe3dca9",
		linebot.NewTemplateMessage("今日文章", template)).Do()



	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	// defer cancel()
	for _, event := range events {
		if event.Type == linebot.EventTypeFollow {
			prof := event.Source.UserID
			follow := linebot.NewButtonsTemplate(
					"", "歡迎您使用本服務!", "你好，我是中央大學的曾怡雯，\n請按以下按鈕做出對應的動作：",
					linebot.NewPostbackTemplateAction("我要輸入我的MAC", "AddMAC",""),
					linebot.NewPostbackTemplateAction("我要更正我的MAC", "ModifyMAC",""),
					linebot.NewURITemplateAction("我想看怡雯畫的阿卡莉", "https://www.pixiv.net/member_illust.php?mode=medium&illust_id=62861397"),
			)

			if _, err := bot.PushMessage(prof, linebot.NewTemplateMessage("Smart AP <3", follow)).Do(); err != nil {
					log.Print(err)
			}
		}
		if event.Type == linebot.EventTypePostback {
			prof := event.Source.UserID
			data := event.Postback.Data
			if data == "AddMAC"{
				bot.PushMessage(prof, linebot.NewTextMessage("現在請輸入你的MAC：")).Do()
			}
			if data == "ModifyMAC"{
				bot.PushMessage(prof, linebot.NewTextMessage("現在請輸入你要更正的MAC：")).Do()
			}
		}
	}
}
