// +build ignore

package gorules

import (
  "github.com/quasilyte/go-ruleguard/dsl"
  "github.com/quasilyte/go-ruleguard/dsl/types"
)

func implementsStringer(ctx *dsl.VarFilterContext) bool {
  stringer := ctx.GetInterface(`fmt.Stringer`)
  return types.Implements(ctx.Type, stringer) ||
    types.Implements(types.NewPointer(ctx.Type), stringer)
}

func implementsWorker(ctx *dsl.VarFilterContext) bool {
  worker := ctx.GetInterface(`github.com/sebastien-rosset/rg1/Worker`)
  return types.Implements(ctx.Type, worker) ||
    types.Implements(types.NewPointer(ctx.Type), worker)
}

func stringerLiteral(m dsl.Matcher) {
  m.Match(`$x{$*_}`).
    Where(m["x"].Filter(implementsStringer)).
    Report("$x implements stringer")
}

func workerLiteral(m dsl.Matcher) {
  m.Match(`$x{$*_}`).
    Where(m["x"].Filter(implementsWorker)).
    Report("$x implements worker")
}

