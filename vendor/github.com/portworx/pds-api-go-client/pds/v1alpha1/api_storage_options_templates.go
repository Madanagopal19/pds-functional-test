/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// StorageOptionsTemplatesApiService StorageOptionsTemplatesApi service
type StorageOptionsTemplatesApiService service

type ApiApiStorageOptionsTemplatesIdDeleteRequest struct {
	ctx context.Context
	ApiService *StorageOptionsTemplatesApiService
	id string
}


func (r ApiApiStorageOptionsTemplatesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiStorageOptionsTemplatesIdDeleteExecute(r)
}

/*
ApiStorageOptionsTemplatesIdDelete Delete StorageOptionsTemplates

Removes a single StorageOptionsTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id StorageOptionsTemplate ID (must be valid UUID)
 @return ApiApiStorageOptionsTemplatesIdDeleteRequest
*/
func (a *StorageOptionsTemplatesApiService) ApiStorageOptionsTemplatesIdDelete(ctx context.Context, id string) ApiApiStorageOptionsTemplatesIdDeleteRequest {
	return ApiApiStorageOptionsTemplatesIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *StorageOptionsTemplatesApiService) ApiStorageOptionsTemplatesIdDeleteExecute(r ApiApiStorageOptionsTemplatesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageOptionsTemplatesApiService.ApiStorageOptionsTemplatesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/storage-options-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiStorageOptionsTemplatesIdGetRequest struct {
	ctx context.Context
	ApiService *StorageOptionsTemplatesApiService
	id string
}


func (r ApiApiStorageOptionsTemplatesIdGetRequest) Execute() (*ModelsStorageOptionsTemplate, *http.Response, error) {
	return r.ApiService.ApiStorageOptionsTemplatesIdGetExecute(r)
}

/*
ApiStorageOptionsTemplatesIdGet Get StorageOptionsTemplate

Fetches a single StorageOptionsTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id StorageOptionsTemplate ID (must be valid UUID)
 @return ApiApiStorageOptionsTemplatesIdGetRequest
*/
func (a *StorageOptionsTemplatesApiService) ApiStorageOptionsTemplatesIdGet(ctx context.Context, id string) ApiApiStorageOptionsTemplatesIdGetRequest {
	return ApiApiStorageOptionsTemplatesIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsStorageOptionsTemplate
func (a *StorageOptionsTemplatesApiService) ApiStorageOptionsTemplatesIdGetExecute(r ApiApiStorageOptionsTemplatesIdGetRequest) (*ModelsStorageOptionsTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsStorageOptionsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageOptionsTemplatesApiService.ApiStorageOptionsTemplatesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/storage-options-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiStorageOptionsTemplatesIdPutRequest struct {
	ctx context.Context
	ApiService *StorageOptionsTemplatesApiService
	id string
	body *ControllersUpdateStorageOptionsTemplateRequest
}

// Request body containing updated template
func (r ApiApiStorageOptionsTemplatesIdPutRequest) Body(body ControllersUpdateStorageOptionsTemplateRequest) ApiApiStorageOptionsTemplatesIdPutRequest {
	r.body = &body
	return r
}

func (r ApiApiStorageOptionsTemplatesIdPutRequest) Execute() (*ModelsStorageOptionsTemplate, *http.Response, error) {
	return r.ApiService.ApiStorageOptionsTemplatesIdPutExecute(r)
}

/*
ApiStorageOptionsTemplatesIdPut Update StorageOptionsTemplate

Updates existing StorageOptionsTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id StorageOptionsTemplate ID (must be valid UUID)
 @return ApiApiStorageOptionsTemplatesIdPutRequest
*/
func (a *StorageOptionsTemplatesApiService) ApiStorageOptionsTemplatesIdPut(ctx context.Context, id string) ApiApiStorageOptionsTemplatesIdPutRequest {
	return ApiApiStorageOptionsTemplatesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsStorageOptionsTemplate
func (a *StorageOptionsTemplatesApiService) ApiStorageOptionsTemplatesIdPutExecute(r ApiApiStorageOptionsTemplatesIdPutRequest) (*ModelsStorageOptionsTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsStorageOptionsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageOptionsTemplatesApiService.ApiStorageOptionsTemplatesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/storage-options-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiTenantsIdStorageOptionsTemplatesGetRequest struct {
	ctx context.Context
	ApiService *StorageOptionsTemplatesApiService
	id string
	sortBy *string
	limit *string
	continuation *string
	id2 *string
	name *string
}

// A given StorageOptionsTemplates attribute to sort results by (one of: id, name, created_at)
func (r ApiApiTenantsIdStorageOptionsTemplatesGetRequest) SortBy(sortBy string) ApiApiTenantsIdStorageOptionsTemplatesGetRequest {
	r.sortBy = &sortBy
	return r
}
// Maximum number of rows to return (could be less)
func (r ApiApiTenantsIdStorageOptionsTemplatesGetRequest) Limit(limit string) ApiApiTenantsIdStorageOptionsTemplatesGetRequest {
	r.limit = &limit
	return r
}
// Use a token returned by a previous query to continue listing with the next batch of rows
func (r ApiApiTenantsIdStorageOptionsTemplatesGetRequest) Continuation(continuation string) ApiApiTenantsIdStorageOptionsTemplatesGetRequest {
	r.continuation = &continuation
	return r
}
// Filter results by StorageOptionsTemplates id
func (r ApiApiTenantsIdStorageOptionsTemplatesGetRequest) Id2(id2 string) ApiApiTenantsIdStorageOptionsTemplatesGetRequest {
	r.id2 = &id2
	return r
}
// Filter results by StorageOptionsTemplates name
func (r ApiApiTenantsIdStorageOptionsTemplatesGetRequest) Name(name string) ApiApiTenantsIdStorageOptionsTemplatesGetRequest {
	r.name = &name
	return r
}

func (r ApiApiTenantsIdStorageOptionsTemplatesGetRequest) Execute() (*ControllersPaginatedStorageOptionsTemplates, *http.Response, error) {
	return r.ApiService.ApiTenantsIdStorageOptionsTemplatesGetExecute(r)
}

/*
ApiTenantsIdStorageOptionsTemplatesGet List StorageOptionsTemplates

Lists StorageOptionsTemplates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdStorageOptionsTemplatesGetRequest
*/
func (a *StorageOptionsTemplatesApiService) ApiTenantsIdStorageOptionsTemplatesGet(ctx context.Context, id string) ApiApiTenantsIdStorageOptionsTemplatesGetRequest {
	return ApiApiTenantsIdStorageOptionsTemplatesGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ControllersPaginatedStorageOptionsTemplates
func (a *StorageOptionsTemplatesApiService) ApiTenantsIdStorageOptionsTemplatesGetExecute(r ApiApiTenantsIdStorageOptionsTemplatesGetRequest) (*ControllersPaginatedStorageOptionsTemplates, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControllersPaginatedStorageOptionsTemplates
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageOptionsTemplatesApiService.ApiTenantsIdStorageOptionsTemplatesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/storage-options-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.continuation != nil {
		localVarQueryParams.Add("continuation", parameterToString(*r.continuation, ""))
	}
	if r.id2 != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id2, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiTenantsIdStorageOptionsTemplatesPostRequest struct {
	ctx context.Context
	ApiService *StorageOptionsTemplatesApiService
	id string
	body *ControllersCreateStorageOptionsTemplatesRequest
}

// Request body containing the storage options template
func (r ApiApiTenantsIdStorageOptionsTemplatesPostRequest) Body(body ControllersCreateStorageOptionsTemplatesRequest) ApiApiTenantsIdStorageOptionsTemplatesPostRequest {
	r.body = &body
	return r
}

func (r ApiApiTenantsIdStorageOptionsTemplatesPostRequest) Execute() (*ModelsStorageOptionsTemplate, *http.Response, error) {
	return r.ApiService.ApiTenantsIdStorageOptionsTemplatesPostExecute(r)
}

/*
ApiTenantsIdStorageOptionsTemplatesPost Create StorageOptionsTemplates

Creates a new StorageOptionsTemplates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdStorageOptionsTemplatesPostRequest
*/
func (a *StorageOptionsTemplatesApiService) ApiTenantsIdStorageOptionsTemplatesPost(ctx context.Context, id string) ApiApiTenantsIdStorageOptionsTemplatesPostRequest {
	return ApiApiTenantsIdStorageOptionsTemplatesPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsStorageOptionsTemplate
func (a *StorageOptionsTemplatesApiService) ApiTenantsIdStorageOptionsTemplatesPostExecute(r ApiApiTenantsIdStorageOptionsTemplatesPostRequest) (*ModelsStorageOptionsTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsStorageOptionsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageOptionsTemplatesApiService.ApiTenantsIdStorageOptionsTemplatesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/storage-options-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
