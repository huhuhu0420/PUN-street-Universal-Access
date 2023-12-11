/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type HistoryDetailStoreOrderInfo struct {
	StoreId int64 `json:"store_id,omitempty"`

	StoreName string `json:"store_name,omitempty"`

	StoreShippingFee int64 `json:"store_shipping_fee,omitempty"`

	ShippingDiscountId int64 `json:"shipping_discount_id,omitempty"`

	ShippingDiscountMaxPrice int64 `json:"shipping_discount_max_price,omitempty"`

	SeasoningDiscountId int64 `json:"seasoning_discount_id,omitempty"`

	SeasoningDiscountStartDate string `json:"seasoning_discount_start_date,omitempty"`

	SeasoningDiscountEndDate string `json:"seasoning_discount_end_date,omitempty"`

	SeasoningDiscountPercentage int64 `json:"seasoning_discount_percentage,omitempty"`

	ProductOrder []ProductOrderInfo `json:"product_order,omitempty"`
}