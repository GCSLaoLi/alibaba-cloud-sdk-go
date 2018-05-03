package drds

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

// ModifyDrdsDBPasswd invokes the drds.ModifyDrdsDBPasswd API synchronously
// api document: https://help.aliyun.com/api/drds/modifydrdsdbpasswd.html
func (client *Client) ModifyDrdsDBPasswd(request *ModifyDrdsDBPasswdRequest) (response *ModifyDrdsDBPasswdResponse, err error) {
	response = CreateModifyDrdsDBPasswdResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDrdsDBPasswdWithChan invokes the drds.ModifyDrdsDBPasswd API asynchronously
// api document: https://help.aliyun.com/api/drds/modifydrdsdbpasswd.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDrdsDBPasswdWithChan(request *ModifyDrdsDBPasswdRequest) (<-chan *ModifyDrdsDBPasswdResponse, <-chan error) {
	responseChan := make(chan *ModifyDrdsDBPasswdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDrdsDBPasswd(request)
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

// ModifyDrdsDBPasswdWithCallback invokes the drds.ModifyDrdsDBPasswd API asynchronously
// api document: https://help.aliyun.com/api/drds/modifydrdsdbpasswd.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDrdsDBPasswdWithCallback(request *ModifyDrdsDBPasswdRequest, callback func(response *ModifyDrdsDBPasswdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDrdsDBPasswdResponse
		var err error
		defer close(result)
		response, err = client.ModifyDrdsDBPasswd(request)
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

// ModifyDrdsDBPasswdRequest is the request struct for api ModifyDrdsDBPasswd
type ModifyDrdsDBPasswdRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
	NewPasswd      string `position:"Query" name:"NewPasswd"`
}

// ModifyDrdsDBPasswdResponse is the response struct for api ModifyDrdsDBPasswd
type ModifyDrdsDBPasswdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyDrdsDBPasswdRequest creates a request to invoke ModifyDrdsDBPasswd API
func CreateModifyDrdsDBPasswdRequest() (request *ModifyDrdsDBPasswdRequest) {
	request = &ModifyDrdsDBPasswdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "ModifyDrdsDBPasswd", "", "")
	return
}

// CreateModifyDrdsDBPasswdResponse creates a response to parse from ModifyDrdsDBPasswd response
func CreateModifyDrdsDBPasswdResponse() (response *ModifyDrdsDBPasswdResponse) {
	response = &ModifyDrdsDBPasswdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
