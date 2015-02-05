Contributing to GoDuckGo
========================

If you love Go and DuckDuckGo, then this is
your place. We're waiting for your Pull Request!

Getting Started
---------------

Before you can do anything, you first need a [GitHub account](https://github.com/signup/free).
This is required because we use GitHub to handle all incoming *Pull Requests* (code modifications)
and *Issues* (bug reports) which cannot be made without a GitHub account.

Submitting a **Bug** or **Suggestion**
--------------------------------------

- Firstly, please make sure the bug is related to the **GoDuckGo** package. If this bug
is about the DuckDuckGo API, or the relevancy of the search results, please visit DuckDuckGo's
feedback page at <https://duckduckgo.com/feedback>.

- Check the **GoDuckGo** [issues](https://github.com/ajanicij/goduckgo/issues) to see if
an issue already exists for the given bug or suggestion:
  - If one doesn't exist, create a GitHub issue in the **GoDuckGo** repository:
    - Clearly describe the bug/improvement, including steps to reproduce when it is a bug.
  - If one already exists, please add any additional comments you have regarding the matter.

If you're submitting a **pull request** (bugfix/addition):
- Fork the **GoDuckGo** repository on GitHub.

Making Changes
--------------

- Before making any changes, make sure your [Go environment](http://golang.org/doc/install) is setup.
- Run `gofmt` and `go vet` commands to clean up your code.
- Make sure your commits are of a reasonable size. They shouldn't be too big.
- Make sure your commit messages effectively explain what changes have been made.

  ```shell
  main.go: Handle error when Icon is empty
  ```

  is much better than:

  ```shell
  annoying error for empty icon is fixed now
  ```

- Make sure you have added the necessary tests for your changes.
- Make sure your change doesn't affect backwards compatibility.

Submitting Changes
------------------

1. Commit your changes.

  ```shell
  git commit -am "My first commit!"
  ```

2. Get your commit history [how you like it](http://book.git-scm.com/4_interactive_rebasing.html).

  ```shell
  git rebase -i origin/master
  ```

  or

  ```shell
  git pull --rebase origin/master
  ```

3. Push your forked repository back to GitHub.

  ```shell
  git push
  ```

4. Add your info to the [AUTHORS.md page](https://github.com/ajanicij/goduckgo/blob/master/AUTHORS.md).

5. Go into GitHub and submit a [pull request!](http://help.github.com/send-pull-requests/) to the
**GoDuckGo** repository.

