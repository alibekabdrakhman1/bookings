package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

func TestForm_Vaild(t *testing.T) {
    r := httptest.NewRequest("POST", "/whatever", nil)
    form := New(r.PostForm)
    isValid := form.Valid()
    if !isValid {
        t.Error("got invalid when should have been valid")
    }
}
func TestForm_Required(t *testing.T) {
    r := httptest.NewRequest("POST", "/whatever", nil)
    form := New(r.PostForm)

    form.Required("a", "b", "c")
    if form.Valid() {
        t.Error("form shows valid when required fields missing")
    }
    postedData := url.Values{}
    postedData.Add("a", "a")
    postedData.Add("c", "c")
    postedData.Add("b", "b")

    r, _ = http.NewRequest("POST", "/whatever", nil) 
    r.PostForm = postedData
    form = New(r.PostForm)
    form.Required("a", "b", "c")
    if !form.Valid() {
        t.Error("shows does not have required fields when it does")
    }


}