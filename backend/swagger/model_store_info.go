/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type StoreInfo struct {
	StoreId int64 `json:"store_id"`

	Description string `json:"description"`

	Name string `json:"name"`

	Address string `json:"address"`

	Email string `json:"email"`

	Rate float32 `json:"rate"`

	RateCount int64 `json:"rate_count"`

	Picture string `json:"picture"`

	Status int64 `json:"status"`
}
