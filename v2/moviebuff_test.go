package moviebuff

import (
	"context"
	"encoding/json"
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
		expectedResp *Calendar
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

			expectedResp: &Calendar{
				ID:   "sample calendar ID",
				Name: "Holidays in India",
				Holidays: []Holiday{
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

func TestMoviebuff_GetMappedCPL(t *testing.T) {
	var testCases = []struct {
		desc         string
		method       string
		respStatus   int
		respBody     string
		expectedErr  error
		expectedResp *MappedCPL
	}{
		{
			desc:       "get mapped cpl 200 response",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody: `{
				"id": 476829,
				"uuid": "ab9754d6-b1cf-4185-8554-4fc505670d7f",
				"content_title_text": "AndhraKingThal_P2_FTR-2D-V3_S_TE-XX_IN-UA_51-Atmos_4K_ASV_20251125_ASV_SMPTE_VF",
				"movie": {
					"id": 99986,
					"name": "Andhra King Taluka",
					"uuid": "4f90f6fa-aab6-4717-a0a0-21d7e59dd1fd",
					"part": {
						"id": 3,
						"uuid": "8928a455-2249-4152-94c8-a31249e05d7c",
						"name": "Part 2"
					}
				}
			}`,
			expectedResp: &MappedCPL{
				ID:               476829,
				UUID:             "ab9754d6-b1cf-4185-8554-4fc505670d7f",
				ContentTitleText: "AndhraKingThal_P2_FTR-2D-V3_S_TE-XX_IN-UA_51-Atmos_4K_ASV_20251125_ASV_SMPTE_VF",
				Movie: MappedCPLMovie{
					ID:   99986,
					Name: "Andhra King Taluka",
					UUID: "4f90f6fa-aab6-4717-a0a0-21d7e59dd1fd",
					Part: &MappedCPLPart{
						ID:   3,
						UUID: "8928a455-2249-4152-94c8-a31249e05d7c",
						Name: "Part 2",
					},
				},
			},
		},
		{
			desc:       "get mapped cpl 200 response without part",
			method:     "GET",
			respStatus: http.StatusOK,
			respBody: `{
				"id": 476830,
				"uuid": "test-uuid-456",
				"content_title_text": "TestMovie_FTR-2D_S_TE-XX_IN-UA_51_4K_ASV_20251125",
				"movie": {
					"id": 99987,
					"name": "Test Movie",
					"uuid": "test-movie-uuid-789"
				}
			}`,
			expectedResp: &MappedCPL{
				ID:               476830,
				UUID:             "test-uuid-456",
				ContentTitleText: "TestMovie_FTR-2D_S_TE-XX_IN-UA_51_4K_ASV_20251125",
				Movie: MappedCPLMovie{
					ID:   99987,
					Name: "Test Movie",
					UUID: "test-movie-uuid-789",
				},
			},
		},
		{
			desc:        "get mapped cpl 403 response",
			method:      "GET",
			respStatus:  http.StatusForbidden,
			respBody:    `{"error": "forbidden"}`,
			expectedErr: ErrInvalidToken,
		},
		{
			desc:        "get mapped cpl 404 response",
			method:      "GET",
			respStatus:  http.StatusNotFound,
			respBody:    `{"error": "not found"}`,
			expectedErr: ErrResourceDoesNotExist,
		},
		{
			desc:        "get mapped cpl 500 response",
			method:      "GET",
			respStatus:  http.StatusInternalServerError,
			respBody:    `{"error": "internal server error"}`,
			expectedErr: ErrResponseNotReceived,
		},
		{
			desc:        "get mapped cpl invalid json response",
			method:      "GET",
			respStatus:  http.StatusOK,
			respBody:    `invalid json`,
			expectedErr: &json.SyntaxError{},
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

			data, err := moviebuffApiServer.GetMappedCPL(context.Background(), "test-cpl-id")
			if testCase.expectedErr != nil {
				// For json errors, just check that an error occurred
				if _, ok := testCase.expectedErr.(*json.SyntaxError); ok {
					assert.Error(err)
				} else {
					assert.EqualValues(testCase.expectedErr, err)
				}
			} else {
				assert.NoError(err)
				assert.Equal(testCase.expectedResp, data)
			}
		})
	}
}
