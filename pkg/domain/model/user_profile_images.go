// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserProfileImage is an object representing the database table.
type UserProfileImage struct { // プロフィール写真ID
	ID int `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ユーザーID
	UserID int `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	// 写真URL
	ImagePath string `boil:"image_path" json:"image_path" toml:"image_path" yaml:"image_path"`
	// プロフィール写真作成日時
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// プロフィール写真更新日時
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	// プロフィール写真論理削除日時
	DeteledAt null.Time `boil:"deteled_at" json:"deteled_at,omitempty" toml:"deteled_at" yaml:"deteled_at,omitempty"`

	R *userProfileImageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userProfileImageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserProfileImageColumns = struct {
	ID        string
	UserID    string
	ImagePath string
	CreatedAt string
	UpdatedAt string
	DeteledAt string
}{
	ID:        "id",
	UserID:    "user_id",
	ImagePath: "image_path",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeteledAt: "deteled_at",
}

var UserProfileImageTableColumns = struct {
	ID        string
	UserID    string
	ImagePath string
	CreatedAt string
	UpdatedAt string
	DeteledAt string
}{
	ID:        "user_profile_images.id",
	UserID:    "user_profile_images.user_id",
	ImagePath: "user_profile_images.image_path",
	CreatedAt: "user_profile_images.created_at",
	UpdatedAt: "user_profile_images.updated_at",
	DeteledAt: "user_profile_images.deteled_at",
}

// Generated where

var UserProfileImageWhere = struct {
	ID        whereHelperint
	UserID    whereHelperint
	ImagePath whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeteledAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "`user_profile_images`.`id`"},
	UserID:    whereHelperint{field: "`user_profile_images`.`user_id`"},
	ImagePath: whereHelperstring{field: "`user_profile_images`.`image_path`"},
	CreatedAt: whereHelpertime_Time{field: "`user_profile_images`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`user_profile_images`.`updated_at`"},
	DeteledAt: whereHelpernull_Time{field: "`user_profile_images`.`deteled_at`"},
}

// UserProfileImageRels is where relationship names are stored.
var UserProfileImageRels = struct {
	User string
}{
	User: "User",
}

// userProfileImageR is where relationships are stored.
type userProfileImageR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userProfileImageR) NewStruct() *userProfileImageR {
	return &userProfileImageR{}
}

func (r *userProfileImageR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// userProfileImageL is where Load methods for each relationship are stored.
type userProfileImageL struct{}

var (
	userProfileImageAllColumns            = []string{"id", "user_id", "image_path", "created_at", "updated_at", "deteled_at"}
	userProfileImageColumnsWithoutDefault = []string{"user_id", "image_path", "deteled_at"}
	userProfileImageColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	userProfileImagePrimaryKeyColumns     = []string{"id"}
	userProfileImageGeneratedColumns      = []string{}
)

type (
	// UserProfileImageSlice is an alias for a slice of pointers to UserProfileImage.
	// This should almost always be used instead of []UserProfileImage.
	UserProfileImageSlice []*UserProfileImage
	// UserProfileImageHook is the signature for custom UserProfileImage hook methods
	UserProfileImageHook func(context.Context, boil.ContextExecutor, *UserProfileImage) error

	userProfileImageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userProfileImageType                 = reflect.TypeOf(&UserProfileImage{})
	userProfileImageMapping              = queries.MakeStructMapping(userProfileImageType)
	userProfileImagePrimaryKeyMapping, _ = queries.BindMapping(userProfileImageType, userProfileImageMapping, userProfileImagePrimaryKeyColumns)
	userProfileImageInsertCacheMut       sync.RWMutex
	userProfileImageInsertCache          = make(map[string]insertCache)
	userProfileImageUpdateCacheMut       sync.RWMutex
	userProfileImageUpdateCache          = make(map[string]updateCache)
	userProfileImageUpsertCacheMut       sync.RWMutex
	userProfileImageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userProfileImageAfterSelectHooks []UserProfileImageHook

var userProfileImageBeforeInsertHooks []UserProfileImageHook
var userProfileImageAfterInsertHooks []UserProfileImageHook

var userProfileImageBeforeUpdateHooks []UserProfileImageHook
var userProfileImageAfterUpdateHooks []UserProfileImageHook

var userProfileImageBeforeDeleteHooks []UserProfileImageHook
var userProfileImageAfterDeleteHooks []UserProfileImageHook

var userProfileImageBeforeUpsertHooks []UserProfileImageHook
var userProfileImageAfterUpsertHooks []UserProfileImageHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserProfileImage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileImageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserProfileImage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileImageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserProfileImage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileImageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserProfileImage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileImageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserProfileImage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileImageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserProfileImage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileImageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserProfileImage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileImageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserProfileImage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileImageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserProfileImage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileImageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserProfileImageHook registers your hook function for all future operations.
func AddUserProfileImageHook(hookPoint boil.HookPoint, userProfileImageHook UserProfileImageHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userProfileImageAfterSelectHooks = append(userProfileImageAfterSelectHooks, userProfileImageHook)
	case boil.BeforeInsertHook:
		userProfileImageBeforeInsertHooks = append(userProfileImageBeforeInsertHooks, userProfileImageHook)
	case boil.AfterInsertHook:
		userProfileImageAfterInsertHooks = append(userProfileImageAfterInsertHooks, userProfileImageHook)
	case boil.BeforeUpdateHook:
		userProfileImageBeforeUpdateHooks = append(userProfileImageBeforeUpdateHooks, userProfileImageHook)
	case boil.AfterUpdateHook:
		userProfileImageAfterUpdateHooks = append(userProfileImageAfterUpdateHooks, userProfileImageHook)
	case boil.BeforeDeleteHook:
		userProfileImageBeforeDeleteHooks = append(userProfileImageBeforeDeleteHooks, userProfileImageHook)
	case boil.AfterDeleteHook:
		userProfileImageAfterDeleteHooks = append(userProfileImageAfterDeleteHooks, userProfileImageHook)
	case boil.BeforeUpsertHook:
		userProfileImageBeforeUpsertHooks = append(userProfileImageBeforeUpsertHooks, userProfileImageHook)
	case boil.AfterUpsertHook:
		userProfileImageAfterUpsertHooks = append(userProfileImageAfterUpsertHooks, userProfileImageHook)
	}
}

// OneG returns a single userProfileImage record from the query using the global executor.
func (q userProfileImageQuery) OneG(ctx context.Context) (*UserProfileImage, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single userProfileImage record from the query.
func (q userProfileImageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserProfileImage, error) {
	o := &UserProfileImage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for user_profile_images")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all UserProfileImage records from the query using the global executor.
func (q userProfileImageQuery) AllG(ctx context.Context) (UserProfileImageSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all UserProfileImage records from the query.
func (q userProfileImageQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserProfileImageSlice, error) {
	var o []*UserProfileImage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to UserProfileImage slice")
	}

	if len(userProfileImageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all UserProfileImage records in the query using the global executor
func (q userProfileImageQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all UserProfileImage records in the query.
func (q userProfileImageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count user_profile_images rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q userProfileImageQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q userProfileImageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if user_profile_images exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserProfileImage) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userProfileImageL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserProfileImage interface{}, mods queries.Applicator) error {
	var slice []*UserProfileImage
	var object *UserProfileImage

	if singular {
		var ok bool
		object, ok = maybeUserProfileImage.(*UserProfileImage)
		if !ok {
			object = new(UserProfileImage)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserProfileImage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserProfileImage))
			}
		}
	} else {
		s, ok := maybeUserProfileImage.(*[]*UserProfileImage)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserProfileImage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserProfileImage))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userProfileImageR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userProfileImageR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userProfileImageAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserProfileImages = append(foreign.R.UserProfileImages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserProfileImages = append(foreign.R.UserProfileImages, local)
				break
			}
		}
	}

	return nil
}

// SetUserG of the userProfileImage to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserProfileImages.
// Uses the global database handle.
func (o *UserProfileImage) SetUserG(ctx context.Context, insert bool, related *User) error {
	return o.SetUser(ctx, boil.GetContextDB(), insert, related)
}

// SetUser of the userProfileImage to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserProfileImages.
func (o *UserProfileImage) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `user_profile_images` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, userProfileImagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userProfileImageR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserProfileImages: UserProfileImageSlice{o},
		}
	} else {
		related.R.UserProfileImages = append(related.R.UserProfileImages, o)
	}

	return nil
}

// UserProfileImages retrieves all the records using an executor.
func UserProfileImages(mods ...qm.QueryMod) userProfileImageQuery {
	mods = append(mods, qm.From("`user_profile_images`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`user_profile_images`.*"})
	}

	return userProfileImageQuery{q}
}

// FindUserProfileImageG retrieves a single record by ID.
func FindUserProfileImageG(ctx context.Context, iD int, selectCols ...string) (*UserProfileImage, error) {
	return FindUserProfileImage(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindUserProfileImage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserProfileImage(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*UserProfileImage, error) {
	userProfileImageObj := &UserProfileImage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `user_profile_images` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userProfileImageObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from user_profile_images")
	}

	if err = userProfileImageObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userProfileImageObj, err
	}

	return userProfileImageObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *UserProfileImage) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserProfileImage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no user_profile_images provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userProfileImageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userProfileImageInsertCacheMut.RLock()
	cache, cached := userProfileImageInsertCache[key]
	userProfileImageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userProfileImageAllColumns,
			userProfileImageColumnsWithDefault,
			userProfileImageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userProfileImageType, userProfileImageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userProfileImageType, userProfileImageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `user_profile_images` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `user_profile_images` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `user_profile_images` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userProfileImagePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into user_profile_images")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userProfileImageMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for user_profile_images")
	}

CacheNoHooks:
	if !cached {
		userProfileImageInsertCacheMut.Lock()
		userProfileImageInsertCache[key] = cache
		userProfileImageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single UserProfileImage record using the global executor.
// See Update for more documentation.
func (o *UserProfileImage) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the UserProfileImage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserProfileImage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userProfileImageUpdateCacheMut.RLock()
	cache, cached := userProfileImageUpdateCache[key]
	userProfileImageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userProfileImageAllColumns,
			userProfileImagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update user_profile_images, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `user_profile_images` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userProfileImagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userProfileImageType, userProfileImageMapping, append(wl, userProfileImagePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update user_profile_images row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for user_profile_images")
	}

	if !cached {
		userProfileImageUpdateCacheMut.Lock()
		userProfileImageUpdateCache[key] = cache
		userProfileImageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q userProfileImageQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q userProfileImageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for user_profile_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for user_profile_images")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UserProfileImageSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserProfileImageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userProfileImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `user_profile_images` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userProfileImagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in userProfileImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all userProfileImage")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *UserProfileImage) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLUserProfileImageUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserProfileImage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no user_profile_images provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userProfileImageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserProfileImageUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userProfileImageUpsertCacheMut.RLock()
	cache, cached := userProfileImageUpsertCache[key]
	userProfileImageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userProfileImageAllColumns,
			userProfileImageColumnsWithDefault,
			userProfileImageColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userProfileImageAllColumns,
			userProfileImagePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("model: unable to upsert user_profile_images, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`user_profile_images`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `user_profile_images` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userProfileImageType, userProfileImageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userProfileImageType, userProfileImageMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for user_profile_images")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userProfileImageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userProfileImageType, userProfileImageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for user_profile_images")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for user_profile_images")
	}

CacheNoHooks:
	if !cached {
		userProfileImageUpsertCacheMut.Lock()
		userProfileImageUpsertCache[key] = cache
		userProfileImageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single UserProfileImage record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *UserProfileImage) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single UserProfileImage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserProfileImage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no UserProfileImage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userProfileImagePrimaryKeyMapping)
	sql := "DELETE FROM `user_profile_images` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from user_profile_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for user_profile_images")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q userProfileImageQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q userProfileImageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no userProfileImageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from user_profile_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for user_profile_images")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o UserProfileImageSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserProfileImageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userProfileImageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userProfileImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `user_profile_images` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userProfileImagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from userProfileImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for user_profile_images")
	}

	if len(userProfileImageAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *UserProfileImage) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("model: no UserProfileImage provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserProfileImage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserProfileImage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserProfileImageSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("model: empty UserProfileImageSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserProfileImageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserProfileImageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userProfileImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `user_profile_images`.* FROM `user_profile_images` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userProfileImagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in UserProfileImageSlice")
	}

	*o = slice

	return nil
}

// UserProfileImageExistsG checks if the UserProfileImage row exists.
func UserProfileImageExistsG(ctx context.Context, iD int) (bool, error) {
	return UserProfileImageExists(ctx, boil.GetContextDB(), iD)
}

// UserProfileImageExists checks if the UserProfileImage row exists.
func UserProfileImageExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `user_profile_images` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if user_profile_images exists")
	}

	return exists, nil
}
