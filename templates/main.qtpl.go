// This file is automatically generated by qtc from "main.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line main.qtpl:1
package templates

//line main.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// Main page template. Implements BasePage methods.
//

//line main.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line main.qtpl:5
type Main struct {
	//CTX *fasthttp.RequestCtx
	Vars map[string]interface{}
}

//line main.qtpl:12
func (p *Main) StreamBody(qw422016 *qt422016.Writer) {
	//line main.qtpl:12
	qw422016.N().S(`
	<h1>Main Page</h1>
	<div>
		Some Body-Content
	</div>
`)
//line main.qtpl:17
}

//line main.qtpl:17
func (p *Main) WriteBody(qq422016 qtio422016.Writer) {
	//line main.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line main.qtpl:17
	p.StreamBody(qw422016)
	//line main.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:17
}

//line main.qtpl:17
func (p *Main) Body() string {
	//line main.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line main.qtpl:17
	p.WriteBody(qb422016)
	//line main.qtpl:17
	qs422016 := string(qb422016.B)
	//line main.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line main.qtpl:17
	return qs422016
//line main.qtpl:17
}