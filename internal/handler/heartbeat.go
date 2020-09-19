package handler

import (
	"encoding/json"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"net/http"
)

//Heartbeat only to check the health of the API
func Heartbeat(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)

	i.TraceIt(spanName)
	defer i.Span.Finish()

	ret, _ := json.Marshal(map[string]string{
		"msg": "hello user",
	})
	_, _ = w.Write(ret)

}
