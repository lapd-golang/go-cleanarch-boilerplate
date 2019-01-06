package repositories

import "github.com/minhduccm/go-cleanarch-boilerplate/entities"

type CategoryMapper struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

type CateRepo struct {
	Database Database
}

func NewCateRepo(db Database) *CateRepo {
	return &CateRepo{Database: db}
}

func (cateRepo *CateRepo) FindCateById(cateId int) (*entities.Category, error) {
	rawCate, err := cateRepo.Database.QueryCateById(cateId)
	if err != nil {
		return nil, err
	}
	return &entities.Category{
		Id:   rawCate.Id,
		Name: rawCate.Name,
	}, nil
}

func (cateRepo *CateRepo) Store(cate *entities.Category) (*entities.Category, error) {
	cateMapper := &CategoryMapper{
		Id:   cate.Id,
		Name: cate.Name,
	}
	err := cateRepo.Database.InsertCate(cateMapper)
	if err != nil {
		return nil, err
	}
	return cate, nil
}
