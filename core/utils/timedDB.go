package utils

import "time"

type dataItem struct {
	data        interface{}
	editTime    time.Time
	lifeTime    time.Duration
	willExpired bool
}

type TimerDB struct {
	tMap map[string]*dataItem
}

func NewTimerDB() *TimerDB {
	return &TimerDB{make(map[string]*dataItem)}
}

func (r *TimerDB) AddItem(key string, data interface{}, lifeTime time.Duration) {
	r.tMap[key] = &dataItem{
		data:        data,
		editTime:    time.Now(),
		lifeTime:    lifeTime,
		willExpired: !(lifeTime == -1),
	}
}

func (r *TimerDB) DeleteItem(key string) {
	_, ok := r.tMap[key]
	if !ok {
		return
	}

	delete(r.tMap, key)
}

func (r *TimerDB) GetItem(key string) interface{} {
	item, ok := r.tMap[key]
	if !ok {
		return nil
	}

	if item.willExpired && time.Since(item.editTime) > item.lifeTime {
		delete(r.tMap, key)
		return nil
	}

	item.editTime = time.Now()
	return item.data
}
