package api

import (
	"context"
	"gorilla/mux"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func Test_decodeSearchContactRequest(t *testing.T) {
	type args struct {
		in0 context.Context
		r   *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"1", args{context.Background(), &http.Request{
			URL: &url.URL{
				RawQuery: "name=isco&page=0",
			},
		}}, searchContactRequest{Page: 0, Name: "isco"}, false},
		{"2", args{context.Background(), &http.Request{
			URL: &url.URL{
				RawQuery: "name=isco",
			},
		}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeSearchContactRequest(tt.args.in0, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeSearchContactRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeSearchContactRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
