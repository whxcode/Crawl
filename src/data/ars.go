package data

import (
	"regexp"
	"strings"
)

type Ars struct {
	Title string `json:"title" db:"title"`
	Icon string `json:"icon" db:"icon"`
	Content string `json:"content" db:"content"`
	Author string `json:"author" db:"author"`
	CTime string `json:"cTime" db:"c_time"`
	Volume string `json:"volume" db:"volume"`
	Fabulous string `json:"fabulous" db:"Fabulous"`
	CommentQuantity string `json:"commentQuantity" db:"comment_quantity"`
}

func CreateArs(d string) *Ars {
	ars := &Ars{}
	regTitle := regexp.MustCompile(`(?s:<h1>(.*?)</h1>)`)
	tits := regTitle.FindStringSubmatch(d)
	ars.Title =  strings.TrimSpace(tits[1])

	regdge := regexp.MustCompile(`<div badge="(\d+)"`)
	dges := regdge.FindAllStringSubmatch(d,-1)

	ars.Fabulous = dges[0][1]
	ars.CommentQuantity = dges[1][1]

	regName := regexp.MustCompile(`<a href="/user/.*?" title="(.*?)">`)
	names := regName.FindStringSubmatch(d)
	ars.Author = names[1] //获取用户名

	regIcon  := regexp.MustCompile(`(?s:<div class="pull-right">.*?<img.*?src="(.*?)".*?>.*?</div>)`)
	icons := regIcon.FindStringSubmatch(d)
	ars.Icon = icons[1] // 获取用户头像

	regTime := regexp.MustCompile(`<span title="(.*?)" class="timeago">`)
	times := regTime.FindStringSubmatch(d)
	ars.CTime = times[1] // 获取时间

	regDiJi := regexp.MustCompile(` <div class="pull-right c9 f11" style="line-height: 12px; padding-top: 3px; text-shadow: 0px 1px 0px #fff;">.*?(\d+).*?次点击.*?</div>`)
	diJis := regDiJi.FindStringSubmatch(d)
	ars.Volume = diJis[1] // 访问量
	regContent := regexp.MustCompile(`(?s:(<div class="box_white" style="overflow: visible;">.*?</div>.*?)<div class="sep20">)`)
	content := regContent.FindStringSubmatch(d)
	ars.Content = content[1] // 获取内容
	return ars
}