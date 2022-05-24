Usage of CreateAttribute methods

Attributes that are available:
- string
- Enum
- Email
- Ip
- Url
- Integer
- Float
- Boolean

Note: for the optional parameters, there is new type created (since GO doesn't have optional or default arguments). The Optional type is defined as a struct with fields:

- Value -> this contains the value
- Specified -> wether or not it is specified (defaults to false)

Optional[T]{} -> replace T with the desired type (string, int, ...)
empty optional = will see as "optional value not set"

Example Use of the CreateAttribute functions:

CreateStringAttribute("unique()", "StringName", 15, false, Optional[string]{Value: "MyDefaultString",Specified: false}, Optional[bool]{}) 
This wil create a String attribute with name "StringName" that is max. 15 characters long and has default value "MyDefaultString". Since the isArray is an empty Optional[bool], the isArray bool will be false -> this is not an array
