package storage

import (
    "github.com/stretchr/testify/assert"
    "testing"
)


// implement all the tests for storage
func TestNew(t *testing.T) {
///
    h := New()
    assert.IsType(t, MemStorage{}, h)
}

func TestGet(t *testing.T) {
   h := New()
   key := "key"
   val := Counter(1)

   h.Data[key] = val


   v, err := h.Get(key)
   assert.Equal(t, val, v)
   assert.NoError(t, err, "Test valid key")

   _, err = h.Get("nosuchkey")
   assert.Error(t, err, "Test invalid key")
}

func TestGetAll(t *testing.T) {
    h := New()
    key := "key"
    val := Counter(1)
    h.Data[key] = val

    expect := map[string]interface{}{
            "key": Counter(1),
    }

    v := h.GetAll()

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
    assert.Equal(t, val + val, v )
}

