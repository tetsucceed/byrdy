package db

import (
	"github.com/gin-gonic/gin"
	"github.com/tetsucceed/byrdy/pkg/utils"
	"net/http"
	"net/url"
)

type Env struct {
	Links LinkDbModel
}

func (model *LinkDbModel) getLast() Structure {
	var r Structure

	model.DB.Last(&Structure{}).Scan(&r)

	return r
}

func (model *LinkDbModel) getByHash(st string) Structure {
	var r Structure

	model.DB.Where("short_link = ?", st).First(&r)

	return r
}

func (model *LinkDbModel) createNewRecord(raw string, hash string) {
	var r Structure

	r = Structure{RawLink: raw, ShortLink: hash}
	model.DB.Create(&r)
}

func (env *Env) GetLink(rc *gin.Context) {
	link := env.Links.getByHash(rc.Query("r"))
	redirectLocation := url.URL{Path: link.RawLink}
	rc.Redirect(http.StatusFound, redirectLocation.RequestURI())
}

func (env *Env) SaveShort(rc *gin.Context) {
	var postModel PostModel
	newShortHash := utils.LinkHash()
	_ = rc.BindJSON(&postModel)
	env.Links.createNewRecord(postModel.OriginLink, newShortHash)

	rc.JSON(http.StatusOK, gin.H{"short_link": newShortHash})
}
