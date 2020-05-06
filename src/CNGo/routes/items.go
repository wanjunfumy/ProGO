package handles

import (
	h "CNGo/routes/handles"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Param func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)

func init() {
	Cells = append(Cells,
		// index根目录
		Cell{
			requestType: http.MethodGet,
			path:        "/",
			handle:      h.Index,
		},

		// users
		Cell{
			requestType: http.MethodGet,
			path:        "/user",
			handle:      h.UserInfo,
		},
	)
}

// 不需要代码来注入，只需要认为处理即可
var Cells []Cell

type Cell struct {
	// 创建一个路由元素，里面有 类型，直径用net/http中的get，post等、路径、一个接口
	requestType string
	path        string
	handle      httprouter.Handle
}

func (c Cell) Init(r *httprouter.Router) {
	r.Handle(c.requestType, c.path, c.handle)
}
