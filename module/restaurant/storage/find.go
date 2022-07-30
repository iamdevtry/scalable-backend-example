package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

func (store *sqlStore) FindRestaurant(
	ctx context.Context,
	cond map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := store.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
