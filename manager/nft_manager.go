package manager

import (
	"fanland/common"
	dao "fanland/db"
	"fanland/db/converter"
	"fanland/model"
)

type NftManager struct {
	nftContractDB *dao.NftContractDB
	nftOrderDB    *dao.NftOrderDB
	options       *common.ServerOptions
}

func (manager *NftManager) InitManager(options *common.ServerOptions) {
	manager.options = options
	manager.nftContractDB = &dao.NftContractDB{}
	manager.nftOrderDB = &dao.NftOrderDB{}

	manager.nftContractDB.InitDB(options.DbName)
	manager.nftOrderDB.InitDB(options.DbName)
}

func (manager *NftManager) AddNFTContract(nft *model.NftContract) error {
	manager.nftContractDB.Open()
	defer manager.nftContractDB.Close()
	nftDO := converter.ConvertToNftContractDO(nft)
	if err := manager.nftContractDB.Insert(nftDO); err != nil {
		return err
	}

	return nil
}

func (manager *NftManager) UpdateNFTContract(nft *model.NftContract) error {
	manager.nftContractDB.Open()
	defer manager.nftContractDB.Close()
	nftDO := converter.ConvertToNftContractDO(nft)
	if err := manager.nftContractDB.Update(nftDO); err != nil {
		return err
	}

	return nil
}

func (manager *NftManager) GetProductContracts(productId uint64) ([]*model.NftContract, error) {
	manager.nftContractDB.Open()
	defer manager.nftContractDB.Close()
	contractDOs, err := manager.nftContractDB.GetListByProductId(productId)
	if err != nil {
		return nil, err
	}

	var contracts []*model.NftContract
	for _, contractDO := range contractDOs {
		contract := converter.ConvertToNftContract(contractDO)
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (manager *NftManager) AddNFTOrder(nftOrder *model.NftOrder) error {
	manager.nftOrderDB.Open()
	defer manager.nftOrderDB.Close()
	nftDO := converter.ConvertToNftOrderDO(nftOrder)
	if err := manager.nftOrderDB.Insert(nftDO); err != nil {
		return err
	}
	return nil
}

func (manager *NftManager) GetProductOrders(productId uint64) ([]*model.NftOrder, error) {
	manager.nftOrderDB.Open()
	defer manager.nftOrderDB.Close()
	orderDOs, err := manager.nftOrderDB.GetListByProductId(productId)
	if err != nil {
		return nil, err
	}

	var orders []*model.NftOrder
	for _, orderDO := range orderDOs {
		order := converter.ConvertToNftOrder(orderDO)
		orders = append(orders, order)
	}

	return orders, nil
}

func (manager *NftManager) UpdateNFTOrder(nftOrder *model.NftOrder) error {
	manager.nftOrderDB.Open()
	defer manager.nftOrderDB.Close()
	nftOrderDO := converter.ConvertToNftOrderDO(nftOrder)
	if err := manager.nftOrderDB.Update(nftOrderDO); err != nil {
		return err
	}

	return nil
}
