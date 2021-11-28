# TODO CLI
Highly configurable, editor-agnostic, todo manager written in [Go](https://go.dev/).
Inspired and cloned from [snitch](https://github.com/tsoding/snitch/). Go, give it a star <3

## How to build?
```
git clone https://github.com/NazgoooAtanasov/todoCLI
cd todoCLI
go build .
```
This will produce a single executable. I suggest adding the repo path to the system `$PATH` for global usage.

## How to configure?
The application supports only one way of extending and that is with a config file `.todocli.yaml`. In the repo there is an example config file [.todocli-example.yaml](https://github.com/NazgoooAtanasov/todoCLI/blob/main/.todocli-example.yaml)

### Configure options
```yaml
directory:
  forbidendirectories:
    - node_modules
    - .git
    - forbiden

pattern:
  keyword: 'custom-keyword-here'
  urgencysuffix: 'custom-urgency-suffix-here'
  commenttype: 'language-specific-comment-type'
  customregex: 'custom-parsing-regex-(WIP)'
```
The `directory` object is responsible for configurations for the directory structure. For now we can just ignore certain directories with the `forbidendirectories` list.

The `pattern` object is responsible for the parsing options:
-   `keyword` field is a `string` value and it is used for determining what word should be used as a start of a todo, defaults to `TODO` if not specified.
-   `urgencysuffix` field is a `string` value and it is used to sort the todos by urgency based on how much times the given suffix occurs after the `keyword`, defaults to `O` if not specified.
-   `commenttype` field is a `string` value and it is used to set the comment type in different languages, example would be C with its comments starting with `//`, and python's staring with `#`
-   `customregex` WIP

With all of the default settings in place, the todos should be in the following format:

```javascript
// TODO: fix that
console.log('beautiful cli');
```

## How does the urgency logic work?
It counts the ammount of times you have written the `urgencysuffix` after the `keyword`. Higher the count - higher the urgency.
```
// TODO: FIXME! - the least urgent todo
// TODOO: FIXME FAST! - one step higher than the least urgent todo
// TODOOO: FIXME FASTER!
// TODOOOO....
```

## How is the todo actually parsed?
It is getting parsed vie a regex expression. The default one looks like this.
```
^\\s*// (TODO)(O*): ([a-zA-Z ]*)$
```

If you specified different `keyword` and/or `urgencysuffix` and/or `commenttype` the regex will be looking a bit differently.
```
^\\s-commenttype- (-keyword-)(-urgencysuffix-*): ([a-zA-Z ]*)$
```
