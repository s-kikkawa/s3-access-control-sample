# s3-access-control-sample
AWS S3 のフォルダ権限が有効か確認するサンプルです

## 事前準備
* IAM ユーザーを2つ作成する
* S3のバケットを作成し、フォルダを2つ作成する
* フォルダに各々のファイルを1つアップロードする
* IAM ユーザーそれぞれにポリシーを設定してどちらかのフォルダにしかアクセスできないようにする
* サンプルプログラムを実行する

## IAMユーザのポリシーサンプル（ユーザ1）

バケット名：`bucket-name` 
フォルダ名：`test1`
の場合
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowUsersToAccessFolder1Only",
            "Effect": "Allow",
            "Action": [
                "s3:GetObject*",
                "s3:PutObject*"
            ],
            "Resource": [
                "arn:aws:s3:::bucket-name/test1/*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:ListBucket*"
            ],
            "Resource": [
                "arn:aws:s3:::bucket-name"
            ],
            "Condition": {
                "StringLike": {
                    "s3:prefix": [
                        "test1/*"
                    ]
                }
            }
        }
    ]
}
```

`https://s3.console.aws.amazon.com/s3/buckets/bucket-name/test1/`に直接アクセスしてファイルが読めることを確認する
また
`https://s3.console.aws.amazon.com/s3/buckets/bucket-name/test2/`に直接アクセスして権限がないことを確認する

## IAMユーザのポリシーサンプル（ユーザ2）

バケット名：`bucket-name` 
フォルダ名：`test2`
の場合
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowUsersToAccessFolder2Only",
            "Effect": "Allow",
            "Action": [
                "s3:GetObject*",
                "s3:PutObject*"
            ],
            "Resource": [
                "arn:aws:s3:::bucket-name/test2/*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:ListBucket*"
            ],
            "Resource": [
                "arn:aws:s3:::bucket-name"
            ],
            "Condition": {
                "StringLike": {
                    "s3:prefix": [
                        "test2/*"
                    ]
                }
            }
        }
    ]
}
```
ユーザ1の逆であることを確認する
