// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package users

// RecordNotFound 根据条件判断数据是否存在
// 有数据返回false
// 没数据 true
func RecordNotFound(where ...interface{}) bool {
	return Model.RecordNotFound(where...)
}

func (m *arModel) RecordNotFound(where ...interface{}) bool {
	r, err := m.M.FindOne(where...)
	if r == nil && err == nil {
		return true
	}
	return false
}
