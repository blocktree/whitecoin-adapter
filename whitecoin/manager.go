/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package whitecoin

import (
	"fmt"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
	bt "github.com/blocktree/whitecoin-adapter/libs/types"
	"github.com/blocktree/whitecoin-adapter/whitecoin_addrdec"
)

type WalletManager struct {
	openwallet.AssetsAdapterBase

	Api             *WalletClient                   // 节点客户端
	Config          *WalletConfig                   // 节点配置
	Decoder         openwallet.AddressDecoder       //地址编码器
	DecoderV2       openwallet.AddressDecoderV2     //地址编码器V2
	TxDecoder       openwallet.TransactionDecoder   //交易单编码器
	Log             *log.OWLogger                   //日志工具
	ContractDecoder openwallet.SmartContractDecoder //智能合约解析器
	Blockscanner    openwallet.BlockScanner         //区块扫描器
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.Config = NewConfig(Symbol)
	wm.Blockscanner = NewBlockScanner(&wm)
	wm.Decoder = NewAddressDecoder(&wm)
	wm.DecoderV2 = whitecoin_addrdec.NewAddressDecoderV2()
	wm.TxDecoder = NewTransactionDecoder(&wm)
	wm.ContractDecoder = NewContractDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())
	return &wm
}

func (wm *WalletManager) GetRequiredFee(ops []bt.Operation, assetID string) ([]bt.AssetAmount, error) {
	resp := make([]bt.AssetAmount, 0)

	if assetID == "1.3.0" {
		//XWC写死1.3.0
		xwcFees := bt.AssetAmount{
			Asset:  bt.AssetIDFromObject(bt.NewAssetID("1.3.0")),
			Amount: bt.Int64(wm.Config.FixFees),
		}

		resp = append(resp, xwcFees)
	}

	if len(resp) == 0 {
		return nil, fmt.Errorf("can not find required fee with asset ID: %s", assetID)
	}

	return resp, nil
}
