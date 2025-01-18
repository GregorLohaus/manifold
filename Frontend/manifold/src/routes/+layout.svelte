<script lang="ts">
	import '@xyflow/svelte/dist/style.css';
  import "../app.css"
  import Button from "$lib/components/ui/button/button.svelte"
  import { SunSolid, MoonSolid } from "flowbite-svelte-icons"
  import { ModeWatcher, toggleMode } from "mode-watcher"
  import { onMount, setContext } from "svelte"
	import { DefaultApiPath } from "$lib/default_api_url";
	import { goto } from "$app/navigation";
  import { page } from "$app/stores"
	import { writable } from 'svelte/store';
  onMount(
    async() => {
      if ($page.route?.id == "/login" || $page.route?.id == "/register" || $page.route?.id == "/verify") {
        return
      }
      fetch(DefaultApiPath("authstatus"),{credentials: "include"}).then((res) => {
        if (!res.ok) {
          goto("/login")
        }
      })
    }
  )
  setContext("graphid",writable(""))
</script>

<title>Manifold</title>
<ModeWatcher/>
<Button on:click={toggleMode} class="absolute bottom-5 right-5 z-40">
    <div class="dark:scale-0"><SunSolid/></div>
    <div class="absolute scale-0 dark:scale-100"><MoonSolid/></div>
</Button>
<div class="flex flex-col flex-wrap justify-center content-center h-screen py-5">
<slot/>
</div>
