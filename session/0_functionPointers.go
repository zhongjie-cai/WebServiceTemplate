package session

import (
	"encoding/json"
	"fmt"
	"net/textproto"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
)

// func pointers for injection / testing: logger.go
var (
	uuidNew                         = uuid.New
	jsonMarshal                     = json.Marshal
	jsonUnmarshal                   = json.Unmarshal
	fmtErrorf                       = fmt.Errorf
	muxVars                         = mux.Vars
	requestGetRequestBody           = request.GetRequestBody
	apperrorGetBadRequestError      = apperror.GetBadRequestError
	apperrorConsolidateAllErrors    = apperror.ConsolidateAllErrors
	textprotoCanonicalMIMEHeaderKey = textproto.CanonicalMIMEHeaderKey
	getFunc                         = Get
	tryUnmarshalFunc                = tryUnmarshal
	getRequestFunc                  = GetRequest
	getAllQueriesFunc               = getAllQueries
	getAllHeadersFunc               = getAllHeaders
)
