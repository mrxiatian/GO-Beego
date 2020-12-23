package main

import (
	"bitcoin/utils"
	"fmt"
)

func main() {
	//res := utils.GetBestBlockHahs()
	//fmt.Println("最高区块的Hash值:", res)
	//blockInfo := utils.GetBlockInfo(res)
	//fmt.Println("最高区块的信息:", blockInfo)

	//chainInfo := utils.GetBlockChainInfoTest()
	//fmt.Println(chainInfo.Softforks.Bip65)
	//fmt.Println(chainInfo.Softforks.Csv)
	//fmt.Println(chainInfo.Softforks.Segwit)
	//fmt.Println(chainInfo.Softforks.Bip34)
	//fmt.Println(chainInfo.Softforks.Bip66)
	//fmt.Println(chainInfo.Prune_target_size)

	//testInfo := utils.GetTest()
	//for i := 0; i < len(testInfo.Feerate_percentiles); i++ {
	//	fmt.Println(testInfo.Feerate_percentiles[i])
	//}
	//
	//str := "000000000000000000055e22b020acc7481293900258dbac75459c9332a93fc7"
	//fmt.Println(str[0])

	//utils.GetTest()
	bc :=utils.GetBC()
	a :=bc.AddMultisigAddress()
	fmt.Println(a)

}

/*
*
 */
