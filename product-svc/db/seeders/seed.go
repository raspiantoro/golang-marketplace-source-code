package seeders

import "gorm.io/gorm"

type SeederFunc func(tx *gorm.DB) (err error)
type Seeders struct {
	db      *gorm.DB
	seeders []SeederFunc
}

type SeedersOption func(s *Seeders)

func WithDefaultSeeders() SeedersOption {
	return func(s *Seeders) {
		s.seeders = append(s.seeders, ProductCategorySeed)
	}
}

func NewSeeders(db *gorm.DB, opts ...SeedersOption) *Seeders {
	s := &Seeders{db: db}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Seeders) Add(fn SeederFunc) {
	s.seeders = append(s.seeders, fn)
}

func (s *Seeders) Seed() error {
	for _, seed := range s.seeders {
		if err := seed(s.db); err != nil {
			return err
		}
	}

	return nil
}
