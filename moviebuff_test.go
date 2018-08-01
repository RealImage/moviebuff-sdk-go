package moviebuff

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoviebuff_GetMovie(t *testing.T) {

	var testCases = []struct {
		desc         string
		data         string
		method       string
		status       int
		expectedErr  error
		expectedResp *Movie
	}{
		{
			desc:   "get movie 200 response",
			data:   (`{"name":"Test_Movie", "type":"movie"}`),
			method: "GET",
			status: http.StatusOK,
			expectedResp: &Movie{
				Name: "Test_Movie",
				Type: "movie",
			},
		}, {
			desc:        "get movie 403 response",
			data:        (`{"name":"Test_Movie", "type":"movie"}`),
			method:      "GET",
			status:      http.StatusForbidden,
			expectedErr: ErrInvalidToken,
		}, {
			desc:        "get movie 404 response",
			data:        (`{"name":"Test_Movie", "type":"movie"}`),
			method:      "GET",
			status:      http.StatusNotFound,
			expectedErr: ErrResourceDoesNotExist,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			assert := assert.New(t)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				w.WriteHeader(testCase.status)
				w.Write([]byte(testCase.data))
			}))

			defer ts.Close()

			moviebuffApiServer := New(Config{
				HostURL:     ts.URL,
				StaticToken: "staticToken",
			})

			data, err := moviebuffApiServer.GetMovie("")
			if err != nil {
				assert.EqualValues(testCase.expectedErr, err)
			} else {
				assert.Equal(testCase.expectedResp, data)
			}
		})
	}
}

func TestMoviebuff_GetPerson(t *testing.T) {

	var testCases = []struct {
		desc         string
		data         string
		method       string
		status       int
		expectedErr  error
		expectedResp *Person
	}{
		{
			desc:   "get person 200 response",
			data:   (`{"name":"Test_Person", "type":"person"}`),
			method: "GET",
			status: http.StatusOK,
			expectedResp: &Person{
				Name: "Test_Person",
				Type: "person",
			},
		}, {
			desc:        "get person 403 response",
			data:        (`{"name":"Test_Person", "type":"person"}`),
			method:      "GET",
			status:      http.StatusForbidden,
			expectedErr: ErrInvalidToken,
		}, {
			desc:        "get person 404 response",
			data:        (`{"name":"Test_Person", "type":"person"}`),
			method:      "GET",
			status:      http.StatusNotFound,
			expectedErr: ErrResourceDoesNotExist,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			assert := assert.New(t)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				w.WriteHeader(testCase.status)
				w.Write([]byte(testCase.data))
			}))

			defer ts.Close()

			moviebuffApiServer := New(Config{
				HostURL:     ts.URL,
				StaticToken: "staticToken",
			})

			data, err := moviebuffApiServer.GetPerson("")
			if err != nil {
				assert.EqualValues(testCase.expectedErr, err)
			} else {
				assert.Equal(testCase.expectedResp, data)
			}
		})
	}
}

func TestMoviebuff_GetEntity(t *testing.T) {

	var testCases = []struct {
		desc         string
		data         string
		method       string
		status       int
		expectedErr  error
		expectedResp *Entity
	}{
		{
			desc:   "get entity 200 response",
			data:   (`{"name":"Test_Entity", "type":"entity"}`),
			method: "GET",
			status: http.StatusOK,
			expectedResp: &Entity{
				Name: "Test_Entity",
				Type: "entity",
			},
		}, {
			desc:        "get entity 403 response",
			data:        (`{"name":"Test_Entity", "type":"entity"}`),
			method:      "GET",
			status:      http.StatusForbidden,
			expectedErr: ErrInvalidToken,
		}, {
			desc:        "get entity 404 response",
			data:        (`{"name":"Test_Entity", "type":"entity"}`),
			method:      "GET",
			status:      http.StatusNotFound,
			expectedErr: ErrResourceDoesNotExist,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			assert := assert.New(t)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				w.WriteHeader(testCase.status)
				w.Write([]byte(testCase.data))
			}))

			defer ts.Close()

			moviebuffApiServer := New(Config{
				HostURL:     ts.URL,
				StaticToken: "staticToken",
			})

			data, err := moviebuffApiServer.GetEntity("")
			if err != nil {
				assert.EqualValues(testCase.expectedErr, err)
			} else {
				assert.Equal(testCase.expectedResp, data)
			}
		})
	}
}

func TestMoviebuff_GetResources(t *testing.T) {

	var testCases = []struct {
		desc         string
		data         string
		method       string
		status       int
		expectedErr  error
		expectedResp *Resources
	}{
		{
			desc:   "get entity 200 response",
			data:   (`{"prev":"1",  "next":"2"}`),
			method: "GET",
			status: http.StatusOK,
			expectedResp: &Resources{
				Prev: "1",
				Next: "2",
			},
		}, {
			desc:        "get entity 403 response",
			data:        (`{"prev":"1", "next":"2"}`),
			method:      "GET",
			status:      http.StatusForbidden,
			expectedErr: ErrInvalidToken,
		}, {
			desc:        "get entity 404 response",
			data:        (`{"prev":"1", "next":"2"}`),
			method:      "GET",
			status:      http.StatusNotFound,
			expectedErr: ErrResourceDoesNotExist,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			assert := assert.New(t)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				w.WriteHeader(testCase.status)
				w.Write([]byte(testCase.data))
			}))

			defer ts.Close()

			moviebuffApiServer := New(Config{
				HostURL:     ts.URL,
				StaticToken: "staticToken",
			})

			data, err := moviebuffApiServer.GetResources("", 0, 0)
			if err != nil {
				assert.EqualValues(testCase.expectedErr, err)
			} else {
				assert.Equal(testCase.expectedResp, data)
			}
		})
	}
}
