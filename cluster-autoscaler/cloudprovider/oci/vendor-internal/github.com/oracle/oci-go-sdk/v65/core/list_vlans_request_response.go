// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVlansRequest wrapper for the ListVlans operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListVlans.go.html to see an example of how to use ListVlansRequest.
type ListVlansRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListVlansSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListVlansSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given lifecycle
	// state. The state value is case-insensitive.
	LifecycleState VlanLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVlansRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVlansRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVlansRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVlansRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVlansRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListVlansSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVlansSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVlansSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVlansSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVlanLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetVlanLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVlansResponse wrapper for the ListVlans operation
type ListVlansResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []Vlan instances
	Items []Vlan `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListVlansResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVlansResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVlansSortByEnum Enum with underlying type: string
type ListVlansSortByEnum string

// Set of constants representing the allowable values for ListVlansSortByEnum
const (
	ListVlansSortByTimecreated ListVlansSortByEnum = "TIMECREATED"
	ListVlansSortByDisplayname ListVlansSortByEnum = "DISPLAYNAME"
)

var mappingListVlansSortByEnum = map[string]ListVlansSortByEnum{
	"TIMECREATED": ListVlansSortByTimecreated,
	"DISPLAYNAME": ListVlansSortByDisplayname,
}

var mappingListVlansSortByEnumLowerCase = map[string]ListVlansSortByEnum{
	"timecreated": ListVlansSortByTimecreated,
	"displayname": ListVlansSortByDisplayname,
}

// GetListVlansSortByEnumValues Enumerates the set of values for ListVlansSortByEnum
func GetListVlansSortByEnumValues() []ListVlansSortByEnum {
	values := make([]ListVlansSortByEnum, 0)
	for _, v := range mappingListVlansSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVlansSortByEnumStringValues Enumerates the set of values in String for ListVlansSortByEnum
func GetListVlansSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListVlansSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVlansSortByEnum(val string) (ListVlansSortByEnum, bool) {
	enum, ok := mappingListVlansSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVlansSortOrderEnum Enum with underlying type: string
type ListVlansSortOrderEnum string

// Set of constants representing the allowable values for ListVlansSortOrderEnum
const (
	ListVlansSortOrderAsc  ListVlansSortOrderEnum = "ASC"
	ListVlansSortOrderDesc ListVlansSortOrderEnum = "DESC"
)

var mappingListVlansSortOrderEnum = map[string]ListVlansSortOrderEnum{
	"ASC":  ListVlansSortOrderAsc,
	"DESC": ListVlansSortOrderDesc,
}

var mappingListVlansSortOrderEnumLowerCase = map[string]ListVlansSortOrderEnum{
	"asc":  ListVlansSortOrderAsc,
	"desc": ListVlansSortOrderDesc,
}

// GetListVlansSortOrderEnumValues Enumerates the set of values for ListVlansSortOrderEnum
func GetListVlansSortOrderEnumValues() []ListVlansSortOrderEnum {
	values := make([]ListVlansSortOrderEnum, 0)
	for _, v := range mappingListVlansSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVlansSortOrderEnumStringValues Enumerates the set of values in String for ListVlansSortOrderEnum
func GetListVlansSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVlansSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVlansSortOrderEnum(val string) (ListVlansSortOrderEnum, bool) {
	enum, ok := mappingListVlansSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
