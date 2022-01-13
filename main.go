package main

import (
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

const (
    REGION = "ap-northeast-1"
    BUCKET_NAME = "bucket-name"
    USER1_ACCESS_KEY = "XXXXXXXXXXXXXXXXXXXX"
    USER1_SECRET_KEY = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    USER2_ACCESS_KEY = "XXXXXXXXXXXXXXXXXXXX"
    USER2_SECRET_KEY = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

)

func main() {
    download(USER1_ACCESS_KEY, USER1_SECRET_KEY, "test1", "test1.txt")
    download(USER1_ACCESS_KEY, USER1_SECRET_KEY, "test2", "test2.txt")
    download(USER2_ACCESS_KEY, USER2_SECRET_KEY, "test1", "test1.txt")
    download(USER2_ACCESS_KEY, USER2_SECRET_KEY, "test2", "test2.txt")
}

func download(accessKey string, secretKey string, folder string, file string) {
    log.Printf("/" + folder + "/" + file + "の取得を試みる")
    creds := credentials.NewStaticCredentials(accessKey, secretKey, "")

    sess := session.Must(session.NewSession(&aws.Config{
        Credentials: creds,
        Region:      aws.String(REGION),
    }))

    svc := s3.New(sess)
    obj, err := svc.GetObject(&s3.GetObjectInput{
                Bucket: aws.String(BUCKET_NAME),
                Key:    aws.String("/" + folder + "/" + file),
    })
        if err != nil {
        log.Printf(err.Error())
        return
    }
    
    rc := obj.Body
    defer rc.Close()
    buf := make([]byte, 9)
    _, err = rc.Read(buf)
    if err != nil {
            log.Printf(err.Error())
    }else{
        log.Printf("%s", buf)
    }
}