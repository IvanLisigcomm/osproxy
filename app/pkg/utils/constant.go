package utils

import "time"

const (
	Scheme                = "http"
	WorkID                = "workId"
	LinkLimit             = 50
	EncryKey              = "*&^@#$storage"
	LocalStore            = "/storage/localstore"
	ServiceRedisPrefix    = "service:proxy"
	ServiceRedisTTl       = time.Second * 3 * 60
	S3StoragePutThreadNum = 10
	MultiPartDownload     = 10
)

// 任务类型
const (
	TaskPartMerge = "partMerge"
)

// 任务状态
const (
	TaskStatusUndo    = 0
	TaskStatusRunning = 1
	TaskStatusFinish  = 2
	TaskStatusError   = 99
)

// worker和队列
const (
	MaxWorker = 100
	MaxQueue  = 200
)
