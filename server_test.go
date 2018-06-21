package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ipochi/api-mock-example/client"
	"github.com/ipochi/api-mock-example/handler"
	"github.com/ipochi/api-mock-example/mocks"
	"github.com/ipochi/api-mock-example/model"
	"github.com/stretchr/testify/assert"
)

type mocktest struct {
	mc      *gomock.Controller
	mockfnc *mocks.MockFunctions
}

func newMocktest(t *testing.T) *mocktest {
	ms := &mocktest{}
	ms.mc = gomock.NewController(t)
	ms.mockfnc = mocks.NewMockFunctions(ms.mc)

	return ms

}
func TestGetCompanies(t *testing.T) {

	ms := newMocktest(t)

	server := handler.New(ms.mockfnc)
	ts := httptest.NewServer(http.HandlerFunc(server.GetCompanies))

	defer ts.Close()
	defer ms.mc.Finish()

	url, _ := url.Parse(ts.URL)
	fmt.Println("What is url ---", url)
	client := &client.Client{
		BaseURL:    url,
		HttpClient: &http.Client{},
	}

	ms.mockfnc.EXPECT().
		GetCompanies().
		Return([]model.Company{
			model.Company{
				Name:  "Lalala",
				Email: "lala@lalaland.com",
				Tel:   "123-lalala",
			},
			model.Company{
				Name:  "Lala",
				Email: "lala@lalaland.com",
				Tel:   "1456-lalala",
			},
		}, nil)

	cmp, err := client.GetCompanies()

	assert.Equal(t, 2, len(cmp))
	fmt.Println("Err --- ", err)
	fmt.Println("What is get companies --- ", cmp)
}
