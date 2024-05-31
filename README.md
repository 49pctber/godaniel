# Go, Daniel

Daniel's daily affirmations.

## About

This project originally started as a joke for my friend Daniel.
He can run this program each day to get some ✨good vibes✨ in the form of three randomly-selected affirmations.

I have updated this project to include any name so that everyone can get some ✨good vibes✨ on the daily.

## Installation

```
go install github.com/49pctber/godaniel/cmd/...@latest
```

This will install two executables: `godaniel` and `godaniel_server`.
`godaniel` is a CLI application that will print your affirmations in the console.
`godaniel_server` starts a server to create a prettier interface that can be viewed in a broswer.

The default name is Daniel, but you can specify any name using `--name` flag.
Note that different names will result in different daily affirmations.
(i.e. Bryan will likely have different affirmations than Daniel even on the same day.)

For the server, you can specify a different port for the server using the `--port` flag.

## Example Output

```
~$ godaniel

🌞 Good morning, Daniel!

  - You are worthy of living your best life.
  - You are on the right path.
  - You are unique and irreplaceable.

🎉 GO, DANIEL! 🎉

```
