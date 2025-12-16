-- 创建捐赠表
CREATE TABLE IF NOT EXISTS donations (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '捐赠用户ID',
    organization_id BIGINT UNSIGNED NULL COMMENT '捐赠给的机构ID',
    type VARCHAR(20) NOT NULL COMMENT '捐赠类型: money/supply/service',
    amount DECIMAL(10,2) DEFAULT 0 COMMENT '资金金额',
    supply_items TEXT NULL COMMENT '物资清单(JSON)',
    service_desc TEXT NULL COMMENT '服务描述',
    message TEXT NULL COMMENT '捐赠留言',
    is_anonymous TINYINT(1) DEFAULT 0 COMMENT '是否匿名',
    status TINYINT DEFAULT 0 COMMENT '状态: 0待确认 1已确认 2已完成 3已取消',
    payment_method VARCHAR(50) NULL COMMENT '支付方式',
    transaction_id VARCHAR(100) NULL COMMENT '交易ID',
    confirmed_at TIMESTAMP NULL COMMENT '确认时间',
    confirmed_by BIGINT UNSIGNED NULL COMMENT '确认人ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_donations_user_id (user_id),
    INDEX idx_donations_org_id (organization_id),
    INDEX idx_donations_type (type),
    INDEX idx_donations_status (status),
    INDEX idx_donations_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='捐赠表';

-- 查看表结构
DESC donations;
