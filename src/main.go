 package main

import (
	"./data"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
)
import "./save"

const path = `https://studygolang.com/topics?p=`
const c = `TY_SESSION_ID=380f71a2-6178-4941-ba75-e491b5dde755; uuid_tt_dd=10_30622525980-1563180031426-412934; dc_session_id=10_1563180031426.319725; smidV2=201908061634212e3f93369529268dde40648a93eb571700845405285580120; UserName=qq_30428653; UserInfo=649707ae4f3944c19cf2b74126155bc9; UserToken=649707ae4f3944c19cf2b74126155bc9; UserNick=qq_30428653; AU=C0E; UN=qq_30428653; BT=1568786396668; p_uid=U000000; Hm_ct_6bcd52f51e9b3dce32bec4a3997715ac=6525*1*10_30622525980-1563180031426-412934!1788*1*PC_VC!5744*1*qq_30428653; __gads=Test; acw_tc=2760828a15736449211416311e0febc1b389f767541cf1fcd21a86f983bf43; Hm_lvt_6bcd52f51e9b3dce32bec4a3997715ac=1573612441,1573644870,1573644883,1573645048; _csdn_newbbs_session=BAh7B0kiDHVzZXJfaWQGOgZFRmkE9RRUA0kiD3Nlc3Npb25faWQGOwBGSSIlZWM0OWMzMGY0OWQzOWI4NDRkZGEwMzA2N2QwMTA3MDYGOwBU--5703855d57519f1eebcd4df67f0b51db98ad1bb2; acw_sc__v2=5dcbfca240f10d728d67ad71d2da0025db29c51d; acw_sc__v3=5dcbfda59f67e010827ef3dc03433ab13ee1fecb; announcement=%257B%2522isLogin%2522%253Atrue%252C%2522announcementUrl%2522%253A%2522https%253A%252F%252Fblogdev.blog.csdn.net%252Farticle%252Fdetails%252F102605809%2522%252C%2522announcementCount%2522%253A0%252C%2522announcementExpire%2522%253A3600000%257D; dc_tos=q0wrze; Hm_lpvt_6bcd52f51e9b3dce32bec4a3997715ac=1573649835`
const url = `https://studygolang.com/topics/`
var cUrls = make(chan string,100)


var t = flag.Bool(`t`,false,"not save")
var start = flag.Int(`s`,1,"start")
var end =flag.Int(`e`,2,"end")

func init() {
	flag.Parse()
}

func fetch(url string,done chan <- bool) {

	defer func() {
		done <- true
	}()
	res,err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	d := readText(res.Body)
	regTitle := regexp.MustCompile(`<a href="/topics/(\d+)".*?>`)
	tits := regTitle.FindAllStringSubmatch(d,-1)
	for _,v := range tits {
		cUrls <- v[1]
	}

}

func readText(at io.ReadCloser) string {
	d,_ := ioutil.ReadAll(at)
	return string(d)
}

func doWork() {
	done := make(chan bool,200)
	for i := *start;i <= *end;i ++ {
		u := path + strconv.Itoa(i)

		go fetch(u,done)
	}
	for i := *start ;i <= *end;i ++ {
		<- done
	}
  close(cUrls)
}

func fetchInfo(url string,v chan <- *data.Ars) {
	  req,err := http.NewRequest("GET",url,nil)
	// req.Header.Set(`Accept`,`text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3`)
	  req.Header.Set(`cookie`,`_ga=GA1.2.1413409530.1563156784; _gid=GA1.2.1724866822.1573647200; Hm_lvt_224c227cd9239761ec770bc8c1fb134c=1573088877,1573180853,1573647200,1573650451; __gads=Test; user=MTU3MzY1MzMzMXxEdi1CQkFFQ180SUFBUkFCRUFBQUt2LUNBQUVHYzNSeWFXNW5EQXNBQ1VsT1JFVllYMVJCUWdaemRISnBibWNNQ1FBSGNHRjViMjVzZVE9PXyAm5bn-jwQvR7am_rij69jIwqxBODhkVAmgeNl1Jxzvw==; Hm_lpvt_224c227cd9239761ec770bc8c1fb134c=1573656708`)
	  req.Header.Set(`User-Agent`,`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36`)
	 res,err := (&http.Client{
	 	Timeout:100 * time.Second,
	 }).Do(req)
		if err != nil {
			log.Println(err)
			return
		}

	 d := readText(res.Body)
	 v <- data.CreateArs(d)

}

func readArs(dxs *[]*data.Ars,doneArs <- chan *data.Ars,wg *sync.WaitGroup) {
		defer wg.Done()
		for v := range doneArs {
			*dxs = append(*dxs,v)
		}
}


func reptile() {
	dxs := []*data.Ars{}
	doneArs := make(chan *data.Ars,20)
	start := time.Now()
	go doWork()

	done := make(chan  bool,2)

	for i := 0;i < 4;i ++ {
		go func() {
			defer func() {
				done <- true
			}()

			for  v := range cUrls {
				u := url + v
				fetchInfo(u,doneArs)
			}
		}()
	}

	go func() {
		for {
			for _,v := range `-\|/` {
				fmt.Printf("\r%c正在爬取数据",v)
				time.Sleep(time.Second / 2)
			}
		}
		fmt.Println()
	}()

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go readArs(&dxs,doneArs,wg)

	for i := 0;i < 4;i ++ {
		<- done
	}

	close(doneArs)
	wg.Wait()

	fs,_ := os.Create(`./data.json`)
	d,_ := json.MarshalIndent(dxs,""," ")
	fs.Write(d)
	fmt.Printf("一共获取 %d 条数据\n",len(dxs))
	fmt.Println(time.Since(start))
}

func main() {
	if(*t) {
		save.DoWork()
	} else {
		reptile()
	}
}
