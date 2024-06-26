package third_party

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

// UploadFile 通过七牛云上传文件
func UploadFile(file *multipart.FileHeader) (string, error) {

	fileSize := file.Size
	f, _ := file.Open()
	buf := make([]byte, file.Size)
	f.Read(buf)
	reader := bytes.NewReader(buf)

	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "n0w13re"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	// TODO: 从环境变量中获取
	mac := auth.New("7h-G7N7OKrSvtVbzrle4yKD5IbzSf5I0fOqjZD-g", "Y_W9CrWpO23k_gV-PE5k7RNqeX38kePTNBIyV-sQ")
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, reader, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return ret.Key, nil
}
