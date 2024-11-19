import { defineConfig } from "vitest/config";
import { sveltekit } from '@sveltejs/kit/vite';
import commonjs from 'vite-plugin-commonjs'

export default defineConfig({
    plugins: [sveltekit(), commonjs()],

    test: {
        include: ['src/**/*.{test,spec}.{js,ts}']
    }
});
