/*
 * job_offer
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type NestTotalPopulation struct {

	// 実績値と推計値の区切り年
	BoundaryYear float32 `json:"boundaryYear,omitempty"`

	Data Nest1TotalPopulation `json:"data,omitempty"`
}
