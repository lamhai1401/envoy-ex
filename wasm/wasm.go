package main

import (
	"math/rand"
	"time"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

const tickMilliseconds uint32 = 1000

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

// Override types.DefaultPluginContext.
func (ctx *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	rand.Seed(time.Now().UnixNano())

	proxywasm.LogInfo("OnPluginStart from Go!")
	if err := proxywasm.SetTickPeriodMilliSeconds(tickMilliseconds); err != nil {
		proxywasm.LogCriticalf("failed to set tick period: %v", err)
	}

	return types.OnPluginStartStatusOK
}

// Override types.DefaultPluginContext.
func (ctx *pluginContext) OnTick() {
	t := time.Now().UnixNano()
	proxywasm.LogInfof("It's %d: random value: %d", t, rand.Uint64())
	proxywasm.LogInfof("OnTick called")
}

// Override types.DefaultPluginContext.
func (ctx *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &helloHttpContext{}
}

type helloHttpContext struct {
	types.DefaultHttpContext
	// headerName and headerValue are the header to be added to response. They are configured via
	// plugin configuration during OnPluginStart.
}

func (ctx *helloHttpContext) OnHttpRequestHeaders(numHeaders int, _ bool) types.Action {
	if numHeaders > 0 {
		headers, err := proxywasm.GetHttpRequestHeaders()
		if err != nil {
			proxywasm.LogErrorf("failed to get request headers with '%v'", err)
			return types.ActionContinue
		}
		proxywasm.LogInfof("request headers: '%+v'", headers)
	}

	return types.ActionContinue
}
