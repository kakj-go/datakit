// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type TaskInfo struct {

    /* 任务id (Optional) */
    TaskId int `json:"taskId"`

    /* 任务操作类型 (Optional) */
    Action string `json:"action"`

    /* 任务状态，pending,running,failed,finished (Optional) */
    TaskStatus string `json:"taskStatus"`

    /* 任务进度，0-100 (Optional) */
    Progress int `json:"progress"`

    /* 额外信息 (Optional) */
    Message string `json:"message"`

    /* 任务创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 任务完成时间 (Optional) */
    FinishedTime string `json:"finishedTime"`
}
