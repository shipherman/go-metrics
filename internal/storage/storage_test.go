package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// implement all the tests for storage
func TestNew(t *testing.T) {
	///
	h := New()
	assert.IsType(t, MemStorage{}, h)
}

func TestGetCounter(t *testing.T) {
	h := New()
	key := "key"
	val := Counter(1)

	h.CounterData[key] = val

	v, err := h.Get(key)
	assert.Equal(t, val, v)
	assert.NoError(t, err, "Test valid key")

	_, err = h.Get("nosuchkey")
	assert.Error(t, err, "Test invalid key")
}

func TestGetAllCounters(t *testing.T) {
	h := New()
	key := "key"
	val := Counter(1)
	h.CounterData[key] = val

	expect := map[string]Counter{
		"key": Counter(1),
	}

	v := h.GetAllCounters()

	assert.Equal(t, expect, v, "Test GetAll")
}

func TestUpdateGauge(t *testing.T) {
	h := New()
	key := "key"
	val := Gauge(1.2)
	h.UpdateGauge(key, val)

	v, _ := h.Get(key)
	assert.Equal(t, val, v)
}

func TestUpdateCounter(t *testing.T) {
	h := New()
	key := "key"
	val := Counter(2)

	h.UpdateCounter(key, val)
	v, _ := h.Get(key)
	assert.Equal(t, val, v)

	h.UpdateCounter(key, val)
	v, _ = h.Get(key)
	assert.Equal(t, val+val, v)
}

func TestLocalFile(t *testing.T) {
	s := MemStorage{
		CounterData: map[string]Counter{
			"c01": 1,
		},
	}
	lf := Localfile{Path: "./file"}
	err := lf.Write(s)
	assert.NoError(t, err)

	err = lf.RestoreData(&s)
	assert.NoError(t, err)
}
