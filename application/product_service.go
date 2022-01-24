package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(ID string) (ProductInterface, error) {
	product, err := s.Persistence.Get(ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Save(product ProductInterface) (ProductInterface, error) {
	panic("not implemented") // TODO: Implement
}
