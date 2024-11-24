# providers.app

This is the client application (web UI) for providers.app, built using [SvelteKit 5](https://svelte.dev/docs/kit). SvelteKit is used as a client-side application only, and not a full-featured framework including the server. We already have a Go app for that. That's why this repo opts out of many of the server-side features SvelteKit offers.

Below are some common commands to get you going.

## Developing

First do a `yarn install` to get dependencies.

To start the dev server:

```bash
yarn run dev

# or start the server and open the app in a new browser tab
yarn run dev -- --open
```

## Building

```bash
yarn run build
```

You can preview the production build with `yarn run preview`.

## Guidelines and best practices

This section consists of some guidance on how to do specific things with Svelte in this repo. The documentation for the library and framework do not give many examples of best practices, and instead leaves those up to the developer.

### UI

For the UI this repo uses [shadcd-svelte](https://www.shadcn-svelte.com/)

### Forms

Forms use a combination of zod and Superforms for validation, as per shadcd-svelte recommendation. However, we do not utilize server-side form actions for API calls in SvelteKit. We simply do that on the client-side.

### Auth checks and rendering

I like to call this "auth check" instead of "authentication" or "authorization", because we don't do validation for user authentication on the front-end. We simply assume if a user has a string in a specific `localStorage` key, that they are logged in. This is because there is absolutely no point in validating the user on the front-end, as anyone can easily manipulate the code on their machine if they'd like to see the UI. However, they will not get access to data on the server without a valid token.

We use bearer tokens (json web tokens, JWT) for authorization. Tokens are stored inside `localStorage` in the browser, which means they are not accessible from the server (in Svelte). You can't call `isAuthenticated()` from `+page.server.ts` file and expect a result. Instead, we have to do auth checks on the client-side. 

> You could possibly do things like create a client-side session ID cookie that's readable on the server with an expiration time that matches the access token, and keep refreshing that every time you refresh the token, but considering I don't plan on leveraging to many server-side features in SvelteKit, I decided to leave out such functionality

Svelte does not provide guidance on how to actually do this, so we just use our own solution. Svelte only has two lifecycle hooks: creation and destruction. The `onMount` runs so late that users will actually see a page get rendered before the code and redirect inside an `onMount` runs. Because of this we have to take care of not rendering anything on a page if the user is not authenticated.

If you want to protect an entire branch of the render-tree (child routes), you would put this in a `+layout.svelte` file at the root of a route group, that is indicated by braces, example. `(user)/+layout.svelte`.

```svelte
let userLoggedIn = isAuthenticated();
onMount(() => {
    if (!userLoggedIn) {
        goto('/');
    }
});
</script>
<div>
    {#if userLoggedIn}
        <!-- Only render UI if user is authenticated -->
        {@render children()}
    {/if}
</div>
```

