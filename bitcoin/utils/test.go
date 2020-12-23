package utils

import (
	"bitcoin/entity"
	"fmt"
	"strings"
)


//获取区块链信息
func GetTest() {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON("getmempoolinfo", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	mempoolInfo := entity.MempoolInfo{}
	//fmt.Printf("%T %v", rpcResult.Data.Result, rpcResult.Data.Result)
	fmt.Printf("")
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		//Entitys(res,"MempoolInfo")
		MyUnmarshal(res, "mempoolInfo")

		//fmt.Println(res)
		mempoolInfo.Loaded = res["loaded"].(bool)
		mempoolInfo.Size = res["size"].(float64)
		mempoolInfo.Bytes = res["bytes"].(float64)
		mempoolInfo.Usage = res["usage"].(float64)
		mempoolInfo.Maxmempool = res["maxmempool"].(float64)
		mempoolInfo.Mempoolminfee = res["mempoolminfee"].(float64)
		mempoolInfo.Minrelaytxfee = res["minrelaytxfee"].(float64)
	}

	//return blockStats
}
