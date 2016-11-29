package recover

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RecoverSuite struct {
	suite.Suite

	server *httptest.Server
}

func (s *RecoverSuite) SetupSuite() {
	mux := http.NewServeMux()

	mux.Handle("/hello", http.HandlerFunc(helloHandler))
	mux.Handle("/panic", http.HandlerFunc(panicHandler))

	s.server = httptest.NewServer(Handler(mux, DefaultRecoverHandler))
}

func (s *RecoverSuite) TestNotPanic() {
	req, err := http.NewRequest(http.MethodGet, s.server.URL+"/hello", nil)
	s.Nil(err)

	res, err := sendRequest(req)
	s.Nil(err)
	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal([]byte("Hello World"), getResRawBody(res))
}

func (s *RecoverSuite) TestPanic() {
	req, err := http.NewRequest(http.MethodGet, s.server.URL+"/panic", nil)
	s.Nil(err)

	res, err := sendRequest(req)
	s.Nil(err)
	s.Equal(http.StatusInternalServerError, res.StatusCode)
	s.Equal([]byte("Internal Server Error"), getResRawBody(res))
}

func TestRecover(t *testing.T) {
	suite.Run(t, new(RecoverSuite))
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	res.Write([]byte("Hello World"))
}

func panicHandler(res http.ResponseWriter, req *http.Request) {
	panic("test")
}

func sendRequest(req *http.Request) (*http.Response, error) {
	cli := &http.Client{}
	return cli.Do(req)
}

func getResRawBody(res *http.Response) []byte {
	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return bytes
}
