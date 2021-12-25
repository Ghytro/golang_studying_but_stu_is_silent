package util

import (
	"errors"
)

type LongURL struct {
	Url             string
	ExpireTimestamp int64
}

type UrlStorage struct {
	storage map[string]LongURL
}

func (s UrlStorage) AddUrl(longUrl string, shortUrl string, expireTimestamp int64) error {
	_, exists := s.storage[shortUrl]
	if exists {
		return errors.New("this short url was already added, try some other url")
	}
	s.storage[shortUrl] = LongURL{longUrl, expireTimestamp}
	return nil
}

func (s UrlStorage) GetLongUrl(shortUrl string) (LongURL, error) {
	result, ok := s.storage[shortUrl]
	if !ok {
		return LongURL{}, errors.New("no such short url in storage")
	}
	return result, nil
}

func NewUrlStorage() *UrlStorage {
	ret := new(UrlStorage)
	ret.storage = make(map[string]LongURL)
	return ret
}
