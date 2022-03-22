package manager

import (
	"fanland/common"
	dao "fanland/db"
	"fanland/db/converter"
	"fanland/model"
)

type NftManager struct {
	nftContractDB *dao.NftContractDB
	nftOrder      *dao.NftOrderDB
	options       *common.ServerOptions
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
	manager.nftOrder.Open()
	defer manager.nftOrder.Close()
	nftDO := converter.ConvertToNftOrderDO(nftOrder)
	if err := manager.nftOrder.Insert(nftDO); err != nil {
		return err
	}
	return nil
}

func (manager *NftManager) GetProductOrders(productId uint64) ([]*model.NftOrder, error) {
	manager.nftOrder.Open()
	defer manager.nftOrder.Close()
	orderDOs, err := manager.nftOrder.GetListByProductId(productId)
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
	manager.nftOrder.Open()
	defer manager.nftOrder.Close()
	nftOrderDO := converter.ConvertToNftOrderDO(nftOrder)
	if err := manager.nftOrder.Update(nftOrderDO); err != nil {
		return err
	}

	return nil
}
