package oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"
	"io/ioutil"
	"os"
	"strconv"
)

// PublishVideoToPublic 将视频保存到本地文件夹中
func PublishVideoToPublic(video []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		klog.Errorf("create %v fail, %v", filePath, err.Error())
		return err
	}
	defer file.Close()
	_, err = file.Write(video)
	if err != nil {
		klog.Errorf("write file fail, %v", err.Error())
		return err
	}
	return nil
}

// PublishVideoToOss 分片将视频上传到Oss
func PublishVideoToOss(objectKey string, filePath string) error {
	err := Bucket.UploadFile(objectKey, filePath, 1024*1024, oss.Routines(3))
	if err != nil {
		klog.Errorf("publish %v to Oss fail, %v ", filePath, err.Error())
		return err
	}
	return nil
}

// QueryOssVideoURL 从oss上获取播放地址
func QueryOssVideoURL(objectKey string) (string, error) {
	signedURL, err := Bucket.SignURL(objectKey, oss.HTTPPut, 60)
	if err != nil {
		klog.Errorf("Query %v Video URL fail, %v", objectKey, err.Error())
		return "", err
	}
	return signedURL, nil
}

// PublishCoverToOss 上传封面到Oss
func PublishCoverToOss(objectKey string, coverReader *bytes.Reader) error {
	err := Bucket.PutObject(objectKey, coverReader)
	if err != nil {
		klog.Errorf("publish %v to Oss fail, %v ", objectKey, err.Error())
		return err
	}
	return nil
}

// QueryOssCoverURL 从oss上获取封面地址
func QueryOssCoverURL(objectKey string) (string, error) {
	signedURL, err := Bucket.SignURL(objectKey, oss.HTTPPut, 60)
	if err != nil {
		klog.Errorf("Query %v Cover URL fail, %v", objectKey, err.Error())
		return "", err
	}
	return signedURL, nil
}

func PublishAvatarInit() error {
	for i := 1; i <= 10; i++ {
		avatarPath := Path + strconv.Itoa(i%10) + "-avatar.png"
		openFile, err := os.Open(avatarPath)
		if err != nil {
			klog.Errorf("open avatar file fail, %v", err.Error())
			return err
		}
		defer openFile.Close()
		avatarData, err := ioutil.ReadAll(openFile)
		if err != nil {
			klog.Errorf("read avatar file fail, %v", err.Error())
			return err
		}

		err = Bucket.PutObject("avatar/"+strconv.Itoa(i%10)+"-avatar.png", bytes.NewReader(avatarData))
		if err != nil {
			klog.Errorf("publish avatar fail, %v", err.Error())
			return err
		}
	}
	return nil
}

// GetAvatar 获取用户头像
func GetAvatar(userId int) string {
	avatarURL := "avatar/" + strconv.Itoa(userId%10) + "-avatar.png"
	signedURL, err := Bucket.SignURL(avatarURL, oss.HTTPPut, 60)
	if err != nil {
		klog.Errorf("Query %v Cover URL fail, %v", avatarURL, err.Error())
		return ""
	}
	return signedURL
}
