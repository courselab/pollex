# Contributing to this project

If you'd like to contribute to this project, please, read this document.

## Essencial notes

 To keep things consistent, we adhere to some standards

 - GitFlow branching strategy [1] (plus a prerelease branch)

 - Semantic versioning 2.0.0 [2]

 - Conventional Commits 1.0.0 [3] (see types bellow)

 - Keep a ChangeLog [4]
 
### Submitting changes

 First create and issue, if one does not exist yet (label it accordingly).
 
 Develop your contribution in a support branch, following the naming scheme
 
 'type/issue-number/short-descriptive-annotation'

 where 'type' is a conventional commit type (see below).
 
 Submit your support branch and mark it as a pull/merge request.
 
 Remember, pull requests
 
   - for regular development should branch off and merge into branch `develop``
   - only hot fixes branch off and merge into `main` (and `develop`)
   - must be reviewed and approved by other developers
 
 ### Commit types
 
 Select the appropriate type.
 
 - fix: fix a bug
 - feat: add new feature
 - build: affect the build 
 - perf: improve code (other than fix and feat)
 - doc: modify internal or external modification
 - test: modify tests
 - tidy: code style, repo organization, standard compliance etc.
 - tmp: a temporary branch for some other purpose

### Further directions

Latest stable releases reside in the _main_ branch.

Prerelease (alpha, beta, release candidates) reside on branch _prerelease_.

If you are a developer and is assigned an issue but you believe you are not
able to handle timely, please, try to reassign it to someone else.

AUTHORS, NEWS and ChangeLog files should be kept up-to-date.

Symbol and comments in the source code must be in English.

If possible, please use English also for issues, discussions etc.

It should be needless to say, but do not commit unnecessary files.

## References

[1] https://nvie.com/posts/a-successful-git-branching-model/

[2] https://semver.org/

[3] https://www.conventionalcommits.org/en/v1.0.0/

[4] https://keepachangelog.com/en/1.0.0/


