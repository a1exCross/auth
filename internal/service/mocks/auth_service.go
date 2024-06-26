// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/a1exCross/auth/internal/service.AuthService -o auth_service.go -n AuthServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/a1exCross/auth/internal/model"
	"github.com/gojuno/minimock/v3"
)

// AuthServiceMock implements service.AuthService
type AuthServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetAccessToken          func(ctx context.Context, s1 string) (s2 string, err error)
	inspectFuncGetAccessToken   func(ctx context.Context, s1 string)
	afterGetAccessTokenCounter  uint64
	beforeGetAccessTokenCounter uint64
	GetAccessTokenMock          mAuthServiceMockGetAccessToken

	funcGetRefreshToken          func(ctx context.Context, s1 string) (s2 string, err error)
	inspectFuncGetRefreshToken   func(ctx context.Context, s1 string)
	afterGetRefreshTokenCounter  uint64
	beforeGetRefreshTokenCounter uint64
	GetRefreshTokenMock          mAuthServiceMockGetRefreshToken

	funcLogin          func(ctx context.Context, l1 model.LoginDTO) (s1 string, err error)
	inspectFuncLogin   func(ctx context.Context, l1 model.LoginDTO)
	afterLoginCounter  uint64
	beforeLoginCounter uint64
	LoginMock          mAuthServiceMockLogin
}

// NewAuthServiceMock returns a mock for service.AuthService
func NewAuthServiceMock(t minimock.Tester) *AuthServiceMock {
	m := &AuthServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetAccessTokenMock = mAuthServiceMockGetAccessToken{mock: m}
	m.GetAccessTokenMock.callArgs = []*AuthServiceMockGetAccessTokenParams{}

	m.GetRefreshTokenMock = mAuthServiceMockGetRefreshToken{mock: m}
	m.GetRefreshTokenMock.callArgs = []*AuthServiceMockGetRefreshTokenParams{}

	m.LoginMock = mAuthServiceMockLogin{mock: m}
	m.LoginMock.callArgs = []*AuthServiceMockLoginParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mAuthServiceMockGetAccessToken struct {
	mock               *AuthServiceMock
	defaultExpectation *AuthServiceMockGetAccessTokenExpectation
	expectations       []*AuthServiceMockGetAccessTokenExpectation

	callArgs []*AuthServiceMockGetAccessTokenParams
	mutex    sync.RWMutex
}

// AuthServiceMockGetAccessTokenExpectation specifies expectation struct of the AuthService.GetAccessToken
type AuthServiceMockGetAccessTokenExpectation struct {
	mock    *AuthServiceMock
	params  *AuthServiceMockGetAccessTokenParams
	results *AuthServiceMockGetAccessTokenResults
	Counter uint64
}

// AuthServiceMockGetAccessTokenParams contains parameters of the AuthService.GetAccessToken
type AuthServiceMockGetAccessTokenParams struct {
	ctx context.Context
	s1  string
}

// AuthServiceMockGetAccessTokenResults contains results of the AuthService.GetAccessToken
type AuthServiceMockGetAccessTokenResults struct {
	s2  string
	err error
}

// Expect sets up expected params for AuthService.GetAccessToken
func (mmGetAccessToken *mAuthServiceMockGetAccessToken) Expect(ctx context.Context, s1 string) *mAuthServiceMockGetAccessToken {
	if mmGetAccessToken.mock.funcGetAccessToken != nil {
		mmGetAccessToken.mock.t.Fatalf("AuthServiceMock.GetAccessToken mock is already set by Set")
	}

	if mmGetAccessToken.defaultExpectation == nil {
		mmGetAccessToken.defaultExpectation = &AuthServiceMockGetAccessTokenExpectation{}
	}

	mmGetAccessToken.defaultExpectation.params = &AuthServiceMockGetAccessTokenParams{ctx, s1}
	for _, e := range mmGetAccessToken.expectations {
		if minimock.Equal(e.params, mmGetAccessToken.defaultExpectation.params) {
			mmGetAccessToken.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetAccessToken.defaultExpectation.params)
		}
	}

	return mmGetAccessToken
}

// Inspect accepts an inspector function that has same arguments as the AuthService.GetAccessToken
func (mmGetAccessToken *mAuthServiceMockGetAccessToken) Inspect(f func(ctx context.Context, s1 string)) *mAuthServiceMockGetAccessToken {
	if mmGetAccessToken.mock.inspectFuncGetAccessToken != nil {
		mmGetAccessToken.mock.t.Fatalf("Inspect function is already set for AuthServiceMock.GetAccessToken")
	}

	mmGetAccessToken.mock.inspectFuncGetAccessToken = f

	return mmGetAccessToken
}

// Return sets up results that will be returned by AuthService.GetAccessToken
func (mmGetAccessToken *mAuthServiceMockGetAccessToken) Return(s2 string, err error) *AuthServiceMock {
	if mmGetAccessToken.mock.funcGetAccessToken != nil {
		mmGetAccessToken.mock.t.Fatalf("AuthServiceMock.GetAccessToken mock is already set by Set")
	}

	if mmGetAccessToken.defaultExpectation == nil {
		mmGetAccessToken.defaultExpectation = &AuthServiceMockGetAccessTokenExpectation{mock: mmGetAccessToken.mock}
	}
	mmGetAccessToken.defaultExpectation.results = &AuthServiceMockGetAccessTokenResults{s2, err}
	return mmGetAccessToken.mock
}

// Set uses given function f to mock the AuthService.GetAccessToken method
func (mmGetAccessToken *mAuthServiceMockGetAccessToken) Set(f func(ctx context.Context, s1 string) (s2 string, err error)) *AuthServiceMock {
	if mmGetAccessToken.defaultExpectation != nil {
		mmGetAccessToken.mock.t.Fatalf("Default expectation is already set for the AuthService.GetAccessToken method")
	}

	if len(mmGetAccessToken.expectations) > 0 {
		mmGetAccessToken.mock.t.Fatalf("Some expectations are already set for the AuthService.GetAccessToken method")
	}

	mmGetAccessToken.mock.funcGetAccessToken = f
	return mmGetAccessToken.mock
}

// When sets expectation for the AuthService.GetAccessToken which will trigger the result defined by the following
// Then helper
func (mmGetAccessToken *mAuthServiceMockGetAccessToken) When(ctx context.Context, s1 string) *AuthServiceMockGetAccessTokenExpectation {
	if mmGetAccessToken.mock.funcGetAccessToken != nil {
		mmGetAccessToken.mock.t.Fatalf("AuthServiceMock.GetAccessToken mock is already set by Set")
	}

	expectation := &AuthServiceMockGetAccessTokenExpectation{
		mock:   mmGetAccessToken.mock,
		params: &AuthServiceMockGetAccessTokenParams{ctx, s1},
	}
	mmGetAccessToken.expectations = append(mmGetAccessToken.expectations, expectation)
	return expectation
}

// Then sets up AuthService.GetAccessToken return parameters for the expectation previously defined by the When method
func (e *AuthServiceMockGetAccessTokenExpectation) Then(s2 string, err error) *AuthServiceMock {
	e.results = &AuthServiceMockGetAccessTokenResults{s2, err}
	return e.mock
}

// GetAccessToken implements service.AuthService
func (mmGetAccessToken *AuthServiceMock) GetAccessToken(ctx context.Context, s1 string) (s2 string, err error) {
	mm_atomic.AddUint64(&mmGetAccessToken.beforeGetAccessTokenCounter, 1)
	defer mm_atomic.AddUint64(&mmGetAccessToken.afterGetAccessTokenCounter, 1)

	if mmGetAccessToken.inspectFuncGetAccessToken != nil {
		mmGetAccessToken.inspectFuncGetAccessToken(ctx, s1)
	}

	mm_params := AuthServiceMockGetAccessTokenParams{ctx, s1}

	// Record call args
	mmGetAccessToken.GetAccessTokenMock.mutex.Lock()
	mmGetAccessToken.GetAccessTokenMock.callArgs = append(mmGetAccessToken.GetAccessTokenMock.callArgs, &mm_params)
	mmGetAccessToken.GetAccessTokenMock.mutex.Unlock()

	for _, e := range mmGetAccessToken.GetAccessTokenMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s2, e.results.err
		}
	}

	if mmGetAccessToken.GetAccessTokenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetAccessToken.GetAccessTokenMock.defaultExpectation.Counter, 1)
		mm_want := mmGetAccessToken.GetAccessTokenMock.defaultExpectation.params
		mm_got := AuthServiceMockGetAccessTokenParams{ctx, s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetAccessToken.t.Errorf("AuthServiceMock.GetAccessToken got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetAccessToken.GetAccessTokenMock.defaultExpectation.results
		if mm_results == nil {
			mmGetAccessToken.t.Fatal("No results are set for the AuthServiceMock.GetAccessToken")
		}
		return (*mm_results).s2, (*mm_results).err
	}
	if mmGetAccessToken.funcGetAccessToken != nil {
		return mmGetAccessToken.funcGetAccessToken(ctx, s1)
	}
	mmGetAccessToken.t.Fatalf("Unexpected call to AuthServiceMock.GetAccessToken. %v %v", ctx, s1)
	return
}

// GetAccessTokenAfterCounter returns a count of finished AuthServiceMock.GetAccessToken invocations
func (mmGetAccessToken *AuthServiceMock) GetAccessTokenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetAccessToken.afterGetAccessTokenCounter)
}

// GetAccessTokenBeforeCounter returns a count of AuthServiceMock.GetAccessToken invocations
func (mmGetAccessToken *AuthServiceMock) GetAccessTokenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetAccessToken.beforeGetAccessTokenCounter)
}

// Calls returns a list of arguments used in each call to AuthServiceMock.GetAccessToken.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetAccessToken *mAuthServiceMockGetAccessToken) Calls() []*AuthServiceMockGetAccessTokenParams {
	mmGetAccessToken.mutex.RLock()

	argCopy := make([]*AuthServiceMockGetAccessTokenParams, len(mmGetAccessToken.callArgs))
	copy(argCopy, mmGetAccessToken.callArgs)

	mmGetAccessToken.mutex.RUnlock()

	return argCopy
}

// MinimockGetAccessTokenDone returns true if the count of the GetAccessToken invocations corresponds
// the number of defined expectations
func (m *AuthServiceMock) MinimockGetAccessTokenDone() bool {
	for _, e := range m.GetAccessTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetAccessTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetAccessTokenCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetAccessToken != nil && mm_atomic.LoadUint64(&m.afterGetAccessTokenCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetAccessTokenInspect logs each unmet expectation
func (m *AuthServiceMock) MinimockGetAccessTokenInspect() {
	for _, e := range m.GetAccessTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthServiceMock.GetAccessToken with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetAccessTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetAccessTokenCounter) < 1 {
		if m.GetAccessTokenMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AuthServiceMock.GetAccessToken")
		} else {
			m.t.Errorf("Expected call to AuthServiceMock.GetAccessToken with params: %#v", *m.GetAccessTokenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetAccessToken != nil && mm_atomic.LoadUint64(&m.afterGetAccessTokenCounter) < 1 {
		m.t.Error("Expected call to AuthServiceMock.GetAccessToken")
	}
}

type mAuthServiceMockGetRefreshToken struct {
	mock               *AuthServiceMock
	defaultExpectation *AuthServiceMockGetRefreshTokenExpectation
	expectations       []*AuthServiceMockGetRefreshTokenExpectation

	callArgs []*AuthServiceMockGetRefreshTokenParams
	mutex    sync.RWMutex
}

// AuthServiceMockGetRefreshTokenExpectation specifies expectation struct of the AuthService.GetRefreshToken
type AuthServiceMockGetRefreshTokenExpectation struct {
	mock    *AuthServiceMock
	params  *AuthServiceMockGetRefreshTokenParams
	results *AuthServiceMockGetRefreshTokenResults
	Counter uint64
}

// AuthServiceMockGetRefreshTokenParams contains parameters of the AuthService.GetRefreshToken
type AuthServiceMockGetRefreshTokenParams struct {
	ctx context.Context
	s1  string
}

// AuthServiceMockGetRefreshTokenResults contains results of the AuthService.GetRefreshToken
type AuthServiceMockGetRefreshTokenResults struct {
	s2  string
	err error
}

// Expect sets up expected params for AuthService.GetRefreshToken
func (mmGetRefreshToken *mAuthServiceMockGetRefreshToken) Expect(ctx context.Context, s1 string) *mAuthServiceMockGetRefreshToken {
	if mmGetRefreshToken.mock.funcGetRefreshToken != nil {
		mmGetRefreshToken.mock.t.Fatalf("AuthServiceMock.GetRefreshToken mock is already set by Set")
	}

	if mmGetRefreshToken.defaultExpectation == nil {
		mmGetRefreshToken.defaultExpectation = &AuthServiceMockGetRefreshTokenExpectation{}
	}

	mmGetRefreshToken.defaultExpectation.params = &AuthServiceMockGetRefreshTokenParams{ctx, s1}
	for _, e := range mmGetRefreshToken.expectations {
		if minimock.Equal(e.params, mmGetRefreshToken.defaultExpectation.params) {
			mmGetRefreshToken.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetRefreshToken.defaultExpectation.params)
		}
	}

	return mmGetRefreshToken
}

// Inspect accepts an inspector function that has same arguments as the AuthService.GetRefreshToken
func (mmGetRefreshToken *mAuthServiceMockGetRefreshToken) Inspect(f func(ctx context.Context, s1 string)) *mAuthServiceMockGetRefreshToken {
	if mmGetRefreshToken.mock.inspectFuncGetRefreshToken != nil {
		mmGetRefreshToken.mock.t.Fatalf("Inspect function is already set for AuthServiceMock.GetRefreshToken")
	}

	mmGetRefreshToken.mock.inspectFuncGetRefreshToken = f

	return mmGetRefreshToken
}

// Return sets up results that will be returned by AuthService.GetRefreshToken
func (mmGetRefreshToken *mAuthServiceMockGetRefreshToken) Return(s2 string, err error) *AuthServiceMock {
	if mmGetRefreshToken.mock.funcGetRefreshToken != nil {
		mmGetRefreshToken.mock.t.Fatalf("AuthServiceMock.GetRefreshToken mock is already set by Set")
	}

	if mmGetRefreshToken.defaultExpectation == nil {
		mmGetRefreshToken.defaultExpectation = &AuthServiceMockGetRefreshTokenExpectation{mock: mmGetRefreshToken.mock}
	}
	mmGetRefreshToken.defaultExpectation.results = &AuthServiceMockGetRefreshTokenResults{s2, err}
	return mmGetRefreshToken.mock
}

// Set uses given function f to mock the AuthService.GetRefreshToken method
func (mmGetRefreshToken *mAuthServiceMockGetRefreshToken) Set(f func(ctx context.Context, s1 string) (s2 string, err error)) *AuthServiceMock {
	if mmGetRefreshToken.defaultExpectation != nil {
		mmGetRefreshToken.mock.t.Fatalf("Default expectation is already set for the AuthService.GetRefreshToken method")
	}

	if len(mmGetRefreshToken.expectations) > 0 {
		mmGetRefreshToken.mock.t.Fatalf("Some expectations are already set for the AuthService.GetRefreshToken method")
	}

	mmGetRefreshToken.mock.funcGetRefreshToken = f
	return mmGetRefreshToken.mock
}

// When sets expectation for the AuthService.GetRefreshToken which will trigger the result defined by the following
// Then helper
func (mmGetRefreshToken *mAuthServiceMockGetRefreshToken) When(ctx context.Context, s1 string) *AuthServiceMockGetRefreshTokenExpectation {
	if mmGetRefreshToken.mock.funcGetRefreshToken != nil {
		mmGetRefreshToken.mock.t.Fatalf("AuthServiceMock.GetRefreshToken mock is already set by Set")
	}

	expectation := &AuthServiceMockGetRefreshTokenExpectation{
		mock:   mmGetRefreshToken.mock,
		params: &AuthServiceMockGetRefreshTokenParams{ctx, s1},
	}
	mmGetRefreshToken.expectations = append(mmGetRefreshToken.expectations, expectation)
	return expectation
}

// Then sets up AuthService.GetRefreshToken return parameters for the expectation previously defined by the When method
func (e *AuthServiceMockGetRefreshTokenExpectation) Then(s2 string, err error) *AuthServiceMock {
	e.results = &AuthServiceMockGetRefreshTokenResults{s2, err}
	return e.mock
}

// GetRefreshToken implements service.AuthService
func (mmGetRefreshToken *AuthServiceMock) GetRefreshToken(ctx context.Context, s1 string) (s2 string, err error) {
	mm_atomic.AddUint64(&mmGetRefreshToken.beforeGetRefreshTokenCounter, 1)
	defer mm_atomic.AddUint64(&mmGetRefreshToken.afterGetRefreshTokenCounter, 1)

	if mmGetRefreshToken.inspectFuncGetRefreshToken != nil {
		mmGetRefreshToken.inspectFuncGetRefreshToken(ctx, s1)
	}

	mm_params := AuthServiceMockGetRefreshTokenParams{ctx, s1}

	// Record call args
	mmGetRefreshToken.GetRefreshTokenMock.mutex.Lock()
	mmGetRefreshToken.GetRefreshTokenMock.callArgs = append(mmGetRefreshToken.GetRefreshTokenMock.callArgs, &mm_params)
	mmGetRefreshToken.GetRefreshTokenMock.mutex.Unlock()

	for _, e := range mmGetRefreshToken.GetRefreshTokenMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s2, e.results.err
		}
	}

	if mmGetRefreshToken.GetRefreshTokenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetRefreshToken.GetRefreshTokenMock.defaultExpectation.Counter, 1)
		mm_want := mmGetRefreshToken.GetRefreshTokenMock.defaultExpectation.params
		mm_got := AuthServiceMockGetRefreshTokenParams{ctx, s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetRefreshToken.t.Errorf("AuthServiceMock.GetRefreshToken got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetRefreshToken.GetRefreshTokenMock.defaultExpectation.results
		if mm_results == nil {
			mmGetRefreshToken.t.Fatal("No results are set for the AuthServiceMock.GetRefreshToken")
		}
		return (*mm_results).s2, (*mm_results).err
	}
	if mmGetRefreshToken.funcGetRefreshToken != nil {
		return mmGetRefreshToken.funcGetRefreshToken(ctx, s1)
	}
	mmGetRefreshToken.t.Fatalf("Unexpected call to AuthServiceMock.GetRefreshToken. %v %v", ctx, s1)
	return
}

// GetRefreshTokenAfterCounter returns a count of finished AuthServiceMock.GetRefreshToken invocations
func (mmGetRefreshToken *AuthServiceMock) GetRefreshTokenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetRefreshToken.afterGetRefreshTokenCounter)
}

// GetRefreshTokenBeforeCounter returns a count of AuthServiceMock.GetRefreshToken invocations
func (mmGetRefreshToken *AuthServiceMock) GetRefreshTokenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetRefreshToken.beforeGetRefreshTokenCounter)
}

// Calls returns a list of arguments used in each call to AuthServiceMock.GetRefreshToken.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetRefreshToken *mAuthServiceMockGetRefreshToken) Calls() []*AuthServiceMockGetRefreshTokenParams {
	mmGetRefreshToken.mutex.RLock()

	argCopy := make([]*AuthServiceMockGetRefreshTokenParams, len(mmGetRefreshToken.callArgs))
	copy(argCopy, mmGetRefreshToken.callArgs)

	mmGetRefreshToken.mutex.RUnlock()

	return argCopy
}

// MinimockGetRefreshTokenDone returns true if the count of the GetRefreshToken invocations corresponds
// the number of defined expectations
func (m *AuthServiceMock) MinimockGetRefreshTokenDone() bool {
	for _, e := range m.GetRefreshTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetRefreshTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetRefreshTokenCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetRefreshToken != nil && mm_atomic.LoadUint64(&m.afterGetRefreshTokenCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetRefreshTokenInspect logs each unmet expectation
func (m *AuthServiceMock) MinimockGetRefreshTokenInspect() {
	for _, e := range m.GetRefreshTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthServiceMock.GetRefreshToken with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetRefreshTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetRefreshTokenCounter) < 1 {
		if m.GetRefreshTokenMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AuthServiceMock.GetRefreshToken")
		} else {
			m.t.Errorf("Expected call to AuthServiceMock.GetRefreshToken with params: %#v", *m.GetRefreshTokenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetRefreshToken != nil && mm_atomic.LoadUint64(&m.afterGetRefreshTokenCounter) < 1 {
		m.t.Error("Expected call to AuthServiceMock.GetRefreshToken")
	}
}

type mAuthServiceMockLogin struct {
	mock               *AuthServiceMock
	defaultExpectation *AuthServiceMockLoginExpectation
	expectations       []*AuthServiceMockLoginExpectation

	callArgs []*AuthServiceMockLoginParams
	mutex    sync.RWMutex
}

// AuthServiceMockLoginExpectation specifies expectation struct of the AuthService.Login
type AuthServiceMockLoginExpectation struct {
	mock    *AuthServiceMock
	params  *AuthServiceMockLoginParams
	results *AuthServiceMockLoginResults
	Counter uint64
}

// AuthServiceMockLoginParams contains parameters of the AuthService.Login
type AuthServiceMockLoginParams struct {
	ctx context.Context
	l1  model.LoginDTO
}

// AuthServiceMockLoginResults contains results of the AuthService.Login
type AuthServiceMockLoginResults struct {
	s1  string
	err error
}

// Expect sets up expected params for AuthService.Login
func (mmLogin *mAuthServiceMockLogin) Expect(ctx context.Context, l1 model.LoginDTO) *mAuthServiceMockLogin {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &AuthServiceMockLoginExpectation{}
	}

	mmLogin.defaultExpectation.params = &AuthServiceMockLoginParams{ctx, l1}
	for _, e := range mmLogin.expectations {
		if minimock.Equal(e.params, mmLogin.defaultExpectation.params) {
			mmLogin.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLogin.defaultExpectation.params)
		}
	}

	return mmLogin
}

// Inspect accepts an inspector function that has same arguments as the AuthService.Login
func (mmLogin *mAuthServiceMockLogin) Inspect(f func(ctx context.Context, l1 model.LoginDTO)) *mAuthServiceMockLogin {
	if mmLogin.mock.inspectFuncLogin != nil {
		mmLogin.mock.t.Fatalf("Inspect function is already set for AuthServiceMock.Login")
	}

	mmLogin.mock.inspectFuncLogin = f

	return mmLogin
}

// Return sets up results that will be returned by AuthService.Login
func (mmLogin *mAuthServiceMockLogin) Return(s1 string, err error) *AuthServiceMock {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &AuthServiceMockLoginExpectation{mock: mmLogin.mock}
	}
	mmLogin.defaultExpectation.results = &AuthServiceMockLoginResults{s1, err}
	return mmLogin.mock
}

// Set uses given function f to mock the AuthService.Login method
func (mmLogin *mAuthServiceMockLogin) Set(f func(ctx context.Context, l1 model.LoginDTO) (s1 string, err error)) *AuthServiceMock {
	if mmLogin.defaultExpectation != nil {
		mmLogin.mock.t.Fatalf("Default expectation is already set for the AuthService.Login method")
	}

	if len(mmLogin.expectations) > 0 {
		mmLogin.mock.t.Fatalf("Some expectations are already set for the AuthService.Login method")
	}

	mmLogin.mock.funcLogin = f
	return mmLogin.mock
}

// When sets expectation for the AuthService.Login which will trigger the result defined by the following
// Then helper
func (mmLogin *mAuthServiceMockLogin) When(ctx context.Context, l1 model.LoginDTO) *AuthServiceMockLoginExpectation {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("AuthServiceMock.Login mock is already set by Set")
	}

	expectation := &AuthServiceMockLoginExpectation{
		mock:   mmLogin.mock,
		params: &AuthServiceMockLoginParams{ctx, l1},
	}
	mmLogin.expectations = append(mmLogin.expectations, expectation)
	return expectation
}

// Then sets up AuthService.Login return parameters for the expectation previously defined by the When method
func (e *AuthServiceMockLoginExpectation) Then(s1 string, err error) *AuthServiceMock {
	e.results = &AuthServiceMockLoginResults{s1, err}
	return e.mock
}

// Login implements service.AuthService
func (mmLogin *AuthServiceMock) Login(ctx context.Context, l1 model.LoginDTO) (s1 string, err error) {
	mm_atomic.AddUint64(&mmLogin.beforeLoginCounter, 1)
	defer mm_atomic.AddUint64(&mmLogin.afterLoginCounter, 1)

	if mmLogin.inspectFuncLogin != nil {
		mmLogin.inspectFuncLogin(ctx, l1)
	}

	mm_params := AuthServiceMockLoginParams{ctx, l1}

	// Record call args
	mmLogin.LoginMock.mutex.Lock()
	mmLogin.LoginMock.callArgs = append(mmLogin.LoginMock.callArgs, &mm_params)
	mmLogin.LoginMock.mutex.Unlock()

	for _, e := range mmLogin.LoginMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmLogin.LoginMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLogin.LoginMock.defaultExpectation.Counter, 1)
		mm_want := mmLogin.LoginMock.defaultExpectation.params
		mm_got := AuthServiceMockLoginParams{ctx, l1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLogin.t.Errorf("AuthServiceMock.Login got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLogin.LoginMock.defaultExpectation.results
		if mm_results == nil {
			mmLogin.t.Fatal("No results are set for the AuthServiceMock.Login")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmLogin.funcLogin != nil {
		return mmLogin.funcLogin(ctx, l1)
	}
	mmLogin.t.Fatalf("Unexpected call to AuthServiceMock.Login. %v %v", ctx, l1)
	return
}

// LoginAfterCounter returns a count of finished AuthServiceMock.Login invocations
func (mmLogin *AuthServiceMock) LoginAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLogin.afterLoginCounter)
}

// LoginBeforeCounter returns a count of AuthServiceMock.Login invocations
func (mmLogin *AuthServiceMock) LoginBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLogin.beforeLoginCounter)
}

// Calls returns a list of arguments used in each call to AuthServiceMock.Login.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLogin *mAuthServiceMockLogin) Calls() []*AuthServiceMockLoginParams {
	mmLogin.mutex.RLock()

	argCopy := make([]*AuthServiceMockLoginParams, len(mmLogin.callArgs))
	copy(argCopy, mmLogin.callArgs)

	mmLogin.mutex.RUnlock()

	return argCopy
}

// MinimockLoginDone returns true if the count of the Login invocations corresponds
// the number of defined expectations
func (m *AuthServiceMock) MinimockLoginDone() bool {
	for _, e := range m.LoginMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoginMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLogin != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		return false
	}
	return true
}

// MinimockLoginInspect logs each unmet expectation
func (m *AuthServiceMock) MinimockLoginInspect() {
	for _, e := range m.LoginMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthServiceMock.Login with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoginMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		if m.LoginMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AuthServiceMock.Login")
		} else {
			m.t.Errorf("Expected call to AuthServiceMock.Login with params: %#v", *m.LoginMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLogin != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		m.t.Error("Expected call to AuthServiceMock.Login")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *AuthServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetAccessTokenInspect()

			m.MinimockGetRefreshTokenInspect()

			m.MinimockLoginInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *AuthServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *AuthServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetAccessTokenDone() &&
		m.MinimockGetRefreshTokenDone() &&
		m.MinimockLoginDone()
}
