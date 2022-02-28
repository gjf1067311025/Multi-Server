package starling_basic_info

import (
	"errors"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/common"
	"github.com/EDDYCJY/go-gin-example/form"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func CreateStarling(req form.BasicInfo) error {
	starlingBasicInfo := models.StarlingBasicInfo{
		Key:                req.Key,
		ChineseText:        req.ChineseText,
		EnglishText:        req.EnglishText,
		Remark:             req.Remark,
		IsMachineTranslate: req.IsMachineTranslate,
	}

	isExist, err := models.ExistStarByKey(req.Key)
	if err != nil {
		logging.Error("models.ExistStarByKey err, ", err)
		return err
	}
	if isExist {
		logging.Error("key ", req.Key, " exists!")
		return errors.New("key already exists!")
	}
	err = models.CreateStarling(&starlingBasicInfo, nil)
	if err != nil {
		logging.Error("models.CreateStarling err, ", err)
		return err
	}
	return nil
}

func BatchCreateStarling(req form.BatchBasicInfo) error {
	batchBasicInfo := make([]models.StarlingBasicInfo, 0)
	for _, basicInfo := range req.BasicInfo {
		starlingBasicInfo := models.StarlingBasicInfo{
			Key:                basicInfo.Key,
			ChineseText:        basicInfo.ChineseText,
			EnglishText:        basicInfo.EnglishText,
			Remark:             basicInfo.Remark,
			IsMachineTranslate: basicInfo.IsMachineTranslate,
		}

		isExist, err := models.ExistStarByKey(basicInfo.Key)
		if err != nil {
			logging.Error("models.ExistStarByKey err, ", err)
			return err
		}
		if isExist {
			logging.Error("key ", basicInfo.Key, " exists!")
			return errors.New("key already exists!")
		}
		batchBasicInfo = append(batchBasicInfo, starlingBasicInfo)
	}
	err := models.CreateBatch(batchBasicInfo)
	if err != nil {
		logging.Error("models.CreateBatch err, ", err)
		return err
	}
	return nil
}

func GetStarlingByCondition(key string, size string, num string, chineseText string, englishText string) ([]*models.StarlingBasicInfo, int, error) {
	pageSize, _ := strconv.Atoi(size)
	pageNum, _ := strconv.Atoi(num)
	condition := form.GetBasicInfoListByCondition{
		Key:         key,
		PageSize:    pageSize,
		PageNum:     pageNum,
		ChineseText: chineseText,
		EnglishText: englishText,
	}
	ret, total, err := models.GetStarlingBasicByCondition(condition)
	if err != nil {
		logging.Error("models.GetStarlingBasicByCondition err, ", err)
		return ret, total, err
	}
	return ret, total, nil
}

func GetStarlingBySearchKey(key string, size string, num string) ([]*models.StarlingBasicInfo, int, error) {
	pageSize, _ := strconv.Atoi(size)
	pageNum, _ := strconv.Atoi(num)
	condition := form.GetBasicInfoListBySearchKey{
		SearchKey: key,
		PageSize:  pageSize,
		PageNum:   pageNum,
	}
	ret, total, err := models.GetStarlingBasicBySearchKey(condition)
	//filterRet := make([]*models.StarlingBasicInfo, 0)
	//if len(ret) > 0 {
	//	filterRet = append(filterRet, ret[0])
	//}
	//for i := 1; i < len(ret); i++ {
	//	if ret[i].ID != ret[i-1].ID {
	//		filterRet = append(filterRet, ret[i])
	//	}
	//}
	if err != nil {
		logging.Error("models.GetStarlingBasicBySearchKey err, ", err)
		return ret, total, err
	}
	return ret, total, nil
}

func GetStarlingByKeyAndLang(req form.KeyLang) (map[string]string, error) {
	ret := make(map[string]string, 0)
	for _, key := range req.Key {
		err, text := models.GetStarBasicByKeyAndLang(key, req.Lang)
		if err != nil {
			logging.Error("models.GetStarBasicByKeyAndLang err, ", err)
			return ret, err
		}
		ret[key] = text

	}
	return ret, nil
}

func UpdateStarling(req form.NewBasicInfo) error {
	err, ret := models.GetStarBasicInfoByID(req.ID)
	if err != nil {
		logging.Error("models.GetStarBasicInfoByID err, ", err)
		return err
	}

	if ret.ID == 0 { // 数据库没有这条数据
		logging.Error("database doesn't exist this data, ", req)
		return errors.New(common.RECORD_NOT_FOUND)
	}

	res, err := models.GetStarByKey(req.Key)
	if err != nil {
		logging.Error("models.ExistStarByKey err, ", err)
		return err
	}

	for _, basic := range res {
		if int(basic.ID) != req.ID {
			logging.Error("key ", req.Key, " exists!")
			return errors.New(common.KEYEXISTS)
		}
	}

	err = models.UpdateStarling(&req)
	if err != nil {
		logging.Error("models.UpdateStarling err, ", err)
		return err
	}
	return nil
}

func DeleteStarling(id int) error {
	err := models.DeleteStarling(id, nil)
	if err != nil {
		logging.Error("models.DeleteStarling err, ", err)
		return err
	}
	return nil
}

func DeleteBatch(idStrList []string) error {
	if len(idStrList) == 0 {
		return nil
	}
	idList := common.String2Int(idStrList)
	err := models.DeleteBatch(idList)
	if err != nil {
		logging.Error("models.DeleteBatch err, ", err)
		return err
	}
	return nil
}

func ExistRepeatKey(key string) (bool, error) {
	isExist, err := models.ExistStarByKey(key)
	if err != nil {
		logging.Error("models.ExistStarByKey err, ", err)
		return true, err
	}
	return isExist, nil
}

func TranslateCh2En(text string) (string, error) {
	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=zh-cn&tl=en&dt=t&q=%s", url.QueryEscape(text))
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//返回的json反序列化比较麻烦, 直接字符串拆解
	ss := string(bs)
	ss = strings.ReplaceAll(ss, "[", "")
	ss = strings.ReplaceAll(ss, "]", "")
	ss = strings.ReplaceAll(ss, "null,", "")
	ss = strings.Trim(ss, `"`)
	ps := strings.Split(ss, `","`)
	return ps[0], nil
}

func TranslateEn2Ch(text string) (string, error) {
	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=en&tl=zh-cn&dt=t&q=%s", url.QueryEscape(text))
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//返回的json反序列化比较麻烦, 直接字符串拆解
	ss := string(bs)
	ss = strings.ReplaceAll(ss, "[", "")
	ss = strings.ReplaceAll(ss, "]", "")
	ss = strings.ReplaceAll(ss, "null,", "")
	ss = strings.Trim(ss, `"`)
	ps := strings.Split(ss, `","`)
	return ps[0], nil
}
