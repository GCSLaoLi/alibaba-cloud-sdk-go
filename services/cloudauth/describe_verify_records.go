package cloudauth

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

// DescribeVerifyRecords invokes the cloudauth.DescribeVerifyRecords API synchronously
// api document: https://help.aliyun.com/api/cloudauth/describeverifyrecords.html
func (client *Client) DescribeVerifyRecords(request *DescribeVerifyRecordsRequest) (response *DescribeVerifyRecordsResponse, err error) {
	response = CreateDescribeVerifyRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVerifyRecordsWithChan invokes the cloudauth.DescribeVerifyRecords API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describeverifyrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerifyRecordsWithChan(request *DescribeVerifyRecordsRequest) (<-chan *DescribeVerifyRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeVerifyRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVerifyRecords(request)
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

// DescribeVerifyRecordsWithCallback invokes the cloudauth.DescribeVerifyRecords API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describeverifyrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerifyRecordsWithCallback(request *DescribeVerifyRecordsRequest, callback func(response *DescribeVerifyRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVerifyRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeVerifyRecords(request)
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

// DescribeVerifyRecordsRequest is the request struct for api DescribeVerifyRecords
type DescribeVerifyRecordsRequest struct {
	*requests.RpcRequest
	StatusList  string           `position:"Query" name:"StatusList"`
	StartDate   string           `position:"Query" name:"StartDate"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	TotalCount  requests.Integer `position:"Query" name:"TotalCount"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	QueryId     string           `position:"Query" name:"QueryId"`
	BizType     string           `position:"Query" name:"BizType"`
	IdCardNum   string           `position:"Query" name:"IdCardNum"`
	EndDate     string           `position:"Query" name:"EndDate"`
	BizId       string           `position:"Query" name:"BizId"`
}

// DescribeVerifyRecordsResponse is the response struct for api DescribeVerifyRecords
type DescribeVerifyRecordsResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	TotalCount  int       `json:"TotalCount" xml:"TotalCount"`
	PageSize    int       `json:"PageSize" xml:"PageSize"`
	CurrentPage int       `json:"CurrentPage" xml:"CurrentPage"`
	QueryId     string    `json:"QueryId" xml:"QueryId"`
	RecordsList []Records `json:"RecordsList" xml:"RecordsList"`
}

// CreateDescribeVerifyRecordsRequest creates a request to invoke DescribeVerifyRecords API
func CreateDescribeVerifyRecordsRequest() (request *DescribeVerifyRecordsRequest) {
	request = &DescribeVerifyRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeVerifyRecords", "", "")
	return
}

// CreateDescribeVerifyRecordsResponse creates a response to parse from DescribeVerifyRecords response
func CreateDescribeVerifyRecordsResponse() (response *DescribeVerifyRecordsResponse) {
	response = &DescribeVerifyRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
