// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base(children ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Sample To-Dos</title><link href=\"/css/tailwind.css\" rel=\"stylesheet\"></head><body class=\"m-0 p-0 w-screen h-screen overflow-hidden\"><div class=\"container mx-auto flex justify-center items-center h-full\"><label class=\"fixed top-3 right-3 flex cursor-pointer gap-2\"><i data-feather=\"sun\"></i> <input type=\"checkbox\" id=\"theme\" name=\"theme\" value=\"synthwave\" class=\"toggle theme-controller\" x-data x-on:change=\"$store.theme.toggle()\" x-data x-bind:checked=\"$store.theme.theme === &#39;synthwave&#39;\"> <i data-feather=\"moon\"></i></label> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, child := range children {
			templ_7745c5c3_Err = child.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><script src=\"https://unpkg.com/feather-icons\"></script><script src=\"//unpkg.com/alpinejs\" defer></script><script>\n\t\t\t\tfeather.replace();\n\n\t\t\t\t// Re-run feather icons after htmx swaps\n\t\t\t\thtmx.on(\"htmx:afterSwap\", function(event) {\n\t\t\t\t\tfeather.replace();\n\t\t\t\t});\n\n\t\t\t\tdocument.addEventListener('alpine:init', () => {\n\t\t\t\t\tAlpine.store('theme', {\n\t\t\t\t\t\tinit() {\n\t\t\t\t\t\t\tlet stored = localStorage.getItem('theme');\n\n\t\t\t\t\t\t\tif (!stored) {\n\t\t\t\t\t\t\t\tstored = 'cupcake'\n\t\t\t\t\t\t\t\tlocalStorage.setItem('theme', stored);\n\t\t\t\t\t\t\t}\n\n\t\t\t\t\t\t\tif (stored === 'cupcake') {\n\t\t\t\t\t\t\t\tthis.theme = 'cupcake';\n\t\t\t\t\t\t\t\tdocument.documentElement.classList.remove('dark');\n\t\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\t\tthis.theme = 'synthwave';\n\t\t\t\t\t\t\t\tdocument.documentElement.classList.add('dark');\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t},\n\n\t\t\t\t\t\ttheme: 'cupcake',\n\n\t\t\t\t\t\ttoggle() {\n\t\t\t\t\t\t\tthis.theme = this.theme === 'synthwave' ? 'cupcake' : 'synthwave';\n\t\t\t\t\t\t\tdocument.documentElement.classList.toggle('dark');\n\t\t\t\t\t\t\tlocalStorage.setItem('theme', this.theme);\n\t\t\t\t\t\t}\n\t\t\t\t\t})\n\t\t\t\t})\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}