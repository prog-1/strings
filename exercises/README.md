# Exercises: Strings

All exercises should contain a runnable application and tests!

Every program must contain tests and must be created in a dedicated folder.
Program names are provided in the task description.

Use the following snippet in your `main` to enter strings from the keyboard:

```go
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
```

## 01 Trim Right

Remove trailing whitespaces. Use " \n\r\t" as a whitespace set. Please note,
it would be better to use [unicode.IsSet] istead, but let's not complicate it.

The program must be named `01_trim_right`, and must contain a function
`TrimRight(string) string`.

Hint: Use `01_trim` program as an example.

[unicode.IsSpace]: https://pkg.go.dev/unicode#IsSpace

## 02 Palindrom

Write a program which checks whether a given string is a palindrom (could be
read left-to-right and right-to-left in the same way).

The program must be named `02_check_palindrom`, and must contain a function
`IsPalidndrom(string) bool`.

## 03 Count byte

Write a program which counts how may times a given by could be found in a
string. E.g. `hello` contains byte `'l'` (ASCII 108 or 0x6c) 2 times.

The program must be named `03_count_byte`, and must contain a function
`CountByte(string) int`.

## 04 Count rune

The same as `03 Count byte`, but counts runes.

The program must be named `03_count_rune`, and must contain a function
`CountRune(string) int`.

# 05 ToLower & ToUpper

Write a program, which ensures a string contains English alphabet letters
all lower, and all upper cases e.g. `"hEllO"` should become `"hello"` and
`"HELLO"`.

The program must be named `05_lower_uppper`, and must contain two functions:
- `ToLower(string) string`
- `ToUpper(string) string`.

# 08 Count words

Write a program which counts words in a string. Words are considered to be
separated by punctuation and/or spaces. See [unicode.IsSpace] and
[unicode.IsPunct] functions.

The function must be named `06_count_words`, and must contain a function
`CountWords(string) int`.

[unicode.IsSpace]: https://pkg.go.dev/unicode#IsSpace
[unicode.IsPunct]: https://pkg.go.dev/unicode#IsPunct

# 07 Title

Write a program, which makes every first letter of every word capital, while
all other letters should become lower case. E.g. `"heLLO WORLD"` should become
`"Hello World"`.

The program must be named `07_title`, and must contain a `Title(string) string`
function.

Note: The program must be implemented using [unicode.ToTitle] and
[unicode.ToLower] functions.

[unicode.ToTitle]: https://pkg.go.dev/unicode#ToUpper
[unicode.ToLower]: https://pkg.go.dev/unicode#ToLower

# 08 Count substrings

Write a program which finds how many times a substring exists in a string with
overlaps. E.g. `"bananas"` contains `"ana"` 2 times.

The program must be named `08_count_substrings`, and must contain a function
`Count(string) int`.

# 09 Dedupe

Write a program which dedupes (removes duplicates) a string by dropping
consequent characters if they are the same e.g. `"000111110001111"` will become
`"0101"`.

The program must be named `09_dedup`, and must contain a function
`Dedup(string) string`.