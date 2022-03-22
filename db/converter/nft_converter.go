package converter

import (
	"fanland/db/dao"
	"fanland/model"
	"fanland/service/request"
)

func ConvertToNftContractDO(nft *model.NftContract) *dao.NftContractDO {
	return nil
}
func ConvertToNftContract(nft *dao.NftContractDO) *model.NftContract {
	return nil
}
func ConvertReqToNftContract(req *request.AddNftContractRequest) *model.NftContract {
	return nil
}

func ConvertReqToUpdateNftContract(req *request.UpdateNftContractRequest) *model.NftContract {
	return nil
}

func ConvertToNftOrderDO(nft *model.NftOrder) *dao.NftOrderDO {
	return nil
}
func ConvertReqToNftOrder(nft *request.AddNftOrderRequest) *model.NftOrder {
	return nil
}
func ConvertToNftOrder(nftDO *dao.NftOrderDO) *model.NftOrder {
	return nil
}

func ConvertReqToUpdateNftOrder(req *request.UpdateNftOrderRequest) *model.NftOrder {
	return nil
}
