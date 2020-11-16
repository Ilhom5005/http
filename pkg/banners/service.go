package banners

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sync"
)

var (
	errItemNotFound = errors.New("item not found")
)

var curID int64 = 0

// Service for control banners
type Service struct {
	mu    sync.RWMutex
	items []*Banner
}

// NewService make service
func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
}

// All return all banners
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items, nil
}

// ByID return banner by ID
func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}
	return nil, errItemNotFound
}

// Save save/updata banner
func (s *Service) Save(ctx context.Context, item *Banner, file multipart.File) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if item.ID == 0 {
		curID++
		item.ID = curID
		item, err := SaveImage(item, file)
		if err != nil {
			return nil, err
		}
		s.items = append(s.items, item)
		return item, nil
	}
	if item.Image != "" {
		var err error
		item, err = SaveImage(item, file)
		if err != nil {
			return nil, err
		}
	}
	for i := 0; i < len(s.items); i++ {
		if s.items[i].ID == item.ID {
			if item.Image == "" {
				item.Image = s.items[i].Image
			}
			s.items[i] = item
			return item, nil
		}
	}
	return nil, errItemNotFound
}

// SaveImage save/update image
func SaveImage(item *Banner, src multipart.File) (*Banner, error) {
	item.Image = fmt.Sprint(item.ID) + "." + item.Image
	file, err := os.Create("./web/banners/" + item.Image)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(file, src)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// RemoveByID remove banner by ID
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, err := s.ByID(ctx, id)
	if err != nil {
		return nil, err
	}
	var index int
	for i := 0; i < len(s.items); i++ {
		if s.items[i].ID == id {
			index = i
			break
		}
	}
	s.items = append(s.items[:index], s.items[index+1:]...)
	return item, nil
}

// Banner represent banner
type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
	Image   string
}
