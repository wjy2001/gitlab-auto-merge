## 使用方法

无图形化版本的执行方法

### 基础配置

> conf.json

```
{
  "parameter": {                              
    "basic_url": "https://gitlab.asants.com", //基础路径
    "token": "user-token"					  //你的token
  }
}
```

### 任务配置

> taskMapInfo.json

```
[
  {
    "project_ids": [											//项目id
      110
    ],
    "group_ids": [												//群组id 能和项目id混用
      0
    ],
    "source_branch": "dev",										//源分支
    "target_branch": "master",									//目标分支
    "title": "auto",											//自动marge的title
    "reviewer_id": [											//审核人的id
      0
    ],
    "interval_time": 3000000000,								//检测间隔(秒)
    "created_time": "2024-06-14T18:08:12.0354972+08:00",		//创建时间(无用)
    "enable": false												//是否开启
  }
]
```

