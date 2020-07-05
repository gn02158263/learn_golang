package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func main() {
	// read tag check data
	req2, _ := http.NewRequest(http.MethodGet, "/v1/tags/"+createTag.Id, nil)
	SetAdminHeader(req2)

	resp2 := httptest.NewRecorder()
	router.ServeHTTP(resp2, req2)
	assert.Equal(t, http.StatusOK, resp2.Code)

	// format data
	result2 := &intercom.JsonResponse{}
	err2 := json.Unmarshal(resp2.Body.Bytes(), result2)
	assert.NoError(t, err2)
	var readTag = &model.Tag{}
	err2 = json.Unmarshal(result2.Result, &readTag)
	assert.NoError(t, err2)
	assert.Equal(t, p.MetaData, readTag.MetaData)
}