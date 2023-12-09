package attribute

type Attribute struct {
	AttributeId string `gorm:"type:varchar;unique;notNull;primaryKey;column:attribute_id;default:uuid_generate_v4()"`
	ProductId   string `gorm:"type:varchar;column:product_id;unique;notNull"`
	Category    string `gorm:"type:varchar;column:category"`
	Color       string `gorm:"type:varchar;column:color"`
	Size        string `gorm:"type:varchar;column:size"`
	Brand       string `gorm:"type:varchar;column:brand"`
}
