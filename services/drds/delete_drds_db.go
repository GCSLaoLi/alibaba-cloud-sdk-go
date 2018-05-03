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

// DeleteDrdsDB invokes the drds.DeleteDrdsDB API synchronously
// api document: https://help.aliyun.com/api/drds/deletedrdsdb.html
func (client *Client) DeleteDrdsDB(request *DeleteDrdsDBRequest) (response *DeleteDrdsDBResponse, err error) {
	response = CreateDeleteDrdsDBResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDrdsDBWithChan invokes the drds.DeleteDrdsDB API asynchronously
// api document: https://help.aliyun.com/api/drds/deletedrdsdb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDrdsDBWithChan(request *DeleteDrdsDBRequest) (<-chan *DeleteDrdsDBResponse, <-chan error) {
	responseChan := make(chan *DeleteDrdsDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDrdsDB(request)
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

// DeleteDrdsDBWithCallback invokes the drds.DeleteDrdsDB API asynchronously
// api document: https://help.aliyun.com/api/drds/deletedrdsdb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDrdsDBWithCallback(request *DeleteDrdsDBRequest, callback func(response *DeleteDrdsDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDrdsDBResponse
		var err error
		defer close(result)
		response, err = client.DeleteDrdsDB(request)
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

// DeleteDrdsDBRequest is the request struct for api DeleteDrdsDB
type DeleteDrdsDBRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// DeleteDrdsDBResponse is the response struct for api DeleteDrdsDB
type DeleteDrdsDBResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteDrdsDBRequest creates a request to invoke DeleteDrdsDB API
func CreateDeleteDrdsDBRequest() (request *DeleteDrdsDBRequest) {
	request = &DeleteDrdsDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "DeleteDrdsDB", "", "")
	return
}

// CreateDeleteDrdsDBResponse creates a response to parse from DeleteDrdsDB response
func CreateDeleteDrdsDBResponse() (response *DeleteDrdsDBResponse) {
	response = &DeleteDrdsDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
