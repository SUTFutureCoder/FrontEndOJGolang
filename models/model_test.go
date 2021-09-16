package models

import (
	"FrontEndOJGolang/pkg/setting"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func SetupTestUnit() {
	setting.Setup("../conf/app.ini")
	setting.Check()
	Setup()
}

func TestDefaultPage(t *testing.T) {
	type args struct {
		page     *int
		pageSize *int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSetup(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestToPager(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		want Pager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPager(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	var list []interface{}
	var labIds []interface{}
	labIds = append(labIds, 1)
	labIds = append(labIds, 2)
	labIds = append(labIds, 3)
	labIds = append(labIds, 4)

	list = append(list, 6)
	list = append(list, labIds...)

	for i := range list {
		println(list[i])
	}

}