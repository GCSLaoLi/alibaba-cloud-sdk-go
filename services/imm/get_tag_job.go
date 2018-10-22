package imm

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

// GetTagJob invokes the imm.GetTagJob API synchronously
// api document: https://help.aliyun.com/api/imm/gettagjob.html
func (client *Client) GetTagJob(request *GetTagJobRequest) (response *GetTagJobResponse, err error) {
	response = CreateGetTagJobResponse()
	err = client.DoAction(request, response)
	return
}

// GetTagJobWithChan invokes the imm.GetTagJob API asynchronously
// api document: https://help.aliyun.com/api/imm/gettagjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTagJobWithChan(request *GetTagJobRequest) (<-chan *GetTagJobResponse, <-chan error) {
	responseChan := make(chan *GetTagJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTagJob(request)
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

// GetTagJobWithCallback invokes the imm.GetTagJob API asynchronously
// api document: https://help.aliyun.com/api/imm/gettagjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTagJobWithCallback(request *GetTagJobRequest, callback func(response *GetTagJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTagJobResponse
		var err error
		defer close(result)
		response, err = client.GetTagJob(request)
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

// GetTagJobRequest is the request struct for api GetTagJob
type GetTagJobRequest struct {
	*requests.RpcRequest
	JobId   string `position:"Query" name:"JobId"`
	Project string `position:"Query" name:"Project"`
}

// GetTagJobResponse is the response struct for api GetTagJob
type GetTagJobResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	JobId      string `json:"JobId" xml:"JobId"`
	SetId      string `json:"SetId" xml:"SetId"`
	SrcUri     string `json:"SrcUri" xml:"SrcUri"`
	Status     string `json:"Status" xml:"Status"`
	Percent    int    `json:"Percent" xml:"Percent"`
	CreateTime string `json:"CreateTime" xml:"CreateTime"`
	FinishTime string `json:"FinishTime" xml:"FinishTime"`
}

// CreateGetTagJobRequest creates a request to invoke GetTagJob API
func CreateGetTagJobRequest() (request *GetTagJobRequest) {
	request = &GetTagJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetTagJob", "imm", "openAPI")
	return
}

// CreateGetTagJobResponse creates a response to parse from GetTagJob response
func CreateGetTagJobResponse() (response *GetTagJobResponse) {
	response = &GetTagJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
