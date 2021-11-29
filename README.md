# Strings

## What are strings?

String is an immutable (cannot be changed) slice of bytes in Go. Strings can be
created by enclosing a set of characters inside double quotes `" "`:

```go
func main() {
  fmt.Println("Hello, world!") // "Hello, world!" is a string!
}
```

`[]byte` operations apply to `strings` as well:

1. Len

    ```go
    const s = "Hello, world!"
    fmt.Printf("len(%s) = %d", s, len(s))
    // Output:
    // len(Hello, world!) = 13
    ```

    A better version would look like this:

    ```go
    fmt.Printf("len(%s) = %d", s, utf8.RuneCountInString(s))
    // Output:
    // len(Hello, world!) = 13
    ```

    Output is the same, but it's counted differently. I will explain that `utf8`
    below.

2. Range

    ```go
    const s = "Hello, world!"
    for _, c := range s {
      fmt.Print(c, " ")
    }
    fmt.Printf("\n% x", s)
    // Output:
    // 72 101 108 108 111 44 32 119 111 114 108 100 33 // decimal codes
    // 48 65 6c 6c 6f 2c 20 77 6f 72 6c 64 21          // hexadecimal codes
    ```

    Please note that iterating over a string yields `runes`, not `bytes`. I
    will explain that further as well:

    ```go
    const s = "čūska"
    for i, c := range s {
      fmt.Print(i, reflect.TypeOf(c), c)
    }
    // Output:
    // 0 int32 269
    // 2 int32 363
    // 4 int32 115
    // 5 int32 107
    // 6 int32 97
    ```

## Characters

Please check this code snippet:

```go
const s = "čūska"
fmt.Printf("len(%s) = %d", s, len(s))
fmt.Printf("% 2x", s)
// Output:
// len(čūska) = 7
// c4 8d c5 ab 73 6b 61
```

In this example we see the length of the strings isn't 5 as we would expect, but
it is 7 instead. Why? To answer this question, we have to understand what those
bytes are.

### ASCII

ASCII abbreviated from American Standard Code for Information Interchange, is a
character encoding standard for electronic communication. ASCII codes represent
text in computers, telecommunications equipment, and other devices. Most modern
character-encoding schemes are based on ASCII, although they support many
additional characters.

256 bytes each represent a symbol e.g. 'A' is 65 (0x41), 'a' is 97 (0x61), '0' is
48 (0x30), ' ' (space) is 32 (0x20), etc. Range [0..31] is reserved for special
control characters e.g. new line 10 (0x0a), tab 9 (0x09), etc. Those characters
used to be recognized by old typewriting machines.

### Unicode

Symbols starting from 128 (0x80) belong to an extended set it could be changed
to contain russian or latvian characters. But how can we have both in the a
single string? It's simple - we use single byte encoded strings which actually
contain symbols above 255 (256 different values) allowed by ASCII.

Unicode, formally the Unicode Standard, is an information technology standard
for the consistent encoding, representation, and handling of text expressed in
most of the world's writing systems. The standard, which is maintained by the
Unicode Consortium, defines 144,697 characters covering 159 modern and historic
scripts, as well as symbols, emoji, and non-visual control and formatting codes.

Go uses UTF-8 encoding, and some of the Go co-authors are also co-authors of
this encoding. So Go does it differently to other languages and it allows to
treat strings as a slice of bytes or a slice of `runes`. These are essentualy
ints that contain symbol codes from different alphabets. E.g. Russian 'А' code
is 1040 (U+0410), and Latvian 'č' is 269 (U+010D).

```go
const s = "čūska"
fmt.Println([]byte(s))
fmt.Println([]rune(s))
// Output:
// [196 141 197 171 115 107 97]
// [269 363 115 107 97]
```

## Creating a string from bytes or runes

```go
s := []byte{65, 66, 67}
fmt.Println(string(s))
// Output: ABC

s := []rune{269, 'ū', 115, 107, 'a'}
fmt.Println(string(s))
// Output: čūska
```
