# Rarity App Backend Http API

## 获取 Address下的Summoners

Request：

```text
HTTP请求

类型: GET

URL：/summoners/{address}

例子： /summoners/0x6D81145629f154dBf07fDAb18D2892818626FeCF

```

Responses：

```json
{
  "result": [
    {
      "ID": 0,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "DeletedAt": null,
      "address": "0x6d81145629f154dbf07fdab18d2892818626fecf",
      "token_id": 2063971,
      "class_id": "8",
      "level": "2",
      "xp": "2000000000000000000000",
      "next_level_xp": "3000000000000000000000",
      "gold": "1000000000000000000000",
      "gold_claimable": "0",
      "adventure_log": "1634149062",
      "dungeon_log": "0",
      "dungeon_scout": "0",
      "strength": 8,
      "dexterity": 18,
      "constitution": 15,
      "intelligence": 8,
      "wisdom": 15,
      "charisma": 8
    }
  ],
  "success": true
}
```

### 特殊说明

在Responses result 中 adventure_log & dungeon_log 可以判断是否是可以进行可冒险、可地牢

```js
if dayjs(new Date(adventure_log * 1000)).isBefore(dayjs(new Date())) {
    // 可冒险
}


if dayjs(new Date(dungeon_log * 1000)).isBefore(dayjs(new Date())) {
    // 可地牢
}
```

## 获取 Address 可以批量跑的任务数量

Request：

```text
HTTP请求

类型: GET

URL：/task/{address}

例子： /tasks/0x6D81145629f154dBf07fDAb18D2892818626FeCF

```

Responses：

```json
{
  "result": [
    {
      "task_type": "levelup",
      "ids": [],
      "count": 0
    },
    {
      "task_type": "goldclaim",
      "ids": [],
      "count": 0
    },
    {
      "task_type": "adventure",
      "ids": [],
      "count": 306
    },
    {
      "task_type": "dungeon",
      "ids": [],
      "count": 286
    }
  ],
  "success": true
}
```

### 特殊说明

Responses result 中 ids 不需要理会

暂时支持这4个批量任务，召唤&加点后续再加

## 判断用户是否给Caller进行授权(SetApprovedForAll权限)

Request：

```text
HTTP请求

类型: POST

URL：/approve/check

例子： 

{
    "address":"0x6D81145629f154dBf07fDAb18D2892818626FeCF"
}

```

Responses：

```json
{
  "result": true,
  "success": true
}
```

## 获取SetApprovedForAll 交易Input（tx data & param）

Request：

```text
HTTP请求

类型: GET

URL：/tx/approve

```

Responses：

```json
{
  "result": {
    "to": "0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb",
    "value": "0",
    "input": "0xa22cb4650000000000000000000000003b8d1f9569ec2b8c3dd76847b944bb1b0c998a380000000000000000000000000000000000000000000000000000000000000001"
  },
  "success": true
}
```

### 特殊说明

Responses result 中的input 就是js 构建tx时所需要的data, 只需要``hex decode``就行了