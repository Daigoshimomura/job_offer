/*
 * job_offer
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API. 
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// GetJobOffer - job_offer
func (s *DefaultApiService) GetJobOffer(prefCode float32, year float32, matter float32, class float32) (interface{}, error) {
	// TODO - update GetJobOffer with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetJobOffer' not implemented")
}

// GetOccupation - occupation
func (s *DefaultApiService) GetOccupation(iscoCode string) (interface{}, error) {
	// TODO - update GetOccupation with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetOccupation' not implemented")
}

// GetTotalPopulation - total_population
func (s *DefaultApiService) GetTotalPopulation(prefCode float32, cityCode float32) (interface{}, error) {
	// TODO - update GetTotalPopulation with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetTotalPopulation' not implemented")
}
