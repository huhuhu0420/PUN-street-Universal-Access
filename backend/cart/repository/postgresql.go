package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type postgresqlCartRepo struct {
	db *sql.DB
}

func NewPostgressqlCartRepo(db *sql.DB) domain.CartRepo {
	return &postgresqlCartRepo{db}
}

func (p *postgresqlCartRepo) GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error) {
	sqlStatement := `
	SELECT orders.store_id, orders.cart_id, orders.order_date, orders.total_price, orders.user_id AS customer_id, orders.status, 
	stores.name AS store_name, stores.picture AS store_picture, stores.rate AS store_rate
	FROM orders LEFT JOIN stores ON orders.store_id = stores.store_id WHERE orders.user_id = $1 AND orders.status = 6;
	`

	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	historyArray := &[]swagger.HistoryInfo{}
	for rows.Next() {
		history := &swagger.HistoryInfo{}
		err := rows.Scan(&history.StoreId, &history.CartId, &history.OrderDate,
			&history.TotalPrice, &history.CustomerId, &history.Status,
			&history.StoreName, &history.StorePicture, &history.StoreRate)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		*historyArray = append(*historyArray, *history)
	}

	return historyArray, nil
}

func (p *postgresqlCartRepo) GetRunOrderByID(ctx context.Context, id int64) (*[]swagger.RunOrderInfo, error) {
	sqlStatement := `
	SELECT orders.store_id, orders.cart_id, orders.user_id, orders.status, 
	stores.name AS store_name, stores.picture AS store_picture
	FROM orders LEFT JOIN stores ON orders.store_id = stores.store_id 
	WHERE orders.user_id = $1 AND 
	orders.status != 0 AND orders.status != 1 AND orders.status != 6;
	`

	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	runOrderArray := &[]swagger.RunOrderInfo{}
	for rows.Next() {
		runOrder := &swagger.RunOrderInfo{}
		err := rows.Scan(&runOrder.StoreId, &runOrder.CartId, &runOrder.UserId, &runOrder.Status,
			&runOrder.StoreName, &runOrder.StorePicture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		*runOrderArray = append(*runOrderArray, *runOrder)
	}

	return runOrderArray, nil
}

func (p *postgresqlCartRepo) DeleteOrder(ctx context.Context, customerId int64, storeId int64) error {
	sqlStatement := `
	DELETE FROM orders
	WHERE user_id = $1 AND 
	cart_id = ( SELECT current_cart_id FROM user_data WHERE user_id = $1) AND 
	store_id = $2;
	`
	_, err := p.db.Exec(sqlStatement, customerId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) IsExitsOrderByStoreCartId(ctx context.Context, customerId int64, storeId int64) (bool, error) {
	sqlStatement := `
	SELECT COUNT(*) > 0 FROM orders 
	WHERE user_id = $1 AND 
	cart_id = ( SELECT current_cart_id FROM user_data WHERE user_id = $1) AND 
	store_id = $2;
	`
	row := p.db.QueryRow(sqlStatement, customerId, storeId)

	exist := false
	err := row.Scan(&exist)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return exist, nil
}

func (p *postgresqlCartRepo) IsExitsCartByStoreCartId(ctx context.Context, customerId int64, storeId int64) (bool, error) {
	sqlStatement := `
	SELECT COUNT(*) > 0 FROM carts 
	WHERE customer_id = $1 AND 
	cart_id = ( SELECT current_cart_id FROM user_data WHERE user_id = $1) AND 
	store_id = $2;
	`
	row := p.db.QueryRow(sqlStatement, customerId, storeId)

	exist := false
	err := row.Scan(&exist)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return exist, nil
}

func (p *postgresqlCartRepo) DeleteProduct(ctx context.Context, customerId int64, productId int64) (int64, error) {
	sqlStatement := `
	DELETE FROM carts
	WHERE customer_id = $1 AND 
	cart_id = ( SELECT current_cart_id FROM user_data WHERE user_id = $1) AND 
	product_id = $2
	RETURNING store_id
	`
	row := p.db.QueryRow(sqlStatement, customerId, productId)

	var storeId int64
	err := row.Scan(&storeId)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return storeId, nil
}
