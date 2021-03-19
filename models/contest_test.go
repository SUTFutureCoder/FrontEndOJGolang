package models

import (
	"FrontEndOJGolang/pkg/utils"
	"reflect"
	"testing"
)

var lastInsertId int64

func TestContest_Insert(t *testing.T) {
	SetupTestUnit()
	type fields struct {
		Model            Model
		ContestName      string
		ContestDesc      string
		ContestStartTime int64
		ContestEndTime   int64
		SignupStartTime  int64
		SignupEndTime    int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{"test", fields{
			Model:            Model{
				CreatorId:  0,
				Creator:    "UnitTest",
				CreateTime: utils.GetMillTime(),
			},
			ContestName:      "UnitTest",
			ContestDesc:      "UnitTest",
			ContestStartTime: utils.GetMillTime(),
			ContestEndTime:   utils.GetMillTime() + 86400 * 1000 * 7,
			SignupStartTime:  utils.GetMillTime(),
			SignupEndTime:    utils.GetMillTime() + 86400 * 1000 * 7,
		}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contest{
				Model:            tt.fields.Model,
				ContestName:      tt.fields.ContestName,
				ContestDesc:      tt.fields.ContestDesc,
				ContestStartTime: tt.fields.ContestStartTime,
				ContestEndTime:   tt.fields.ContestEndTime,
				SignupStartTime:  tt.fields.SignupStartTime,
				SignupEndTime:    tt.fields.SignupEndTime,
			}
			got, err := c.Insert()
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= 0 {
				t.Errorf("Insert() got = %v, want >= 0", got)
			}
			lastInsertId = got
		})
	}
}

func TestContest_GetByIds(t *testing.T) {
	type fields struct {
		Model            Model
		ContestName      string
		ContestDesc      string
		ContestStartTime int64
		ContestEndTime   int64
		SignupStartTime  int64
		SignupEndTime    int64
	}
	type args struct {
		contestIds []interface{}
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Contest
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contest{
				Model:            tt.fields.Model,
				ContestName:      tt.fields.ContestName,
				ContestDesc:      tt.fields.ContestDesc,
				ContestStartTime: tt.fields.ContestStartTime,
				ContestEndTime:   tt.fields.ContestEndTime,
				SignupStartTime:  tt.fields.SignupStartTime,
				SignupEndTime:    tt.fields.SignupEndTime,
			}
			if got := c.GetByIds(tt.args.contestIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContest_GetList(t *testing.T) {
	type fields struct {
		Model            Model
		ContestName      string
		ContestDesc      string
		ContestStartTime int64
		ContestEndTime   int64
		SignupStartTime  int64
		SignupEndTime    int64
	}
	type args struct {
		page   Pager
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Contest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contest{
				Model:            tt.fields.Model,
				ContestName:      tt.fields.ContestName,
				ContestDesc:      tt.fields.ContestDesc,
				ContestStartTime: tt.fields.ContestStartTime,
				ContestEndTime:   tt.fields.ContestEndTime,
				SignupStartTime:  tt.fields.SignupStartTime,
				SignupEndTime:    tt.fields.SignupEndTime,
			}
			got, err := c.GetList(tt.args.page, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
