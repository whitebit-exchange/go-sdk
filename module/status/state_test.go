package status

import (
	"testing"

	"github.com/whitebit-exchange/go-sdk/tests/mocks"
)

type testSuite struct {
	client   *mocks.Client
	service  *Service
	endpoint *maintenanceStatus
}

func (s *testSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newMaintenanceStatusEndpoint()
}

func TestState(t *testing.T) {
	s := testSuite{}
	s.SetupTest()

	s.client.On("SendRequest", s.endpoint).Return([]byte(`{"status": 1}`), nil).Once()

	r, err := s.service.GetMaintenanceStatus()
	if err != nil {
		t.Error(err)
	}
	println(r.Status)
}
