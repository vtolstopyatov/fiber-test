// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type newsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *newsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("news").
func (v *newsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *newsTableType) Columns() []string {
	return []string{
		"id",
		"title",
		"content",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *newsTableType) NewStruct() reform.Struct {
	return new(News)
}

// NewRecord makes a new record for that table.
func (v *newsTableType) NewRecord() reform.Record {
	return new(News)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *newsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// NewsTable represents news view or table in SQL database.
var NewsTable = &newsTableType{
	s: parse.StructInfo{
		Type:    "News",
		SQLName: "news",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int64", Column: "id"},
			{Name: "Title", Type: "string", Column: "title"},
			{Name: "Content", Type: "string", Column: "content"},
		},
		PKFieldIndex: 0,
	},
	z: new(News).Values(),
}

// String returns a string representation of this struct or record.
func (s News) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Title: " + reform.Inspect(s.Title, true)
	res[2] = "Content: " + reform.Inspect(s.Content, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *News) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Title,
		s.Content,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *News) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Title,
		&s.Content,
	}
}

// View returns View object for that struct.
func (s *News) View() reform.View {
	return NewsTable
}

// Table returns Table object for that record.
func (s *News) Table() reform.Table {
	return NewsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *News) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *News) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *News) HasPK() bool {
	return s.ID != NewsTable.z[NewsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *News) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = NewsTable
	_ reform.Struct = (*News)(nil)
	_ reform.Table  = NewsTable
	_ reform.Record = (*News)(nil)
	_ fmt.Stringer  = (*News)(nil)
)

type newsCategoriesTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *newsCategoriesTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("news_categories").
func (v *newsCategoriesTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *newsCategoriesTableType) Columns() []string {
	return []string{
		"news_id",
		"category_id",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *newsCategoriesTableType) NewStruct() reform.Struct {
	return new(NewsCategories)
}

// NewRecord makes a new record for that table.
func (v *newsCategoriesTableType) NewRecord() reform.Record {
	return new(NewsCategories)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *newsCategoriesTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// NewsCategoriesTable represents news_categories view or table in SQL database.
var NewsCategoriesTable = &newsCategoriesTableType{
	s: parse.StructInfo{
		Type:    "NewsCategories",
		SQLName: "news_categories",
		Fields: []parse.FieldInfo{
			{Name: "NewsId", Type: "int64", Column: "news_id"},
			{Name: "CategoryId", Type: "int64", Column: "category_id"},
		},
		PKFieldIndex: 0,
	},
	z: new(NewsCategories).Values(),
}

// String returns a string representation of this struct or record.
func (s NewsCategories) String() string {
	res := make([]string, 2)
	res[0] = "NewsId: " + reform.Inspect(s.NewsId, true)
	res[1] = "CategoryId: " + reform.Inspect(s.CategoryId, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *NewsCategories) Values() []interface{} {
	return []interface{}{
		s.NewsId,
		s.CategoryId,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *NewsCategories) Pointers() []interface{} {
	return []interface{}{
		&s.NewsId,
		&s.CategoryId,
	}
}

// View returns View object for that struct.
func (s *NewsCategories) View() reform.View {
	return NewsCategoriesTable
}

// Table returns Table object for that record.
func (s *NewsCategories) Table() reform.Table {
	return NewsCategoriesTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *NewsCategories) PKValue() interface{} {
	return s.NewsId
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *NewsCategories) PKPointer() interface{} {
	return &s.NewsId
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *NewsCategories) HasPK() bool {
	return s.NewsId != NewsCategoriesTable.z[NewsCategoriesTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.NewsId = pk.
func (s *NewsCategories) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = NewsCategoriesTable
	_ reform.Struct = (*NewsCategories)(nil)
	_ reform.Table  = NewsCategoriesTable
	_ reform.Record = (*NewsCategories)(nil)
	_ fmt.Stringer  = (*NewsCategories)(nil)
)

func init() {
	parse.AssertUpToDate(&NewsTable.s, new(News))
	parse.AssertUpToDate(&NewsCategoriesTable.s, new(NewsCategories))
}
