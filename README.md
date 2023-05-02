# Denoken

A simple, config based design token generator.

#### Works On

-   Intel Macs
-   Apple Silicon Macs
-   Linux

## What is Denoken?

Denoken is a css file generator that builds out parts of a design system you'd use in a project. By using a simple `JSON` config file you can set the parameters that'll generate fleshed out parts of a system for you, instead of doing a lot of the work manually each time you start a project. For now, denoken has a single `generate` command that generates a css file containing font variables which you can then use within your project.

## How to Install

Install denoken via the homebrew command below.

```
brew install denoken
```

## Configuration

To configure denoken create a `denoken.config.json` file in the root of your project, or within a `/config` folder at the root. An example config looks like this:

```json
{
    "path": "global.css",
    "typeScale": {
        "base": 16,
        "multiplier": 1.414,
        "shrink": 0.6,
        "steps": {
            "small": 2,
            "large": 5
        }
    }
}
```

You can find an example css file at `/global.example.css` in the root of this repo.

**Options**

-   `path`: Where the generated css file will be placed when using the `generate` command. You don't need to put a leading `/` in this path, and you can also nest the file in a folder by using a path like `styles/global.css`.
-   `typeScale`: Configuration for generating font variables.
    -   `base`: The base font size for the type.
    -   `multiplier`: The scale used to create font sizes from the base.
    -   `shrink`: Used to calculate the smallest font size in css clamps. Using the example config, fonts will stop getting smaller when they hit 60% of their original size.
    -   `steps`: The small and large steps tell the generator how many font sizes you want that are smaller and larger than your base size.

## Commands

These are the commands you can use with Denoken and what they do.

**_Generate_**

```
denoken generate
```

This command generates your css file at the location of your `denoken.config.json` files' `path` variable . It uses the `typeScale` object within the config to generate a css file with fleshed out font variables using css custom properties and css clamp.

## Why Denoken?

I built denoken because I wanted an easy way to build out all of the starter variables I use when designing my side projects. I always found myself going from web app to web app generating these pieces. For example, most projects I build have some kind of frontend, and so I always went to type-scale.com to get the font sizes and then spend 15-30 minutes building out usable css variables for fonts. My motivation for denoken is to streamline that process down to a config file and a terminal command. Having denoken as a global command installed via homebrew makes it easy to reuse across my projects since it just needs a config file.
