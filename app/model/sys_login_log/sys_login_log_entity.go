// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package sys_login_log

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table sys_login_log.
type Entity struct {
    InfoId        int64  `orm:"info_id,primary" json:"info_id"`        // è®¿é—®ID                                 
    LoginName     string `orm:"login_name"      json:"login_name"`     // ç™»å½•è´¦å·                              
    Ipaddr        string `orm:"ipaddr"          json:"ipaddr"`         // ç™»å½•IPåœ°å€                            
    LoginLocation string `orm:"login_location"  json:"login_location"` // ç™»å½•åœ°ç‚¹                             
    Browser       string `orm:"browser"         json:"browser"`        // æµè§ˆå™¨ç±»åž‹                           
    Os            string `orm:"os"              json:"os"`             // æ“ä½œç³»ç»Ÿ                               
    Status        int    `orm:"status"          json:"status"`         // ç™»å½•çŠ¶æ€ï¼ˆ0æˆåŠŸ 1å¤±è´¥ï¼‰  
    Msg           string `orm:"msg"             json:"msg"`            // æç¤ºæ¶ˆæ¯                                
    LoginTime     int64  `orm:"login_time"      json:"login_time"`     // è®¿é—®æ—¶é—´                       
    Module        string `orm:"module"          json:"module"`         // ç™»å½•æ¨¡å—                             
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