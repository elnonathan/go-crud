package person_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-crud/person"
	"net/http"
	"net/http/httptest"
	"testing"
)

import (
	testdb "go-crud/test"
	testrouter "go-crud/test"
)

func TestGetPerson(testing *testing.T) {
	database, err := testdb.OpenTestDatabase()
	require.NoError(testing, err)

	tested := person.Person{
		Name:  "Jonathan",
		Email: "jonathan@test.com",
	}

	require.NoError(testing, database.Create(&tested).Error)
	require.NotNil(testing, tested.Id)

	router := testrouter.SetupTestRouter(database)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/person/%d", *tested.Id), nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(testing, http.StatusOK, recorder.Code)
	assert.Contains(testing, recorder.Body.String(), "Jonathan")
}

func TestListPerson(testing *testing.T) {
	database, err := testdb.OpenTestDatabase()
	require.NoError(testing, err)

	person_a := person.Person{
		Name:  "Person L1",
		Email: "l1@person_testing.io",
	}

	require.NoError(testing, database.Create(&person_a).Error)
	require.NotNil(testing, person_a.Id)

	person_b := person.Person{
		Name:  "Person L2",
		Email: "l2@person_testing.io",
	}

	require.NoError(testing, database.Create(&person_b).Error)
	require.NotNil(testing, person_b.Id)

	router := testrouter.SetupTestRouter(database)
	request, _ := http.NewRequest("GET", "/person/list?email=person_testing", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var response []person.Person
	err = json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.NoError(testing, err)
	assert.Equal(testing, len(response), 2)
}

/*
func TestCreatePerson(testing *testing.T) {
	database, err := testdb.OpenTestDatabase()
	require.NoError(testing, err)

	tested := person.Person{
		Name:  "Person C1",
		Email: "c1@person_testing.io",
	}

	require.NoError(testing, database.Create(&tested).Error)
	require.NotNil(testing, tested.Id)

	router := testrouter.SetupTestRouter(database)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/person/%d", *tested.Id), nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(testing, http.StatusOK, recorder.Code)
	assert.Contains(testing, recorder.Body.String(), "Jonathan")
}
*/
