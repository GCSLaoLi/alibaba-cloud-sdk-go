package cs

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

// PauseComponentUpgrade invokes the cs.PauseComponentUpgrade API synchronously
// api document: https://help.aliyun.com/api/cs/pausecomponentupgrade.html
func (client *Client) PauseComponentUpgrade(request *PauseComponentUpgradeRequest) (response *PauseComponentUpgradeResponse, err error) {
	response = CreatePauseComponentUpgradeResponse()
	err = client.DoAction(request, response)
	return
}

// PauseComponentUpgradeWithChan invokes the cs.PauseComponentUpgrade API asynchronously
// api document: https://help.aliyun.com/api/cs/pausecomponentupgrade.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PauseComponentUpgradeWithChan(request *PauseComponentUpgradeRequest) (<-chan *PauseComponentUpgradeResponse, <-chan error) {
	responseChan := make(chan *PauseComponentUpgradeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PauseComponentUpgrade(request)
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

// PauseComponentUpgradeWithCallback invokes the cs.PauseComponentUpgrade API asynchronously
// api document: https://help.aliyun.com/api/cs/pausecomponentupgrade.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PauseComponentUpgradeWithCallback(request *PauseComponentUpgradeRequest, callback func(response *PauseComponentUpgradeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PauseComponentUpgradeResponse
		var err error
		defer close(result)
		response, err = client.PauseComponentUpgrade(request)
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

// PauseComponentUpgradeRequest is the request struct for api PauseComponentUpgrade
type PauseComponentUpgradeRequest struct {
	*requests.RoaRequest
	Componentid string `position:"Path" name:"componentid"`
	Clusterid   string `position:"Path" name:"clusterid"`
}

// PauseComponentUpgradeResponse is the response struct for api PauseComponentUpgrade
type PauseComponentUpgradeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePauseComponentUpgradeRequest creates a request to invoke PauseComponentUpgrade API
func CreatePauseComponentUpgradeRequest() (request *PauseComponentUpgradeRequest) {
	request = &PauseComponentUpgradeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "PauseComponentUpgrade", "/clusters/[clusterid]/components/[componentid]/pause", "", "")
	request.Method = requests.POST
	return
}

// CreatePauseComponentUpgradeResponse creates a response to parse from PauseComponentUpgrade response
func CreatePauseComponentUpgradeResponse() (response *PauseComponentUpgradeResponse) {
	response = &PauseComponentUpgradeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
