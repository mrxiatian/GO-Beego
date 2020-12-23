package utils

import (
	"bitcoin/bitServices"
	"bitcoin/entity"
	"strings"
)

type btcSer struct {
	bitServices.BlockChahin
	bitServices.Control
	bitServices.Generating
	bitServices.Mining
	bitServices.Network
	bitServices.Util
	bitServices.Wallet
	bitServices.Zmq
}

func GetBC() btcSer {
	return btcSer{}
}

//获取最高区块Hash,成功返回最高区块Hash,否则返回空字符串
func (bc btcSer) GetBestBlockHahs() string {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETBESTBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}
	return ""
}

//根据区块Hash获取区块信息
func (bc btcSer) GetBlockInfoByHash(blockHash string) entity.BlockInfo {
	paramsSlice := []interface{}{blockHash}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETBLOCK, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	blockInfo := entity.BlockInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockInfo.Time = resBytes["time"].(float64)
		blockInfo.Bits = resBytes["bits"].(string)
		blockInfo.NTx = resBytes["nTx"].(float64)
		blockInfo.Previousblockhash = resBytes["previousblockhash"].(string)
		blockInfo.Strippedsize = resBytes["strippedsize"].(float64)
		blockInfo.Size = resBytes["size"].(float64)
		blockInfo.Merkleroot = resBytes["merkleroot"].(string)
		blockInfo.Weight = resBytes["weight"].(float64)
		blockInfo.Version = resBytes["version"].(float64)
		blockInfo.VersionHex = resBytes["versionHex"].(string)
		blockInfo.Nonce = resBytes["nonce"].(float64)
		blockInfo.Difficulty = resBytes["difficulty"].(float64)
		blockInfo.Hash = resBytes["hash"].(string)
		blockInfo.Height = resBytes["height"].(float64)
		blockInfo.Mediantime = resBytes["mediantime"].(float64)
		blockInfo.Chainwork = resBytes["chainwork"].(string)
		blockInfo.Confirmations = resBytes["confirmations"].(float64)
		blockInfo.Tx = resBytes["tx"].([]interface{})
	}

	return blockInfo
}

//根据区块hash获取区块Info
func (bc btcSer) GetBlockByHeight(height int) entity.BlockInfo {
	blockInfo := entity.BlockInfo{}
	if float64(height) > bc.GetBlockCount() {
		return blockInfo
	}

	paramsSlice := []interface{}{height}

	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	hash, ok := rpcResult.Data.Result.(string)
	if ok {
		return bc.GetBlockInfoByHash(hash)
	}

	return blockInfo
}

//获取区块链信息
func (bc btcSer) GetBlockChainInfo() entity.BlockChainInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETBLOCKCHAININFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	blockChainInfo := entity.BlockChainInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockChainInfo.Size_on_disk = resBytes["size_on_disk"].(float64)
		blockChainInfo.Blocks = resBytes["blocks"].(float64)
		blockChainInfo.Mediantime = resBytes["mediantime"].(float64)
		blockChainInfo.Chainwork = resBytes["chainwork"].(string)
		blockChainInfo.Automatic_pruning = resBytes["automatic_pruning"].(bool)
		blockChainInfo.Pruneheight = resBytes["pruneheight"].(float64)
		blockChainInfo.Bestblockhash = resBytes["bestblockhash"].(string)
		blockChainInfo.Difficulty = resBytes["difficulty"].(float64)
		blockChainInfo.Initialblockdownload = resBytes["initialblockdownload"].(bool)
		blockChainInfo.Pruned = resBytes["pruned"].(bool)
		blockChainInfo.Prune_target_size = resBytes["prune_target_size"].(float64)
		blockChainInfo.Chain = resBytes["chain"].(string)
		blockChainInfo.Warnings = resBytes["warnings"].(string)
		blockChainInfo.Headers = resBytes["headers"].(float64)

		blockChainInfo.Softforks_, ok = resBytes["softforks"].(map[string]interface{})
		if ok {
			blockChainInfo.Softforks.Bip34_, ok = blockChainInfo.Softforks_["bip34"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Bip34.Height = blockChainInfo.Softforks.Bip34_["height"].(float64)
				blockChainInfo.Softforks.Bip34.Type = blockChainInfo.Softforks.Bip34_["type"].(string)
				blockChainInfo.Softforks.Bip34.Active = blockChainInfo.Softforks.Bip34_["active"].(bool)
			}

			blockChainInfo.Softforks.Bip65_, ok = blockChainInfo.Softforks_["bip65"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Bip65.Height = blockChainInfo.Softforks.Bip65_["height"].(float64)
				blockChainInfo.Softforks.Bip65.Type = blockChainInfo.Softforks.Bip65_["type"].(string)
				blockChainInfo.Softforks.Bip65.Active = blockChainInfo.Softforks.Bip65_["active"].(bool)
			}

			blockChainInfo.Softforks.Bip66_, ok = blockChainInfo.Softforks_["Bip66"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Bip66.Height = blockChainInfo.Softforks.Bip66_["height"].(float64)
				blockChainInfo.Softforks.Bip66.Type = blockChainInfo.Softforks.Bip66_["type"].(string)
				blockChainInfo.Softforks.Bip66.Active = blockChainInfo.Softforks.Bip66_["active"].(bool)
			}

			blockChainInfo.Softforks.Segwit_, ok = blockChainInfo.Softforks_["Segwit"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Segwit.Height = blockChainInfo.Softforks.Segwit_["height"].(float64)
				blockChainInfo.Softforks.Segwit.Type = blockChainInfo.Softforks.Segwit_["type"].(string)
				blockChainInfo.Softforks.Segwit.Active = blockChainInfo.Softforks.Segwit_["active"].(bool)
			}

			blockChainInfo.Softforks.Csv_, ok = blockChainInfo.Softforks_["Csv"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Csv.Height = blockChainInfo.Softforks.Csv_["height"].(float64)
				blockChainInfo.Softforks.Csv.Type = blockChainInfo.Softforks.Csv_["type"].(string)
				blockChainInfo.Softforks.Csv.Active = blockChainInfo.Softforks.Csv_["active"].(bool)
			}

		}
	}

	return blockChainInfo
}

//获取网络信息
func (bc btcSer) GetNetWorkInfo() entity.NetWorkInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETNETWORKINFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	netWorkInfo := entity.NetWorkInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		netWorkInfo.Relayfee = resBytes["relayfee"].(float64)
		netWorkInfo.Warnings = resBytes["warnings"].(string)
		netWorkInfo.Localrelay = resBytes["localrelay"].(bool)
		netWorkInfo.Networks_, ok = resBytes["networks"].([]interface{})
		if ok {
			for i := 0; i < len(netWorkInfo.Networks_); i++ {
				mapValue, ok := netWorkInfo.Networks_[i].(map[string]interface{})
				if ok {
					var network entity.NetWork
					network.Name = mapValue["name"].(string)
					network.Limited = mapValue["limited"].(bool)
					network.Reachable = mapValue["reachable"].(bool)
					network.Proxy = mapValue["proxy"].(string)
					network.Proxy_randomize_credentials = mapValue["proxy_randomize_credentials"].(bool)
					netWorkInfo.Networks = append(netWorkInfo.Networks, network)
				}
			}
		}

		netWorkInfo.Version = resBytes["version"].(float64)
		netWorkInfo.Subversion = resBytes["subversion"].(string)
		netWorkInfo.Protocolversion = resBytes["protocolversion"].(float64)
		netWorkInfo.Timeoffset = resBytes["timeoffset"].(float64)
		netWorkInfo.Networkactive = resBytes["networkactive"].(bool)
		netWorkInfo.Localaddresses = resBytes["localaddresses"].([]interface{})
		netWorkInfo.Localservices = resBytes["localservices"].(string)
		netWorkInfo.Localservicesnames = resBytes["localservicesnames"].([]interface{})
		netWorkInfo.Connections = resBytes["connections"].(float64)
		netWorkInfo.Incrementalfee = resBytes["incrementalfee"].(float64)
	}

	return netWorkInfo
}

//获取区块总数
func (bc btcSer) GetBlockCount() float64 {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETBLOCKCOUNT, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -1
}

//根据区块高度获取区块的hash
func (bc btcSer) GetBlockHash(height int) string {
	if float64(height) > bc.GetBlockCount() {
		return ""
	}

	paramsSlice := []interface{}{height}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	hash, ok := rpcResult.Data.Result.(string)
	if ok {
		return hash
	}

	return ""
}

//根据区块Hash获取区块头信息
func (bc btcSer) GetBlockHeaderInfoByHash(hash string) entity.BlockHeaderInfo {
	paramsSlice := []interface{}{hash}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETBLOCKHEADER, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列华操作
	blockHeadInfo := entity.BlockHeaderInfo{}

	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockHeadInfo.Height = resBytes["height"].(float64)
		blockHeadInfo.Mediantime = resBytes["mediantime"].(float64)
		blockHeadInfo.Chainwork = resBytes["chainwork"].(string)
		blockHeadInfo.NTx = resBytes["nTx"].(float64)
		blockHeadInfo.Previousblockhash = resBytes["previousblockhash"].(string)
		blockHeadInfo.Hash = resBytes["hash"].(string)
		blockHeadInfo.Version = resBytes["version"].(float64)
		blockHeadInfo.Merkleroot = resBytes["merkleroot"].(string)
		blockHeadInfo.Time = resBytes["time"].(float64)
		blockHeadInfo.Difficulty = resBytes["difficulty"].(float64)
		blockHeadInfo.Confirmations = resBytes["confirmations"].(float64)
		blockHeadInfo.VersionHex = resBytes["versionHex"].(string)
		blockHeadInfo.Nonce = resBytes["nonce"].(float64)
		blockHeadInfo.Bits = resBytes["bits"].(string)
	}

	return blockHeadInfo
}

//根据区块Height获取区块头信息
func (bc btcSer) GetBlockHeaderInfoByHeight(height float64) entity.BlockHeaderInfo {
	blockHeaderInfo := entity.BlockHeaderInfo{}
	if float64(height) > bc.GetBlockCount() {
		return blockHeaderInfo
	}

	paramsSlice := []interface{}{height}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	hash, ok := rpcResult.Data.Result.(string)
	if ok {
		return bc.GetBlockHeaderInfoByHash(hash)
	}

	return blockHeaderInfo
}

//根据区块高度获取区块状态
func (bc btcSer) GetBlockStatsInfoByHeight(height float64) entity.BlockStats {
	blockStats := entity.BlockStats{}
	if height > bc.GetBlockCount() {
		return blockStats
	}

	paramsSlice := []interface{}{height}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON("getblockstats", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockStats.Subsidy = res["subsidy"].(float64)
		blockStats.Total_weight = res["total_weight"].(float64)
		blockStats.Height = res["height"].(float64)
		blockStats.Medianfee = res["medianfee"].(float64)
		blockStats.Mediantxsize = res["mediantxsize"].(float64)
		blockStats.Totalfee = res["totalfee"].(float64)
		blockStats.Avgfee = res["avgfee"].(float64)
		blockStats.Avgfeerate = res["avgfeerate"].(float64)
		blockStats.Minfee = res["minfee"].(float64)
		blockStats.Swtxs = res["swtxs"].(float64)
		blockStats.Time = res["time"].(float64)
		blockStats.Total_size = res["total_size"].(float64)
		blockStats.Blockhash = res["blockhash"].(string)
		blockStats.Feerate_percentiles_, ok = res["feerate_percentiles"].([]interface{})
		if ok {
			for i := 0; i < len(blockStats.Feerate_percentiles_); i++ {
				value, ok := blockStats.Feerate_percentiles_[i].(float64)
				if ok {
					blockStats.Feerate_percentiles = append(blockStats.Feerate_percentiles, value)
				}
			}
		}

		blockStats.Ins = res["ins"].(float64)
		blockStats.Minfeerate = res["minfeerate"].(float64)
		blockStats.Avgtxsize = res["avgtxsize"].(float64)
		blockStats.Maxfeerate = res["maxfeerate"].(float64)
		blockStats.Mintxsize = res["mintxsize"].(float64)
		blockStats.Outs = res["outs"].(float64)
		blockStats.Utxo_increase = res["utxo_increase"].(float64)
		blockStats.Mediantime = res["mediantime"].(float64)
		blockStats.Swtotal_weight = res["swtotal_weight"].(float64)
		blockStats.Txs = res["txs"].(float64)
		blockStats.Utxo_size_inc = res["utxo_size_inc"].(float64)
		blockStats.Maxfee = res["maxfee"].(float64)
		blockStats.Swtotal_size = res["swtotal_size"].(float64)
		blockStats.Total_out = res["total_out"].(float64)
	}

	return blockStats
}

//根据区块Hash获取区块状态
func (bc btcSer) GetBlockStatsInfoByHash(hash string) entity.BlockStats {
	blockStats := entity.BlockStats{}
	if len(hash) != 64 || hash[0] != 48 {
		return blockStats
	}

	var height float64 = -1
	if hash == bc.GetBlockHash(0) {
		return bc.GetBlockStatsInfoByHeight(0)
	}

	blockInfo := bc.GetBlockInfoByHash(hash)
	height = blockInfo.Height

	if height > 0 {
		return bc.GetBlockStatsInfoByHeight(height)
	}

	return blockStats
}

//获取区块链Tip信息
func (bc btcSer) GetChainTips() entity.ChainTips {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETCHAINTIPS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	chainTips := entity.ChainTips{}

	res, ok := rpcResult.Data.Result.([]interface{})
	if ok {
		for i := 0; i < len(res); i++ {
			var tip entity.Tip
			m, ok := res[i].(map[string]interface{})
			if ok {
				tip.Height = m["height"].(float64)
				tip.Hash = m["hash"].(string)
				tip.Branchlen = m["branchlen"].(float64)
				tip.Status = m["status"].(string)
			}
			chainTips.Tips = append(chainTips.Tips, tip)
		}
	}

	return chainTips
}

//获取区块链的交易状态
func (bc btcSer) GetChainTxStats() entity.ChainTxStats {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETCHAINTXSTATS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	chainTxStats := entity.ChainTxStats{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		chainTxStats.Window_final_block_hash = res["window_final_block_hash"].(string)
		chainTxStats.Window_final_block_height = res["window_final_block_height"].(float64)
		chainTxStats.Window_block_count = res["window_block_count"].(float64)
		chainTxStats.Window_tx_count = res["window_tx_count"].(float64)
		chainTxStats.Window_interval = res["window_interval"].(float64)
		chainTxStats.Txrate = res["txrate"].(float64)
		chainTxStats.Time = res["time"].(float64)
		chainTxStats.Txcount = res["txcount"].(float64)
	}

	return chainTxStats
}

//获取当前挖矿难度,成功返回挖矿难度,否则返回-1
func (bc btcSer) GetDifficulty() float64 {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETDIFFICULTY, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -0
}

//获取回收内存信息
func (bc btcSer) GetMempoolInfo() entity.MempoolInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON("getmempoolinfo", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	mempoolInfo := entity.MempoolInfo{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		mempoolInfo.Loaded = res["loaded"].(bool)
		mempoolInfo.Size = res["size"].(float64)
		mempoolInfo.Bytes = res["bytes"].(float64)
		mempoolInfo.Usage = res["usage"].(float64)
		mempoolInfo.Maxmempool = res["maxmempool"].(float64)
		mempoolInfo.Mempoolminfee = res["mempoolminfee"].(float64)
		mempoolInfo.Minrelaytxfee = res["minrelaytxfee"].(float64)
	}

	return mempoolInfo
}

//==============2020.12.22============//
//===============Wallet===============//
//============返回是否停止重新扫描==============
func (bc btcSer) Abortrescan() bool {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(ABORTRESCAN, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(bool)
	if ok {
		return res
	}

	return false
}

//==========添加多重账户地址===========
func (bc btcSer) AddMultisigAddress(nrequried int64,keys []string)entity.AddMultisgAddressInfo {
	paramsSlice := []interface{}{nrequried,keys}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(ADDMULTISIGADDRESS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
	multisigAddress := entity.AddMultisgAddressInfo{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		multisigAddress.Address = res["address"].(string)
		multisigAddress.RedeemScript = res["redemscript"].(string)
		multisigAddress.Descriptor = res["descriptor"].(string)
	}
	return multisigAddress
}
//=============转至备份钱包==========///
//func (bc btcSer) BackUpWallet(destination string)entity.UpWallet {
//	paramsSlice := []interface{}{destination}
//	//RPC通信标椎格JSON式数据
//	rpcNormJson := PrepareJSON(ABORTRESCAN, paramsSlice)
//
//	//bitcoin Core 响应的结果
//	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
//
//	res, ok := rpcResult.Data.Result.()
//	if ok {
//		return null
//	}
//
//	return -1
//}
//============查询撞的费用==========//
func (bc btcSer) BumpFee(txId string)entity.Bumpfee {
	paramsSlice := []interface{}{txId}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(BUMPFEE, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
	bumpfeeIfo := entity.Bumpfee{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		bumpfeeIfo.Psbt = res["psbt"].(string)
		bumpfeeIfo.TxId = res["redemscript"].(string)
		bumpfeeIfo.Origfee = res["origfee"].(int64)
		bumpfeeIfo.Fee = res["origfee"].(int64)
		bumpfeeIfo.Errors = res["errorx"].([]string)
		bumpfeeIfo.Str = res["str"].([]string)
	}
	return bumpfeeIfo
}
//=============创建钱包==============//
func (bc btcSer) CreateWallet(wallet_name string,passphrase string)entity.Createwallet {
	paramsSlice := []interface{}{wallet_name,passphrase}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(CREATEWALLET, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
		createWallet := entity.Createwallet{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		createWallet.Name = res["name"].(string)
		createWallet.Warning = res["warning"].(string)
	}
	return createWallet
}
//==============转储私钥============//
func (bc btcSer) DumpPrivkey(adress string)string{
	paramsSlice := []interface{}{adress}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(DUMPPRIVKEY, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res

	}
	return ""
}
//==============转储钱包============//
func (bc btcSer) DumpWallet(filename string)entity.Dumpwallet {
	paramsSlice := []interface{}{filename}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON("dumpwallet", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
	dumpWallet := entity.Dumpwallet{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		dumpWallet.Filename = res["filename"].(string)

	}
	return dumpWallet
}
//===============加密钱包============//
func (bc btcSer) EncyptWallet(passphrase string)entity.EncyptWallet {
	paramsSlice := []interface{}{passphrase}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON("encryptwallet", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
	encryptWallet := entity.EncyptWallet{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		encryptWallet.Passphrase = res["str"].(string)

	}
	return encryptWallet
}
//===============通过标签获取地址============//
func (bc btcSer) GetAddressesByLabel(label string)entity.Getaddressesbylabel {
	paramsSlice := []interface{}{label}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON("getaddressesbylabel", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
	getAddressesByLabel := entity.Getaddressesbylabel{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		getAddressesByLabel.Address = res["address"].(string)
		getAddressesByLabel.Purpose = res["purpose"].(string)

	}
	return getAddressesByLabel
}
//===============获取地址信息============//
func (bc btcSer) GetAddressinfo(address string)entity.Getaddressinfo{
	paramsSlice := []interface{}{address}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON("getaddressinfo", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
	getAddressInfo := entity.Getaddressinfo{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		getAddressInfo.Address = res["address"].(string)
		getAddressInfo.ScriptPubKey = res["scriptPubKey"].(string)
		getAddressInfo.Ismine = res["ismine"].(bool)
		getAddressInfo.Solvable = res["solvable"].(bool)
		getAddressInfo.Desc = res["desc"].(string)
		getAddressInfo.Iswatchonly = res["iswatchonly"].(bool)
		getAddressInfo.Isscript = res["isscript"].(bool)
		getAddressInfo.Iswitness = res["iswitness"].(bool)
		getAddressInfo.Script = res["script"].(string)
		getAddressInfo.Hex = res["hex"].(string)
		getAddressInfo.Pubkey = res["pubkey"].(string)
		getAddressInfo.Ischange = res["ischange"].(bool)
		getAddressInfo.Timestamp = res["timestamp"].(int64)
		getAddressInfo.Hdkeypath = res["hdkeypath"].(string)
		getAddressInfo.Hdseedid = res["hdseedid"].(string)
		getAddressInfo.Hdmasterfingerprint = res["hdmasterfingerprint"].(string)
		getAddressInfo.Labels = res["labels"].([]string)

		getAddressInfo.Embedded_, ok = res["Embedded"].(map[string]interface{})
		if ok {
			getAddressInfo.Embedded.Isscript = res["isscript"].(bool)
			getAddressInfo.Embedded.Iswitness = res["iswitness"].(bool)
			getAddressInfo.Embedded.Witness_version = res["witness_version"].(int64)
			getAddressInfo.Embedded.Witness_program = res["witness_program"].(string)
			getAddressInfo.Embedded.Pubkey = res["pubkey"].(string)
			getAddressInfo.Embedded.Address = res["address"].(string)
			getAddressInfo.Embedded.ScriptPubKey = res["scriptpubkey"].(string)
		}

	}
	return getAddressInfo
}
//==================返回这个钱包收到的比特币总数==========//
func (bc btcSer) GetBalance()entity.Getbalance{
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(ABORTRESCAN, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
	getBalance := entity.Getbalance{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		getBalance.N = res["n"].(float64)
	}

	return getBalance
}
//===========返回一个BTC中所有余额的对象。============//
func (bc btcSer) GetBalances()entity.Getbalances{
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(ABORTRESCAN, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))
	getBalances := entity.Getbalances{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		getBalances.Mine_, ok = res["Mine"].(map[string]interface{})
		if ok {
			getBalances.Mine.Trusted = res["trusted"].(float64)
			getBalances.Mine.Untrusted_pending = res["untrusted_pending"].(float64)
			getBalances.Mine.Immature = res["immature"].(float64)
			getBalances.Mine.Used = res["used"].(float64)
		}
		getBalances.Watchonly_, ok = res["Watchonly"].(map[string]interface{})
		if ok {
			getBalances.Mine.Trusted = res["trusted"].(float64)
			getBalances.Mine.Untrusted_pending = res["untrusted_pending"].(float64)
			getBalances.Mine.Immature = res["immature"].(float64)
		}
	}

	return getBalances
}
//=========返回新地址============//
func (bc btcSer) GetNewAddress()string{
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON(GETNEWADDRESS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}

	return ""
}
//===========获取原始更改地址========//
