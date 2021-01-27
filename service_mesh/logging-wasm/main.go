// Copyright 2020 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetNewRootContext(newRootContext)
	proxywasm.SetNewHttpContext(newHttpContext)
	proxywasm.SetNewStreamContext(newNetworkContext)
}

type httpLogRootContext struct {
	proxywasm.DefaultRootContext
	contextID uint32
}

func newRootContext(contextID uint32) proxywasm.RootContext {
	return &httpLogRootContext{contextID: contextID}
}

// override
func (ctx *httpLogRootContext) OnVMStart(vmConfigurationSize int) bool {
	path := []string{"cluster_name"}
	result, err := proxywasm.GetProperty(path)
	if err != nil {
		proxywasm.LogCritical("failed to get property in envoy.")
	} else {
		proxywasm.LogInfo("Node Info  : " + string(result))
	}
	return true
}

type logHttpContext struct {
	proxywasm.DefaultHttpContext
}

func newHttpContext(uint32, uint32) proxywasm.HttpContext {
	return &logHttpContext{}
}

// override
func (ctx *logHttpContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	pathMethod := []string{"request", "method"}
	methodResult, _ := proxywasm.GetProperty(pathMethod)
	proxywasm.LogInfo("Request Method  : " + string(methodResult))

	connectionPath := []string{"source", "address"}
	connectionResult, _ := proxywasm.GetProperty(connectionPath)
	proxywasm.LogInfo("Source Address :" + string(connectionResult))

	return types.ActionContinue
}

type networkContext struct {
	// you must embed the default context so that you need not to reimplement all the methods by yourself
	proxywasm.DefaultStreamContext
}

func newNetworkContext(rootContextID, contextID uint32) proxywasm.StreamContext {
	return &networkContext{}
}

func (ctx *networkContext) OnUpstreamData(dataSize int, endOfStream bool) types.Action {
	if dataSize == 0 {
		return types.ActionContinue
	}

	ret, err := proxywasm.GetProperty([]string{"upstream", "address"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get upstream data: %v", err)
		return types.ActionContinue
	}

	proxywasm.LogInfof("remote address: %s", string(ret))

	data, err := proxywasm.GetUpstreamData(0, dataSize)
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogCritical(err.Error())
	}

	proxywasm.LogInfof("<<<<<< upstream data received <<<<<<\n%s", string(data))
	return types.ActionContinue
}
