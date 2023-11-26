package profile

import (
	"net/http"

	"github.com/POABOB/go-discord/apps/app/api/internal/logic/profile"
	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func GetUniqueProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := profile.NewGetUniqueProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetUniqueProfile()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
