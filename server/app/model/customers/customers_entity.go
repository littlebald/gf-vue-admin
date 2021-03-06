// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package customers

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table customers.
type Entity struct {
    Id                 uint        `orm:"id,primary"            json:"id"`                    // 自增ID        
    CreateAt           *gtime.Time `orm:"create_at"             json:"create_at"`             // 创建时间      
    UpdateAt           *gtime.Time `orm:"update_at"             json:"update_at"`             // 更新时间      
    DeleteAt           *gtime.Time `orm:"delete_at"             json:"delete_at"`             // 删除时间      
    CustomerName       string      `orm:"customer_name"         json:"customer_name"`         // 客户名        
    CustomerPhoneData  string      `orm:"customer_phone_data"   json:"customer_phone_data"`   // 客户电话      
    SysUserId          uint        `orm:"sys_user_id"           json:"sys_user_id"`           // 负责员工id    
    SysUserAuthorityId string      `orm:"sys_user_authority_id" json:"sys_user_authority_id"` // 负责员工角色  
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}