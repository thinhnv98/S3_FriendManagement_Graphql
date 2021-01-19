// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBlocks(t *testing.T) {
	t.Parallel()

	query := Blocks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBlocksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Blocks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BlockExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Block exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BlockExists to return true, but got false.")
	}
}

func testBlocksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	blockFound, err := FindBlock(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if blockFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBlocksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Blocks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBlocksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Blocks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBlocksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = blockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Blocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBlocksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = blockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func blockBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Block) error {
	*o = Block{}
	return nil
}

func blockAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Block) error {
	*o = Block{}
	return nil
}

func blockAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Block) error {
	*o = Block{}
	return nil
}

func blockBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Block) error {
	*o = Block{}
	return nil
}

func blockAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Block) error {
	*o = Block{}
	return nil
}

func blockBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Block) error {
	*o = Block{}
	return nil
}

func blockAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Block) error {
	*o = Block{}
	return nil
}

func blockBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Block) error {
	*o = Block{}
	return nil
}

func blockAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Block) error {
	*o = Block{}
	return nil
}

func testBlocksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Block{}
	o := &Block{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, blockDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Block object: %s", err)
	}

	AddBlockHook(boil.BeforeInsertHook, blockBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	blockBeforeInsertHooks = []BlockHook{}

	AddBlockHook(boil.AfterInsertHook, blockAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	blockAfterInsertHooks = []BlockHook{}

	AddBlockHook(boil.AfterSelectHook, blockAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	blockAfterSelectHooks = []BlockHook{}

	AddBlockHook(boil.BeforeUpdateHook, blockBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	blockBeforeUpdateHooks = []BlockHook{}

	AddBlockHook(boil.AfterUpdateHook, blockAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	blockAfterUpdateHooks = []BlockHook{}

	AddBlockHook(boil.BeforeDeleteHook, blockBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	blockBeforeDeleteHooks = []BlockHook{}

	AddBlockHook(boil.AfterDeleteHook, blockAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	blockAfterDeleteHooks = []BlockHook{}

	AddBlockHook(boil.BeforeUpsertHook, blockBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	blockBeforeUpsertHooks = []BlockHook{}

	AddBlockHook(boil.AfterUpsertHook, blockAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	blockAfterUpsertHooks = []BlockHook{}
}

func testBlocksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlocksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(blockColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlockToOneUseremailUsingRequestor(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Block
	var foreign Useremail

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, useremailDBTypes, false, useremailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Useremail struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RequestorID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Requestor().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BlockSlice{&local}
	if err = local.L.LoadRequestor(ctx, tx, false, (*[]*Block)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Requestor == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Requestor = nil
	if err = local.L.LoadRequestor(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Requestor == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBlockToOneUseremailUsingTarget(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Block
	var foreign Useremail

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, useremailDBTypes, false, useremailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Useremail struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TargetID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Target().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BlockSlice{&local}
	if err = local.L.LoadTarget(ctx, tx, false, (*[]*Block)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Target == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Target = nil
	if err = local.L.LoadTarget(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Target == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBlockToOneSetOpUseremailUsingRequestor(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Block
	var b, c Useremail

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, useremailDBTypes, false, strmangle.SetComplement(useremailPrimaryKeyColumns, useremailColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, useremailDBTypes, false, strmangle.SetComplement(useremailPrimaryKeyColumns, useremailColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Useremail{&b, &c} {
		err = a.SetRequestor(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Requestor != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RequestorBlocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RequestorID != x.ID {
			t.Error("foreign key was wrong value", a.RequestorID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RequestorID))
		reflect.Indirect(reflect.ValueOf(&a.RequestorID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RequestorID != x.ID {
			t.Error("foreign key was wrong value", a.RequestorID, x.ID)
		}
	}
}
func testBlockToOneSetOpUseremailUsingTarget(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Block
	var b, c Useremail

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, useremailDBTypes, false, strmangle.SetComplement(useremailPrimaryKeyColumns, useremailColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, useremailDBTypes, false, strmangle.SetComplement(useremailPrimaryKeyColumns, useremailColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Useremail{&b, &c} {
		err = a.SetTarget(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Target != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TargetBlocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TargetID != x.ID {
			t.Error("foreign key was wrong value", a.TargetID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TargetID))
		reflect.Indirect(reflect.ValueOf(&a.TargetID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TargetID != x.ID {
			t.Error("foreign key was wrong value", a.TargetID, x.ID)
		}
	}
}

func testBlocksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBlocksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBlocksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Blocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	blockDBTypes = map[string]string{`ID`: `bigint`, `RequestorID`: `bigint`, `TargetID`: `bigint`}
	_            = bytes.MinRead
)

func testBlocksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(blockAllColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, blockDBTypes, true, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBlocksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(blockAllColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, blockDBTypes, true, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(blockAllColumns, blockPrimaryKeyColumns) {
		fields = blockAllColumns
	} else {
		fields = strmangle.SetComplement(
			blockAllColumns,
			blockPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := BlockSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBlocksUpsert(t *testing.T) {
	t.Parallel()

	if len(blockAllColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Block{}
	if err = randomize.Struct(seed, &o, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, blockDBTypes, false, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err = Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
