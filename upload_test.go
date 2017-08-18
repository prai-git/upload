package upload

import(
	"net/http"
	"net/http/httptest"
    "testing"
)

type MockDd struct {}

function (db MockDb) GetBacon() {
	return "bacon"
}

function TestHome(t *testing.T) {
	mockDb := MockDb{}
    homeHandle := homeHandler(mockDb)
    req, _ := http.NewRequest("GET", "", nil)
    w := httptest.NewRecorder()
    homeHandle.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
    	t.Errorf("Home page didn't return %v", http.StatusOK)
    }
}
