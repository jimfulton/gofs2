package main

import (
	"io/ioutil"
	"log"
	"me/j1m/fs2"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	tdir, err := ioutil.TempDir("", "test")
	if (err != nil) { log.Fatal(err) }
	defer os.RemoveAll(tdir)

	err = os.Chdir(tdir)
	if (err != nil) { log.Fatal(err) }
	os.Exit(m.Run())
}


func TestCreate(t *testing.T) {
	// Create new:
	store, err := fs2.New("t")
	if err != nil { t.Error(err); return }
	if store.Path != "t" { t.Error(store.Path); return }

	err = store.Close()
	if err != nil { t.Error(err); return }

	// Reopen
	store, err = fs2.New("t")
	if err != nil { t.Error(err); return }
	if store.Path != "t" { t.Error(store.Path); return }


	err = store.Close()
	if err != nil { t.Error(err); return }
}

// func TestStore(t *testing.T) {
// 	store, _ := fs2.New("t")
// 	if err := store.Store("1", nil, "data1"); err != nil {
// 		t.Error(err)
// 		return
// 	}
// }
