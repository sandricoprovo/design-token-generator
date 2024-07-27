Fran is undergoing a rebuild, so excuse the mess!

# Fran CLI

## About the Project

Fran is a CLI tool built in Go with the goal of making it easier to generate and update design tokens in our front end projects. The problem Fran tries to solve is enabling you to create tokens and make updates without having to reach for multiple third party web apps/services. They're great, but sometime it's nice to stay on the command line.

For an example use case, lets say you have 25 hex colours you want to convert to rgb. Maybe your editor has a built in color changer like VScode, or you know a good web app. Instead of copying and pasting or changing each color one by one, Fran can walk your file(s) and convert those colour codes in one command.

## How to Install

## Commands

```
fran convert
```

-   Converts any supported color found in a file or directory to the targeted format.
    -   Supported colours:
        -   hex
        -   rgb
        -   oklch
    -   Supported File Types:
        -   `.css`
        -   `.scss`
        -   `.svelte`
    -   Flags:
        -   `--format / -f`: _Required_. The target format that other colours will be converted to.
        -   `--path / -p`: _Optional_. The target file. If this flag is omitted, fran will traverse your entire project directory and convert colours within the supported file formats.

##### Tips - Buy me (or my dog) A Snack 👨🏾‍💻
