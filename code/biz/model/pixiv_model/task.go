package pixiv_model

import (
	"time"
)

type DownloadTaskStatus int64

const (
	MaxRetryTime = 3
)

const (
	DownloadTaskStatusPending = 1 // 待执行
	DownloadTaskStatusRunning = 2 // 执行中
	DownloadTaskStatusFail    = 3 // 失败
	DownloadTaskStatusDead    = 4 // 死信
	DownloadTaskStatusSuccess = 5 // 成功
)

type DownloadTask struct {
	IllustId     string             `bson:"illust_id" json:"illust_id"`
	R18Index     int                `bson:"r18_index" json:"r18_index"`
	Status       DownloadTaskStatus `bson:"status" json:"status"`
	CreateTime   time.Time          `bson:"create_time" json:"create_time"`
	UpdateTime   time.Time          `bson:"update_time" json:"update_time"`
	RetryTime    int                `bson:"retry_time" json:"retry_time"`
	LastRunTime  time.Time          `bson:"last_run_time" json:"last_run_time"`
	LastRunError string             `bson:"last_run_error" json:"last_run_error"`
}

func NewDownloadTask(illustId string, r18Index int) *DownloadTask {
	return &DownloadTask{
		IllustId:   illustId,
		R18Index:   r18Index,
		Status:     DownloadTaskStatusPending,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

func (d *DownloadTask) StartToDownload() *DownloadTask {
	d.LastRunError = ""
	d.LastRunTime = time.Now()
	d.UpdateTime = time.Now()
	d.RetryTime++
	d.Status = DownloadTaskStatusRunning
	return d
}

func (d *DownloadTask) Success() *DownloadTask {
	d.LastRunError = ""
	d.UpdateTime = time.Now()
	d.Status = DownloadTaskStatusSuccess
	return d
}

func (d *DownloadTask) Fail(err error) *DownloadTask {
	if err != nil {
		d.LastRunError = err.Error()
	}
	d.UpdateTime = time.Now()
	if d.RetryTime >= MaxRetryTime {
		d.Status = DownloadTaskStatusDead
	} else {
		d.Status = DownloadTaskStatusFail
	}
	return d
}
