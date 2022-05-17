package dao

import (
	"context"
	"reflect"
	"testing"

	"github.com/thoohv5/template/internal/data/dao/model"
	"github.com/thoohv5/template/pkg/dbx/standard"
)

func TestMain(m *testing.M) {

}

func TestTest_Add(t *testing.T) {
	type fields struct {
		dao dao
	}
	type args struct {
		ctx  context.Context
		data *model.Test
		opts []Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &test{
				dao: tt.fields.dao,
			}
			if err := m.Add(tt.args.ctx, tt.args.data, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTest_Detail(t *testing.T) {
	type fields struct {
		dao dao
	}
	type args struct {
		ctx       context.Context
		condition *TestCondition
		opts      []Option
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantInfo *model.Test
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &test{
				dao: tt.fields.dao,
			}
			gotInfo, err := m.Detail(tt.args.ctx, tt.args.condition, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Detail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Detail() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func TestTest_List(t *testing.T) {
	type fields struct {
		dao dao
	}
	type args struct {
		ctx       context.Context
		condition *TestCondition
		opts      []Option
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantList []*model.Test
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &test{
				dao: tt.fields.dao,
			}
			gotList, err := m.List(tt.args.ctx, tt.args.condition, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("List() gotList = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}

func TestTest_Update(t *testing.T) {
	type fields struct {
		dao dao
	}
	type args struct {
		ctx        context.Context
		condition  *TestCondition
		updateData map[string]interface{}
		opts       []Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &test{
				dao: tt.fields.dao,
			}
			if err := m.Update(tt.args.ctx, tt.args.condition, tt.args.updateData, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTest_buildFilterQuery(t *testing.T) {
	type fields struct {
		dao dao
	}
	type args struct {
		ctx       context.Context
		condition *TestCondition
		opts      []Option
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantBuild standard.IBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &test{
				dao: tt.fields.dao,
			}
			if gotBuild := m.buildFilterQuery(tt.args.ctx, tt.args.condition, tt.args.opts...); !reflect.DeepEqual(gotBuild, tt.wantBuild) {
				t.Errorf("buildFilterQuery() = %v, want %v", gotBuild, tt.wantBuild)
			}
		})
	}
}

func TestWithBaseRequest(t *testing.T) {
	type args struct {
		br *BaseRequest
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithBaseRequest(tt.args.br); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithBaseRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithBuild(t *testing.T) {
	type args struct {
		build standard.IBuilder
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithBuild(tt.args.build); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithBuild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLimit(t *testing.T) {
	type args struct {
		limit int32
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLimit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dao_Close(t *testing.T) {
	type fields struct {
		build standard.IBuilder
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dao{
				build: tt.fields.build,
			}
			d.Close()
		})
	}
}

func Test_dao_baseConn(t *testing.T) {
	type fields struct {
		build standard.IBuilder
	}
	type args struct {
		build        standard.IBuilder
		condition    *BaseRequest
		appointLimit int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   standard.IBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dao{
				build: tt.fields.build,
			}
			if got := d.baseConn(tt.args.build, tt.args.condition, tt.args.appointLimit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dao_buildFilterQuery(t *testing.T) {
	type fields struct {
		build standard.IBuilder
	}
	type args struct {
		ctx  context.Context
		opts []Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   standard.IBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dao{
				build: tt.fields.build,
			}
			if got := d.buildFilterQuery(tt.args.ctx, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildFilterQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_optionFunc_apply(t *testing.T) {
	type args struct {
		o *options
	}
	tests := []struct {
		name string
		f    optionFunc
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.apply(tt.args.o)
		})
	}
}
