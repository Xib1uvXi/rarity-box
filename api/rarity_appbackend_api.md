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
        },
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
