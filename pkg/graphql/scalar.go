package graphql_pkg

//
//import (
//	"github.com/99designs/gqlgen/graphql"
//	"github.com/vektah/gqlparser/v2/ast"
//	"time"
//)
//
//var JSONScalar = graphql.NewScalar(graphql.ScalarConfig{
//	Name:        "JSON",
//	Description: "The `JSON` scalar type represents JSON values.",
//	Serialize: func(value interface{}) interface{} {
//		return value
//	},
//	ParseValue: func(value interface{}) interface{} {
//		return value
//	},
//	ParseLiteral: func(value ast.Value) interface{} {
//		switch value := value.(type) {
//		case *ast.StringValue:
//			return value.Value
//		default:
//			return nil
//		}
//	},
//})
//
//var DurationScalar = graphql.NewScalar(graphql.ScalarConfig{
//	Name:        "Duration",
//	Description: "The `Duration` scalar type represents a duration in nanoseconds.",
//	Serialize: func(value interface{}) interface{} {
//		if duration, ok := value.(time.Duration); ok {
//			return duration.Nanoseconds()
//		}
//		return nil
//	},
//	ParseValue: func(value interface{}) interface{} {
//		if nanoseconds, ok := value.(int64); ok {
//			return time.Duration(nanoseconds)
//		}
//		return nil
//	},
//	ParseLiteral: func(value ast.Value) interface{} {
//		if intValue, ok := value.(*ast.IntValue); ok {
//			nanoseconds, err := intValue.Value.Int64()
//			if err == nil {
//				return time.Duration(nanoseconds)
//			}
//		}
//		return nil
//	},
//})
//
//var TimeScalar = graphql.NewScalar(graphql.ScalarConfig{
//	Name:        "Time",
//	Description: "The `Time` scalar type represents a point in time.",
//	Serialize: func(value interface{}) interface{} {
//		if t, ok := value.(time.Time); ok {
//			return t.Format(time.RFC3339)
//		}
//		return nil
//	},
//	ParseValue: func(value interface{}) interface{} {
//		if str, ok := value.(string); ok {
//			t, err := time.Parse(time.RFC3339, str)
//			if err == nil {
//				return t
//			}
//		}
//		return nil
//	},
//	ParseLiteral: func(value ast.Value) interface{} {
//		if strValue, ok := value.(*ast.StringValue); ok {
//			t, err := time.Parse(time.RFC3339, strValue.Value)
//			if err == nil {
//				return t
//			}
//		}
//		return nil
//	},
//})
//
