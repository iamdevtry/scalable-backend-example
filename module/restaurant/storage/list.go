package restaurantstorage

import (
	"context"
	"food-delivery-service/common"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

func (store sqlStore) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []restaurantmodel.Restaurant

	db := store.db

	if v := filter.OwnerId; v > 0 {
		db = db.Where("owner_id = ?", v)
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Count(&paging.Total).
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
