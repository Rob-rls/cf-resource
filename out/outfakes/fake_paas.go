// Code generated by counterfeiter. DO NOT EDIT.
package outfakes

import (
	"sync"

	"github.com/concourse/cf-resource/out"
)

type FakePAAS struct {
	LoginStub        func(api string, username string, password string, clientID string, clientSecret string, insecure bool) error
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		api          string
		username     string
		password     string
		clientID     string
		clientSecret string
		insecure     bool
	}
	loginReturns struct {
		result1 error
	}
	loginReturnsOnCall map[int]struct {
		result1 error
	}
	TargetStub        func(organization string, space string) error
	targetMutex       sync.RWMutex
	targetArgsForCall []struct {
		organization string
		space        string
	}
	targetReturns struct {
		result1 error
	}
	targetReturnsOnCall map[int]struct {
		result1 error
	}
	PushAppStub        func(manifest string, path string, currentAppName string, dockerUser string, showLogs bool) error
	pushAppMutex       sync.RWMutex
	pushAppArgsForCall []struct {
		manifest       string
		path           string
		currentAppName string
		dockerUser     string
		showLogs       bool
	}
	pushAppReturns struct {
		result1 error
	}
	pushAppReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePAAS) Login(api string, username string, password string, clientID string, clientSecret string, insecure bool) error {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		api          string
		username     string
		password     string
		clientID     string
		clientSecret string
		insecure     bool
	}{api, username, password, clientID, clientSecret, insecure})
	fake.recordInvocation("Login", []interface{}{api, username, password, clientID, clientSecret, insecure})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub(api, username, password, clientID, clientSecret, insecure)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.loginReturns.result1
}

func (fake *FakePAAS) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakePAAS) LoginArgsForCall(i int) (string, string, string, string, string, bool) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return fake.loginArgsForCall[i].api, fake.loginArgsForCall[i].username, fake.loginArgsForCall[i].password, fake.loginArgsForCall[i].clientID, fake.loginArgsForCall[i].clientSecret, fake.loginArgsForCall[i].insecure
}

func (fake *FakePAAS) LoginReturns(result1 error) {
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePAAS) LoginReturnsOnCall(i int, result1 error) {
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePAAS) Target(organization string, space string) error {
	fake.targetMutex.Lock()
	ret, specificReturn := fake.targetReturnsOnCall[len(fake.targetArgsForCall)]
	fake.targetArgsForCall = append(fake.targetArgsForCall, struct {
		organization string
		space        string
	}{organization, space})
	fake.recordInvocation("Target", []interface{}{organization, space})
	fake.targetMutex.Unlock()
	if fake.TargetStub != nil {
		return fake.TargetStub(organization, space)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.targetReturns.result1
}

func (fake *FakePAAS) TargetCallCount() int {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return len(fake.targetArgsForCall)
}

func (fake *FakePAAS) TargetArgsForCall(i int) (string, string) {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return fake.targetArgsForCall[i].organization, fake.targetArgsForCall[i].space
}

func (fake *FakePAAS) TargetReturns(result1 error) {
	fake.TargetStub = nil
	fake.targetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePAAS) TargetReturnsOnCall(i int, result1 error) {
	fake.TargetStub = nil
	if fake.targetReturnsOnCall == nil {
		fake.targetReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.targetReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePAAS) PushApp(manifest string, path string, currentAppName string, dockerUser string, showLogs bool) error {
	fake.pushAppMutex.Lock()
	ret, specificReturn := fake.pushAppReturnsOnCall[len(fake.pushAppArgsForCall)]
	fake.pushAppArgsForCall = append(fake.pushAppArgsForCall, struct {
		manifest       string
		path           string
		currentAppName string
		dockerUser     string
		showLogs       bool
	}{manifest, path, currentAppName, dockerUser, showLogs})
	fake.recordInvocation("PushApp", []interface{}{manifest, path, currentAppName, dockerUser, showLogs})
	fake.pushAppMutex.Unlock()
	if fake.PushAppStub != nil {
		return fake.PushAppStub(manifest, path, currentAppName, dockerUser, showLogs)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pushAppReturns.result1
}

func (fake *FakePAAS) PushAppCallCount() int {
	fake.pushAppMutex.RLock()
	defer fake.pushAppMutex.RUnlock()
	return len(fake.pushAppArgsForCall)
}

func (fake *FakePAAS) PushAppArgsForCall(i int) (string, string, string, string, bool) {
	fake.pushAppMutex.RLock()
	defer fake.pushAppMutex.RUnlock()
	return fake.pushAppArgsForCall[i].manifest, fake.pushAppArgsForCall[i].path, fake.pushAppArgsForCall[i].currentAppName, fake.pushAppArgsForCall[i].dockerUser, fake.pushAppArgsForCall[i].showLogs
}

func (fake *FakePAAS) PushAppReturns(result1 error) {
	fake.PushAppStub = nil
	fake.pushAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePAAS) PushAppReturnsOnCall(i int, result1 error) {
	fake.PushAppStub = nil
	if fake.pushAppReturnsOnCall == nil {
		fake.pushAppReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pushAppReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePAAS) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	fake.pushAppMutex.RLock()
	defer fake.pushAppMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePAAS) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ out.PAAS = new(FakePAAS)
