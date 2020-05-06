package handles

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not protected!\n")
}

func UserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "START!\nwanjunfu nshima 有的是！")

}