// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// McontentText is an object representing the database table.
type McontentText struct {
	ID         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt  null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt  null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Lang       string      `boil:"lang" json:"lang" toml:"lang" yaml:"lang"`
	Title      string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Subtitle   null.String `boil:"subtitle" json:"subtitle,omitempty" toml:"subtitle" yaml:"subtitle,omitempty"`
	Detail     null.String `boil:"detail" json:"detail,omitempty" toml:"detail" yaml:"detail,omitempty"`
	Xtext1     null.String `boil:"xtext1" json:"xtext1,omitempty" toml:"xtext1" yaml:"xtext1,omitempty"`
	CreatorID  null.Int    `boil:"creator_id" json:"creator_id,omitempty" toml:"creator_id" yaml:"creator_id,omitempty"`
	McontentID int         `boil:"mcontent_id" json:"mcontent_id" toml:"mcontent_id" yaml:"mcontent_id"`

	R *mcontentTextR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mcontentTextL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var McontentTextColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	Lang       string
	Title      string
	Subtitle   string
	Detail     string
	Xtext1     string
	CreatorID  string
	McontentID string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	Lang:       "lang",
	Title:      "title",
	Subtitle:   "subtitle",
	Detail:     "detail",
	Xtext1:     "xtext1",
	CreatorID:  "creator_id",
	McontentID: "mcontent_id",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var McontentTextWhere = struct {
	ID         whereHelperint
	CreatedAt  whereHelpernull_Time
	UpdatedAt  whereHelpernull_Time
	DeletedAt  whereHelpernull_Time
	Lang       whereHelperstring
	Title      whereHelperstring
	Subtitle   whereHelpernull_String
	Detail     whereHelpernull_String
	Xtext1     whereHelpernull_String
	CreatorID  whereHelpernull_Int
	McontentID whereHelperint
}{
	ID:         whereHelperint{field: `id`},
	CreatedAt:  whereHelpernull_Time{field: `created_at`},
	UpdatedAt:  whereHelpernull_Time{field: `updated_at`},
	DeletedAt:  whereHelpernull_Time{field: `deleted_at`},
	Lang:       whereHelperstring{field: `lang`},
	Title:      whereHelperstring{field: `title`},
	Subtitle:   whereHelpernull_String{field: `subtitle`},
	Detail:     whereHelpernull_String{field: `detail`},
	Xtext1:     whereHelpernull_String{field: `xtext1`},
	CreatorID:  whereHelpernull_Int{field: `creator_id`},
	McontentID: whereHelperint{field: `mcontent_id`},
}

// McontentTextRels is where relationship names are stored.
var McontentTextRels = struct {
	Mcontent string
}{
	Mcontent: "Mcontent",
}

// mcontentTextR is where relationships are stored.
type mcontentTextR struct {
	Mcontent *Mcontent
}

// NewStruct creates a new relationship struct
func (*mcontentTextR) NewStruct() *mcontentTextR {
	return &mcontentTextR{}
}

// mcontentTextL is where Load methods for each relationship are stored.
type mcontentTextL struct{}

var (
	mcontentTextColumns               = []string{"id", "created_at", "updated_at", "deleted_at", "lang", "title", "subtitle", "detail", "xtext1", "creator_id", "mcontent_id"}
	mcontentTextColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "title", "subtitle", "detail", "xtext1", "creator_id", "mcontent_id"}
	mcontentTextColumnsWithDefault    = []string{"id", "lang"}
	mcontentTextPrimaryKeyColumns     = []string{"id"}
)

type (
	// McontentTextSlice is an alias for a slice of pointers to McontentText.
	// This should generally be used opposed to []McontentText.
	McontentTextSlice []*McontentText
	// McontentTextHook is the signature for custom McontentText hook methods
	McontentTextHook func(context.Context, boil.ContextExecutor, *McontentText) error

	mcontentTextQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mcontentTextType                 = reflect.TypeOf(&McontentText{})
	mcontentTextMapping              = queries.MakeStructMapping(mcontentTextType)
	mcontentTextPrimaryKeyMapping, _ = queries.BindMapping(mcontentTextType, mcontentTextMapping, mcontentTextPrimaryKeyColumns)
	mcontentTextInsertCacheMut       sync.RWMutex
	mcontentTextInsertCache          = make(map[string]insertCache)
	mcontentTextUpdateCacheMut       sync.RWMutex
	mcontentTextUpdateCache          = make(map[string]updateCache)
	mcontentTextUpsertCacheMut       sync.RWMutex
	mcontentTextUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mcontentTextBeforeInsertHooks []McontentTextHook
var mcontentTextBeforeUpdateHooks []McontentTextHook
var mcontentTextBeforeDeleteHooks []McontentTextHook
var mcontentTextBeforeUpsertHooks []McontentTextHook

var mcontentTextAfterInsertHooks []McontentTextHook
var mcontentTextAfterSelectHooks []McontentTextHook
var mcontentTextAfterUpdateHooks []McontentTextHook
var mcontentTextAfterDeleteHooks []McontentTextHook
var mcontentTextAfterUpsertHooks []McontentTextHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *McontentText) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mcontentTextBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *McontentText) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mcontentTextBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *McontentText) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mcontentTextBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *McontentText) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mcontentTextBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *McontentText) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mcontentTextAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *McontentText) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mcontentTextAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *McontentText) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mcontentTextAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *McontentText) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mcontentTextAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *McontentText) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mcontentTextAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMcontentTextHook registers your hook function for all future operations.
func AddMcontentTextHook(hookPoint boil.HookPoint, mcontentTextHook McontentTextHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		mcontentTextBeforeInsertHooks = append(mcontentTextBeforeInsertHooks, mcontentTextHook)
	case boil.BeforeUpdateHook:
		mcontentTextBeforeUpdateHooks = append(mcontentTextBeforeUpdateHooks, mcontentTextHook)
	case boil.BeforeDeleteHook:
		mcontentTextBeforeDeleteHooks = append(mcontentTextBeforeDeleteHooks, mcontentTextHook)
	case boil.BeforeUpsertHook:
		mcontentTextBeforeUpsertHooks = append(mcontentTextBeforeUpsertHooks, mcontentTextHook)
	case boil.AfterInsertHook:
		mcontentTextAfterInsertHooks = append(mcontentTextAfterInsertHooks, mcontentTextHook)
	case boil.AfterSelectHook:
		mcontentTextAfterSelectHooks = append(mcontentTextAfterSelectHooks, mcontentTextHook)
	case boil.AfterUpdateHook:
		mcontentTextAfterUpdateHooks = append(mcontentTextAfterUpdateHooks, mcontentTextHook)
	case boil.AfterDeleteHook:
		mcontentTextAfterDeleteHooks = append(mcontentTextAfterDeleteHooks, mcontentTextHook)
	case boil.AfterUpsertHook:
		mcontentTextAfterUpsertHooks = append(mcontentTextAfterUpsertHooks, mcontentTextHook)
	}
}

// One returns a single mcontentText record from the query.
func (q mcontentTextQuery) One(ctx context.Context, exec boil.ContextExecutor) (*McontentText, error) {
	o := &McontentText{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mcontent_texts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all McontentText records from the query.
func (q mcontentTextQuery) All(ctx context.Context, exec boil.ContextExecutor) (McontentTextSlice, error) {
	var o []*McontentText

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to McontentText slice")
	}

	if len(mcontentTextAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all McontentText records in the query.
func (q mcontentTextQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mcontent_texts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mcontentTextQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mcontent_texts exists")
	}

	return count > 0, nil
}

// Mcontent pointed to by the foreign key.
func (o *McontentText) Mcontent(mods ...qm.QueryMod) mcontentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.McontentID),
	}

	queryMods = append(queryMods, mods...)

	query := Mcontents(queryMods...)
	queries.SetFrom(query.Query, "\"mcontents\"")

	return query
}

// LoadMcontent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (mcontentTextL) LoadMcontent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMcontentText interface{}, mods queries.Applicator) error {
	var slice []*McontentText
	var object *McontentText

	if singular {
		object = maybeMcontentText.(*McontentText)
	} else {
		slice = *maybeMcontentText.(*[]*McontentText)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &mcontentTextR{}
		}
		args = append(args, object.McontentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &mcontentTextR{}
			}

			for _, a := range args {
				if a == obj.McontentID {
					continue Outer
				}
			}

			args = append(args, obj.McontentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`mcontents`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Mcontent")
	}

	var resultSlice []*Mcontent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Mcontent")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for mcontents")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for mcontents")
	}

	if len(mcontentTextAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Mcontent = foreign
		if foreign.R == nil {
			foreign.R = &mcontentR{}
		}
		foreign.R.McontentTexts = append(foreign.R.McontentTexts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.McontentID == foreign.ID {
				local.R.Mcontent = foreign
				if foreign.R == nil {
					foreign.R = &mcontentR{}
				}
				foreign.R.McontentTexts = append(foreign.R.McontentTexts, local)
				break
			}
		}
	}

	return nil
}

// SetMcontent of the mcontentText to the related item.
// Sets o.R.Mcontent to related.
// Adds o to related.R.McontentTexts.
func (o *McontentText) SetMcontent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Mcontent) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"mcontent_texts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"mcontent_id"}),
		strmangle.WhereClause("\"", "\"", 2, mcontentTextPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.McontentID = related.ID
	if o.R == nil {
		o.R = &mcontentTextR{
			Mcontent: related,
		}
	} else {
		o.R.Mcontent = related
	}

	if related.R == nil {
		related.R = &mcontentR{
			McontentTexts: McontentTextSlice{o},
		}
	} else {
		related.R.McontentTexts = append(related.R.McontentTexts, o)
	}

	return nil
}

// McontentTexts retrieves all the records using an executor.
func McontentTexts(mods ...qm.QueryMod) mcontentTextQuery {
	mods = append(mods, qm.From("\"mcontent_texts\""))
	return mcontentTextQuery{NewQuery(mods...)}
}

// FindMcontentText retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMcontentText(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*McontentText, error) {
	mcontentTextObj := &McontentText{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mcontent_texts\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mcontentTextObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mcontent_texts")
	}

	return mcontentTextObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *McontentText) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mcontent_texts provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mcontentTextColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mcontentTextInsertCacheMut.RLock()
	cache, cached := mcontentTextInsertCache[key]
	mcontentTextInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mcontentTextColumns,
			mcontentTextColumnsWithDefault,
			mcontentTextColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mcontentTextType, mcontentTextMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mcontentTextType, mcontentTextMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mcontent_texts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mcontent_texts\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into mcontent_texts")
	}

	if !cached {
		mcontentTextInsertCacheMut.Lock()
		mcontentTextInsertCache[key] = cache
		mcontentTextInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the McontentText.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *McontentText) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mcontentTextUpdateCacheMut.RLock()
	cache, cached := mcontentTextUpdateCache[key]
	mcontentTextUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mcontentTextColumns,
			mcontentTextPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mcontent_texts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mcontent_texts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mcontentTextPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mcontentTextType, mcontentTextMapping, append(wl, mcontentTextPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update mcontent_texts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mcontent_texts")
	}

	if !cached {
		mcontentTextUpdateCacheMut.Lock()
		mcontentTextUpdateCache[key] = cache
		mcontentTextUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mcontentTextQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mcontent_texts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mcontent_texts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o McontentTextSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mcontentTextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mcontent_texts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mcontentTextPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mcontentText slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mcontentText")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *McontentText) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mcontent_texts provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mcontentTextColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	mcontentTextUpsertCacheMut.RLock()
	cache, cached := mcontentTextUpsertCache[key]
	mcontentTextUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mcontentTextColumns,
			mcontentTextColumnsWithDefault,
			mcontentTextColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			mcontentTextColumns,
			mcontentTextPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert mcontent_texts, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(mcontentTextPrimaryKeyColumns))
			copy(conflict, mcontentTextPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mcontent_texts\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(mcontentTextType, mcontentTextMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mcontentTextType, mcontentTextMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert mcontent_texts")
	}

	if !cached {
		mcontentTextUpsertCacheMut.Lock()
		mcontentTextUpsertCache[key] = cache
		mcontentTextUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single McontentText record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *McontentText) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no McontentText provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mcontentTextPrimaryKeyMapping)
	sql := "DELETE FROM \"mcontent_texts\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mcontent_texts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mcontent_texts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mcontentTextQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mcontentTextQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mcontent_texts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mcontent_texts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o McontentTextSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no McontentText slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(mcontentTextBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mcontentTextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mcontent_texts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mcontentTextPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mcontentText slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mcontent_texts")
	}

	if len(mcontentTextAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *McontentText) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMcontentText(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *McontentTextSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := McontentTextSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mcontentTextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mcontent_texts\".* FROM \"mcontent_texts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mcontentTextPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in McontentTextSlice")
	}

	*o = slice

	return nil
}

// McontentTextExists checks if the McontentText row exists.
func McontentTextExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mcontent_texts\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mcontent_texts exists")
	}

	return exists, nil
}