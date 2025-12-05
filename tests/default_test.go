package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/beego/beego/v2/core/logs"

	_ "jkl-finance/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	// Test Status Code Should Be 200
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Test The Result Should Not Be Empty
	if w.Body.Len() == 0 {
		t.Error("Expected response body to not be empty")
	}
}

