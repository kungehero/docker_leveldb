package dockerleveldb

import (
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelResource struct {
	LevelDB *leveldb.DB
}

func NewLevelResource(db *leveldb.DB) LevelResource {
	return LevelResource{LevelDB: db}
}

func (level *LevelResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/leveldb").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	tags := []string{"goleveldb"}
	ws.Route(ws.PUT("/set").To(level.Put).
		Doc("put level data").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads("").
		Returns(http.StatusOK, "ok", nil).
		Returns(http.StatusBadRequest, "bad request", nil).
		Returns(http.StatusInternalServerError, "bad server", nil))
	return ws
}

func (level *LevelResource) Put(request *restful.Request, response *restful.Response) {
	err := level.LevelDB.Put([]byte("key"), []byte("value"), nil)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}
	response.WriteHeaderAndEntity(http.StatusCreated, "create ok!")
}
