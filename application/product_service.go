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

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()

	product.Name = name
	product.Price = price

	if _, err := product.IsValid(); err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	if err := product.Enable(); err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, nil
	}

	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	if err := product.Disable(); err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, nil
	}

	return result, nil
}
