package moviebuff

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/RealImage/moviebuff-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestMoviebuff_GetMovie(t *testing.T) {

	var testCases = []struct {
		desc         string
		method       string
		respStatus   int
		respBody     string
		expectedErr  error
		expectedResp *v1.Movie
	}{
		{
			desc:       "get movie 200 response",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody:   (`{"name":"Test_Movie", "type":"movie"}`),
			expectedResp: &v1.Movie{
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

			data, err := moviebuffApiServer.GetMovie(context.Background(), "")
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
		expectedResp *v1.Person
	}{
		{
			desc:       "get person 200 response",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody:   (`{"name":"Test_Person", "type":"person"}`),
			expectedResp: &v1.Person{
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

			data, err := moviebuffApiServer.GetPerson(context.Background(), "")
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
		expectedResp *v1.Entity
	}{
		{
			desc:       "get entity 200 response",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody:   (`{"name":"Test_Entity", "type":"entity"}`),
			expectedResp: &v1.Entity{
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

			data, err := moviebuffApiServer.GetEntity(context.Background(), "")
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
		expectedResp *v1.Resources
	}{
		{
			desc:       "get entity 200 response",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody:   (`{"prev":"1",  "next":"2"}`),
			expectedResp: &v1.Resources{
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

			data, err := moviebuffApiServer.GetResources(context.Background(), "", 0, 0)
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
		expectedResp []v1.Certification
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
			expectedResp: []v1.Certification{
				v1.Certification{
					ChildSafe: false,
					UUID:      "3dd9313d-005f-4be9-9342-91d72d5d6896",
					Code:      "A",
					Country: v1.Country{
						Name: "India",
						Code: "IN",
						UUID: "16a917fb-e15e-43b4-8ee9-5c3e822eb332",
					},
				},
				v1.Certification{ChildSafe: true,
					UUID: "892557a5-aa1b-4e0c-8c07-69208cb8003c",
					Code: "X",
					Country: v1.Country{
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
			expectedResp: []v1.Certification{
				v1.Certification{
					ChildSafe: false,
					UUID:      "3dd9313d-005f-4be9-9342-91d72d5d6896",
					Code:      "A",
					Country: v1.Country{
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

			data, err := mb.GetCertifications(context.Background(), testCase.country)
			if testCase.expectedErr != nil {
				assert.EqualValues(testCase.expectedErr, err)
			} else {
				assert.NoError(err)
				assert.Equal(testCase.expectedResp, data)
			}
		})
	}
}

func TestMoviebuff_GetHolidayCalendar(t *testing.T) {
	var testCases = []struct {
		countryID    string
		calendarID   string
		respBody     string
		expectedResp *v1.Calendar
		respStatus   int
	}{
		{

			countryID:  "16a917fb-e15e-43b4-8ee9-5c3e822eb332",
			calendarID: "en.indian#holiday@group.v.calendar.google.com",

			respBody: `{
				"calendarId": "sample calendar ID",
				"name": "Holidays in India",
				"holidays": [
					{
						"id": "20180414_60o30d9h60o30c1g60o30dr568",
						"name": "Vaisakhi",
						"date": "2018-04-14"
					},
					{
						"id": "20181106_60o30d9h60o30c1g60o30dr565",
						"name": "Diwali",
						"date": "2018-11-06"
					}
				],
				"syncToken": "sample sync token",
				"timeZone": "Asia/Calcutta"
				}`,

			expectedResp: &v1.Calendar{
				ID:   "sample calendar ID",
				Name: "Holidays in India",
				Holidays: []v1.Holiday{
					{
						ID:   "20180414_60o30d9h60o30c1g60o30dr568",
						Name: "Vaisakhi",
						Date: "2018-04-14",
					},
					{
						ID:   "20181106_60o30d9h60o30c1g60o30dr565",
						Name: "Diwali",
						Date: "2018-11-06",
					},
				},
				SyncToken: "sample sync token",
				TimeZone:  "Asia/Calcutta",
			},

			respStatus: 200,
		},
	}

	for _, testCase := range testCases {
		t.Run("Get Holiday Calendar", func(t *testing.T) {
			assert := assert.New(t)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				w.WriteHeader(testCase.respStatus)
				w.Write([]byte(testCase.respBody))
			}))

			mb := New(Config{
				HostURL:     ts.URL,
				StaticToken: "staticToken",
			})

			data, err := mb.GetHolidayCalendar(context.Background(), testCase.countryID)
			assert.Equal(testCase.expectedResp, data)
			assert.NoError(err)
		})
	}

}
