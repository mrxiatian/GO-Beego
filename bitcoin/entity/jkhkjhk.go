package entity

//====================Apply Wallet============//
//=============== begin:AbandontranSactionInfo ===============//
type AbandontranSactionInfo struct {
}
//=============== end:AbandontranSactionInfo ===============//

//============中止重新扫描==========
type AbortRescan struct {
}
//==========添加多重账户地址===========
type AddMultisgAddressInfo struct {
	Address string
	RedeemScript string
	Descriptor string
}
//=============转至备份钱包==========///

type  UpWallet struct {

}
//============查询撞的费用==========//
//!!!!!!!!!!!1
type Bumpfee struct {
	Psbt string
	TxId string
	Origfee int64
	Fee int64
	Errors []string//json array类型
	Str []string//可能为空
}
//=============创建钱包==============//
type Createwallet struct {
	Name string
	Warning string
}
//==============转储私钥============//
type Dumpprivkey struct {
	Str string  //The private key
}
//==============转储钱包============//
type Dumpwallet struct {
	Filename string
}
//===============加密钱包============//
type EncyptWallet struct {
	Passphrase string
}
//===============通过标签获取地址============//
type Getaddressesbylabel struct {
	Address string //!!!!!!(json object) json object with information about address
	Purpose string
}
type Getaddressinfo struct {
	Address string
	ScriptPubKey string
	Ismine bool
	Solvable bool
	Desc string
	Iswatchonly bool
	Isscript bool
	Iswitness bool
	Script string
	Hex string
	Pubkey string
	Ischange bool
	Timestamp int64
	Hdkeypath string
	Hdseedid string
	Hdmasterfingerprint string
	Labels	[]string
	Embedded Embedded
	Embedded_ map[string]interface{}
}
type Embedded struct {

	Isscript bool
	Iswitness bool
	Witness_version int64
	Witness_program string
	Pubkey string
	Address string
	ScriptPubKey string
}
//==============GetBalance==========//
type Getbalance struct {
	N     float64
}
//============end GetBalance=========//
//==============begin:getbalances======//
type Getbalances struct {
	Mine Mine
	Mine_  map[string]interface{}
	Watchonly Watchonly
	Watchonly_ map[string]interface{}
}
type Mine struct {
	Trusted float64
	Untrusted_pending float64
	Immature float64
	Used float64
}
type Watchonly struct {
	Trusted float64
	Untrusted_pending float64
	Immature float64
}
//========end: getbalances========//




