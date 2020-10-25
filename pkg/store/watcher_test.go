package store

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewWatcher(t *testing.T) {
	tests := []struct {
		name string
		want *Watcher
	}{
		{
			name: "normal",
			want: &Watcher{
				chMap: make(map[string][]chan Notification),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWatcher(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatcher_Add(t *testing.T) {
	type fields struct {
		chMap map[string][]chan Notification
	}
	type args struct {
		key string
		fn  NotifyFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				chMap: make(map[string][]chan Notification),
			},
			args: args{
				key: "hello",
				fn: func(nch <-chan Notification) {
					fmt.Printf("%v\n", <-nch)
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Watcher{
				chMap: tt.fields.chMap,
			}
			if err := w.Add(tt.args.key, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWatcher_Notify(t *testing.T) {
	type fields struct {
		chMap map[string][]chan Notification
	}
	type args struct {
		action int
		key    string
		oldVal string
		newVal string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				chMap: make(map[string][]chan Notification),
			},
			args: args{
				action: SET,
				key:    "hello",
				oldVal: "",
				newVal: "world",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Watcher{
				chMap: tt.fields.chMap,
			}
			if err := w.Notify(tt.args.action, tt.args.key, tt.args.oldVal, tt.args.newVal); (err != nil) != tt.wantErr {
				t.Errorf("Notify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWatcher(t *testing.T) {
	t.Run("watcher", func(t *testing.T) {
		watcher.Add("hello", listenChange)
		watcher.Add("hello", listenChange)
		evolvest.Set("hello", "world")

		watcher.Add("hello", listenChange)
		evolvest.Set("hello", "newWorld")

		watcher.Add("hello", listenChange)
		evolvest.Del("hello")
		watcher.Add("hello", listenChange)
		evolvest.Del("hello")

	})
}

func listenChange(ch <-chan Notification) {
	n := <-ch
	fmt.Printf("notification: %v\n", n)
}