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
		method       string
		respStatus   int
		respBody     string
		expectedErr  error
		expectedResp *Movie
	}{
		{
			desc:       "get movie 200 response",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody:   (`{"name":"Test_Movie", "type":"movie"}`),
			expectedResp: &Movie{
				Name: "Test_Movie",
				Type: "movie",
			},
		}, {
			desc:        "get movie 403 response",
			method:      "GET",
			respStatus:  http.StatusForbidden,
			respBody:    (`{"name":"Test_Movie", "type":"movie"}`),
			expectedErr: ErrInvalidToken,
		}, {
			desc:        "get movie 404 response",
			method:      "GET",
			respStatus:  http.StatusNotFound,
			respBody:    (`{"name":"Test_Movie", "type":"movie"}`),
			expectedErr: ErrResourceDoesNotExist,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			assert := assert.New(t)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				w.WriteHeader(testCase.respStatus)
				w.Write([]byte(testCase.respBody))
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
		method       string
		respStatus   int
		respBody     string
		expectedErr  error
		expectedResp *Person
	}{
		{
			desc:       "get person 200 response",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody:   (`{"name":"Test_Person", "type":"person"}`),
			expectedResp: &Person{
				Name: "Test_Person",
				Type: "person",
			},
		}, {
			desc:        "get person 403 response",
			method:      "GET",
			respStatus:  http.StatusForbidden,
			respBody:    (`{"name":"Test_Person", "type":"person"}`),
			expectedErr: ErrInvalidToken,
		}, {
			desc:        "get person 404 response",
			method:      "GET",
			respStatus:  http.StatusNotFound,
			respBody:    (`{"name":"Test_Person", "type":"person"}`),
			expectedErr: ErrResourceDoesNotExist,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			assert := assert.New(t)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				w.WriteHeader(testCase.respStatus)
				w.Write([]byte(testCase.respBody))
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
		method       string
		respStatus   int
		respBody     string
		expectedErr  error
		expectedResp *Entity
	}{
		{
			desc:       "get entity 200 response",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody:   (`{"name":"Test_Entity", "type":"entity"}`),
			expectedResp: &Entity{
				Name: "Test_Entity",
				Type: "entity",
			},
		}, {
			desc:        "get entity 403 response",
			method:      "GET",
			respStatus:  http.StatusForbidden,
			respBody:    (`{"name":"Test_Entity", "type":"entity"}`),
			expectedErr: ErrInvalidToken,
		}, {
			desc:        "get entity 404 response",
			method:      "GET",
			respStatus:  http.StatusNotFound,
			respBody:    (`{"name":"Test_Entity", "type":"entity"}`),
			expectedErr: ErrResourceDoesNotExist,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			assert := assert.New(t)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				w.WriteHeader(testCase.respStatus)
				w.Write([]byte(testCase.respBody))
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
		method       string
		respStatus   int
		respBody     string
		expectedErr  error
		expectedResp *Resources
	}{
		{
			desc:       "get entity 200 response",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody:   (`{"prev":"1",  "next":"2"}`),
			expectedResp: &Resources{
				Prev: "1",
				Next: "2",
			},
		}, {
			desc:        "get entity 403 response",
			method:      "GET",
			respStatus:  http.StatusForbidden,
			respBody:    (`{"prev":"1", "next":"2"}`),
			expectedErr: ErrInvalidToken,
		}, {
			desc:        "get entity 404 response",
			method:      "GET",
			respStatus:  http.StatusNotFound,
			respBody:    (`{"prev":"1", "next":"2"}`),
			expectedErr: ErrResourceDoesNotExist,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			assert := assert.New(t)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				w.WriteHeader(testCase.respStatus)
				w.Write([]byte(testCase.respBody))
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
