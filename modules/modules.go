package modules

import "time"

//[db_apiinfo]
type ApiInfo struct {
	Id         int    `gorm:"primary_key"`
	apiName    string `gorm:"name:api_name;size:255"`
	Protocol   string `gorm:"name:protocol;sizee:8"`
	Hostname   string `gorm:"size:255"`
	Port       string `gorm:"size:255"`
	CreateBy   string `gorm:"size:255"`
	CreateDate *time.Time
	UpdateBy   string `gorm:"size:255"`
	UpdateDate *time.Time
	OfficeId   string `gorm:"size:64"`
	IsDelete   bool   `gorm:"size:1"`
	Remark     string `gorm:"size:255"`
}

func (ApiInfo) TableName() string { return "db_apiinfo" }

//[db_datacontrol]
type DataControl struct {
	Id             string `gorm:"size:64"`
	FormId         string `gorm:"size:64"`
	ControlId      string `gorm:"size:64"`
	DatasetId      string `gorm:"size:64"`
	DatasourceId   string `gorm:"size:64"`
	DatasetField   string `gorm:"size:255"`
	GlobalParmname string `gorm:"size:255"`
	DataFromType   string `gorm:"size:1"`
	PolicyType     string `gorm:"size:2"`
	IsDelete       string `gorm:"size:2"`
	BindAttribute  string `gorm:"size:255"`
	FieldTable     string `gorm:"size:255"`
	AnotherName    string `gorm:"size:255"`
	DatacontrolSql string `gorm:"size:255"`
	DatabaseName   string `gorm:"size:255"`
	FieldAlias     string `gorm:"size:255"`
	OfficeId       string `gorm:"size:64"`
	UpdateDate     *time.Time
	UpdateBy       string `gorm:"size:255"`
	CreateDate     *time.Time
	CreateBy       string `gorm:"size:255"`
}

func (DataControl) TableName() string { return "db_datacontrol" }

//[db_dataset]
type DataSet struct {
	Id           string `gorm:"size:64"`
	DatasetName  string `gorm:"size:64"`
	DatasetTable string `gorm:"size:255"`
	PrimaryField string `gorm:"size:255"`
	DatasourceId string `gorm:"size:64"`
	DatabaseName string `gorm:"size:255"`
	IsDelete     string `gorm:"size:1"`
	OfficeId     string `gorm:"size:64"`
	CreateBy     string `gorm:"size:255"`
	CreateDate   *time.Time
	UpdateBy     string `gorm:"size:255"`
	UpdateDate   *time.Time
	FormId       string `gorm:"size:64"`
}

func (DataSet) TableName() string { return "db_dataset" }

//[db_datasource]
type DataSource struct {
	Id          string `gorm:"size:64"`
	Name        string `gorm:"size:255"`
	Type        int    `gorm:"size:11"`
	IsDelete    string `gorm:"size:1"`
	Ip          string `gorm:"size:255"`
	Port        int    `gorm:"size:11"`
	DbName      string `gorm:"size:255"`
	Username    string `gorm:"size:255"`
	Password    string `gorm:"size:255"`
	Description string `gorm:"size:255"`
	CreateBy    string `gorm:"size:255"`
	CreateDate  *time.Time
	UpdateBy    string `gorm:"size:255"`
	UpdateDate  *time.Time
	OfficeId    string `gorm:"size:64"`
}

func (DataSource) TableName() string { return "db_datasource" }

//[form_application_detail]
type ApplicationDetail struct {
	Id              string `gorm:"size:64"`
	Code            string `gorm:"size:64"`
	Html            string
	Css             string
	Js              string
	Json            string
	Script          string
	ApplicationType string `gorm:"size:1"`
}

func (ApplicationDetail) TableName() string { return "form_application_detail" }

//[form_application_summary]
type ApplicationSummary struct {
	Id              string `gorm:"size:64"`
	Code            string `gorm:"size:255"`
	IsDelete        string `gorm:"size:1"`
	CreateBy        string `gorm:"size:255"`
	CreateDate      *time.Time
	UpdateBy        string `gorm:"size:255"`
	UpdateDate      *time.Time
	OfficeId        string `gorm:"size:255"`
	ApplicationType string `gorm:"size:1"`
	Name            string `gorm:"size:255"`
	Remark          string `gorm:"size:1024"`
}

func (ApplicationSummary) TableName() string { return "form_application_summary" }

//[form_associated_apiinfo]
type ApplicationApiinfo struct {
	Id        string `gorm:"size:64"`
	ApiInfoId string `gorm:"size:64"`
	FormId    string `gorm:"size:64"`
}

func (ApplicationApiinfo) TableName() string { return "form_associated_apiinfo" }

//[form_associated_datasource]
type ApplicationDatasource struct {
	Id           string `gorm:"size:64"`
	FormId       string `gorm:"size:64"`
	DatasourceId string `gorm:"size:64"`
}

func (ApplicationDatasource) TableName() string { return "form_associated_datasource" }

//[form_associated_group]
type AssociatedGroup struct {
	Id            string `gorm:"size:64"`
	FormGroupCode string `gorm:"size:64"`
	FormId        string `gorm:"size:64"`
	Sort          int    `gorm:"size:11"`
	DefaultExpand string `gorm:"size:1"`
	DefaultLoad   string `gorm:"size:1"`
}

func (AssociatedGroup) TableName() string { return "form_associated_group" }

//[form_bussiness_val]
type BussinessVal struct {
	Id         string `gorm:"size:64"`
	FormId     string `gorm:"size:64"`
	ControlKey string `gorm:"size:64"`
	ControlVal string `gorm:"size:2048"`
	TaskInstId string `gorm:"size:64"`
	ProcInstId string `gorm:"size:64"`
	CreateBy   string `gorm:"size:64"`
	CreateDate *time.Time
	UpdateBy   string `gorm:"size:64"`
	UpdateDate *time.Time
	IsDelete   string `gorm:"size:1"`
	OfficeId   string `gorm:"size:64"`
	Code       string `gorm:"size:64"`
	CompanyId  string `gorm:"size:64"`
}

func (BussinessVal) TableName() string { return "form_bussiness_val" }

//[form_group]
type Group struct {
	Id          string `gorm:"size:64"`
	Code        string `gorm:"size:64"`
	GroupName   string `gorm:"size:255"`
	GroupStyle  string `gorm:"size:1"`
	TabPosition string `gorm:"size:1"`
	CreateBy    string `gorm:"size:255"`
	CreateDate  *time.Time
	UpdateBy    string `gorm:"size:255"`
	UpdateDate  *time.Time
	IsDelete    string `gorm:"size:1"`
	Officeid    string `gorm:"size:255"`
}

func (Group) TableName() string { return "form_group" }

//[form_page_detail]
type PageDetail struct {
	Id      string `gorm:"size:64"`
	Code    string `gorm:"size:64"`
	Html    string
	Css     string
	Js      string
	Json    string
	Script  string
	Version int `gorm:"size:11"`
}

func (PageDetail) TableName() string { return "form_page_detail" }

//[form_page_summary]
type PageSummary struct {
	Id         string `gorm:"size:64"`
	Code       string `gorm:"size:64"`
	IsDelete   string `gorm:"size:1"`
	CreateBy   string `gorm:"size:255"`
	CreateDate *time.Time
	UpdateBy   string `gorm:"size:255"`
	UpdateDate *time.Time
	OfficeId   string `gorm:"size:64"`
	Name       string `gorm:"size:255"`
	Remark     string `gorm:"size:1024"`
	Category   string `gorm:"size:255"`
	Version    int    `gorm:"size:11"`
	IsNew      string `gorm:"size:2"`
}

func (PageSummary) TableName() string { return "form_page_summary" }

//[sys_usertest]
type UserTest struct {
	Id    string `gorm:"size:64"`
	Name  string `gorm:"size:255"`
	Age   int    `gorm:"size:11"`
	Email string `gorm:"size:255"`
	Phone string `gorm:"size:255"`
}

func (UserTest) TableName() string { return "sys_usertest" }
