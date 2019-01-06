package repositories

type Database interface {
	QueryBooksByCateId(int) ([]*BookMapper, error)
	QueryCateById(int) (*CategoryMapper, error)
	InsertBook(*BookMapper) error
	InsertCate(*CategoryMapper) error
}
