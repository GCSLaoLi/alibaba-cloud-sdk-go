package bssopenapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryRedeem invokes the bssopenapi.QueryRedeem API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryredeem.html
func (client *Client) QueryRedeem(request *QueryRedeemRequest) (response *QueryRedeemResponse, err error) {
	response = CreateQueryRedeemResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRedeemWithChan invokes the bssopenapi.QueryRedeem API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryredeem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRedeemWithChan(request *QueryRedeemRequest) (<-chan *QueryRedeemResponse, <-chan error) {
	responseChan := make(chan *QueryRedeemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRedeem(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryRedeemWithCallback invokes the bssopenapi.QueryRedeem API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryredeem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRedeemWithCallback(request *QueryRedeemRequest, callback func(response *QueryRedeemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRedeemResponse
		var err error
		defer close(result)
		response, err = client.QueryRedeem(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QueryRedeemRequest is the request struct for api QueryRedeem
type QueryRedeemRequest struct {
	*requests.RpcRequest
	ExpiryTimeEnd   string           `position:"Query" name:"ExpiryTimeEnd"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ExpiryTimeStart string           `position:"Query" name:"ExpiryTimeStart"`
	PageNum         requests.Integer `position:"Query" name:"PageNum"`
	EffectiveOrNot  requests.Boolean `position:"Query" name:"EffectiveOrNot"`
}

// QueryRedeemResponse is the response struct for api QueryRedeem
type QueryRedeemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryRedeemRequest creates a request to invoke QueryRedeem API
func CreateQueryRedeemRequest() (request *QueryRedeemRequest) {
	request = &QueryRedeemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryRedeem", "", "")
	return
}

// CreateQueryRedeemResponse creates a response to parse from QueryRedeem response
func CreateQueryRedeemResponse() (response *QueryRedeemResponse) {
	response = &QueryRedeemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}