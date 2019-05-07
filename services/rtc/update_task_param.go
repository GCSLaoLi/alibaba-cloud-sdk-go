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

package rtc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateTaskParam invokes the rtc.UpdateTaskParam API synchronously
// api document: https://help.aliyun.com/api/rtc/updatetaskparam.html
func (client *Client) UpdateTaskParam(request *UpdateTaskParamRequest) (response *UpdateTaskParamResponse, err error) {
	response = CreateUpdateTaskParamResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTaskParamWithChan invokes the rtc.UpdateTaskParam API asynchronously
// api document: https://help.aliyun.com/api/rtc/updatetaskparam.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateTaskParamWithChan(request *UpdateTaskParamRequest) (<-chan *UpdateTaskParamResponse, <-chan error) {
	responseChan := make(chan *UpdateTaskParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTaskParam(request)
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

// UpdateTaskParamWithCallback invokes the rtc.UpdateTaskParam API asynchronously
// api document: https://help.aliyun.com/api/rtc/updatetaskparam.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateTaskParamWithCallback(request *UpdateTaskParamRequest, callback func(response *UpdateTaskParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTaskParamResponse
		var err error
		defer close(result)
		response, err = client.UpdateTaskParam(request)
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

// UpdateTaskParamRequest is the request struct for api UpdateTaskParam
type UpdateTaskParamRequest struct {
	*requests.RpcRequest
	OwnerId    requests.Integer           `position:"Query" name:"OwnerId"`
	AppId      string                     `position:"Query" name:"AppId"`
	ChannelId  string                     `position:"Query" name:"ChannelId"`
	TemplateId requests.Integer           `position:"Query" name:"TemplateId"`
	TaskId     requests.Integer           `position:"Query" name:"TaskId"`
	MixPanes   *[]UpdateTaskParamMixPanes `position:"Query" name:"MixPanes" type:"Repeated"`
}

type UpdateTaskParamMixPanes struct {
	PaneId     requests.Integer `name:"PaneId"`
	UserId     string           `name:"UserId"`
	SourceType string           `name:"SourceType"`
}

// UpdateTaskParamResponse is the response struct for api UpdateTaskParam
type UpdateTaskParamResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
}

// CreateUpdateTaskParamRequest creates a request to invoke UpdateTaskParam API
func CreateUpdateTaskParamRequest() (request *UpdateTaskParamRequest) {
	request = &UpdateTaskParamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "UpdateTaskParam", "rtc", "openAPI")
	return
}

// CreateUpdateTaskParamResponse creates a response to parse from UpdateTaskParam response
func CreateUpdateTaskParamResponse() (response *UpdateTaskParamResponse) {
	response = &UpdateTaskParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
