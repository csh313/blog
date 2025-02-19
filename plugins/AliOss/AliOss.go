package AliOss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"server/global"
)

var client *oss.Client // 全局变量用来存储OSS客户端实例
var endpoint string
var accessKeyId string
var accessKeySecret string
var bucketName string
var region string

func InitAliOss() {
	if !global.Config.AliyunOSS.Enable {
		log.Printf("未开启阿里云Oss")
		return
	}
	endpoint = global.Config.AliyunOSS.Endpoint
	accessKeyId = global.Config.AliyunOSS.AccessKeyID
	accessKeySecret = global.Config.AliyunOSS.AccessKeySecret
	bucketName = global.Config.AliyunOSS.BucketName
	region = global.Config.AliyunOSS.Region
	var err error
	client, err = oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// 输出客户端信息。
	log.Printf("Client: %#v\n", client)
	//log.Printf("Bucket: %#v\n", bucketName)
	//获取存储空间。
	//bucket, err = client.Bucket(bucketName)
	//if err != nil {
	//	handleError(err)
	//}
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。

}

// handleError 用于处理不可恢复的错误，并记录错误信息后终止程序。
func HandleError(err error) {
	log.Fatalf("Error: %v", err)
}

// createBucket 用于创建一个新的OSS存储空间。
// 参数：
//
//	bucketName - 存储空间名称。
//	endpoint - Bucket对应的Endpoint。
//
// 如果成功，记录成功日志；否则，返回错误。
func createBucket(bucketName string) error {
	// 创建存储空间。
	err := client.CreateBucket(bucketName)
	if err != nil {
		return err
	}

	// 存储空间创建成功后，记录日志。
	log.Printf("Bucket created successfully: %s", bucketName)
	return nil
}

// uploadFile 用于将本地文件上传到OSS存储桶。
// 参数：
//
//	bucketName - 存储空间名称。
//	objectName - Object完整路径，完整路径中不包含Bucket名称。
//	localFileName - 本地文件的完整路径。
//	endpoint - Bucket对应的Endpoint。
//
// 如果成功，记录成功日志；否则，返回错误。
func UploadFile(bucketName, objectName, localFileName string) (string, error) {

	if client == nil {
		return "client为空", fmt.Errorf("OSS client is not initialized")
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "bucket获取失败", err
	}

	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return "oss上传失败", err
	}

	// 文件上传成功后，记录日志。
	ossURL := fmt.Sprintf("https://%s.%s.aliyuncs.com/%s", bucketName, region, objectName)
	log.Printf("File uploaded successfully to %s/%s", bucketName, objectName)
	return ossURL, nil
}

// downloadFile 用于从OSS存储桶下载一个文件到本地路径。
// 参数：
//
//	bucketName - 存储空间名称。
//	objectName - Object完整路径，完整路径中不能包含Bucket名称。
//	downloadedFileName - 本地文件的完整路径。
//	endpoint - Bucket对应的Endpoint。
//
// 如果成功，记录成功日志；否则，返回错误。
func DownloadFile(bucketName, objectName, downloadedFileName string) error {
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 下载文件。
	err = bucket.GetObjectToFile(objectName, downloadedFileName)
	if err != nil {
		return err
	}

	// 文件下载成功后，记录日志。
	log.Printf("File downloaded successfully to %s", downloadedFileName)
	return nil
}

// listObjects 用于列举OSS存储空间中的所有对象。
// 参数：
//
//	bucketName - 存储空间名称。
//	endpoint - Bucket对应的Endpoint。
//
// 如果成功，打印所有对象；否则，返回错误。
func ListObjects(bucketName string) error {
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 列举文件。
	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			return err
		}

		// 打印列举文件，默认情况下一次返回100条记录。
		for _, object := range lsRes.Objects {
			log.Printf("Object: %s", object.Key)
		}

		if !lsRes.IsTruncated {
			break
		}
		marker = lsRes.NextMarker
	}

	return nil
}

// deleteObject 用于删除OSS存储空间中的一个对象。
// 参数：
//
//	bucketName - 存储空间名称。
//	objectName - 要删除的对象名称。
//	endpoint - Bucket对应的Endpoint。
//
// 如果成功，记录成功日志；否则，返回错误。
func DeleteObject(bucketName, objectName string) error {
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 删除文件。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		return err
	}

	// 文件删除成功后，记录日志。
	log.Printf("Object deleted successfully: %s/%s", bucketName, objectName)
	return nil
}
