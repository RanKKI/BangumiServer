// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNamePersonSubjects = "chii_person_cs_index"

// PersonSubjects mapped from table <chii_person_cs_index>
type PersonSubjects struct {
	PrsnType      string  `gorm:"column:prsn_type;type:enum('prsn','crt');primaryKey"`
	PersonID      uint32  `gorm:"column:prsn_id;type:mediumint(9) unsigned;primaryKey"`
	PrsnPosition  uint16  `gorm:"column:prsn_position;type:smallint(5) unsigned;primaryKey"` // 监督，原案，脚本,..
	SubjectID     uint32  `gorm:"column:subject_id;type:mediumint(9) unsigned;primaryKey"`
	SubjectTypeID uint8   `gorm:"column:subject_type_id;type:tinyint(4) unsigned;not null"`
	Summary       string  `gorm:"column:summary;type:mediumtext;not null"`
	PrsnAppearEps string  `gorm:"column:prsn_appear_eps;type:mediumtext;not null"` // 可选，人物参与的章节
	Subject       Subject `gorm:"foreignKey:subject_id;references:subject_id" json:"subject"`
	Person        Person  `gorm:"foreignKey:prsn_id;references:prsn_id" json:"person"`
}

// TableName PersonSubjects's table name
func (*PersonSubjects) TableName() string {
	return TableNamePersonSubjects
}
