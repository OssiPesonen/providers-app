This is a list of things and technologies I wish to try out:

| Technology / Task               | Done                |
|---------------------------------|---------------------|
| gRPC                            | :white_check_mark:  |
| Authentication                  | :white_check_mark:  |
| Tests (and transition to TDD)   | :hammer:            |
| Front-end                       | :hammer:            |

## Notes

- I can consider authentication complete as I've completed access + refresh tokens, with the ability to revoke them. I wanted to use sessions for simplicity, but gRPC doesn't work for that. If I wanted additional layers of protection, I would use a third party platform that offers things like automatic reuse detection. I can mitigate the window of malicious use by lowering access token lifetime and offering the user a button to log out of all devices, which would revoke their active refresh tokens.
- ~A monorepo setup and a build tool for a project of this size is overkill. It introduces unnecessary complexity and I don't have the need to share code in one repo. I could change the git branching strategy and makefile to build specific apps on changes to specific branches, but currently I just want to try those tools out.~ I burned out my capacity to prototype more tools with grpc-web and trying to make a successful request from the browser. To do this I had to install Envoy proxy, because browsers do not natively support gRPC trailers or have an API to force HTTP/2, and for some reason the network requests in Windows do not go through to the upstream service, but I got them working on OS X. Not really sure what's happening there. Envoy also did not work on OS X with Podman due to some permission issue that seemed to have no solutions, so I installed Rancher Desktop and that worked.
- ~I am planning on trying Svelte, but a better medium for this would be a mobile app. I am just a bit bored of Flutter, and this is just for fun, so I'm most likely not gonna do it.~ I had to toss the idea of using Svelte because it now runs with Vite and I just did not have the patience at that moment to start fighting with it to support CommonJS. Unfortunately proto buffer files for the web are compiled using commonjs and not ESM. The support is still not there. I quickly switched to Next.js and got it working in 5 minutes.