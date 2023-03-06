# file_populator

A quick tool to take a json list and create a list of .md files from it. Essentially the headache of Glossary made me do it. 🥲

### Example of a json file

```json
[
    {
        "title": "My First File",
        "body": "My first Body"
    },
    {
        "title": "My Second File",
        "body": "My Second Body"
    },
    {
        "title": "really weird n@me!",
        "alt-title": "my-third-file.md",
        "body": "My Third Body"
    }
]
```

The structure of each entry is as follows:

* `title`: Required field
* `body`: Required field
* `alt-title`: Optional field for titles that are structured weirdly, this only affects the file name and not the content

### Usage

* `-file`: The json file with the list of the items to add
* `-folder`: The folder you want all the files to go into, should not end with a slash. Just the name of the folder