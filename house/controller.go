package house

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/codegangsta/martini"
)

type HouseRec struct {
	HouseId    string `json:"house_id"`
	Price      string `json:"price"`
	TotalPrice int    `json:"total_price"`
	FetchTime  int    `json:"fetch_time"`
	Href       string `json:"href"`
	Pos        string `json:"pos"`
	Cnt        string `json:"cnt,omitempty"`
	Title      string `json:"title,omitempty"`
	Img        string `json:"img,omitempty"`
}

type HouseList []HouseRec

/**
下面的这些服务已经被包含在核心Martini中: martini.Classic():

*log.Logger - Martini的全局日志.
martini.Context - http request context （请求上下文）.
martini.Params - map[string]string of named params found by route matching. （名字和参数键值对的参数列表）
martini.Routes - Route helper service. （路由协助处理）
http.ResponseWriter - http Response writer interface. (响应结果的流接口)
*http.Request - http Request. （http请求)
**/

func List(params martini.Params, logger *log.Logger, req *http.Request, res http.ResponseWriter, conn *sql.DB) (int, string) {
	logger.Println(params)

	if len(strings.TrimSpace(params["position"])) > 0 {
		sqlRes := QueryPositon(params["position"], conn)
		logger.Println(sqlRes)

		jsonRes, err := json.Marshal(sqlRes)
		if err != nil {
			panic(err)
		}
		logger.Println(string(jsonRes))

		return 200, string(jsonRes)

	} else {
		logger.Println("no position")
	}

	return 200, ""
}

func ListNew(params martini.Params, logger *log.Logger, req *http.Request, res http.ResponseWriter, conn *sql.DB) (int, string) {
	logger.Println(params)

	if len(strings.TrimSpace(params["position"])) > 0 {
		sqlRes := QueryPositonNew(params["position"], 0, conn)
		logger.Println(sqlRes)

		jsonRes, err := json.Marshal(sqlRes)
		if err != nil {
			panic(err)
		}
		logger.Println(string(jsonRes))

		return 200, string(jsonRes)

	} else {
		logger.Println("no position")
	}

	return 200, ""
}

func ListHistory(params martini.Params, logger *log.Logger, req *http.Request, res http.ResponseWriter, conn *sql.DB) (int, string) {
	logger.Println(params)

	if len(strings.TrimSpace(params["position"])) > 0 && len(strings.TrimSpace(params["date"])) > 0 {

		var iDate int64
		var err error
		iDate, err = strconv.ParseInt(params["date"], 10, 64)

		if err != nil {
			panic(err)
		}

		sqlRes := QueryPositonNew(params["position"], iDate, conn)
		logger.Println(sqlRes)

		jsonRes, err := json.Marshal(sqlRes)
		if err != nil {
			panic(err)
		}
		logger.Println(string(jsonRes))

		return 200, string(jsonRes)

	} else {
		logger.Println("error")
	}

	return 200, ""
}

func ListChanged(params martini.Params, logger *log.Logger, req *http.Request, res http.ResponseWriter, db *sql.DB) (int, string) {
	logger.Println(params)

	if len(strings.TrimSpace(params["position"])) > 0 {
		sqlRes := QueryPositonChanged(params["position"], db)
//		logger.Println(sqlRes)

		jsonRes, err := json.Marshal(sqlRes)
		if err != nil {
			panic(err)
		}
//		logger.Println(string(jsonRes))

		return 200, string(jsonRes)

	} else {
		logger.Println("no position")
	}

	return 200, ""
}
