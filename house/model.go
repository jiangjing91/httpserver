package house

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func QueryPositonNew(position string, nowTime int64) HouseList {
	conn, err := sql.Open("mysql", "root:root@/house")

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//defer fmt.Println("db close")

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	sql := "select house_id, price, total_price, fetch_time, href, position, count(house_id) as cnt from house_uniq where position like ? group by house_id having cnt = 1 and fetch_time > ? order by fetch_time desc"

	stmt, err := conn.Prepare(sql)
	if err != nil {
		panic(err)
	}

	var timeCondition string

	if nowTime == 0 {
		nowTime = time.Now().Unix()
		//fmt.Println(nowTime, (nowTime+3600*8)%86400)
		nowTime = nowTime - (nowTime+3600*8)%86400
		//	fmt.Println(nowTime)
		beginTime := nowTime - 86400*3
		timeCondition = time.Unix(beginTime, 0).Format("2006010215")
	} else {
		timeCondition = strconv.FormatInt(nowTime, 10)
	}

	//fmt.Println(nowTime, beginTime, timeCondition)
	fmt.Println("new house from[", timeCondition, "]pos like[", position, "]")
	rows, err := stmt.Query(position+"%", timeCondition)
	if err != nil {
		panic(err)
	}

	res := HouseList{}

	for rows.Next() {
		var hr HouseRec
		var house_id string
		var price string
		var total_price int
		var fetch_time int
		var href string
		var pos string
		var cnt string
		err = rows.Scan(&house_id, &price, &total_price, &fetch_time, &href, &pos, &cnt)

		if err != nil {
			panic(err)
		}

		hr.HouseId = house_id
		hr.Price = price
		hr.TotalPrice = total_price
		hr.FetchTime = fetch_time
		hr.Href = href
		hr.Pos = pos
		hr.Cnt = cnt
		//fmt.Println(house_id, price, total_price, fetch_time, href, pos)
		res = append(res, hr)
	}
	return res
}

func QueryPositon(position string) []HouseRec {
	conn, err := sql.Open("mysql", "root:root@/house")

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//defer fmt.Println("db close")

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	sql := "select house_id, price, min(total_price), min(fetch_time), href, position, count(house_id) as cnt from house_uniq where position like ? group by house_id order by fetch_time desc"

	stmt, err := conn.Prepare(sql)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(position + "%")
	if err != nil {
		panic(err)
	}

	res := []HouseRec{}

	for rows.Next() {
		var hr HouseRec
		var house_id string
		var price string
		var total_price int
		var fetch_time int
		var href string
		var pos string
		var cnt string
		err = rows.Scan(&house_id, &price, &total_price, &fetch_time, &href, &pos, &cnt)

		if err != nil {
			panic(err)
		}

		hr.HouseId = house_id
		hr.Price = price
		hr.TotalPrice = total_price
		hr.FetchTime = fetch_time
		hr.Href = href
		hr.Pos = pos
		hr.Cnt = cnt
		//fmt.Println(house_id, price, total_price, fetch_time, href, pos)
		res = append(res, hr)
	}
	return res
}

func QueryPositonChanged(position string) HouseList {
	conn, err := sql.Open("mysql", "root:root@/house")

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//defer fmt.Println("db close")

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	sql := "select a.house_id, a.price, a.total_price, a.fetch_time, a.href, a.position from house_uniq a join (select house_uniq.house_id, count(*) as cnt from house_uniq where position like ? group by house_uniq.house_id having cnt > 1) b on a.house_id = b.house_id order by a.house_id,fetch_time"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println(position + "%")
	rows, err := stmt.Query(position + "%")
	if err != nil {
		panic(err)
	}
	var lastHouseId string
	var lastPrice string
	var lastTotalPrice int
	var lastFetchTime int
	mapHouseRec := make(map[string]HouseRec)
	mapHref := make(map[string]string)
	mapPosition := make(map[string]string)
	mapInList := make(map[string]bool)

	var hlist HouseList

	for rows.Next() {
		var rec HouseRec
		var house_id string
		var price string
		var total_price int
		var fetch_time int
		var href string
		var pos string
		err = rows.Scan(&house_id, &price, &total_price, &fetch_time, &href, &pos)

		fmt.Println(house_id)

		if err != nil {
			panic(err)
		}

		//存href,position

		mapHref[house_id] = href
		mapPosition[house_id] = pos
		//检查调价
		if lastTotalPrice == 0 {
			lastHouseId = house_id
			lastPrice = price
			lastTotalPrice = total_price
			lastFetchTime = fetch_time

			_, keyExsit := mapHouseRec[house_id]
			if !keyExsit {
				rec.HouseId = house_id
				rec.Price = price
				rec.TotalPrice = total_price
				rec.FetchTime = fetch_time
				rec.Href = mapHref[house_id]
				rec.Pos = mapPosition[house_id]
				mapHouseRec[house_id] = rec
			}
		}

		if lastHouseId == house_id {
			if price != lastPrice || lastTotalPrice != total_price {
				fmt.Println(house_id, lastPrice, lastTotalPrice, lastFetchTime)
				fmt.Println(house_id, price, total_price, fetch_time, mapHref[house_id], mapPosition[house_id])

				rec.HouseId = house_id
				rec.Price = price
				rec.TotalPrice = total_price
				rec.FetchTime = fetch_time
				rec.Href = mapHref[house_id]
				rec.Pos = mapPosition[house_id]

				if !mapInList[house_id] {
					hlist = append(hlist, mapHouseRec[house_id])
					mapInList[house_id] = true
				}

				hlist = append(hlist, rec)

				lastPrice = price
				lastTotalPrice = total_price
				lastFetchTime = fetch_time
			}
		} else {
			lastHouseId = house_id
			lastFetchTime = fetch_time
			lastPrice = price
			lastTotalPrice = total_price

			_, keyExsit := mapHouseRec[house_id]
			if !keyExsit {
				rec.HouseId = house_id
				rec.Price = price
				rec.TotalPrice = total_price
				rec.FetchTime = fetch_time
				rec.Href = mapHref[house_id]
				rec.Pos = mapPosition[house_id]
				mapHouseRec[house_id] = rec
			}
		}

	}

	return hlist

}
