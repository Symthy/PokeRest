/*
 * PokeRest
 *
 * PokeRest API
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TypeCompatibility struct {

	CompatibilityTable [][]float32 `json:"compatibilityTable"`

	// テーブル行列のタイプ順序
	TypeOrder []string `json:"typeOrder"`
}