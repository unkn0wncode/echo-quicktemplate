// This file is automatically generated by qtc from "base.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line base.qtpl:1
package templates

//line base.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// This is a base page template. All the other template pages implement this interface.
//

//line base.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line base.qtpl:4
type Page interface {
	//line base.qtpl:4
	Body() string
	//line base.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
	//line base.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line base.qtpl:4
}

// Page prints a page implementing Page interface.

//line base.qtpl:11
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
	//line base.qtpl:11
	qw422016.N().S(`
<html>
	<head>
		<title>Echo+Quicktemplate</title>
	</head>
	<body>
		<div>
			Some fancy menu
		</div>
		`)
	//line base.qtpl:20
	p.StreamBody(qw422016)
	//line base.qtpl:20
	qw422016.N().S(`
	</body>
	<footer>
		Some fancy footer
	</footer>
</html>
`)
//line base.qtpl:26
}

//line base.qtpl:26
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
	//line base.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line base.qtpl:26
	StreamPageTemplate(qw422016, p)
	//line base.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:26
}

//line base.qtpl:26
func PageTemplate(p Page) string {
	//line base.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
	//line base.qtpl:26
	WritePageTemplate(qb422016, p)
	//line base.qtpl:26
	qs422016 := string(qb422016.B)
	//line base.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
	//line base.qtpl:26
	return qs422016
//line base.qtpl:26
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line base.qtpl:31
type Base struct{}

//line base.qtpl:32
func (p *Base) StreamBody(qw422016 *qt422016.Writer) {
//line base.qtpl:32
qw422016.N().S(`This is a base body`) }

//line base.qtpl:32
//line base.qtpl:32
func (p *Base) WriteBody(qq422016 qtio422016.Writer) {
	//line base.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line base.qtpl:32
	p.StreamBody(qw422016)
	//line base.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:32
}

//line base.qtpl:32
func (p *Base) Body() string {
	//line base.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
	//line base.qtpl:32
	p.WriteBody(qb422016)
	//line base.qtpl:32
	qs422016 := string(qb422016.B)
	//line base.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
	//line base.qtpl:32
	return qs422016
//line base.qtpl:32
}
