package model

import (
	"testing"
)

func TestBook_Create(t *testing.T) {
	TestConnDB(t)

	type fields struct {
		ID           int64
		Name         string
		Author       string
		Description  string
		ExtensionURL string
		Keywords     string
		Category     string
		LastChapter  int64
		ChapterCount int64
		Tags         string
		Status       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "正常写入",
			fields: fields{
				ID:           5,
				Name:         "测试书名",
				Author:       "作者名",
				Description:  "描述",
				ExtensionURL: "封面URL",
				Keywords:     "关键字描述",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				ID:           tt.fields.ID,
				Name:         tt.fields.Name,
				Author:       tt.fields.Author,
				Description:  tt.fields.Description,
				ExtensionURL: tt.fields.ExtensionURL,
				Keywords:     tt.fields.Keywords,
				Category:     tt.fields.Category,
				LastChapter:  tt.fields.LastChapter,
				ChapterCount: tt.fields.ChapterCount,
				Tags:         tt.fields.Tags,
				Status:       tt.fields.Status,
			}
			if err := b.Create(); (err != nil) != tt.wantErr {
				t.Errorf("Book.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBook_Get(t *testing.T) {
	TestConnDB(t)

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    *Book
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "正常请求",
			args: args{
				id: 4,
			},
			want1:   false,
			wantErr: false,
		},
		{
			name: "为空的请求",
			args: args{
				id: 2,
			},
			want1:   true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{}
			_, got1, err := b.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Book.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want1 {
				t.Errorf("Book.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
