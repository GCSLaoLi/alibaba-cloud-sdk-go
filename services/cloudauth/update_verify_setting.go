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

// UpdateVerifySetting invokes the cloudauth.UpdateVerifySetting API synchronously
// api document: https://help.aliyun.com/api/cloudauth/updateverifysetting.html
func (client *Client) UpdateVerifySetting(request *UpdateVerifySettingRequest) (response *UpdateVerifySettingResponse, err error) {
	response = CreateUpdateVerifySettingResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateVerifySettingWithChan invokes the cloudauth.UpdateVerifySetting API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/updateverifysetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateVerifySettingWithChan(request *UpdateVerifySettingRequest) (<-chan *UpdateVerifySettingResponse, <-chan error) {
	responseChan := make(chan *UpdateVerifySettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateVerifySetting(request)
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

// UpdateVerifySettingWithCallback invokes the cloudauth.UpdateVerifySetting API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/updateverifysetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateVerifySettingWithCallback(request *UpdateVerifySettingRequest, callback func(response *UpdateVerifySettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateVerifySettingResponse
		var err error
		defer close(result)
		response, err = client.UpdateVerifySetting(request)
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

// UpdateVerifySettingRequest is the request struct for api UpdateVerifySetting
type UpdateVerifySettingRequest struct {
	*requests.RpcRequest
	GuideStep   requests.Boolean `position:"Query" name:"GuideStep"`
	ResultStep  requests.Boolean `position:"Query" name:"ResultStep"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Solution    string           `position:"Query" name:"Solution"`
	BizName     string           `position:"Query" name:"BizName"`
	BizType     string           `position:"Query" name:"BizType"`
	PrivacyStep requests.Boolean `position:"Query" name:"PrivacyStep"`
}

// UpdateVerifySettingResponse is the response struct for api UpdateVerifySetting
type UpdateVerifySettingResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	BizType   string   `json:"BizType" xml:"BizType"`
	BizName   string   `json:"BizName" xml:"BizName"`
	Solution  string   `json:"Solution" xml:"Solution"`
	StepList  []string `json:"StepList" xml:"StepList"`
}

// CreateUpdateVerifySettingRequest creates a request to invoke UpdateVerifySetting API
func CreateUpdateVerifySettingRequest() (request *UpdateVerifySettingRequest) {
	request = &UpdateVerifySettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "UpdateVerifySetting", "", "")
	return
}

// CreateUpdateVerifySettingResponse creates a response to parse from UpdateVerifySetting response
func CreateUpdateVerifySettingResponse() (response *UpdateVerifySettingResponse) {
	response = &UpdateVerifySettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
