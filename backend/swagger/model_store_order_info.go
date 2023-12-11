/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type StoreOrderInfo struct {
	StoreId int64 `json:"store_id,omitempty"`

	StoreName string `json:"store_name,omitempty"`

	StoreShippingFee int64 `json:"store_shipping_fee,omitempty"`

	ProductOrder []ProductOrderInfo `json:"product_order,omitempty"`
}
