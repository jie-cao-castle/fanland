package manager

import (
	"fanland/common"
	dao "fanland/db"
	"fanland/db/converter"
	"fanland/model"
)

type CategoryManager struct {
	productCategoryDB *dao.ProductCategoryDB
	options           *common.ServerOptions
}

func (manager *CategoryManager) InitManager(options *common.ServerOptions) {
	manager.options = options
	manager.productCategoryDB.InitDB(options.DbName)
}

func (manager *CategoryManager) GetCategories(offset uint64, limit uint64) ([]*model.ProductCategory, error) {
	manager.productCategoryDB.Open()
	defer manager.productCategoryDB.Close()
	categoryDOs, err := manager.productCategoryDB.GetList(offset, limit)
	if err != nil {
		return nil, err
	}

	var categories []*model.ProductCategory
	for _, categoryDO := range categoryDOs {
		category := converter.ConvertToProductCategory(categoryDO)
		categories = append(categories, category)
	}

	return categories, nil
}
