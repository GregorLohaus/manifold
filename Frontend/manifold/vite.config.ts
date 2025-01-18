import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import dns from 'node:dns'

dns.setDefaultResultOrder('ipv4first')
export default defineConfig({
	plugins: [sveltekit()]
});
