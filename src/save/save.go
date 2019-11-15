package save

import (
	"../model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)
import "../data"
var c = make(chan *data.Ars,10)
func readData() []*data.Ars {
	dx,err := os.Open(`./data.json`)
	if err != nil {
		log.Println(err)
		return nil
	}
	ars := []*data.Ars{}
	err = json.NewDecoder(dx).Decode(&ars)
	return ars
}

func DoWork() {
  db := model.GetDbx()
	wg := new(sync.WaitGroup)
	defer func() func(){
		start := time.Now()
		return func() {
			fmt.Println("用时:",time.Since(start))
		}
	}()()
	for i := 0;i < 10;i ++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				db.MustExec(`insert into ars(title,icon,content,author,c_time) values(?,?,?,?,?)`,
					v.Title,v.Icon,v.Content,v.Author,v.CTime,
					)
			}
		}()
	}

	go func() {
		for {
			for _,v := range `-\|/` {
				fmt.Printf("\r%c正在写入数据",v)
				time.Sleep(time.Second / 2)
			}
		}

	}()

	ds := readData()
	for _,v := range ds {
		c <- v
	}

	close(c)
	wg.Wait()
	fmt.Println()
	fmt.Printf("写入数据完成 %d条记录\n",len(ds))
}

