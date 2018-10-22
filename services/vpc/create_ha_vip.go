package vpc

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

// CreateHaVip invokes the vpc.CreateHaVip API synchronously
// api document: https://help.aliyun.com/api/vpc/createhavip.html
func (client *Client) CreateHaVip(request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
	response = CreateCreateHaVipResponse()
	err = client.DoAction(request, response)
	return
}

// CreateHaVipWithChan invokes the vpc.CreateHaVip API asynchronously
// api document: https://help.aliyun.com/api/vpc/createhavip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateHaVipWithChan(request *CreateHaVipRequest) (<-chan *CreateHaVipResponse, <-chan error) {
	responseChan := make(chan *CreateHaVipResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateHaVip(request)
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

// CreateHaVipWithCallback invokes the vpc.CreateHaVip API asynchronously
// api document: https://help.aliyun.com/api/vpc/createhavip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateHaVipWithCallback(request *CreateHaVipRequest, callback func(response *CreateHaVipResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateHaVipResponse
		var err error
		defer close(result)
		response, err = client.CreateHaVip(request)
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

// CreateHaVipRequest is the request struct for api CreateHaVip
type CreateHaVipRequest struct {
	*requests.RpcRequest
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	IpAddress            string           `position:"Query" name:"IpAddress"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateHaVipResponse is the response struct for api CreateHaVip
type CreateHaVipResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HaVipId   string `json:"HaVipId" xml:"HaVipId"`
}

// CreateCreateHaVipRequest creates a request to invoke CreateHaVip API
func CreateCreateHaVipRequest() (request *CreateHaVipRequest) {
	request = &CreateHaVipRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateHaVip", "vpc", "openAPI")
	return
}

// CreateCreateHaVipResponse creates a response to parse from CreateHaVip response
func CreateCreateHaVipResponse() (response *CreateHaVipResponse) {
	response = &CreateHaVipResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
