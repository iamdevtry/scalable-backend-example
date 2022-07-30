package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

func (store *sqlStore) DeleteRestaurant(
	ctx context.Context,
	cond map[string]interface{},
) error {
	if err := store.db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where(cond).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
