/**
* This is an auto generated code. This code should not be modified since the file can be overwritten
* if new genezio commands are executed.
*/

package genezioSdk




type Hello struct {
	remote *Remote
}

func NewHello() *Hello {
	return &Hello{remote: &Remote{URL: "https://eedtfxekwcz52fo66v65ldhnjm0rdeca.lambda-url.eu-central-1.on.aws/" }}
}
func (f *Hello) SayHello() (string, error) {
    result, err := f.remote.Call("Hello.SayHello")
    var castResult string
    if err != nil {
        return castResult, err
    }
    castResult = result.(string)
	return castResult, err
}
func (f *Hello) SayHelloTo(name string) (string, error) {
    result, err := f.remote.Call("Hello.SayHelloTo", name)
    var castResult string
    if err != nil {
        return castResult, err
    }
    castResult = result.(string)
	return castResult, err
}

