This is a list of things and technologies I wish to try out:

| Technology / Task               | Done                |
|---------------------------------|---------------------|
| gRPC                            | :white_check_mark:  |
| Authentication                  | :white_check_mark:  |
| Tests (and transition to TDD)   | :hammer:            |
| Monorepo setup (new tool)       |                     |
| Front-end with maybe Svelte (?) |                     |

## Notes

- I can consider authentication complete as I've completed access + refresh tokens, with the ability to revoke them. I wanted to use sessions for simplicity, but gRPC doesn't work for that. If I wanted additional layers of protection, I would use a third party platform that offers things like automatic reuse detection. I can mitigate the window of malicious use by lowering access token lifetime and offering the user a button to log out of all devices, which would revoke their active refresh tokens.
- A monorepo setup and a build tool for a project of this size is overkill. It introduces unnecessary complexity and I don't have the need to share code in one repo. I could change the git branching strategy and makefile to build specific apps on changes to specific branches, but currently I just want to try those tools out.
- I am planning on trying Svelte, but a better medium for this would be a mobile app. I am just a bit bored of Flutter, and this is just for fun, so I'm most likely not gonna do it.