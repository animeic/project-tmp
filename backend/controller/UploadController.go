package controller

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"animeii.tech/dd"
	"animeii.tech/response"
	"animeii.tech/utils"

	"github.com/gin-gonic/gin"
)

type UploadController struct{}

var UD = new(UploadController)

func (uD *UploadController) Image(c *gin.Context) {
	mfor, _ := c.MultipartForm()
	file := mfor.File
	header, ok := file["image"]
	if !ok {
		response.Fail(c, response.ImageParamErr)
		return
	}
	// 确定是单文件
	hd := header[0]
	if hd.Size > 10240000 {
		response.Fail(c, response.ImageSizeErr)
		return
	}
	strs := strings.Split(hd.Filename, ".")
	if len(strs) < 2 {
		response.Fail(c, response.UploadFileErr)
		return
	}

	suffix := strs[len(strs)-1]
	suffixs := map[string]struct{}{
		"jpeg": {},
		"jpg":  {},
		"png":  {},
		"webp": {},
	}
	if _, ok := suffixs[suffix]; !ok {
		response.Fail(c, response.ImageSuffixErr)
		return

	}

	// 生成图片名
	key := fmt.Sprintf("%s%d", hd.Filename, time.Now().Unix())
	fmt.Println(key)
	prefix := utils.RandStr(utils.Sha256(key))
	fnm := prefix + "." + suffix
	fmt.Println(fnm)

	f, err := hd.Open()
	if err != nil {
		panic(err)
	}

	defer f.Close()
	upload_root_dir := os.Getenv("UPLOAD_ROOT_DIR")

	if upload_root_dir == "" {
		upload_root_dir = "./images"
	}
	dd.Log.Info(upload_root_dir)
	localPath := upload_root_dir + "/" + fnm
	fmt.Println(localPath)
	out, err := os.Create(localPath)
	if err != nil {
		panic(err)

	}
	defer out.Close()
	_, err = io.Copy(out, f)
	if err != nil {
		panic(err)
	}
	imagurl := dd.Cf.Asset.PrefixHost + "/" + fnm

	response.Success(c, response.UploadImagSucc, imagurl)

}
