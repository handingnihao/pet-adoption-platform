#!/bin/bash
# 修复宠物模块类型问题的脚本

echo "正在修复类型不匹配问题..."

# 说明：将Pet模型中的UserID等字段从uint改为int64以匹配User.ID
echo "
Pet模型中的类型已修改为：
- UserID: int64 (原uint)
- AdoptedBy: *int64 (原*uint)  
- ReviewedBy: *int64 (原*uint)

Service和DAO层的函数签名也需要相应修改。

请手动执行以下SQL修改数据库表结构：

ALTER TABLE pets MODIFY COLUMN user_id BIGINT NOT NULL COMMENT '发布用户ID';
ALTER TABLE pets MODIFY COLUMN adopted_by BIGINT COMMENT '领养人ID';
ALTER TABLE pets MODIFY COLUMN reviewed_by BIGINT COMMENT '审核人ID';
"
