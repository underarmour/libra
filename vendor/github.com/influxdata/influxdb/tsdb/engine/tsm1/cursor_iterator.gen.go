// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: cursor_iterator.gen.go.tmpl

package tsm1

import (
	"context"

	"github.com/influxdata/influxdb/models"
	"github.com/influxdata/influxdb/query"
	"github.com/influxdata/influxdb/tsdb"
)

// buildFloatBatchCursor creates a batch cursor for a float field.
func (q *cursorIterator) buildFloatBatchCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.FloatBatchCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Float == nil {
			q.asc.Float = newFloatAscendingBatchCursor()
		}
		q.asc.Float.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.Float
	} else {
		if q.desc.Float == nil {
			q.desc.Float = newFloatDescendingBatchCursor()
		}
		q.desc.Float.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.Float
	}
}

// buildIntegerBatchCursor creates a batch cursor for a integer field.
func (q *cursorIterator) buildIntegerBatchCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.IntegerBatchCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Integer == nil {
			q.asc.Integer = newIntegerAscendingBatchCursor()
		}
		q.asc.Integer.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.Integer
	} else {
		if q.desc.Integer == nil {
			q.desc.Integer = newIntegerDescendingBatchCursor()
		}
		q.desc.Integer.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.Integer
	}
}

// buildUnsignedBatchCursor creates a batch cursor for a unsigned field.
func (q *cursorIterator) buildUnsignedBatchCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.UnsignedBatchCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Unsigned == nil {
			q.asc.Unsigned = newUnsignedAscendingBatchCursor()
		}
		q.asc.Unsigned.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.Unsigned
	} else {
		if q.desc.Unsigned == nil {
			q.desc.Unsigned = newUnsignedDescendingBatchCursor()
		}
		q.desc.Unsigned.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.Unsigned
	}
}

// buildStringBatchCursor creates a batch cursor for a string field.
func (q *cursorIterator) buildStringBatchCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.StringBatchCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.String == nil {
			q.asc.String = newStringAscendingBatchCursor()
		}
		q.asc.String.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.String
	} else {
		if q.desc.String == nil {
			q.desc.String = newStringDescendingBatchCursor()
		}
		q.desc.String.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.String
	}
}

// buildBooleanBatchCursor creates a batch cursor for a boolean field.
func (q *cursorIterator) buildBooleanBatchCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.BooleanBatchCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Boolean == nil {
			q.asc.Boolean = newBooleanAscendingBatchCursor()
		}
		q.asc.Boolean.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.Boolean
	} else {
		if q.desc.Boolean == nil {
			q.desc.Boolean = newBooleanDescendingBatchCursor()
		}
		q.desc.Boolean.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.Boolean
	}
}
