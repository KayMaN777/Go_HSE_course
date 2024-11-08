package store

import (
	"task1/internal/idgenerator"
	"task1/internal/model"
)

type MapStore struct {
	bookMap map[uint32]model.Book
}

func NewMapStore() *MapStore {
	return &MapStore{
		bookMap: make(map[uint32]model.Book),
	}
}

func (store *MapStore) Add(id uint32, book model.Book) {
	book.Id = id
	store.bookMap[id] = book
}

func (store *MapStore) Search(id uint32) (model.Book, bool) {
	book, found := store.bookMap[id]
	return book, found
}

func (store *MapStore) Remove(id uint32) {
	delete(store.bookMap, id)
}

func (store *MapStore) Regenerate(generator idgenerator.BookIdGenerator) {
	newMap := make(map[uint32]model.Book)
	for _, book := range store.bookMap {
		newMap[generator.GenerateID(book.Title)] = book
	}
	store.bookMap = newMap
}
