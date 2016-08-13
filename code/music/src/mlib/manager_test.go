package mlib

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty")
	}

	m0 := &MusicEntry{
		"1", "My Heart Will Go On", "Celion Dion", "http://music.163.com/#/m/song?id=2308499", "mp3",
	}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed")
	} else {
		t.Log("MusicManager.Add() succ ,", mm.Len())
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed.Found item mismatch.")
	} else {
		t.Log("MusicManager.Find() succ ,", m.Name)
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	} else {
		t.Log("MusicManager.Get() succ ,", m.Name)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed", err)
	} else {
		t.Log("MusicManager.Remove() succ ,", m.Name, mm.Len())
	}
}
