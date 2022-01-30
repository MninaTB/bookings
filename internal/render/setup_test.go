package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/MninaTB/bookings/internal/config"
	"github.com/MninaTB/bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	session = scs.New()
	// the time the session will last
	session.Lifetime = 24 * time.Hour
	// the cookies should persist, after the user closes the browser window
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// should the cookie be encrypted - only works with https
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
