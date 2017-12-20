package arms

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

func (client *Client) ARMSQueryDataSet(request *ARMSQueryDataSetRequest) (response *ARMSQueryDataSetResponse, err error) {
	response = CreateARMSQueryDataSetResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ARMSQueryDataSetWithChan(request *ARMSQueryDataSetRequest) (<-chan *ARMSQueryDataSetResponse, <-chan error) {
	responseChan := make(chan *ARMSQueryDataSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ARMSQueryDataSet(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) ARMSQueryDataSetWithCallback(request *ARMSQueryDataSetRequest, callback func(response *ARMSQueryDataSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ARMSQueryDataSetResponse
		var err error
		defer close(result)
		response, err = client.ARMSQueryDataSet(request)
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

type ARMSQueryDataSetRequest struct {
	*requests.RpcRequest
	DateStr       string                          `position:"Query" name:"DateStr"`
	MinTime       string                          `position:"Query" name:"MinTime"`
	ReduceTail    string                          `position:"Query" name:"ReduceTail"`
	MaxTime       string                          `position:"Query" name:"MaxTime"`
	OptionalDims  *[]ARMSQueryDataSetOptionalDims `position:"Query" name:"OptionalDims"  type:"Repeated"`
	Measures      *[]string                       `position:"Query" name:"Measures"  type:"Repeated"`
	IntervalInSec string                          `position:"Query" name:"IntervalInSec"`
	IsDrillDown   string                          `position:"Query" name:"IsDrillDown"`
	HungryMode    string                          `position:"Query" name:"HungryMode"`
	OrderByKey    string                          `position:"Query" name:"OrderByKey"`
	Limit         string                          `position:"Query" name:"Limit"`
	DatasetId     string                          `position:"Query" name:"DatasetId"`
	RequiredDims  *[]ARMSQueryDataSetRequiredDims `position:"Query" name:"RequiredDims"  type:"Repeated"`
	Dimensions    *[]ARMSQueryDataSetDimensions   `position:"Query" name:"Dimensions"  type:"Repeated"`
}

type ARMSQueryDataSetOptionalDims struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}
type ARMSQueryDataSetRequiredDims struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}
type ARMSQueryDataSetDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}

type ARMSQueryDataSetResponse struct {
	*responses.BaseResponse
	Data string `json:"Data"`
}

func CreateARMSQueryDataSetRequest() (request *ARMSQueryDataSetRequest) {
	request = &ARMSQueryDataSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2016-11-25", "ARMSQueryDataSet", "", "")
	return
}

func CreateARMSQueryDataSetResponse() (response *ARMSQueryDataSetResponse) {
	response = &ARMSQueryDataSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
