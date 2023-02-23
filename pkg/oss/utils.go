package oss

import (
	"bytes"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"
	"io/ioutil"
	"math/rand"
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
		avatarPath := Path + "/public/avatar/" + strconv.Itoa(i%10) + "-avatar.png"
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

func PublishBackgroundInit() error {
	for i := 1; i <= 10; i++ {
		avatarPath := Path + "/public/background/" + strconv.Itoa(i%10) + "-background.jpg"
		openFile, err := os.Open(avatarPath)
		if err != nil {
			klog.Errorf("open background file fail, %v", err.Error())
			return err
		}
		defer openFile.Close()
		avatarData, err := ioutil.ReadAll(openFile)
		if err != nil {
			klog.Errorf("read background file fail, %v", err.Error())
			return err
		}

		err = Bucket.PutObject("background/"+strconv.Itoa(i%10)+"-background.jpg", bytes.NewReader(avatarData))
		if err != nil {
			klog.Errorf("publish background fail, %v", err.Error())
			return err
		}
	}
	return nil
}

// GetAvatar 获取用户头像
func GetAvatar() string {
	avatarURL := "https://" + constants.OssBucket + "." + constants.OssEndPoint + "/avatar/" + strconv.Itoa(rand.Intn(10)) + "-avatar.png"
	return avatarURL
}

// GetBackground 获取用户头像
func GetBackground() string {
	avatarURL := "https://" + constants.OssBucket + "." + constants.OssEndPoint + "/background/" + strconv.Itoa(rand.Intn(10)) + "-background.jpg"
	return avatarURL
}
