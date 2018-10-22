package vod

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

// ListAIVideoCategoryJob invokes the vod.ListAIVideoCategoryJob API synchronously
// api document: https://help.aliyun.com/api/vod/listaivideocategoryjob.html
func (client *Client) ListAIVideoCategoryJob(request *ListAIVideoCategoryJobRequest) (response *ListAIVideoCategoryJobResponse, err error) {
	response = CreateListAIVideoCategoryJobResponse()
	err = client.DoAction(request, response)
	return
}

// ListAIVideoCategoryJobWithChan invokes the vod.ListAIVideoCategoryJob API asynchronously
// api document: https://help.aliyun.com/api/vod/listaivideocategoryjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAIVideoCategoryJobWithChan(request *ListAIVideoCategoryJobRequest) (<-chan *ListAIVideoCategoryJobResponse, <-chan error) {
	responseChan := make(chan *ListAIVideoCategoryJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAIVideoCategoryJob(request)
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

// ListAIVideoCategoryJobWithCallback invokes the vod.ListAIVideoCategoryJob API asynchronously
// api document: https://help.aliyun.com/api/vod/listaivideocategoryjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAIVideoCategoryJobWithCallback(request *ListAIVideoCategoryJobRequest, callback func(response *ListAIVideoCategoryJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAIVideoCategoryJobResponse
		var err error
		defer close(result)
		response, err = client.ListAIVideoCategoryJob(request)
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

// ListAIVideoCategoryJobRequest is the request struct for api ListAIVideoCategoryJob
type ListAIVideoCategoryJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       string `position:"Query" name:"ResourceOwnerId"`
	AIVideoCategoryJobIds string `position:"Query" name:"AIVideoCategoryJobIds"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               string `position:"Query" name:"OwnerId"`
}

// ListAIVideoCategoryJobResponse is the response struct for api ListAIVideoCategoryJob
type ListAIVideoCategoryJobResponse struct {
	*responses.BaseResponse
	RequestId                     string                        `json:"RequestId" xml:"RequestId"`
	NonExistAIVideoCategoryJobIds NonExistAIVideoCategoryJobIds `json:"NonExistAIVideoCategoryJobIds" xml:"NonExistAIVideoCategoryJobIds"`
	AIVideoCategoryJobList        AIVideoCategoryJobList        `json:"AIVideoCategoryJobList" xml:"AIVideoCategoryJobList"`
}

// CreateListAIVideoCategoryJobRequest creates a request to invoke ListAIVideoCategoryJob API
func CreateListAIVideoCategoryJobRequest() (request *ListAIVideoCategoryJobRequest) {
	request = &ListAIVideoCategoryJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoCategoryJob", "vod", "openAPI")
	return
}

// CreateListAIVideoCategoryJobResponse creates a response to parse from ListAIVideoCategoryJob response
func CreateListAIVideoCategoryJobResponse() (response *ListAIVideoCategoryJobResponse) {
	response = &ListAIVideoCategoryJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
