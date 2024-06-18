## 使用方法

无图形化版本的执行方法

### 基础配置

> conf.json

```json
{
  "parameter": {
    "basic_url": "https://gitlab.asants.com",
    "token": "your-token"
  },
  "project_blacklist": {
    "101": "bas-common-url-fuzzy",
    "102": "bas-common-url-spider",
    "153": "bas-upgrade",
    "166": "rule-upgrade-tool",
    "170": "privatization",
    "183": "server-upgrade-tool",
    "277": "bas-openapi-gateway",
    "279": "privatization"
  }
}
```

| 字段              | 说明                                 |
| ----------------- | ------------------------------------ |
| basic_url         | 基础路径                             |
| token             | 你的token                            |
| project_blacklist | 项目黑名单map  key:项目id value:备注 |

### 任务配置

> taskMapInfo.json

```json
[
  {
    "project_ids": [										
      110
    ],
    "group_ids": [],
    "source_branch": "dev",										
    "target_branch": "test",									
    "title": "auto",											
    "reviewer_id": [											
      102
    ],
    "interval_time": 60,								
    "created_time": "2024-06-14T18:08:12.0354972+08:00",
    "remove_source_branch": false,
    "enable": false												
  }
]
```

| 字段                 | 说明              |
| -------------------- |-----------------|
| project_ids          | 项目id            |
| group_ids            | 群组id 能和项目id混用   |
| source_branch        | 源分支             |
| target_branch        | 目标分支            |
| title                | 自动创建merge的title |
| reviewer_id          | 审核人的id          |
| interval_time        | 检测间隔(秒)         |
| created_time         | 创建时间(无用)        |
| enable               | 是否开启            |
| remove_source_branch | 是否删除源分支         |

