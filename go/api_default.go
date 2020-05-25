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
	//"encoding/json"
	"net/http"
	"strconv"
	"strings"
	//"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{service: s}
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"GetJobOffer",
			strings.ToUpper("Get"),
			"/job_offer",
			c.GetJobOffer,
		},
		{
			"GetOccupation",
			strings.ToUpper("Get"),
			"/occupation",
			c.GetOccupation,
		},
		{
			"GetTotalPopulation",
			strings.ToUpper("Get"),
			"/total_population",
			c.GetTotalPopulation,
		},
	}
}

// GetJobOffer - job_offer
func (c *DefaultApiController) GetJobOffer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	prefCode := query.Get("prefCode")
	year := query.Get("year")
	matter := query.Get("matter")
	class := query.Get("class")

	result, err := c.service.GetJobOffer(prefCode, year, matter, class)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// GetOccupation - occupation
func (c *DefaultApiController) GetOccupation(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	iscoCode := query.Get("iscoCode")
	result, err := c.service.GetOccupation(iscoCode)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// GetTotalPopulation - total_population
func (c *DefaultApiController) GetTotalPopulation(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	//stringからfloat32に変換
	sprefCode, _ := strconv.ParseFloat(query.Get("prefCode"), 32)
	var prefCode float32 = float32(sprefCode)

	scityCode, _ := strconv.ParseFloat(query.Get("cityCode"), 32)
	var cityCode float32 = float32(scityCode)

	result, err := c.service.GetTotalPopulation(prefCode, cityCode)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
