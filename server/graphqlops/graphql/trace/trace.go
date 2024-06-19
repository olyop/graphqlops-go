// The trace package provides tracing functionality.
// Deprecated: this package has been deprecated. Use package trace/tracer instead.
package trace

import (
	"github.com/olyop/graphqlops-go/graphqlops/graphql/errors"
	"github.com/olyop/graphqlops-go/graphqlops/graphql/trace/noop"
	"github.com/olyop/graphqlops-go/graphqlops/graphql/trace/opentracing"
	"github.com/olyop/graphqlops-go/graphqlops/graphql/trace/tracer"
)

// Deprecated: this type has been deprecated. Use [tracer.QueryFinishFunc] instead.
type TraceQueryFinishFunc = func([]*errors.QueryError)

// Deprecated: this type has been deprecated. Use [tarcer.FieldFinishFunc] instead.
type TraceFieldFinishFunc = func(*errors.QueryError)

// Deprecated: this interface has been deprecated. Use [tracer.Tracer] instead.
type Tracer = tracer.Tracer

// Deprecated: this type has been deprecated. Use [opentracing.Tracer] instead.
type OpenTracingTracer = opentracing.Tracer

// Deprecated: this type has been deprecated. Use [noop.Tracer] instead.
type NoopTracer = noop.Tracer
