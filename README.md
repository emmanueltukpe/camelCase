
# Camel Case

The program converts snake case strings to camel case, pascalCase or uppercase.

## Examples


camelCase("foo_bar"); //=> "fooBar"

camelCase("foo_bar"); //=> "fooBar"

camelCase("Foo-Bar"); //=> "fooBar"

camelCase("Foo-Bar", { pascalCase: true }); //=> "FooBar"

camelCase("--foo.bar", { pascalCase: false }); //=> "fooBar"

camelCase("Foo-BAR", { uppercase: true }); //=> "fooBAR"

camelCase("fooBAR", { pascalCase: true, uppercase: true })); //=> "FooBAR"

camelCase("foo bar"); //=> "fooBar"
