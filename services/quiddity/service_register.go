/* For license and copyright information please see LEGAL file in repository */

package quiddity

import (
	"../../libgo/ganjine"
	"../../libgo/http"
	"../../libgo/protocol"
	"../../libgo/service"
	"../../libgo/srpc"
	"../../libgo/uuid"
	org "../../protocol"
)

var RegisterService registerService

func init() {
	RegisterService.Service.MediaType.Init(org.FindQuiddityByOrgIDURN, "", protocol.Software_PreAlpha, 1604939795)
	RegisterService.Service.MediaType.SetDetail(protocol.LanguageEnglish,
		"Register Quiddity",
		"",
		[]string{"Quiddity"})
	RegisterService.Service.SetAuthorization(protocol.CRUDCreate, protocol.UserType_Org)
}

type registerService struct {
	service.Service
}

func (ser *registerService) Process(st protocol.Stream, req org.Quiddity_Service_Register_Request) (err protocol.Error) {
	err = st.Authorize()
	if err != nil {
		return
	}

	if req.URI != "" {
		err = checkQuiddityURI(st, req.URI)
		if err != nil {
			return
		}
	}

	// err = checkOrgName(st, req.Name)
	// if err != nil {
	// 	return
	// }
	// err = checkOrgDomain(st, req.Domain)
	// if err != nil {
	// 	return
	// }

	var q = Quiddity{
		AppInstanceID:    achaemenid.App.Nodes.LocalNode.InstanceID,
		UserConnectionID: st.Connection.ID,
		ID:               uuid.Random32Byte(),
		OrgID:            st.Connection.UserID,

		Language: req.Language,
		URI:      req.URI,
		Title:    req.Title,
		Status:   data.QuiddityStatusRegister,
	}
	err = storage.Save(&q)
	if err != nil {
		return
	}

	res = &registerQuiddityRes{
		ID: q.ID,
	}

	return
}
func (ser *registerService) ServeSRPC(st protocol.Stream) (err protocol.Error) {
	var req RegisterRequest
	err = req.FromSyllab(srpc.GetPayload(st.IncomePayload))
	if err != nil {
		return
	}

	err = ser.Process(st, &req)
	return
}
func (ser *registerService) ServeHTTP(st protocol.Stream, httpReq protocol.HTTPRequest, httpRes protocol.HTTPResponse) (err protocol.Error) {
	var req RegisterRequest
	err = req.FromJSON(httpReq.Body().Marshal())
	if err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	err = ser.Process(st, &req)
	// Check if any error occur in business logic
	if err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	return
}
func (ser *registerService) validateRequest(req org.Quiddity_Service_Register_Request) (err protocol.Error) {
	// TODO::: auto generate
	return
}

func checkQuiddityURI(st protocol.Stream, uri string) (err protocol.Error) {
	var findQuiddityByURIReq = findQuiddityByURIReq{
		URI:    uri,
		Offset: 18446744073709551615,
		Limit:  1,
	}
	var findQuiddityByURIRes *findQuiddityByURIRes
	findQuiddityByURIRes, err = findQuiddityByURI(st, &findQuiddityByURIReq)
	if err.Equal(ganjine.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return
	}

	var getQuiddityLanguagesReq = getQuiddityLanguagesReq{
		ID: findQuiddityByURIRes.IDs[0],
	}
	var getQuiddityLanguagesRes *getQuiddityLanguagesRes
	getQuiddityLanguagesRes, err = getQuiddityLanguages(st, &getQuiddityLanguagesReq)
	if err.Equal(ganjine.ErrRecordNotFound) {
		// TODO::: how it is possible???
		return nil
	}
	if err != nil {
		return
	}

	var getQuiddityReq = getQuiddityReq{
		ID:       findQuiddityByURIRes.IDs[0],
		Language: getQuiddityLanguagesRes.Languages[0],
	}
	var getQuiddityRes *getQuiddityRes
	getQuiddityRes, err = getQuiddity(st, &getQuiddityReq)
	if err.Equal(ganjine.ErrRecordNotFound) {
		// TODO::: how it is possible???
		return nil
	}
	if err != nil {
		return
	}
	if getQuiddityRes.URI == uri {
		return ErrQuiddityURIRegistered
	}
	return
}
