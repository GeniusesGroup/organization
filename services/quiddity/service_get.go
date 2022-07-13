/* For license and copyright information please see LEGAL file in repository */

package quiddity

import (
	"../../libgo/http"
	"../../libgo/json"
	"../../libgo/protocol"
	"../../libgo/service"
	"../../libgo/srpc"
	org "../../protocol"
)

var GetService getService

func init() {
	GetService.Service.MediaType.Init("", "", protocol.Software_PreAlpha, 1605026701)
	GetService.Service.MediaType.SetDetail(protocol.LanguageEnglish,
		"Get Quiddity",
		"",
		[]string{"Quiddity"})
	GetService.Service.SetAuthorization(protocol.CRUDRead, protocol.UserType_All)
}

type getService struct {
	service.Service
}

func (ser *getService) Process(st protocol.Stream, req org.Quiddity_Service_Get_Request) (res org.Quiddity_Service_Get_Response, err protocol.Error) {
	// TODO::: authorization??
	res, err = storage.Get(req.QuiddityID())
	return
}
func (ser *getService) ServeSRPC(st protocol.Stream) (err protocol.Error) {
	// TODO::: auto generate
	var req = &GetRequest{}
	err = req.FromSyllab(srpc.GetPayload(st.IncomePayload))
	if err != nil {
		return
	}

	var res *getQuiddityRes
	res, err = ser.Process(st, req)
	// Check if any error occur in business logic
	if err != nil {
		return
	}

	st.OutcomePayload = make([]byte, srpc.MinLength+res.LenAsSyllab())
	res.ToSyllab(srpc.GetPayload(st.OutcomePayload))
	return
}
func (ser *getService) ServeHTTP(st protocol.Stream, httpReq protocol.HTTPRequest, httpRes protocol.HTTPResponse) (err protocol.Error) {
	// TODO::: auto generate
	var req = &GetRequest{}
	err = req.FromJSON(httpReq.Body().Marshal())
	if err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	var res *getQuiddityRes
	res, err = ser.Process(st, req)
	// Check if any error occur in business logic
	if err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	httpRes.SetBody(json.NewCodec(res))
	return
}

func (ser *getService) Do(req org.Quiddity_Service_Get_Request) (res org.Quiddity_Service_Get_Response, err protocol.Error) {
	// TODO::: auto generate
	return
}
func (ser *getService) doSRPC(req org.Quiddity_Service_Get_Request) (res org.Quiddity_Service_Get_Response, err protocol.Error) {
	// TODO::: auto generate
	return
}
func (ser *getService) doHTTP(req org.Quiddity_Service_Get_Request) (res org.Quiddity_Service_Get_Response, err protocol.Error) {
	// TODO::: auto generate
	return
}

func (ser *getService) validateRequest(req org.Quiddity_Service_Get_Request) (err protocol.Error) {
	// TODO::: auto generate
	return
}
