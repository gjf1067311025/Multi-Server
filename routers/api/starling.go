package api

import (
	"github.com/EDDYCJY/go-gin-example/common"
	"github.com/EDDYCJY/go-gin-example/form"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/service/starling_basic_info"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Post Starling
// @Produce  json
// @Param key query string true "Key"
// @Param ChineseText query string true "chinese_text"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /starling [post]
func CreateStarling(c *gin.Context) {
	appG := app.Gin{C: c}
	var req form.BasicInfo
	err := c.ShouldBindJSON(&req)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = starling_basic_info.CreateStarling(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_STAR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Post Starling
// @Produce  json
// @Param key query string true "key"
// @Param ChineseText query string true "chinese_text"
// @Param EnglishText query string true "english_text"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /starling/batch [post]
func BatchCreateStarling(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	var req form.BatchBasicInfo
	err := c.ShouldBind(&req)
	if err != nil && len(req.BasicInfo) == 0 {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = starling_basic_info.BatchCreateStarling(req)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_BATCH_CREATE_STAR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 分页搜索
// @Description 根据字段分页返回
// @Accept json
// @Produce  json
// @Param key path string true "key"
// @Param page_size path string true "page_size"
// @Param page_num path string true "page_num"
// @Param chinese_text path string true "chinese_text"
// @Param english_text path string true "english_text"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /starling [get]
func GetStarlingListByCondition(c *gin.Context) {
	appG := app.Gin{C: c}
	key := c.DefaultQuery("key", "")
	size := c.DefaultQuery("page_size", "10")
	num := c.DefaultQuery("page_num", "1")
	chinese := c.DefaultQuery("chinese_text", "")
	english := c.DefaultQuery("english_text", "")
	ret, total, err := starling_basic_info.GetStarlingByCondition(key, size, num, chinese, english)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_SEARCH, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"total":    total,
		"starList": ret,
	})
}

func GetStarlingListBySearchKey(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	key := c.DefaultQuery("search_key", "")
	size := c.DefaultQuery("page_size", "10")
	num := c.DefaultQuery("page_num", "1")
	ret, total, err := starling_basic_info.GetStarlingBySearchKey(key, size, num)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_SEARCH, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"total":    total,
		"starList": ret,
	})
}

// @Summary Post Starling
// @Produce  json
// @Param ID query string true "id"
// @Param key query string true "key"
// @Param ChineseText query string true "chinese_text"
// @Param englishText query string true "english_text"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /starling [put]
func UpdateStarling(c *gin.Context) {
	appG := app.Gin{C: c}
	var req form.NewBasicInfo
	err := c.ShouldBind(&req)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = starling_basic_info.UpdateStarling(req)
	if err != nil {
		if err.Error() == common.RECORD_NOT_FOUND {
			appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_STARLING, nil)
			return
		}
		if err.Error() == common.KEYEXISTS {
			appG.Response(http.StatusInternalServerError, e.ERROR_EXISTS, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 删除镜像迁移任务
// @Description 删除镜像迁移任务
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /starling [delete]
func DeleteStarling(c *gin.Context) {
	appG := app.Gin{C: c}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = starling_basic_info.DeleteStarling(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 删除镜像迁移任务
// @Description 删除镜像迁移任务
// @Accept  json
// @Produce  json
// @Param ids query string true "ids"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /starling [delete]
func DeleteBatchStarling(c *gin.Context) {
	appG := app.Gin{C: c}
	idStrs := c.DefaultQuery("ids", "")
	idStrsList := strings.Split(idStrs, ",")
	err := starling_basic_info.DeleteBatch(idStrsList)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetTextByKeyAndLang(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	var req form.KeyLang
	err := c.ShouldBindJSON(&req)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	ret, err := starling_basic_info.GetStarlingByKeyAndLang(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, ret)
}

func JudgeRepeatKey(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	key := c.DefaultQuery("key", "")
	if key == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	isExist, err := starling_basic_info.ExistRepeatKey(key)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]bool{
		"is_exist": isExist,
	})
}

func TranslateText(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	text := c.DefaultQuery("text", "")
	if text == "" {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	translateText, err := starling_basic_info.TranslateCh2En(text)
	if err != nil {
		logging.Error("TranslateCh2En err,", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"text": translateText,
	})
}
