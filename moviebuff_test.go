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

func TestMoviebuff_GetCertifications(t *testing.T) {
	testCases := []struct {
		desc         string
		method       string
		country      string
		respStatus   int
		respBody     string
		expectedErr  error
		expectedResp []Certification
	}{
		{
			desc:       "get all certifications 200 response",
			method:     http.MethodGet,
			country:    "",
			respStatus: 200,
			respBody: `{"data": [
				{
					"childSafe": false,
					"uuid": "3dd9313d-005f-4be9-9342-91d72d5d6896",
					"code": "A",
					"country": {
							"name": "India",
							"code": "IN",
							"uuid": "16a917fb-e15e-43b4-8ee9-5c3e822eb332"
					}
			},
			{
				"childSafe": true,
				"uuid": "892557a5-aa1b-4e0c-8c07-69208cb8003c",
				"code": "X",
				"country": {
						"name": "Bulgaria",
						"code": "BG",
						"uuid": "ed431020-0858-4d2d-a368-e54ba99b0284"
				}
			}
			]}`,
			expectedErr: nil,
			expectedResp: []Certification{
				Certification{
					ChildSafe: false,
					UUID:      "3dd9313d-005f-4be9-9342-91d72d5d6896",
					Code:      "A",
					Country: Country{
						Name: "India",
						Code: "IN",
						UUID: "16a917fb-e15e-43b4-8ee9-5c3e822eb332",
					},
				},
				Certification{ChildSafe: true,
					UUID: "892557a5-aa1b-4e0c-8c07-69208cb8003c",
					Code: "X",
					Country: Country{
						Name: "Bulgaria",
						Code: "BG",
						UUID: "ed431020-0858-4d2d-a368-e54ba99b0284",
					},
				},
			},
		},
		{
			desc:       "get country certifications 200 response",
			method:     http.MethodGet,
			country:    "3dd9313d-005f-4be9-9342-91d72d5d6896",
			respStatus: 200,
			respBody: `{"data": [
				{
					"childSafe": false,
					"uuid": "3dd9313d-005f-4be9-9342-91d72d5d6896",
					"code": "A",
					"country": {
							"name": "India",
							"code": "IN",
							"uuid": "16a917fb-e15e-43b4-8ee9-5c3e822eb332"
					}
			}
			]}`,
			expectedErr: nil,
			expectedResp: []Certification{
				Certification{
					ChildSafe: false,
					UUID:      "3dd9313d-005f-4be9-9342-91d72d5d6896",
					Code:      "A",
					Country: Country{
						Name: "India",
						Code: "IN",
						UUID: "16a917fb-e15e-43b4-8ee9-5c3e822eb332",
					},
				},
			},
		},
		{
			desc:         "get certifications 403 response",
			method:       http.MethodGet,
			country:      "",
			respStatus:   403,
			respBody:     "",
			expectedErr:  ErrInvalidToken,
			expectedResp: nil,
		},
		{
			desc:         "get certifications 404 response",
			method:       http.MethodGet,
			country:      "",
			respStatus:   404,
			respBody:     "",
			expectedErr:  ErrResourceDoesNotExist,
			expectedResp: nil,
		},
		{
			desc:         "get certifications 500 response",
			method:       http.MethodGet,
			country:      "",
			respStatus:   500,
			respBody:     "",
			expectedErr:  ErrResponseNotReceived,
			expectedResp: nil,
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

			mb := New(Config{
				HostURL:     ts.URL,
				StaticToken: "staticToken",
			})

			data, err := mb.GetCertifications(testCase.country)
			if testCase.expectedErr != nil {
				assert.EqualValues(testCase.expectedErr, err)
			} else {
				assert.NoError(err)
				assert.Equal(testCase.expectedResp, data)
			}
		})
	}
}
