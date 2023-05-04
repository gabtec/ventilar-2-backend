package e2e_test

// go get github.com/stretchr/testify

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gabtec/ventilar-2-backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type E2ESuite struct {
	suite.Suite
}

func TestE2ESuite(t *testing.T) {
	suite.Run(t, new(E2ESuite))
}

func (s *E2ESuite) TestGetWards(t *testing.T) {

	c := http.Client{}

	r, _ := c.Get("http://localhost:4000/api/wards")

	assert.Equal(t, http.StatusOK, r.StatusCode)

	wards := []types.Ward{}
	// data, _ := ioutil.ReadAll(r.Body)
	json.NewDecoder(r.Body).Decode(&wards)
	
	// fmt.Println(string(data))
	// assert.JSONEq(t, "[{\"ward_id\":1,\"name\":\"HSA_Intensiva\",\"belongs_to\":\"HSA\",\"is_park\":true,\"created_at\":\"2023-05-04T10:26:04.149605+01:00\",\"updated_at\":\"2023-05-04T10:26:04.149605+01:00\"}]", string(data))
	assert.Equal(t, len(wards), 1)

	assert.Equal(t, wards[0].ID, uint(1))
	assert.Equal(t, wards[0].Name, "HSA_Intensiva")
}