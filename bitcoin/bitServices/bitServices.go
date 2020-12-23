package bitServices

import "bitcoin/entity"

//区块链
type BlockChahin interface {
	//获取最高区块Hash
	GetBestBlockHahs() string

	//根据区块Hash获取区块信息
	GetBlockInfoByHash(blockHash string) entity.BlockInfo

	//根据区块hash获取区块Info
	GetBlockByHeight(height int) entity.BlockInfo

	//获取区块链信息
	GetBlockChainInfo() entity.BlockChainInfo

	//获取网络信息
	GetNetWorkInfo() entity.NetWorkInfo

	//获取区块总数
	GetBlockCount() float64

	//根据区块高度获取区块hash
	GetBlockHash(height int) string

	//根据区块Hash获取区块头信息
	GetBlockHeaderInfoByHash(hash string) entity.BlockHeaderInfo

	//根据区块Height获取区块头信息
	GetBlockHeaderInfoByHeight(height float64) entity.BlockHeaderInfo

	//根据区块高度获取区块状态
	GetBlockStatsInfoByHeight(height float64) entity.BlockStats

	//根据区块Hash获取区块状态
	GetBlockStatsInfoByHash(hash string) entity.BlockStats

	//获取区块链Tips
	GetChainTips() entity.ChainTips

	//获取区块链的交易状态
	GetChainTxStats() entity.ChainTxStats

	//获取当前挖矿难度
	GetDifficulty() float64

	//根据TxId获取到内存池的祖先
	GetMempoolAncestors(txId string) entity.MempoolAncestorsInfo

	//根据TxId获取到内存池的后代
	GetMempoolDescendants(txId string) entity.MempoolDescendantsInfo

	//根据TxId获取内存池数据
	GetMempoolEntry(txId string) entity.MempoolEntryInfo

	//获取回收内存信息
	GetMempoolInfo() entity.MempoolInfo
}


//一个人完成以下四个接口
//控制
type Control interface {

}

//生产
type Generating interface {

}

//矿工
type Mining interface {

}

//网络
type Network interface {

}


//一个人完成以下两个接口
//原始交易
type Rawtransactions interface {

}

//工具
type Util interface {

}


//两人合作完成下面一个接口
//钱包
type Wallet interface {
	//==============放弃交易=================
	Abandontransaction(txId string) entity.AbandontranSactionInfo
	//===========中止重新扫描===========!!!直接返回
	Abortrescan() entity.AbortRescan
	//============添加多重账户地址=======//
	AddMultisigAddress(nrequried int64,keys []string)entity.AddMultisgAddressInfo
	//=============转至备份钱包==========///
	BackUpWallet(destination string)entity.UpWallet
	//============查询撞的费用==========//
	BumpFee(txId string)entity.Bumpfee
	//=============创建钱包==============//
	CreateWallet(wallet_name string,passphrase string)entity.Createwallet
	//==============转储私钥============//
	DumpPrivkey(adress string)entity.Dumpprivkey
	//==============转储钱包============//
	DumpWallet(filename string)entity.Dumpwallet
	//===============加密钱包============//
	EncyptWallet(passphrase string)entity.EncyptWallet
	//===============通过标签获取地址============//
	GetAddressesByLabel(label string)entity.Getaddressesbylabel
	//===============获取地址信息============//
	GetAddressinfo(address string)entity.Getaddressinfo
	//==================返回这个钱包收到的比特币总数==========//
	GetBalance()entity.Getbalance
	//===========返回一个BTC中所有余额的对象。============//
	GetBalances()entity.Getbalances
	//=========返回新地址============//
	GetNewAddress()
	//===========获取原始更改地址========//
	GetRawChangeAddress()
}


//消息队列
type Zmq interface {

}
