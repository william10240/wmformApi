package modules

import "time"

type ApiInfo struct {
	Id         int        `gorm:"primary_key"`
	apiName    string     `gorm:"name:api_name;size:255"`
	Protocol   string     `gorm:"name:protocol;sizee:8"`
	Hostname   string     // VARCHAR(255) NULL DEFAULT NULL COMMENT '主机名称，可以是域名，或者IP' COLLATE 'utf8mb4_general_ci',
	Port       string     // VARCHAR(255) NULL DEFAULT NULL COMMENT '端口' COLLATE 'utf8mb4_general_ci',
	CreateBy   string     // VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	CreateDate string     // DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
	UpdateBy   string     // VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	UpdateDate *time.Time // DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
	OfficeId   string     // VARCHAR(64) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	IsDelete   bool       // VARCHAR(1) NULL DEFAULT NULL COMMENT '删除标识0未删除，1 删除' COLLATE 'utf8mb4_general_ci',
	Remark     string     // VARCHAR(255) NULL DEFAULT NULL COMMENT '接口备注' COLLATE 'utf8mb4_general_ci',
}

func (ApiInfo) TableName() string { return "db_apiinfo" }
