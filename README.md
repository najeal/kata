# kata

You can find exercises **Karate Chop**, **Anagrams** and **Word Chains**<br/>
All given commands are assuming you are in root folder

## run tests

```go test -v ./...```

## run benchmark for karatechop implementations

```go test -bench=. ./pkg/karatechop```

## run somes examples


Install kata

```go install```

Anagrams example with litle data set

```kata anagrams```

Anagrams example with big data set

```kata anagrams -f ./example_files/wordlist.txt```

Word Chains

```kata chains word city```

```kata chains cat dog```

```kata chains ruby code```
