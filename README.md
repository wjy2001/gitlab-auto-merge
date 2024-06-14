## 使用方法

无图形化版本的执行方法

### 基础配置

> conf.json

```json
{
  "parameter": {                              
    "basic_url": "https://gitlab.asants.com", 
    "token": "user-token"					  
  }
}
```

| basic_url | 基础路径  |
| --------- | --------- |
| token     | 你的token |



### 任务配置

> taskMapInfo.json

```json
[
  {
    "project_ids": [										
      110
    ],
    "group_ids": [											
      0
    ],
    "source_branch": "dev",										
    "target_branch": "master",									
    "title": "auto",											
    "reviewer_id": [											
      0
    ],
    "interval_time": 3,								
    "created_time": "2024-06-14T18:08:12.0354972+08:00",		
    "enable": false												
  }
]
```

| project_ids   | 项目id                |
| ------------- | --------------------- |
| group_ids     | 群组id 能和项目id混用 |
| source_branch | 源分支                |
| target_branch | 目标分支              |
| title         | 自动marge的title      |
| reviewer_id   | 审核人的id            |
| interval_time | 检测间隔(秒)          |
| created_time  | 创建时间(无用)        |
| enable        | 是否开启              |
|               |                       |

