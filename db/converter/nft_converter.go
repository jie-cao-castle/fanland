package converter

import (
	"fanland/db/dao"
	"fanland/model"
	"fanland/service/request"
)

func ConvertToNftContractDO(contract *model.NftContract) *dao.NftContractDO {
	contractDO := &dao.NftContractDO{
		Id:              contract.Id,
		TokenName:       contract.TokenName,
		TokenSymbol:     contract.TokenSymbol,
		ChainId:         contract.ChainId,
		ChainCode:       contract.ChainCode,
		ContractAddress: contract.ContractAddress,
		Status:          contract.Status,
		ProductId:       contract.ProductId,
		TokenAmount:     contract.TokenAmount,
		NextTokenId:     contract.NextTokenId,
		CreateTime:      contract.CreateTime,
		UpdateTime:      contract.UpdateTime,
	}
	return contractDO
}
func ConvertToNftContract(contractDO *dao.NftContractDO) *model.NftContract {
	contract := &model.NftContract{
		Id:              contractDO.Id,
		TokenName:       contractDO.TokenName,
		TokenSymbol:     contractDO.TokenSymbol,
		ChainId:         contractDO.ChainId,
		ChainCode:       contractDO.ChainCode,
		ContractAddress: contractDO.ContractAddress,
		Status:          contractDO.Status,
		ProductId:       contractDO.ProductId,
		TokenAmount:     contractDO.TokenAmount,
		NextTokenId:     contractDO.NextTokenId,
		CreateTime:      contractDO.CreateTime,
		UpdateTime:      contractDO.UpdateTime,
	}
	return contract
}
func ConvertReqToNftContract(req *request.AddNftContractRequest) *model.NftContract {
	contract := &model.NftContract{
		TokenName:       req.TokenName,
		TokenSymbol:     req.TokenSymbol,
		ChainId:         req.ChainId,
		ChainCode:       req.ChainCode,
		ContractAddress: req.ContractAddress,
		Status:          req.Status,
		ProductId:       req.ProductId,
		TokenAmount:     req.TokenAmount,
		NextTokenId:     req.NextTokenId,
	}
	return contract
}

func ConvertReqToUpdateNftContract(req *request.UpdateNftContractRequest) *model.NftContract {
	contract := &model.NftContract{
		Id:     req.Id,
		Status: req.Status,
	}
	return contract
}

func ConvertToNftOrderDO(nft *model.NftOrder) *dao.NftOrderDO {
	nftDO := &dao.NftOrderDO{
		Id:              nft.Id,
		Price:           nft.Price,
		PriceUnit:       nft.PriceUnit,
		ChainId:         nft.ProductId,
		ChainCode:       nft.ChainCode,
		NftKey:          nft.NftKey,
		Status:          nft.Status,
		ProductId:       nft.ProductId,
		TransactionHash: nft.TransactionHash,
		CreateTime:      nft.CreateTime,
		UpdateTime:      nft.UpdateTime,
		Amount:          nft.Amount,
		ToUserName:      nft.ToUserName,
		ToUserId:        nft.ToUserId,
	}
	return nftDO
}
func ConvertReqToNftOrder(req *request.AddNftOrderRequest) *model.NftOrder {
	nft := &model.NftOrder{
		Price:           req.Price,
		PriceUnit:       req.PriceUnit,
		ChainId:         req.ProductId,
		ChainCode:       req.ChainCode,
		NftKey:          req.NftKey,
		Status:          req.Status,
		ProductId:       req.ProductId,
		TransactionHash: req.TransactionHash,
		Amount:          req.Amount,
		ToUserId:        req.ToUserId,
	}
	return nft
}
func ConvertToNftOrder(nftDO *dao.NftOrderDO) *model.NftOrder {
	nft := &model.NftOrder{
		Id:              nftDO.Id,
		Price:           nftDO.Price,
		PriceUnit:       nftDO.PriceUnit,
		ChainId:         nftDO.ProductId,
		ChainCode:       nftDO.ChainCode,
		NftKey:          nftDO.NftKey,
		Status:          nftDO.Status,
		ProductId:       nftDO.ProductId,
		TransactionHash: nftDO.TransactionHash,
		CreateTime:      nftDO.CreateTime,
		UpdateTime:      nftDO.UpdateTime,
		Amount:          nftDO.Amount,
		ToUserName:      nftDO.ToUserName,
		ToUserId:        nftDO.ToUserId,
	}
	return nft
}

func ConvertReqToUpdateNftOrder(req *request.UpdateNftOrderRequest) *model.NftOrder {
	nft := &model.NftOrder{
		Id:              req.Id,
		Price:           req.Price,
		PriceUnit:       req.PriceUnit,
		ChainId:         req.ProductId,
		ChainCode:       req.ChainCode,
		NftKey:          req.NftKey,
		Status:          req.Status,
		ProductId:       req.ProductId,
		TransactionHash: req.TransactionHash,
		Amount:          req.Amount,
	}
	return nft
}
