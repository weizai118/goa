package design

import "github.com/goadesign/goa/eval"

type (
	// ServiceExpr describes a set of related endpoints.
	ServiceExpr struct {
		// DSLFunc contains the DSL used to initialize the expression.
		eval.DSLFunc
		// Name of endpoint group.
		Name string
		// Description of endpoint group for consumption by humans.
		Description string
		// Docs points to external documentation
		Docs *DocsExpr
		// Endpoints is the list of service endpoints.
		Endpoints []*EndpointExpr
		// DefaultType is the service default response type. The default
		// type attributes also define the default properties for
		// request attributes with identical names.
		DefaultType UserType
	}
)

// EvalName returns the generic expression name used in error messages.
func (s *ServiceExpr) EvalName() string {
	if s.Name == "" {
		return "unnamed service"
	}
	return "service " + s.Name
}
